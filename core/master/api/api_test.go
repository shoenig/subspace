// Author hoenig

package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/stretchr/testify/require"
)

func Test_Client_CreateStream(t *testing.T) {
	a := API{
		store: &MockStore{},
	}
	now := time.Date(2016, 04, 17, 17, 25, 0, 0, time.UTC)

	recorder := httptest.NewRecorder()
	stream := stream.Metadata{
		Name:    "testsub",
		Owner:   "devops",
		Created: now.Unix(),
	}

	js, err := stream.JSON()
	require.NoError(t, err, "failed to jsonify stream")
	request, err := http.NewRequest("POST", "127.0.0.1:2000/v1/stream/create", strings.NewReader(js))
	require.NoError(t, err, "failed to create request")
	a.NewStream(recorder, request)

	// recorder should capture something
	require.Equal(t, 201, recorder.Code)
}
