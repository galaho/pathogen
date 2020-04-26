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
	"time"
)

// Context represents a context.
type Context struct {

	// Variables define the variables for the context.
	Variables map[string]string
}

// Functions returns the function map for the context.
func (c *Context) Functions() map[string]interface{} {
	return map[string]interface{}{
		"now":      now(),
		"variable": variable(c.Variables),
	}
}

// now returns a function that returns the current time as a formatted string.
func now() func(string) string {
	present := time.Now()
	return func(format string) string {
		return present.Format(format)
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
