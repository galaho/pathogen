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

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Decode provides a function interface for decoding variable maps from an io.Reader.
type Decode func(io.Reader) (map[string]string, error)

// DecodeJSON implements a Decode function for decoding values from JSON.
func DecodeJSON(reader io.Reader) (map[string]string, error) {

	var values map[string]string

	err := json.NewDecoder(reader).Decode(&values)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding variables")
	}

	return values, nil
}

// DecodeYAML implements a Decode function for decoding values from YAML.
func DecodeYAML(reader io.ReadCloser) (map[string]string, error) {

	var values map[string]string

	err := yaml.NewDecoder(reader).Decode(&values)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding variables")
	}

	return values, nil
}
