// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func MustEncodeJSON(t *testing.T, in interface{}) string {
	var b bytes.Buffer
	require.NoError(t, json.NewEncoder(&b).Encode(in))
	return b.String()
}
