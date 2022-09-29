package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	result := Response(nil, 200, "success", nil)

	// result.Error = errors.New("err")
	assert.Nil(t, result.Error, result.Error)
}
