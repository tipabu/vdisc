// Copyright © 2019 NVIDIA Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package driver

import (
	"io"
)

// AnonymousObject represents a read-only, fixed size, random access object.
type AnonymousObject interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Size() int64
}

// Object represents a AnonymousObject with a URL
type Object interface {
	AnonymousObject
	URL() string
}

type XattrObject interface {
	Object
	GetXattr(name string) ([]byte, error)
}
