package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSessionInput(t *testing.T) {

	region := "some_region"

	out, err := NewSessionInput(region)

	assert.NoError(t, err)
	assert.Equal(t, region, out.region)

	_, shouldBeNoRegionProvidedErr := NewSessionInput("")

	assert.Equal(t, ErrNoRegionProvided, shouldBeNoRegionProvidedErr.Error())

}
