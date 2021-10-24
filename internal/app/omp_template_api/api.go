package omp_template_api

import (
	desc "github.com/ozonmp/omp-grpc-template/pkg/omp-template-api"
)

type Implementation struct {
	desc.UnimplementedOmpTemplateApiServiceServer
}

func NewTemplateAPI() desc.OmpTemplateApiServiceServer {
	return &Implementation{}
}
