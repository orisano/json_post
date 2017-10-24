package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type yesman struct{}

func (*yesman) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

var smallObject map[string]interface{} = map[string]interface{}{
	"description": "Benchmark Sample Request",
	"author":      "@orisano",
	"created_at":  "2017-10-25 01:07:20",
}

var middleObject map[string]interface{} = map[string]interface{}{
	"text": strings.Repeat("1", 30*1024),
}

var largeObject map[string]interface{} = map[string]interface{}{
	"text": strings.Repeat("1", 30*1024*1024),
}

func benchmarkRequest(b *testing.B, name string, fn func(string, interface{}) error, rawurl string, data interface{}) {
	b.Run(name, func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			err := fn(rawurl, data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSmall(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := smallObject
	benchmarkRequest(b, "RequestMarshal              ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "RequestEncodeDefaultBuffer  ", RequestEncodeDefaultBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewBuffer      ", RequestEncodeNewBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewNilBuffer   ", RequestEncodeNewNilBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeReservedBuffer ", RequestEncodeReservedBuffer(1024), s.URL, o)
	benchmarkRequest(b, "RequestEncodeCheatBuffer    ", RequestEncodeReservedBuffer(128), s.URL, o)
	benchmarkRequest(b, "RequestEncodePipe           ", RequestEncodePipe, s.URL, o)
}

func BenchmarkMiddle(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := middleObject
	benchmarkRequest(b, "RequestMarshal             ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "RequestEncodeDefaultBuffer ", RequestEncodeDefaultBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewBuffer     ", RequestEncodeNewBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewNilBuffer  ", RequestEncodeNewNilBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeReservedBuffer", RequestEncodeReservedBuffer(50*1024), s.URL, o)
	benchmarkRequest(b, "RequestEncodeCheatBuffer   ", RequestEncodeReservedBuffer(32*1024), s.URL, o)
	benchmarkRequest(b, "RequestEncodePipe          ", RequestEncodePipe, s.URL, o)
}

func BenchmarkLarge(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := largeObject
	benchmarkRequest(b, "RequestMarshal              ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "RequestEncodeDefaultBuffer  ", RequestEncodeDefaultBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewBuffer      ", RequestEncodeNewBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeNewNilBuffer   ", RequestEncodeNewNilBuffer, s.URL, o)
	benchmarkRequest(b, "RequestEncodeReservedBuffer ", RequestEncodeReservedBuffer(40*1024*1024), s.URL, o)
	benchmarkRequest(b, "RequestEncodeCheatBuffer    ", RequestEncodeReservedBuffer(32*1024*1024), s.URL, o)
	benchmarkRequest(b, "RequestEncodePipe           ", RequestEncodePipe, s.URL, o)
}
