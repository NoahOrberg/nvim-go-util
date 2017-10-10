package util

import "github.com/neovim/go-client/nvim"

func (v *nvim.Nvim) Echom(s string) error {
	return v.Command("echom '" + s + "'")
}

func (v *nvim.Nvim) Echo(s string) error {
	return v.Command("echo '" + s + "'")
}
