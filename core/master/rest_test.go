// Author hoenig

package master

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/stretchr/testify/require"
)

func Test_Client_CreateStream(t *testing.T) {
	a := API{}
	recorder := httptest.NewRecorder()
	creation := stream.Creation{
		Name:  "testsub",
		Owner: "devops",
	}
	js, err := creation.JSON()
	require.NoError(t, err, "failed to jsonify creation")
	request, err := http.NewRequest("POST", "127.0.0.1:2000/v1/stream/create", strings.NewReader(js))
	require.NoError(t, err, "failed to create request")
	a.CreateStream(recorder, request)

	// recorder should capture something
	body := recorder.Body.String()
	t.Log("body:", body)
	require.Contains(t, body, "create stream")
}
