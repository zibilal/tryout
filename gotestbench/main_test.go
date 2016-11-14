package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHandleStructAdd(t *testing.T) {
	r := request(t, "/struct?first=20&second=30")

	rw := httptest.NewRecorder()
	handleStructAdd(rw, r)

	if rw.Code == 500 {
		t.Fatal("Internal server Error: " + rw.Body.String())
	}

	if rw.Body.String() != "<h2>Here is the sum 50</h2>" {
		t.Fatal("Expected ", rw.Body.String())
	}
}

func BenchmarkHandleStructAdd(b *testing.B) {
	b.ReportAllocs()
	r := request(b, "/struct?first=20&second=30")

	for i:=0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleStructAdd(rw, r)
	}
}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		t.Fatal(err)
	}

	return req
}
