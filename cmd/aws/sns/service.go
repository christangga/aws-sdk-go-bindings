package sns

import "github.com/andream16/aws-sdk-go-bindings/pkg/aws/sns"

// Publish publishes a given input to in a given targetArn
func (svc *SNS) Publish(input interface{}, targetArn string) error {

	out, outErr := sns.NewPublishInput(
		input,
		targetArn,
	)
	if outErr != nil {
		return outErr
	}

	pubErr := svc.SnsPublish(out)
	if pubErr != nil {
		return pubErr
	}

	return nil

}
