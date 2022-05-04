package generate

import (
	"testing"

	command "github.com/pPrecel/cloud-agent/cmd"
	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	c := NewCmd(&command.Options{})

	assert.Equal(t, 2, len(c.Commands()))
}
