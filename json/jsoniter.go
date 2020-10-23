//+build !jsonstd
// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package json

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	compatible = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gin/json package.
	Marshal = compatible.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = compatible.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = compatible.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = compatible.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = compatible.NewEncoder
)
