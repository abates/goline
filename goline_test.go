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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"os"
)

var _ = Describe("Goline", func() {
	Describe("Reading a line", func() {
		var lineEditor *LineEditor
		var stdin io.Writer
		var stdout io.Reader

		BeforeEach(func() {
			var stdin_r io.Reader
			var stdout_wr io.Writer
			stdin_r, stdin, _ = os.Pipe()
			stdout, stdout_wr, _ = os.Pipe()

			lineEditor = NewLineEditor()
			lineEditor.SetStdout(stdout_wr)
			lineEditor.SetStdin(stdin_r)
		})

		It("Should display the given prompt", func() {
			prompt := "prompt"
			io.WriteString(stdin, "\n")
			_, err := lineEditor.Readline(prompt)
			Expect(err).To(Succeed())
			expected := make([]byte, len(prompt))
			_, _ = stdout.Read(expected)
			Expect(string(expected)).To(Equal(prompt))
		})

		It("Should return the line read from stdin", func() {
			line := "line"
			io.WriteString(stdin, line+"\n")
			line, err := lineEditor.Readline("")
			Expect(err).To(Succeed())

			expected := make([]byte, len(line))
			_, _ = stdout.Read(expected)
			Expect(string(expected)).To(Equal(line))
		})
	})
})
