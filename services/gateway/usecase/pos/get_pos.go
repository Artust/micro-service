package pos

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"fmt"
)

type GetPosInput struct {
	Id int64 `json:"id"`
}

func GetPos(
	input *GetPosInput,
	posClient pb.POSClient,
	cfg *config.Environment,
) (*CreatePosOutput, error) {
	getPosInput := &pb.GetByIdRequest{
		Id: input.Id,
	}
	response, err := posClient.GetPos(context.Background(), getPosInput)
	if err != nil {
		return nil, err
	}
	backgroundName := []string{}
	for _, background := range response.Backgrounds {
		backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
	}
	output := &CreatePosOutput{
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

	return output, nil
}
