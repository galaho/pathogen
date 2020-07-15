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

package commands

import (
	"os"

	"github.com/galaho/pathogen/pkg/repositories"
	"github.com/galaho/pathogen/pkg/resolvers"
	"github.com/galaho/pathogen/pkg/templates"
	"github.com/pkg/errors"
)

// Generate generates filesystem entries from a template.
func Generate(input string, repo string, configFile string, dest string) error {

	repository, err := repositories.Open(repo, configFile)
	if err != nil {
		return errors.Wrapf(err, "error fetching repository [%s]", repo)
	}

	defer repository.Close()

	var resolver resolvers.Resolver

	resolver = resolvers.NewIOResolver(os.Stdin, os.Stdout)

	if input != "" {
		resolver = resolvers.NewFileResolver(input)
	}

	variables, err := resolver.Resolve(repository.Variables)
	if err != nil {
		return errors.Wrap(err, "error resolving variables")
	}

	context := &templates.Context{Scripts: repository.Scripts, Variables: variables}

	err = repository.Render(dest, context)

	if err != nil {
		return errors.Wrap(err, "error walking repository")
	}

	return nil
}
