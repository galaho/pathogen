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

package generation

import (
	"github.com/galaho/pathogen/pkg/repositories"
	"github.com/galaho/pathogen/pkg/resolvers"
	"github.com/galaho/pathogen/pkg/templates"
	"github.com/pkg/errors"
)

// Generate generates filesystem entries from a template.
func Generate(source string, destination string, config string, resolver resolvers.Resolver) error {

	repository, err := repositories.Open(source, config)
	if err != nil {
		return errors.Wrapf(err, "error fetching source [%s]", source)
	}

	defer repository.Close()

	variables, err := resolver.Resolve(repository.Variables)
	if err != nil {
		return errors.Wrap(err, "error resolving variables")
	}

	context := &templates.Context{Scripts: repository.Scripts, Variables: variables}

	err = repository.Render(destination, context)

	if err != nil {
		return errors.Wrap(err, "error rendering repository")
	}

	return nil
}

