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

package generate

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/galaho/pathogen/repositories"
	"github.com/galaho/pathogen/resolvers"
	"github.com/galaho/pathogen/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Generate returns a command that generates filesystem entries from a template.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "generate REPOSITORY DESTINATION",
		Short: "Generate filesystem entries from a template",
		RunE: func(command *cobra.Command, args []string) error {

			repository, err := repositories.Fetch(args[0])
			if err != nil {
				return errors.Wrapf(err, "error fetching repository [%s]", args[0])
			}

			defer repository.Close()

			input, err := command.Flags().GetString("input")
			if err != nil {
				return errors.Wrap(err, "error determining input file")
			}

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

			err = repository.Walk(func(file *repositories.File) error {

				path, err := templates.Render(file.Path, context)
				if err != nil {
					return errors.Wrapf(err, "error rendering template path [%s]", file.Path)
				}

				content, err := templates.Render(string(file.Bytes), context)
				if err != nil {
					return errors.Wrapf(err, "error rendering template file [%s]", file.Path)
				}

				directory := filepath.Dir(filepath.Join(args[1], path))

				err = os.MkdirAll(directory, 0777)
				if err != nil {
					return errors.Wrapf(err, "error creating directory [%s]", directory)
				}

				err = ioutil.WriteFile(filepath.Join(args[1], path), []byte(content), file.Info.Mode())
				if err != nil {
					return errors.Wrapf(err, "error creating file [%s]", path)
				}

				return nil
			})

			if err != nil {
				return errors.Wrap(err, "error walking repository")
			}

			return nil
		},
	}

	command.Flags().StringP("input", "i", "", "file for non-interactive variable resolution")

	return command
}
