/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package minimum

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/utils/ptr"
)

func TestBasicStruct(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&BasicStruct{
		// all zero values
		IntPtrField:     ptr.To(0),
		UintPtrField:    ptr.To(uint(0)),
		TypedefPtrField: ptr.To(IntType(0)),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("intField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("intPtrField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("int16Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("int32Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("int64Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("uintField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("uintPtrField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("uint16Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("uint32Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("uint64Field"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("typedefField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("typedefPtrField"), nil, "").WithOrigin("minimum"),
	})

	// Test validation ratcheting
	st.Value(&BasicStruct{
		IntPtrField:     ptr.To(0),
		UintPtrField:    ptr.To(uint(0)),
		TypedefPtrField: ptr.To(IntType(0)),
	}).OldValue(&BasicStruct{
		IntPtrField:     ptr.To(0),
		UintPtrField:    ptr.To(uint(0)),
		TypedefPtrField: ptr.To(IntType(0)),
	}).ExpectValid()

	st.Value(&BasicStruct{
		IntField:        1,
		IntPtrField:     ptr.To(1),
		Int16Field:      1,
		Int32Field:      1,
		Int64Field:      1,
		UintField:       1,
		Uint16Field:     1,
		Uint32Field:     1,
		Uint64Field:     1,
		UintPtrField:    ptr.To(uint(1)),
		TypedefField:    IntType(1),
		TypedefPtrField: ptr.To(IntType(1)),
	}).ExpectValid()
}

func TestOptionalStruct(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&OptionalStruct{
		// zero values
		OptionalIntPtrField: ptr.To(0),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("optionalIntPtrField"), nil, "").WithOrigin("minimum"),
	})

	st.Value(&OptionalStruct{
		OptionalIntField:    1,
		OptionalIntPtrField: ptr.To(1),
	}).ExpectValid()

	st.Value(&OptionalStruct{
		OptionalIntField:    -1,
		OptionalIntPtrField: ptr.To(-1),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("optionalIntField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("optionalIntPtrField"), nil, "").WithOrigin("minimum"),
	})
}

func TestRequiredStruct(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&RequiredStruct{
		// zero values
		RequiredIntField:    0,
		RequiredIntPtrField: ptr.To(0),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Required(field.NewPath("requiredIntField"), ""),
		field.Invalid(field.NewPath("requiredIntPtrField"), nil, "").WithOrigin("minimum"),
	})

	st.Value(&RequiredStruct{
		RequiredIntField:    0,
		RequiredIntPtrField: ptr.To(0),
	}).OldValue(&RequiredStruct{
		RequiredIntField:    0,
		RequiredIntPtrField: ptr.To(0),
	}).ExpectValid()

	st.Value(&RequiredStruct{
		RequiredIntField:    1,
		RequiredIntPtrField: ptr.To(1),
	}).ExpectValid()

	st.Value(&RequiredStruct{
		RequiredIntField:    -1,
		RequiredIntPtrField: ptr.To(-1),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("requiredIntField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("requiredIntPtrField"), nil, "").WithOrigin("minimum"),
	})
}

func TestNegativeMinimumStruct(t *testing.T) {
	st := localSchemeBuilder.Test(t)

	st.Value(&NegativeMinimumStruct{
		// zero values (valid for -10)
		NegativeMinimumPtrField:         ptr.To(0),
		OptionalNegativeMinimumPtrField: ptr.To(0),
		RequiredNegativeMinimumField:    0,
		RequiredNegativeMinimumPtrField: ptr.To(0),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Required(field.NewPath("requiredNegativeMinimumField"), ""),
	})

	st.Value(&NegativeMinimumStruct{
		NegativeMinimumField:            -10,
		NegativeMinimumPtrField:         ptr.To(-10),
		OptionalNegativeMinimumField:    -10,
		OptionalNegativeMinimumPtrField: ptr.To(-10),
		RequiredNegativeMinimumField:    -10,
		RequiredNegativeMinimumPtrField: ptr.To(-10),
	}).ExpectValid()

	st.Value(&NegativeMinimumStruct{
		NegativeMinimumField:            -11,
		NegativeMinimumPtrField:         ptr.To(-11),
		OptionalNegativeMinimumField:    -11,
		OptionalNegativeMinimumPtrField: ptr.To(-11),
		RequiredNegativeMinimumField:    -11,
		RequiredNegativeMinimumPtrField: ptr.To(-11),
	}).ExpectMatches(field.ErrorMatcher{}.ByType().ByField().ByDetailSubstring().ByOrigin(), field.ErrorList{
		field.Invalid(field.NewPath("negativeMinimumField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("negativeMinimumPtrField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("optionalNegativeMinimumField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("optionalNegativeMinimumPtrField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("requiredNegativeMinimumField"), nil, "").WithOrigin("minimum"),
		field.Invalid(field.NewPath("requiredNegativeMinimumPtrField"), nil, "").WithOrigin("minimum"),
	})
}
