// Author hoenig

package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidateBundle_no_name(t *testing.T) {
	require.Error(t, ValidateBundle(Bundle{
		Name:    "",
		Path:    "/tmp/abc",
		Owner:   "devops",
		Comment: "a nice little comment",
	}))
}

func Test_ValidateBundle_bad_name(t *testing.T) {
	require.Error(t, ValidateBundle(Bundle{
		Name:    "MrBigglesworth",
		Path:    "/tmp/abc",
		Owner:   "devops",
		Comment: "a nice little comment",
	}))
}

func Test_ValidateBundle_no_path(t *testing.T) {
	require.Error(t, ValidateBundle(Bundle{
		Name:    "foobar",
		Path:    "",
		Owner:   "devops",
		Comment: "a nice little comment",
	}))
}

func Test_ValidateBundle_bad_owner(t *testing.T) {
	require.Error(t, ValidateBundle(Bundle{
		Name:    "foobar",
		Path:    "/tmp/abc",
		Owner:   "dev.ops",
		Comment: "a nice little comment",
	}))
}

func Test_ValidateBundle_ok(t *testing.T) {
	require.NoError(t, ValidateBundle(Bundle{
		Name:    "foobar",
		Path:    "/tmp/abc",
		Owner:   "devops",
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
