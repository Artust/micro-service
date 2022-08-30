package service_template

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"context"
	"fmt"
)

type GetListServiceTemplateInput struct {
	ServiceTemplateCategoryId int64 `form:"serviceTemplateCategoryId"`
	CorporationId             int64 `form:"corporationId"`
	Page                      int64 `form:"page"`
	PerPage                   int64 `form:"perPage"`
}
type GetListServiceTemplateOutput struct {
	ServiceTemplates []*CreateServiceTemplateOutput `json:"results"`
}

func GetListServiceTemplate(
	input *GetListServiceTemplateInput,
	centerClient pb.CenterClient,
	cfg *config.Environment,
) (*GetListServiceTemplateOutput, error) {
	ctx := context.Background()
	getServiceTemplateByTalkSessionInput := &pb.GetListServiceTemplateRequest{
		CorporationId:             input.CorporationId,
		Page:                      input.Page,
		PerPage:                   input.PerPage,
		ServiceTemplateCategoryId: input.ServiceTemplateCategoryId,
	}
	ServiceTemplates, err := centerClient.GetListServiceTemplate(ctx, getServiceTemplateByTalkSessionInput)
	if err != nil {
		return nil, err
	}
	ServiceTemplateOutput := &GetListServiceTemplateOutput{}
	ServiceTemplateOutput.ServiceTemplates = make([]*CreateServiceTemplateOutput, 0)
	for _, serviceTemplate := range ServiceTemplates.ListServiceTemplate {
		backgroundName := []string{}
		for _, background := range serviceTemplate.Backgrounds {
			backgroundName = append(backgroundName, fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, background))
		}
		ServiceTemplate := &CreateServiceTemplateOutput{
			Id:                      serviceTemplate.Id,
			Name:                    serviceTemplate.Name,
			Detail:                  serviceTemplate.Detail,
			CorporationId:           serviceTemplate.CorporationId,
			DefaultMaleRoutineIds:   serviceTemplate.DefaultMaleRoutineIds,
			MaleRoutineIds:          serviceTemplate.MaleRoutineIds,
			DefaultFemaleRoutineIds: serviceTemplate.DefaultFemaleRoutineIds,
			FemaleRoutineIds:        serviceTemplate.FemaleRoutineIds,
			DefaultMaleAvatarId:     serviceTemplate.DefaultMaleAvatarId,
			DefaultFemaleAvatarId:   serviceTemplate.DefaultFemaleAvatarId,
			AvatarIds:               serviceTemplate.AvatarIds,
			ServiceTemplateCategory: serviceTemplate.ServiceTemplateCategory,
			Backgrounds:             backgroundName,
			CreatedBy:               serviceTemplate.CreatedBy,
			DefaultBackground:       fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.BackgroundBucketName, serviceTemplate.DefaultBackground),
			AvatarUri:               "http://13.113.115.189:4566/routine/defaultRoutineImage.png",
			CreatedAt:               serviceTemplate.CreatedAt,
			UpdatedAt:               serviceTemplate.UpdatedAt,
		}
		ServiceTemplateOutput.ServiceTemplates = append(ServiceTemplateOutput.ServiceTemplates, ServiceTemplate)
	}
	return ServiceTemplateOutput, nil
}
