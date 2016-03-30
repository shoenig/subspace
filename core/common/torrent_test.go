// Author hoenig

package common

// func Test_CreateBundle_file(t *testing.T) {
// 	content := []byte("this is the content of a file!")
// 	tmpfile, err := ioutil.TempFile("", "state-")
// 	require.NoError(t, err, "could not create tempfile")
// 	defer os.Remove(tmpfile.Name())
// 	_, err = tmpfile.Write(content)
// 	require.NoError(t, err, "could not write to tmpfile")
// 	err = tmpfile.Close()
// 	require.NoError(t, err, "could not close tmpfile")

// 	t.Log("path of tmpfile", tmpfile.Name())

// 	minfo, err := CreateBundle(tmpfile.Name())
// 	require.NoError(t, err, "error in create bundle")

// 	t.Log("metainfo %v", minfo)
// }
