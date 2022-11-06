package servers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"dns/requests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	sectorID float64 = 10.0

	server *Server
)

func TestMain(m *testing.M) {
	code := 1
	defer func() {
		os.Exit(code)
	}()
	server = New(sectorID)
	code = m.Run()
}

func TestCalculate(t *testing.T) {
	req := requests.Request{
		CoordX:   "0",
		CoordY:   "0",
		CoordZ:   "0",
		Velocity: "0",
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)
	require.NotNil(t, b)

	res := handleRequest(t, http.MethodPost, "/calculate", b)
	require.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"loc":0}`, res.Body.String())
}

func TestCalculateWithMomCorp(t *testing.T) {
	req := requests.Request{
		CoordX:   "0",
		CoordY:   "0",
		CoordZ:   "0",
		Velocity: "0",
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)
	require.NotNil(t, b)

	res := handleRequest(t, http.MethodPost, "/calculate?resp=mom", b)
	require.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, `{"location":0}`, res.Body.String())
}

func TestCalculateEmptyRequest(t *testing.T) {
	req := requests.Request{}
	b, err := json.Marshal(req)
	require.NoError(t, err)
	require.NotNil(t, b)

	res := handleRequest(t, http.MethodPost, "/calculate", b)
	require.Equal(t, http.StatusBadRequest, res.Code)
	exp := `{"error":"invalid X coordinate value: ()","success":false}`
	assert.Equal(t, exp, res.Body.String())
}

func TestCalculateWithInvalidXCoord(t *testing.T) {
	req := requests.Request{
		CoordX: "hello",
	}
	b, err := json.Marshal(req)
	require.NoError(t, err)
	require.NotNil(t, b)

	res := handleRequest(t, http.MethodPost, "/calculate", b)
	require.Equal(t, http.StatusBadRequest, res.Code)
	exp := `{"error":"invalid X coordinate value: (hello)","success":false}`
	assert.Equal(t, exp, res.Body.String())
}

func TestCalculateWithInvalidJSON(t *testing.T) {
	b := []byte(`
	{"x":"0","y":"0","z":"0","vel":"0",}
	`)

	res := handleRequest(t, http.MethodPost, "/calculate", b)
	require.Equal(t, http.StatusBadRequest, res.Code)
}

func handleRequest(t *testing.T, method, path string, payload []byte) *httptest.ResponseRecorder {
	t.Helper()
	w := httptest.NewRecorder()
	var (
		req *http.Request
		err error
	)
	if len(payload) > 0 {
		fmt.Printf("\nRequest: %s\n\n", payload)
		req, err = http.NewRequest(method, path, bytes.NewBuffer(payload))
	} else {
		req, err = http.NewRequest(method, path, nil)
	}
	require.NoError(t, err)
	server.Router.ServeHTTP(w, req)

	fmt.Printf("\nResponse: %s\n", w.Body.String())

	return w
}
