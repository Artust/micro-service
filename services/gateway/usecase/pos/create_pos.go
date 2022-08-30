package pos

import (
	"avatar/services/gateway/config"
	pbCenter "avatar/services/gateway/protos/center"
	pbPos "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

type CreatePosInput struct {
	Name                    string   `json:"name" binding:"required"`
	ServiceName             string   `json:"serviceName" binding:"required"`
	ServiceCategory         string   `json:"serviceCategory" binding:"required"`
	ServiceDetail           string   `json:"serviceDetail" binding:"required"`
	ShopId                  int64    `json:"shopId" binding:"required"`
	CenterId                int64    `json:"centerId" binding:"required"`
	ServiceTemplateId       int64    `json:"serviceTemplateId" binding:"required"`
	MaleRoutineIds          []int64  `json:"maleRoutineIds" binding:"required"`
	DefaultMaleRoutineIds   []int64  `json:"defaultMaleRoutineIds" binding:"required"`
	FemaleRoutineIds        []int64  `json:"femaleRoutineIds" binding:"required"`
	DefaultFemaleRoutineIds []int64  `json:"defaultFemaleRoutineIds" binding:"required"`
	DefaultAvatarId         int64    `json:"defaultAvatarId" binding:"required"`
	UseServiceTemplate      bool     `json:"useServiceTemplate"`
	Backgrounds             []string `form:"backgrounds" binding:"required"`
	DefaultBackground       string   `form:"defaultBackground" binding:"required"`
}

type CreatePosOutput struct {
	Id                      int64    `json:"id"`
	Name                    string   `json:"name"`
	ServiceName             string   `json:"serviceName"`
	ServiceCategory         string   `json:"serviceCategory"`
	ServiceDetail           string   `json:"serviceDetail"`
	ShopId                  int64    `json:"shopId"`
	CenterId                int64    `json:"centerId"`
	ServiceTemplateId       int64    `json:"serviceTemplateId"`
	MaleRoutineIds          []int64  `json:"maleRoutineIds"`
	DefaultMaleRoutineIds   []int64  `json:"defaultMaleRoutineIds"`
	FemaleRoutineIds        []int64  `json:"femaleRoutineIds"`
	DefaultFemaleRoutineIds []int64  `json:"defaultFemaleRoutineIds"`
	DefaultAvatarId         int64    `json:"defaultAvatarId"`
	UseServiceTemplate      bool     `json:"useServiceTemplate"`
	Backgrounds             []string `json:"backgrounds"`
	DefaultBackground       string   `json:"defaultBackground"`
	CreatedBy               int64    `json:"createdBy"`
	CreatedAt               string   `json:"createdAt"`
	UpdatedAt               string   `json:"updatedAt,omitempty"`
}

func CreatePos(
	input *CreatePosInput,
	posClient pbPos.POSClient,
	centerClient pbCenter.CenterClient,
	cfg *config.Environment,
) (*CreatePosOutput, error) {
	ctx := context.Background()
	var backgrounds []string
	for _, val := range input.Backgrounds {
		splBackground := strings.LastIndex(val, "/")
		backgrounds = append(backgrounds, val[splBackground+1:])
	}
	data := &pbPos.CreatePosRequest{
		Name:                    input.Name,
		ServiceName:             input.ServiceName,
		ServiceCategory:         input.ServiceCategory,
		ServiceDetail:           input.ServiceDetail,
		ShopId:                  input.ShopId,
		CenterId:                input.CenterId,
		ServiceTemplateId:       input.ServiceTemplateId,
		MaleRoutineIds:          []int64{},
		DefaultMaleRoutineIds:   []int64{},
		FemaleRoutineIds:        []int64{},
		DefaultFemaleRoutineIds: []int64{},
		DefaultAvatarId:         input.DefaultAvatarId,
		CreatedBy:               1,
		Active:                  false,
		UseServiceTemplate:      input.UseServiceTemplate,
		Backgrounds:             backgrounds,
		DefaultBackground:       input.DefaultBackground[strings.LastIndex(input.DefaultBackground, "/")+1:],
	}
	response, err := posClient.CreatePos(ctx, data)
	if err != nil {
		return nil, err
	}
	routineIds := input.FemaleRoutineIds
	routineIds = append(routineIds, input.MaleRoutineIds...)
	routines, err := centerClient.GetListRoutine(ctx, &pbCenter.GetListRoutineRequest{
		Ids: routineIds,
	})
	if err != nil {
		return nil, err
	}
	var createManyRoutineRequest pbPos.CreateManyRoutineRequest
	for _, routine := range routines.ListRoutine {
		var createRoutineRequest pbPos.CreateRoutineRequest
		err = copier.Copy(&createRoutineRequest, routine)
		if err != nil {
			return nil, err
		}
		createRoutineRequest.ServiceTemplateRoutineId = routine.Id
		createManyRoutineRequest.Routines = append(createManyRoutineRequest.Routines, &createRoutineRequest)
	}

	newRoutine, err := posClient.CreateManyRoutine(ctx, &createManyRoutineRequest)
	if err != nil {
		return nil, err
	}
	newRoutineIds := make(map[int64]int64)
	for _, routine := range newRoutine.Routines {
		newRoutineIds[routine.ServiceTemplateRoutineId] = routine.Id
	}
	data.MaleRoutineIds = filterPosRoutine(newRoutineIds, input.MaleRoutineIds)
	data.FemaleRoutineIds = filterPosRoutine(newRoutineIds, input.FemaleRoutineIds)
	data.DefaultMaleRoutineIds = filterPosRoutine(newRoutineIds, input.DefaultMaleRoutineIds)
	data.DefaultFemaleRoutineIds = filterPosRoutine(newRoutineIds, input.DefaultFemaleRoutineIds)
	var updatePosInput pbPos.UpdatePosRequest
	err = copier.Copy(&updatePosInput, data)
	if err != nil {
		return nil, err
	}
	updatePosInput.Id = response.Id
	response, err = posClient.UpdatePos(ctx, &updatePosInput)
	if err != nil {
		return nil, err
	}
	var output CreatePosOutput
	err = copier.Copy(&output, response)
	if err != nil {
		return nil, err
	}
	output.Backgrounds = []string{}
	for _, background := range response.Backgrounds {
		output.Backgrounds = append(output.Backgrounds, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
	}
	output.DefaultBackground = fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, response.DefaultBackground)
	return &output, nil
}

func filterPosRoutine(newRoutineIds map[int64]int64, routines []int64) (Ids []int64) {
	mapRoutine := make(map[int64]bool)
	for _, val := range routines {
		mapRoutine[val] = true
	}
	for oldId, newId := range newRoutineIds {
		if mapRoutine[oldId] {
			Ids = append(Ids, newId)
		}
	}
	return
}
