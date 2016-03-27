// Author hoenig

package state

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Validate(t *testing.T) {
	require.Error(t, Validate(Bundle{
		Name:    "",
		Path:    "/tmp/abc",
		Owner:   "devops",
		Comment: "a nice little comment",
	}))

	require.Error(t, Validate(Bundle{
		Name:    "foobar",
		Path:    "",
		Owner:   "devops",
		Comment: "a nice little comment",
	}))

	require.NoError(t, Validate(Bundle{
		Name:    "foobar",
		Path:    "/tmp/abc",
		Owner:   "",
		Comment: "",
	}))
}

func Test_String(t *testing.T) {
	b := Bundle{
		Name:    "foobar",
		Path:    "/tmp/abc",
		Owner:   "devops",
		Comment: "a nice little comment",
	}
	str := b.String()
	require.Contains(t, str, "foobar")
	require.Contains(t, str, "/tmp/abc")
	require.Contains(t, str, "devops")
	require.Contains(t, str, "a nice little comment")
}
