/*
Copyright 2016 The Kubernetes Authors.

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

package generators

import (
	"github.com/stretchr/testify/assert"
	"k8s.io/code-generator/cmd/client-gen/types"
	genutil "k8s.io/code-generator/pkg/util"
	"testing"
)

func TestGroupPackage(t *testing.T) {

	t.Run("with dash in package name should be remove in go package", func(t *testing.T) {
		generatedPackage := groupPackage(genutil.EmptyString, types.GroupVersions{
			PackageName: "code-gen",
			Group:       genutil.EmptyString,
		}, nil)

		assert.NotNil(t, generatedPackage)
		assert.Equal(t, "codegen", generatedPackage.Name())
	})

}
