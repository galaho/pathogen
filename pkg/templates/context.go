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

package templates

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/pkg/errors"
)

// Context represents a context.
type Context struct {

	// Scripts define the custom scripts for the context.
	Scripts map[string]string

	// Variables define the variables for the context.
	Variables map[string]string
}

// Functions returns the function map for the context.
func (c *Context) Functions() map[string]interface{} {
	functions := map[string]interface{}{
		"environment": environment(),
		"lower":       lower(),
		"now":         now(),
		"split":       split(),
		"upper":       upper(),
		"variable":    variable(c.Variables),
	}
	for name, content := range c.Scripts {
		functions[name] = script(name, content)
	}
	return functions
}

// script returns a function that returns the value of a custom script.
func script(name string, content string) interface{} {
	return func(args ...interface{}) (interface{}, error) {
		machine := goja.New()
		machine.Set("args", args)
		value, err := machine.RunScript(name, content)
		if err != nil {
			return nil, errors.Wrap(err, "error executing custom script")
		}
		return value.Export(), nil
	}
}

// environment returns a function that returns the value of an environment variable.
func environment() func(string) string {
	return func(name string) string {
		return os.Getenv(name)
	}
}

// lower returns a function that returns the result of calling strings.ToLower on a string.
func lower() func(string) string {
	return func(value string) string {
		return strings.ToLower(value)
	}
}

// now returns a function that returns the current time as a formatted string.
func now() func(string) string {
	present := time.Now()
	return func(format string) string {
		return present.Format(format)
	}
}

// split returns a function that returns the slice produced by calling strings.Split on a string.
func split() func(string, string) []string {
	return func(delimiter string, value string) []string {
		return strings.Split(value, delimiter)
	}
}

// variable returns a function that returns a variable from a map.
func variable(variables map[string]string) func(string) (string, error) {
	return func(name string) (string, error) {
		value, exists := variables[name]
		if !exists {
			return "", fmt.Errorf("error fetching variable with name [%s]", name)
		}
		return value, nil
	}
}

// upper returns a function that returns the result of calling strings.ToUpper on a string.
func upper() func(string) string {
	return func(value string) string {
		return strings.ToUpper(value)
	}
}
