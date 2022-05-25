package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructures(t *testing.T) {
	err := ErrorResponse{}

	resp := Response{}

	assert.Empty(t, err)
	assert.Empty(t, resp)
}
