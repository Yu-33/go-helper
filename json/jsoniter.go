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

var (
	_ = Marshal
	_ = Unmarshal
	_ = MarshalIndent
	_ = NewDecoder
	_ = NewEncoder
)
