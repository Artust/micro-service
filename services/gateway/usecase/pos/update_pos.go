package pos

import (
	"avatar/services/gateway/config"
	pbCenter "avatar/services/gateway/protos/center"
	pb "avatar/services/gateway/protos/pos"
	pbPos "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
	"strings"
)

type UpdatePosInput struct {
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
	Active                  bool     `json:"active"`
	Backgrounds             []string `form:"backgrounds"`
	DefaultBackground       string   `form:"defaultBackground"`
}

func UpdatePos(
	input *UpdatePosInput,
	posClient pb.POSClient,
	centerClient pbCenter.CenterClient,
	cfg *config.Environment,
) (*CreatePosOutput, error) {
	ctx := context.Background()
	routineIds := input.FemaleRoutineIds
	routineIds = append(routineIds, input.MaleRoutineIds...)
	routineIds = append(routineIds, input.DefaultMaleRoutineIds...)
	routineIds = append(routineIds, input.FemaleRoutineIds...)
	routineIds = append(routineIds, input.DefaultFemaleRoutineIds...)
	listRoutines, err := centerClient.GetListRoutineByCategory(ctx, &pbCenter.GetListRoutineByCategoryRequest{
		Ids: routineIds,
	})
	if err != nil {
		return nil, err
	}
	newRoutineIds := make(map[int64]int64)
	for _, category := range listRoutines.GetListRoutineByCategoryResponse {
		for _, routine := range category.Routine {
			newRoutine, err := posClient.CreateRoutine(ctx, &pbPos.CreateRoutineRequest{
				Name:                     routine.Name,
				Detail:                   routine.Detail,
				AnimationFile:            routine.AnimationFile,
				SoundFile:                routine.SoundFile,
				ImageFile:                routine.ImageFile,
				StartDate:                routine.StartDate,
				EndDate:                  routine.EndDate,
				CategoryId:               routine.CategoryId,
				PosId:                    input.Id,
				ServiceTemplateRoutineId: routine.Id,
			})
			if err != nil {
				return nil, err
			}
			newRoutineIds[routine.Id] = newRoutine.Id
		}
	}
	var backgrounds []string
	for _, val := range input.Backgrounds {
		splBackground := strings.LastIndex(val, "/")
		backgrounds = append(backgrounds, val[splBackground+1:])
	}
	updatePosRequest := &pb.UpdatePosRequest{
		Id:                      input.Id,
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
		Active:                  input.Active,
		UseServiceTemplate:      input.UseServiceTemplate,
		Backgrounds:             backgrounds,
		DefaultBackground:       input.DefaultBackground[strings.LastIndex(input.DefaultBackground, "/")+1:],
	}

	updatePosRequest.MaleRoutineIds = append(updatePosRequest.MaleRoutineIds, dropServiceTemplateRoutine(newRoutineIds, input.MaleRoutineIds)...)
	updatePosRequest.FemaleRoutineIds = append(updatePosRequest.FemaleRoutineIds, dropServiceTemplateRoutine(newRoutineIds, input.FemaleRoutineIds)...)
	updatePosRequest.DefaultMaleRoutineIds = append(updatePosRequest.DefaultMaleRoutineIds, dropServiceTemplateRoutine(newRoutineIds, input.DefaultMaleRoutineIds)...)
	updatePosRequest.DefaultFemaleRoutineIds = append(updatePosRequest.DefaultFemaleRoutineIds, dropServiceTemplateRoutine(newRoutineIds, input.DefaultFemaleRoutineIds)...)

	updatePosRequest.MaleRoutineIds = append(updatePosRequest.MaleRoutineIds, filterPosRoutine(newRoutineIds, input.MaleRoutineIds)...)
	updatePosRequest.FemaleRoutineIds = append(updatePosRequest.FemaleRoutineIds, filterPosRoutine(newRoutineIds, input.FemaleRoutineIds)...)
	updatePosRequest.DefaultMaleRoutineIds = append(updatePosRequest.DefaultMaleRoutineIds, filterPosRoutine(newRoutineIds, input.DefaultMaleRoutineIds)...)
	updatePosRequest.DefaultFemaleRoutineIds = append(updatePosRequest.DefaultFemaleRoutineIds, filterPosRoutine(newRoutineIds, input.DefaultFemaleRoutineIds)...)

	response, err := posClient.UpdatePos(ctx, updatePosRequest)
	if err != nil {
		return nil, err
	}
	backgroundName := []string{}
	for _, backgrounds := range response.Backgrounds {
		backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, backgrounds))
	}
	result := &CreatePosOutput{
		Id:                      response.Id,
		Name:                    response.Name,
		ServiceName:             response.ServiceName,
		ServiceCategory:         response.ServiceCategory,
		ServiceDetail:           response.ServiceDetail,
		ShopId:                  response.ShopId,
		CenterId:                response.CenterId,
		ServiceTemplateId:       response.ServiceTemplateId,
		MaleRoutineIds:          response.MaleRoutineIds,
		DefaultMaleRoutineIds:   response.DefaultMaleRoutineIds,
		FemaleRoutineIds:        response.FemaleRoutineIds,
		DefaultFemaleRoutineIds: response.DefaultFemaleRoutineIds,
		DefaultAvatarId:         response.DefaultAvatarId,
		UseServiceTemplate:      response.UseServiceTemplate,
		Backgrounds:             backgroundName,
		DefaultBackground:       fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, response.DefaultBackground),
		CreatedBy:               response.CreatedBy,
		CreatedAt:               response.CreatedAt,
		UpdatedAt:               response.UpdatedAt,
	}
	return result, nil
}

func dropServiceTemplateRoutine(newRoutineIds map[int64]int64, routines []int64) (Ids []int64) {
	mapRoutine := make(map[int64]bool)
	for oldId, _ := range newRoutineIds {
		mapRoutine[oldId] = true
	}
	for _, val := range routines {
		if !mapRoutine[val] {
			Ids = append(Ids, val)
		}
	}
	return
}
