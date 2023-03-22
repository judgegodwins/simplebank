package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/judgegodwins/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func requireBodyMatches[D db.Account | []db.Account | db.User](t *testing.T, body *bytes.Buffer, account D) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getData D
	err = json.Unmarshal(data, &getData)

	require.NoError(t, err)
	require.Equal(t, account, getData)
}

func setupServerAndRecorder(t *testing.T, server *Server, url string, request *http.Request) *httptest.ResponseRecorder {
	var err error
	recorder := httptest.NewRecorder()
	if request == nil {
		request, err = http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)
	}

	server.router.ServeHTTP(recorder, request)

	return recorder
}
