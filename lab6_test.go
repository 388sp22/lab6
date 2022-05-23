package lab6

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoEcho(t *testing.T) {
	body := "test=value&type=form&command=noecho&you=smart&keys=unsorted"
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/echo", strings.NewReader(body))
	require.NoError(t, err)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	echo(w, r)
	assert.Empty(t, w.Body.String())
}

func TestEcho(t *testing.T) {
	body := "test=value&type=form&command=echo&you=smart&keys=unsorted"
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/echo", strings.NewReader(body))
	require.NoError(t, err)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	echo(w, r)
	assert.Equal(t, body, w.Body.String())
}
