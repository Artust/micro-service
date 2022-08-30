package handler

import (
	pb "avatar/services/center/protos"
	"avatar/services/center/usecase/service_template"
	"context"
)

func (c *CenterServer) CreateServiceTemplate(ctx context.Context, input *pb.CreateServiceTemplateRequest) (*pb.CreateServiceTemplateResponse, error) {
	return service_template.Create(ctx, c.neo4jDriver, c.serviceTemplateRepository, input)
}

func (c *CenterServer) GetServiceTemplate(ctx context.Context, input *pb.GetByIdRequest) (*pb.CreateServiceTemplateResponse, error) {
	return service_template.GetById(ctx, c.neo4jDriver, c.serviceTemplateRepository, input)
}

func (c *CenterServer) GetListServiceTemplate(ctx context.Context, input *pb.GetListServiceTemplateRequest) (*pb.GetListServiceTemplateResponse, error) {
	return service_template.GetList(ctx, c.neo4jDriver, c.serviceTemplateRepository, input)
}

func (c *CenterServer) UpdateServiceTemplate(ctx context.Context, input *pb.UpdateServiceTemplateRequest) (*pb.CreateServiceTemplateResponse, error) {
	return service_template.Update(ctx, c.neo4jDriver, c.serviceTemplateRepository, input)
}

func (c *CenterServer) DeleteServiceTemplate(ctx context.Context, input *pb.DeleteByIdRequest) (*pb.DeleteResponse, error) {
	return service_template.Delete(ctx, c.neo4jDriver, c.serviceTemplateRepository, input)
}
