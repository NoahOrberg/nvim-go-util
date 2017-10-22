package util

import (
	"fmt"
	"regexp"

	"github.com/NoahOrberg/go-client/nvim"
)

type cmdString string

type Command struct {
	cmd string
}

type Commander interface {
	ToCmd() string
	Exec(*nvim.Nvim) error
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) ToCmd() string {
	return c.cmd
}

func (c *Command) Exec(v *nvim.Nvim) error {
	return v.Command(c.cmd)
}

func (c *Command) Echom() *Command {
	nc := &Command{
		cmd: fmt.Sprintf("%s", "echom"),
	}
	return nc
}

func (c *Command) Echo() *Command {
	nc := &Command{
		cmd: fmt.Sprintf("%s", "echo"),
	}
	return nc
}

func (c *Command) Args(args []string) *Command {
	var s string
	for i, arg := range args {
		regSp := regexp.MustCompile(" ")
		if regSp.MatchString(arg) {
			if i == 0 {
				s = fmt.Sprintf("'%s'", arg)
			} else {
				s = fmt.Sprintf("%s '%s'", s, arg)
			}
		} else {
			if i == 0 {
				s = fmt.Sprintf("%s", arg)
			} else {
				s = fmt.Sprintf("%s %s", s, arg)
			}
		}
	}
	nc := &Command{
		cmd: fmt.Sprintf("%s %s", c.cmd, s),
	}
	return nc
}

func (c *Command) NewBuffer() *Command {
	nc := &Command{
		fmt.Sprintf("%s", "enew"),
	}
	return nc
}

func (c *Command) UnixCommand(ucmd string) *Command {
	nc := &Command{
		fmt.Sprintf("!%s", ucmd),
	}
	return nc
}
