/*
Copyright 2021 The Knative Authors

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

package experimental

import (
	"testing"

	"github.com/stretchr/testify/require"
	. "knative.dev/pkg/configmap/testing"
	_ "knative.dev/pkg/system/testing"
)

func TestFlags_IsEnabled_NilMap(t *testing.T) {
	require.False(t, Flags(nil).IsEnabled("myflag"))
}

func TestFlags_IsEnabled_EmptyMap(t *testing.T) {
	require.False(t, Flags{}.IsEnabled("myflag"))
}

func TestFlags_IsEnabled_ContainingFlag(t *testing.T) {
	require.True(t, Flags{
		"myflag": true,
	}.IsEnabled("myflag"))
	require.False(t, Flags{
		"myflag": false,
	}.IsEnabled("myflag"))
}

func TestGetFlags(t *testing.T) {
	_, example := ConfigMapsFromTestFile(t, FlagsConfigName)
	flags, err := NewFlagsConfigFromConfigMap(example)
	require.NoError(t, err)

	require.True(t, flags.IsEnabled("my-true-flag"))
	require.False(t, flags.IsEnabled("my-false-flag"))
	require.False(t, flags.IsEnabled("non-existing-flag"))
}
