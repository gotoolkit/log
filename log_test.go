package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New(WithOutputFile("test.log", 10, 10, 7, true))
	assert.NotNil(t, l)
	l.Println("test")
}

func TestSTD(t *testing.T) {
	SetOutputFile("test1.log", 10, 10, 7, true)
	Println("test")
}
