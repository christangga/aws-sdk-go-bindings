package sqs

import (
	"github.com/andream16/aws-sdk-go-bindings/cmd/aws"
	"github.com/andream16/aws-sdk-go-bindings/pkg/aws/sqs"
)

// SQS embeds *sqs.SQS and is used to call sqs methods on high level
type SQS struct {
	*sqs.SQS
}

// New returns a *SQS given a *aws.Session
func New(svc *aws.Session) (*SQS, error) {

	snsSvc, err := sqs.New(svc.Session)
	if err != nil {
		return nil, err
	}

	newSqsSvc := new(SQS)
	newSqsSvc.SQS = snsSvc

	return newSqsSvc, nil

}
