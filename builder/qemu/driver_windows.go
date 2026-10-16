// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build windows

package qemu

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func setNewProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windows.CREATE_NEW_PROCESS_GROUP}
}

func terminateProcess(p *os.Process) error {
	return windows.GenerateConsoleCtrlEvent(windows.CTRL_C_EVENT, uint32(p.Pid))
}
