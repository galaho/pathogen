// Copyright 2020 Joshua Mark Rutherford
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
