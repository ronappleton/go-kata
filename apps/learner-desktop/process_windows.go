//go:build windows

package main

import "os/exec"

func prepareCommand(_ *exec.Cmd) {}

func sendInterrupt(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}
	return cmd.Process.Kill()
}

func forceKill(cmd *exec.Cmd) error {
	if cmd == nil || cmd.Process == nil {
		return nil
	}
	return cmd.Process.Kill()
}

func isNoSuchProcess(_ error) bool {
	return false
}
