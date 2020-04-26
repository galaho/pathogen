// Copyright (c) 2020 Joshua Mark Rutherford
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package variables

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// IOResolver implements a variable resolver that resolves using input from an io.Reader and outputs to an io.Writer.
type IOResolver struct {
	reader io.Reader
	writer io.Writer
}

// NewIOResolver returns a new instance of an IOResolver.
func NewIOResolver(reader io.Reader, writer io.Writer) *IOResolver {
	return &IOResolver{
		reader: reader,
		writer: writer,
	}
}

// Resolve resolves variables.
func (r *IOResolver) Resolve(variables []Variable) (map[string]string, error) {

	scanner := bufio.NewScanner(r.reader)
	resolved := make(map[string]string)

	for _, variable := range variables {

		prompt := fmt.Sprintf("Enter %s [%s]: ", variable.Description, variable.Value)

		_, err := r.writer.Write([]byte(prompt))
		if err != nil {
			return nil, errors.Wrapf(err, "error requesting value for variable [%s]", variable.Name)
		}

		scanner.Scan()

		if scanner.Err() != nil {
			return nil, errors.Wrapf(err, "error reading value for variable [%s]", variable.Name)
		}

		if scanner.Text() == "" {
			resolved[variable.Name] = variable.Value
		} else {
			resolved[variable.Name] = scanner.Text()
		}
	}

	return resolved, nil
}
