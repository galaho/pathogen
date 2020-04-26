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

package templates

import (
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// Render renders a template with a context and returns the rendered value.
func Render(source string, context *Context) (string, error) {

	writer := &strings.Builder{}

	parsed, err := template.New("").Funcs(context.Functions()).Parse(source)
	if err != nil {
		return "", errors.Wrap(err, "error parsing template")
	}

	err = parsed.Execute(writer, nil)
	if err != nil {
		return "", errors.Wrap(err, "error executing template")
	}

	return writer.String(), nil
}
