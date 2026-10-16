// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build unix

package qemu

import (
	"os"
	"os/exec"
)

func setNewProcessGroup(cmd *exec.Cmd) {}

func terminateProcess(p *os.Process) error {
	return p.Signal(os.Interrupt)
}
