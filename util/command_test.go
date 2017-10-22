package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteUnixCommand(t *testing.T) {
	cmd := NewCommand().
		ExecuteUnixCommand("echo").
		Args([]string{"ok, I think so."}).
		ToCmd()

	assert.Equal(t, "!echo 'ok, I think so.'", cmd)
}

func TestCommand_Echom(t *testing.T) {
	cmd := NewCommand().
		Echom().
		Args([]string{"Hello, World!"}).
		ToCmd()

	assert.Equal(t, "echom 'Hello, World!'", cmd)
}

func TestCommand_NewBuffer(t *testing.T) {
	cmd := NewCommand().
		NewBuffer().
		ToCmd()

	assert.Equal(t, "enew", cmd)
}
