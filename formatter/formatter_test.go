/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package formatter

import (
	"fmt"
	"io"
	"testing"

	"go.uber.org/zap/buffer"
	"gotest.tools/assert"
)

type testStruct struct {
	Name   string
	Status string
}

// Print prints formatted lists in different formats
func TestPrint(t *testing.T) {
	testList := []testStruct{
		{
			Name:   "myName",
			Status: "myStatus",
		},
	}

	b := &buffer.Buffer{}
	assert.NilError(t, Print(testList, PRETTY, b, func(w io.Writer) {
		for _, t := range testList {
			_, _ = fmt.Fprintf(w, "%s\t%s\n", t.Name, t.Status)
		}
	}, "NAME", "STATUS"))
	assert.Equal(t, b.String(), "NAME                STATUS\nmyName              myStatus\n")

	b.Reset()
	assert.NilError(t, Print(testList, JSON, b, func(w io.Writer) {
		for _, t := range testList {
			_, _ = fmt.Fprintf(w, "%s\t%s\n", t.Name, t.Status)
		}
	}, "NAME", "STATUS"))
	assert.Equal(t, b.String(), `[
    {
        "Name": "myName",
        "Status": "myStatus"
    }
]`)
}
