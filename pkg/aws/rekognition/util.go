package rekognition

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// CompareFacesInput embeds *rekognition.CompareFacesInput
type CompareFacesInput struct {
	*rekognition.CompareFacesInput
}

// CompareFacesOutput embeds *rekognition.CompareFacesOutput
type CompareFacesOutput struct {
	*rekognition.CompareFacesOutput
}

// DetectFacesInput embeds *rekognition.DetectFacesInput
type DetectFacesInput struct {
	*rekognition.DetectFacesInput
}

// DetectFacesOutput embeds *rekognition.DetectFacesOutput
type DetectFacesOutput struct {
	*rekognition.DetectFacesOutput
}

// DetectTextInput embeds *rekognition.DetectTextInput
type DetectTextInput struct {
	*rekognition.DetectTextInput
}

// DetectTextOutput embeds *rekognition.DetectTextOutput
type DetectTextOutput struct {
	*rekognition.DetectTextOutput
}

// NewCompareFacesInput builds a CompareFacesInput starting from the two images, their bucket and a similarity threshold
func NewCompareFacesInput(source, target []byte, similarity float64) *CompareFacesInput {

	out := new(CompareFacesInput)

	out.CompareFacesInput = &rekognition.CompareFacesInput{
		SimilarityThreshold: aws.Float64(similarity),
		SourceImage:         newInputImage(source),
		TargetImage:         newInputImage(target),
	}

	return out

}

// NewDetectFacesInput builds a DetectFacesInput starting from the image
func NewDetectFacesInput(source []byte) *DetectFacesInput {

	out := new(DetectFacesInput)

	out.DetectFacesInput = &rekognition.DetectFacesInput{
		Image: newInputImage(source),
	}

	return out

}

// NewDetectTextInput builds a DetectTextInput starting from the image
func NewDetectTextInput(source []byte) *DetectTextInput {

	out := new(DetectTextInput)

	out.DetectTextInput = &rekognition.DetectTextInput{
		Image: newInputImage(source),
	}

	return out

}

// newInputImage returns a *rekognition.Image given an S3 image []byte encoded
func newInputImage(image []byte) *rekognition.Image {

	out := new(rekognition.Image)

	out.Bytes = image

	return out

}