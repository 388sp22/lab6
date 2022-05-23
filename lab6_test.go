package lab6

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoEcho(t *testing.T) {
	body := "test=value&type=form&command=noecho&you=smart&keys=unsorted&nonce=" + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000))
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/echo", strings.NewReader(body))
	require.NoError(t, err)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	echo(w, r)
	assert.Empty(t, w.Body.String())
}

func TestEcho(t *testing.T) {
	body := "test=value&type=form&command=echo&you=smart&keys=unsorted&nonce=" + strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000))
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/echo", strings.NewReader(body))
	require.NoError(t, err)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	echo(w, r)
	assert.Equal(t, body, w.Body.String())
}
