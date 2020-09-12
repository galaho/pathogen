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
	"io"

	"github.com/galaho/pathogen/pkg/repositories"
	"github.com/pkg/errors"
)

// Resolver implements a variable resolver that resolves using an external mechanism. The resolver works by encoding
// the variables requiring resolution and writing those to an io.Writer then reading and decoding the resolved
// variables from an io.Reader.
type Resolver struct {
	decode Decode
	encode Encode
	reader io.Reader
	writer io.Writer
}

// NewResolver returns a new instance of an Resolver.
func NewResolver(decode Decode, encode Encode, reader io.Reader, writer io.Writer) *Resolver {
	return &Resolver{
		decode: decode,
		encode: encode,
		reader: reader,
		writer: writer,
	}
}

// Resolve resolves variables.
func (r *Resolver) Resolve(variables []repositories.Variable) (map[string]string, error) {

	err := r.encode(variables, r.writer)
	if err != nil {
		return nil, errors.Wrap(err, "error encoding variables")
	}

	values, err := r.decode(r.reader)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding variables")
	}

	return values, nil
}
