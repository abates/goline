/**
 * Copyright 2015 Andrew Bates
 *
 * Licensed under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with the
 * License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package goline

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Line struct {
	screenWidth int
	prompt      string
	head        string
	tail        string
}

type LineEditor struct {
	stdin  *bufio.Reader
	stdout io.Writer
	stderr io.Writer
}

func NewLineEditor() *LineEditor {
	return &LineEditor{
		stdin:  bufio.NewReader(os.Stdin),
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

func (this *LineEditor) SetStdin(stdin io.Reader) error {
	if stdin == nil {
		return ErrorNilReader
	}
	this.stdin = bufio.NewReader(stdin)
	return nil
}

func (this *LineEditor) SetStdout(stdout io.Writer) error {
	if stdout == nil {
		return ErrorNilWriter
	}
	this.stdout = stdout
	return nil
}

func (this *LineEditor) SetStderr(stderr io.Writer) error {
	if stderr == nil {
		return ErrorNilWriter
	}
	this.stderr = stderr
	return nil
}

func (this *LineEditor) Readline(prompt string) (string, error) {
	line := ""
	fmt.Fprintf(this.stdout, "%v", prompt)
	for l, isPrefix, err := this.stdin.ReadLine(); isPrefix; {
		if err != nil {
			return "", err
		}
		line += string(l)
	}
	return line, nil
}
