package aws

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	uuid "github.com/satori/go.uuid"

	"github.com/elastic/beats/libbeat/logp"
)

type opCreateCloudFormation struct {
	log         *logp.Logger
	svc         *cloudformation.CloudFormation
	templateURL string
	stackName   string
}

func newOpCreateCloudFormation(
	log *logp.Logger,
	cfg aws.Config,
	templateURL, stackName string,
) *opCreateCloudFormation {
	return &opCreateCloudFormation{
		log:         log,
		svc:         cloudformation.New(cfg),
		templateURL: templateURL,
		stackName:   stackName,
	}
}

func (o *opCreateCloudFormation) Execute(ctx *executorContext) error {
	o.log.Debug("Creating CloudFormation create request")
	uuid := uuid.NewV4()
	input := &cloudformation.CreateStackInput{
		ClientRequestToken: aws.String(uuid.String()),
		StackName:          aws.String(o.stackName),
		TemplateURL:        aws.String(o.templateURL),
		Capabilities: []cloudformation.Capability{
			cloudformation.CapabilityCapabilityNamedIam,
		},
	}

	req := o.svc.CreateStackRequest(input)
	resp, err := req.Send()
	if err != nil {
		o.log.Debugf("could not create the cloud formation stack request, resp: %v", resp)
		return err
	}
	return nil
}
