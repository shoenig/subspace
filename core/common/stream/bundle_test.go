// Author hoenig

package stream

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidateBundle_no_name(t *testing.T) {
	require.Error(t, Bundle{
		Info: Info{
			Name:  "",
			Owner: "devops",
		},
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_bad_name(t *testing.T) {
	require.Error(t, Bundle{
		Info: Info{
			Name:  "MrBigglesworth",
			Owner: "devops",
		},
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_no_path(t *testing.T) {
	require.Error(t, Bundle{
		Info: Info{
			Name:  "foobar",
			Owner: "devops",
		},
		Path:    "",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_bad_owner(t *testing.T) {
	require.Error(t, Bundle{
		Info: Info{
			Name:  "foobar",
			Owner: "dev.ops",
		},
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_ok(t *testing.T) {
	require.NoError(t, Bundle{
		Info: Info{
			Name:  "foobar",
			Owner: "devops",
		},
		Path:    "/tmp/abc",
		Comment: "",
	}.valid())
}

func Test_String(t *testing.T) {
	b := Bundle{
		Info: Info{
			Name:  "foobar",
			Owner: "devops",
		},
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}
	str, err := b.JSON()
	require.NoError(t, err)
	require.Contains(t, str, "foobar")
	require.Contains(t, str, "/tmp/abc")
	require.Contains(t, str, "devops")
	require.Contains(t, str, "a nice little comment")
}
