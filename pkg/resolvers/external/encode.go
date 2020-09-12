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

package external

import (
	"encoding/json"
	"io"

	"github.com/galaho/pathogen/pkg/repositories"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Encode provides a function interface for encoding variable slices to an io.Writer.
type Encode func([]repositories.Variable, io.Writer) error

// EncodeJSON implements an Encode function for encoding variables to JSON.
func EncodeJSON(variables []repositories.Variable, writer io.Writer) error {

	err := json.NewEncoder(writer).Encode(variables)
	if err != nil {
		return errors.Wrap(err, "error encoding variables")
	}

	return nil
}

// EncodeYAML implements an Encode function for encoding variables to YAML.
func EncodeYAML(variables []repositories.Variable, writer io.Writer) error {

	err := yaml.NewEncoder(writer).Encode(variables)
	if err != nil {
		return errors.Wrap(err, "error encoding variables")
	}

	return nil
}
