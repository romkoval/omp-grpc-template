package omp_template_api

import (
	desc "github.com/ozonmp/omp-template-api/pkg/omp-template-api"
)

type Implementation struct {
	desc.UnimplementedOmpTemplateApiServiceServer
}

func NewTemplateAPI() desc.OmpTemplateApiServiceServer {
	return &Implementation{}
}
