// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"strings"

	"go.uber.org/zap"
)

type sanitizedString string

func (s sanitizedString) String() string {
	return strings.ReplaceAll(string(s), "\n", "\\n")
}

// UserString constructs a field with the given key and the value stripped of
// newlines. The value is sanitized lazily.
func UserString(key, val string) zap.Field {
	return zap.Stringer(key, sanitizedString(val))
}
