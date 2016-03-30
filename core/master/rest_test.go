// Author hoenig

package master

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shoenig/subspace/core/common/subscription"
	"github.com/stretchr/testify/require"
)

func Test_Client_CreateSubscription(t *testing.T) {
	a := API{}
	recorder := httptest.NewRecorder()
	creation := subscription.Creation{
		Name:  "testsub",
		Owner: "devops",
	}
	js, err := creation.JSON()
	require.NoError(t, err, "failed to jsonify creation")
	request, err := http.NewRequest("POST", "127.0.0.1:2000/v1/subscription/create", strings.NewReader(js))
	require.NoError(t, err, "failed to create request")
	a.CreateSubscription(recorder, request)

	// recorder should capture something
	body := recorder.Body.String()
	t.Log("body:", body)
	require.Contains(t, body, "create subscription")
}
