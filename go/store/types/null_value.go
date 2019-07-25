// Copyright 2019 Liquidata, Inc.
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
//
// This file incorporates work covered by the following copyright and
// permission notice:
//
// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"context"

	"github.com/liquidata-inc/dolt/go/store/hash"
)

var NullValue Null

// IsNull returns true if the value is nil, or if the value is of kind NULLKind
func IsNull(val Value) bool {
	return val == nil || val.Kind() == NullKind
}

// Int is a Noms Value wrapper around the primitive int32 type.
type Null byte

// Value interface
func (v Null) Value(ctx context.Context) Value {
	return v
}

func (v Null) Equals(other Value) bool {
	return other.Kind() == NullKind
}

func (v Null) Less(nbf *NomsBinFormat, other LesserValuable) bool {
	return NullKind < other.Kind()
}

func (v Null) Hash(nbf *NomsBinFormat) hash.Hash {
	return getHash(NullValue, nbf)
}

func (v Null) WalkValues(ctx context.Context, cb ValueCallback) {
}

func (v Null) WalkRefs(nbf *NomsBinFormat, cb RefCallback) {
}

func (v Null) typeOf() *Type {
	return NullType
}

func (v Null) Kind() NomsKind {
	return NullKind
}

func (v Null) valueReadWriter() ValueReadWriter {
	return nil
}

func (v Null) writeTo(w nomsWriter, nbf *NomsBinFormat) {
	NullKind.writeTo(w, nbf)
}

func (v Null) valueBytes(nbf *NomsBinFormat) []byte {
	buff := make([]byte, 1)
	w := binaryNomsWriter{buff, 0}
	v.writeTo(&w, nbf)
	return buff
}