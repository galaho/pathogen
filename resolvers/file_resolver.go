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

package resolvers

import (
	"io/ioutil"

	"github.com/galaho/pathogen/repositories"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// FileResolver implements a variable resolver that resolves using a yaml file containg key and value pairs.
type FileResolver struct {
	path string
}

// NewFileResolver returns a new instance of a FileResolver.
func NewFileResolver(path string) *FileResolver {
	return &FileResolver{
		path: path,
	}
}

// Resolve resolves variables.
func (r *FileResolver) Resolve(variables []repositories.Variable) (map[string]string, error) {

	var unmarshalled map[string]string

	bytes, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading variable file")
	}

	err = yaml.Unmarshal(bytes, &unmarshalled)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling variable file")
	}

	resolved := make(map[string]string)

	for _, variable := range variables {

		value, exists := unmarshalled[variable.Name]
		if !exists {
			value = variable.Value
		}

		resolved[variable.Name] = value
	}

	return resolved, nil
}
