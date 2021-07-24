// Copyright (C) 2021 Akira Tanimura (@autopp)
//
// Licensed under the Apache License, Version 2.0 (the “License”);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an “AS IS” BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s program [args...]\n", os.Args[0])
		os.Exit(1)
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Run()

	ps := cmd.ProcessState

	if ps.Exited() {
		fmt.Printf("exited with %d\n", ps.ExitCode())
		os.Exit(0)
	}

	ws, ok := ps.Sys().(syscall.WaitStatus)
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown (*ProcessState).Sys(): %T %#v", ps.Sys(), ps.Sys())
	}

	if ws.Signaled() {
		s := ws.Signal()
		fmt.Printf("signaled with %d (%s)\n", s, s.String())
	}
}
