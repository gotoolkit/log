package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	err := Setup(WithOutputFile("test.log", 10, 10, 7, true))
	Println("test")
	assert.NoError(t, err)
}
