//go:build linux || darwin
// +build linux darwin

package main

import "github.com/piigyy/gostrap/cmd"

func main() {
	cmd.Execute()
}
