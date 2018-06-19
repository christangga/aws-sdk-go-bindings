package sns

import (
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sns"
)

// SNS embeds *sns.Session and is used to call sns methods on high level
type SNS struct {
	*sns.SNS
}

// New returns a *SNS given a *aws.Session
func New(svc *aws.Session) (*SNS, error) {

	snsSvc, err := sns.New(svc)
	if err != nil {
		return nil, err
	}

	newSnsSvc := new(SNS)
	newSnsSvc.SNS = snsSvc

	return newSnsSvc, nil

}
