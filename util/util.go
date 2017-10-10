package util

import "github.com/neovim/go-client/nvim"

func Echom(v *nvim.Nvim, s string) error {
	return v.Command("echom '" + s + "'")
}

func Echo(v *nvim.Nvim, s string) error {
	return v.Command("echo '" + s + "'")
}
