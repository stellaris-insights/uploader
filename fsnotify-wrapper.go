// Copyright © 2018 C45tr0 <william.the.developer+stellaris@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package uploader

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type FSNotify interface {
	Close() error
	Add(string) error
	Remove(string) error

	Events() chan fsnotify.Event
	Errors() chan error
}

type FSNotifyWrapper struct {
	watcher *fsnotify.Watcher
}

func NewFSNotifyWrapper() FSNotifyWrapper {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return FSNotifyWrapper {
		watcher: watcher,
	}
}

func (fsn FSNotifyWrapper) Close() error {
	return fsn.watcher.Close();
}

func (fsn FSNotifyWrapper) Add(path string) error {
	return fsn.watcher.Add(path);
}

func (fsn FSNotifyWrapper) Remove(path string) error {
	return fsn.watcher.Remove(path);
}

func (fsn FSNotifyWrapper) Events() chan fsnotify.Event {
	return fsn.watcher.Events;
}

func (fsn FSNotifyWrapper) Errors() chan error {
	return fsn.watcher.Errors;
}