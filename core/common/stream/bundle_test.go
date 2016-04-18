// Author hoenig

package stream

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidateBundle_no_name(t *testing.T) {
	require.Error(t, Bundle{
		Stream:  "",
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_bad_name(t *testing.T) {
	require.Error(t, Bundle{
		Stream:  "MrBigglesworth",
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_no_path(t *testing.T) {
	require.Error(t, Bundle{
		Stream:  "foobar",
		Path:    "",
		Comment: "a nice little comment",
	}.valid())
}

func Test_ValidateBundle_ok(t *testing.T) {
	require.NoError(t, Bundle{
		Stream:  "foobar",
		Path:    "/tmp/abc",
		Comment: "",
	}.valid())
}

func Test_String(t *testing.T) {
	b := Bundle{
		Stream:  "foobar",
		Path:    "/tmp/abc",
		Comment: "a nice little comment",
	}
	str, err := b.JSON()
	require.NoError(t, err)
	require.Contains(t, str, "foobar")
	require.Contains(t, str, "/tmp/abc")
}
