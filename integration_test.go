package main

import (
	"bytes"
	"flag"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var shouldRun = flag.Bool("integration", false, "run integration tests")

func TestPing(t *testing.T) {
	flag.Parse()
	if *shouldRun {
		resp, err := http.Get("http://localhost:8080/ping")
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		str := buf.String()
		assert.Equal(t, "\"pong\"", str)
	}
}
