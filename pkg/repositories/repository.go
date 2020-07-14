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

package repositories

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/galaho/pathogen/pkg/templates"
	"github.com/hashicorp/go-getter"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Repository represents a template repository.
type Repository struct {
	directory string
	ignore    []*regexp.Regexp

	// Scripts defines the scripts for the repository.
	Scripts map[string]string

	// Variables defines the variables for the repository.
	Variables []Variable
}

// Close frees up all resources associated within the repository.
func (r *Repository) Close() error {
	return os.RemoveAll(r.directory)
}

// Open returns a repository ready for processing.
func Open(repositoryURL string, configurationFile string) (*Repository, error) {

	directory := filepath.Join(os.TempDir(), random())

	err := getter.Get(directory, repositoryURL)
	if err != nil {
		return nil, errors.Wrapf(err, "error fetching repository [%s]", repositoryURL)
	}

	configuration, err := load(filepath.Join(directory, configurationFile))
	if err != nil {
		return nil, errors.Wrapf(err, "error loading configuration file [%s]", configurationFile)
	}

	ignore, err := compile(configuration.Ignore)
	if err != nil {
		return nil, errors.Wrapf(err, "error compiling ignore regular expressions")
	}

	repository := &Repository{
		directory: directory,
		ignore: ignore,
		Scripts: configuration.Scripts,
		Variables: configuration.Variables,
	}

	return repository, nil
}

// Render renders the templates from this repository to disk.
func (r *Repository) Render(destination string, context *templates.Context) error {

	return r.walk(func(file *File) error {

		path, err := templates.Render(file.Path, context)
		if err != nil {
			return errors.Wrapf(err, "error rendering template path [%s]", file.Path)
		}

		content, err := templates.Render(string(file.Bytes), context)
		if err != nil {
			return errors.Wrapf(err, "error rendering template file [%s]", file.Path)
		}

		directory := filepath.Dir(filepath.Join(destination, path))

		err = os.MkdirAll(directory, 0777)
		if err != nil {
			return errors.Wrapf(err, "error creating directory [%s]", directory)
		}

		err = ioutil.WriteFile(filepath.Join(destination, path), []byte(content), file.Info.Mode())
		if err != nil {
			return errors.Wrapf(err, "error creating file [%s]", path)
		}

		return nil
	})
}

// walk invokes a callback for every file within the repository.
func (r *Repository) walk(callback func(*File) error) error {

	evaluated, err := filepath.EvalSymlinks(r.directory)
	if err != nil {
		return errors.Wrapf(err, "error evaluating symlinks [%s]", r.directory)
	}

	err = filepath.Walk(evaluated, func(absolute string, info os.FileInfo, err error) error {

		relative, err := filepath.Rel(evaluated, absolute)
		if err != nil {
			return errors.Wrapf(err, "error determining relative path from [%s] to [%s]", evaluated, absolute)
		}

		if info.Mode().IsRegular() && !matches(relative, r.ignore) {

			bytes, err := ioutil.ReadFile(absolute)
			if err != nil {
				return fmt.Errorf("error reading file [%s]", absolute)
			}

			err = callback(&File{Bytes: bytes, Info: info, Path: relative})
			if err != nil {
				return errors.Wrapf(err, "error processing file [%s]", absolute)
			}
		}

		return nil
	})

	return err
}

// compile returns compiled regular expressions for patterns.
func compile(patterns []string) ([]*regexp.Regexp, error) {

	regexes := make([]*regexp.Regexp, len(patterns))
	for index, pattern := range patterns {

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, errors.Wrapf(err, "error compiling pattern to regexp [%s]", pattern)
		}

		regexes[index] = regex
	}

	return regexes, nil
}

// load returns a configuration from a file.
func load(path string) (Configuration, error) {

	var configuration Configuration

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return configuration, errors.Wrapf(err, "error reading configuration file [%s]", path)
	}

	err = yaml.Unmarshal(bytes, &configuration)
	if err != nil {
		return configuration, errors.Wrapf(err, "error unmarshalling configuration file [%s]", path)
	}

	return configuration, nil
}

// matches returns true if a value matches any regexes.
func matches(value string, regexes []*regexp.Regexp) bool {

	for _, regex := range regexes {
		if regex.MatchString(value) {
			return true
		}
	}

	return false
}

// random returns a random hexidecimal string.
func random() string {
	bytes := make([]byte, 16)
	rand.New(rand.NewSource(time.Now().Unix())).Read(bytes)
	return hex.EncodeToString(bytes)
}
