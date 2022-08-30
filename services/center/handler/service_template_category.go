package handler

import (
	pb "avatar/services/center/protos"
	"avatar/services/center/usecase/service_template_category"
	"context"
)

func (c *CenterServer) CreateServiceTemplateCategory(
	ctx context.Context,
	input *pb.CreateServiceTemplateCategoryRequest,
) (*pb.CreateServiceTemplateCategoryResponse, error) {
	return service_template_category.Create(ctx, c.neo4jDriver, c.serviceTemplateCategoryRepository, input)
}

func (c *CenterServer) GetServiceTemplateCategory(
	ctx context.Context,
	input *pb.GetByIdRequest,
) (*pb.CreateServiceTemplateCategoryResponse, error) {
	return service_template_category.GetById(ctx, c.neo4jDriver, c.serviceTemplateCategoryRepository, input)
}

func (c *CenterServer) GetListServiceTemplateCategory(
	ctx context.Context,
	input *pb.GetListServiceTemplateCategoryRequest,
) (*pb.GetListServiceTemplateCategoryResponse, error) {
	return service_template_category.GetList(ctx, c.neo4jDriver, c.serviceTemplateCategoryRepository, input)
}

func (c *CenterServer) UpdateServiceTemplateCategory(
	ctx context.Context,
	input *pb.UpdateServiceTemplateCategoryRequest,
) (*pb.CreateServiceTemplateCategoryResponse, error) {
	return service_template_category.Update(ctx, c.neo4jDriver, c.serviceTemplateCategoryRepository, input)
}

func (c *CenterServer) DeleteServiceTemplateCategory(
	ctx context.Context,
	input *pb.DeleteByIdRequest,
) (*pb.DeleteResponse, error) {
	return service_template_category.Delete(ctx, c.neo4jDriver, c.serviceTemplateCategoryRepository, input)
}
