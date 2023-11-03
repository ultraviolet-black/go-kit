package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ProviderOption func(*provider)

func WithDynamoDBEndpoint(endpoint string, region string) ProviderOption {
	return func(p *provider) {
		p.dynamodbEndpoint = endpoint
		p.dynamodbRegion = region
	}
}

type Provider interface {
	GetConfig() aws.Config
}

func NewProvider(opts ...ProviderOption) Provider {

	p := &provider{}

	for _, opt := range opts {
		opt(p)
	}

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if len(p.dynamodbEndpoint) > 0 && service == dynamodb.ServiceID {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           p.dynamodbEndpoint,
				SigningRegion: p.dynamodbRegion,
			}, nil
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfgOpts := []func(*config.LoadOptions) error{
		config.WithEndpointResolver(customResolver),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), cfgOpts...)
	if err != nil {
		panic(err)
	}

	p.config = cfg

	return p

}
