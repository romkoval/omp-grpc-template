package omp_template_api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/ozonmp/omp-grpc-template/pkg/omp-template-api"
)

func (i *Implementation) DescribeTemplateV1(
	ctx context.Context,
	req *desc.DescribeTemplateV1Request,
) (*desc.DescribeTemplateV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.DescribeTemplateV1Response{
		Value: &desc.Template{
			// Id:  template.ID,
			// Foo: template.Foo,
		},
	}, nil
}
