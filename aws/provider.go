package aws

import "github.com/aws/aws-sdk-go-v2/aws"

type provider struct {
	dynamodbEndpoint string
	dynamodbRegion   string

	config aws.Config
}

func (p *provider) GetConfig() aws.Config {
	return p.config
}
