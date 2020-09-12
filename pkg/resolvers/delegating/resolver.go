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

package delegating

import (
	"github.com/galaho/pathogen/pkg/repositories"
	"github.com/galaho/pathogen/pkg/resolvers"
	"github.com/pkg/errors"
)

// Resolver implements a variable resolver that delegates variable resolution to one or more resolvers. The variables
// returned by the first non-erroring resolver are returned.
type Resolver struct {
	resolvers []resolvers.Resolver
}

// NewResolver returns a new instance of an Resovler.
func NewResolver(resolver ...resolvers.Resolver) *Resolver {
	return &Resolver{
		resolvers: resolver[:],
	}
}

// Resolve resolves variables.
func (r *Resolver) Resolve(variables []repositories.Variable) (map[string]string, error) {
	for index := range r.resolvers {
		variables, err := r.resolvers[index].Resolve(variables)
		if err == nil {
			return variables, nil
		}
	}
	return nil, errors.New("unable to resolve variables with all resolvers")
}
