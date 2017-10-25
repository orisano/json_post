package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

type simpleStruct struct {
	ID       string  `json:"id"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Weight   float64 `json:"weight"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type nestedStruct struct {
	Locations []Location `json:"locations"`
	Name      string     `json:"name"`
	StartTime time.Time  `json:"start_time"`
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
	benchmarkRequest(b, "MarshalE              ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBuffer   ", RequestEncodeDefaultBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBufferE  ", RequestEncodeDefaultBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewBuffer       ", RequestEncodeNewBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewBufferE      ", RequestEncodeNewBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBuffer    ", RequestEncodeNewNilBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBufferE   ", RequestEncodeNewNilBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBuffer  ", RequestEncodeReservedBuffer(1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBufferE ", RequestEncodeReservedBuffer(1024, true), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBuffer     ", RequestEncodeReservedBuffer(128, false), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBufferE    ", RequestEncodeReservedBuffer(128, true), s.URL, o)
	benchmarkRequest(b, "EncodePipe            ", RequestEncodePipe(false), s.URL, o)
	benchmarkRequest(b, "EncodePipeE           ", RequestEncodePipe(true), s.URL, o)
}

func BenchmarkMiddle(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := middleObject
	benchmarkRequest(b, "MarshalE             ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBuffer  ", RequestEncodeDefaultBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBufferE ", RequestEncodeDefaultBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewBuffer      ", RequestEncodeNewBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewBufferE     ", RequestEncodeNewBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBuffer   ", RequestEncodeNewNilBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBufferE  ", RequestEncodeNewNilBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBuffer ", RequestEncodeReservedBuffer(50*1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBufferE", RequestEncodeReservedBuffer(50*1024, true), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBuffer    ", RequestEncodeReservedBuffer(32*1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBufferE   ", RequestEncodeReservedBuffer(32*1024, true), s.URL, o)
	benchmarkRequest(b, "EncodePipe           ", RequestEncodePipe(false), s.URL, o)
	benchmarkRequest(b, "EncodePipeE          ", RequestEncodePipe(true), s.URL, o)
}

func BenchmarkLarge(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := largeObject
	benchmarkRequest(b, "MarshalE              ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBuffer   ", RequestEncodeDefaultBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBufferE  ", RequestEncodeDefaultBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewBuffer       ", RequestEncodeNewBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewBufferE      ", RequestEncodeNewBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBuffer    ", RequestEncodeNewNilBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBufferE   ", RequestEncodeNewNilBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBuffer  ", RequestEncodeReservedBuffer(40*1024*1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBufferE ", RequestEncodeReservedBuffer(40*1024*1024, true), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBuffer     ", RequestEncodeReservedBuffer(31*1024*1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBufferE    ", RequestEncodeReservedBuffer(31*1024*1024, true), s.URL, o)
	benchmarkRequest(b, "EncodePipe            ", RequestEncodePipe(false), s.URL, o)
	benchmarkRequest(b, "EncodePipeE           ", RequestEncodePipe(true), s.URL, o)
}

func BenchmarkSimple(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := simpleStruct{
		ID:       "test",
		Password: "password",
		Email:    "test@example.com",
		Age:      18,
		Weight:   45.0,
	}
	benchmarkRequest(b, "MarshalE             ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBuffer  ", RequestEncodeDefaultBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBufferE ", RequestEncodeDefaultBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewBuffer      ", RequestEncodeNewBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewBufferE     ", RequestEncodeNewBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBuffer   ", RequestEncodeNewNilBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBufferE  ", RequestEncodeNewNilBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBuffer ", RequestEncodeReservedBuffer(1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBufferE", RequestEncodeReservedBuffer(1024, true), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBuffer    ", RequestEncodeReservedBuffer(128, false), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBufferE   ", RequestEncodeReservedBuffer(128, true), s.URL, o)
	benchmarkRequest(b, "EncodePipe           ", RequestEncodePipe(false), s.URL, o)
	benchmarkRequest(b, "EncodePipeE          ", RequestEncodePipe(true), s.URL, o)
}

func BenchmarkNested(b *testing.B) {
	s := httptest.NewServer(&yesman{})
	defer s.Close()
	o := nestedStruct{
		Locations: []Location{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
			{9, 10},
			{11, 12},
			{13, 14},
		},
		Name:      "Tokyo",
		StartTime: time.Unix(10000000, 1000),
	}
	benchmarkRequest(b, "MarshalE             ", RequestMarshal, s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBuffer  ", RequestEncodeDefaultBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeDefaultBufferE ", RequestEncodeDefaultBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewBuffer      ", RequestEncodeNewBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewBufferE     ", RequestEncodeNewBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBuffer   ", RequestEncodeNewNilBuffer(false), s.URL, o)
	benchmarkRequest(b, "EncodeNewNilBufferE  ", RequestEncodeNewNilBuffer(true), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBuffer ", RequestEncodeReservedBuffer(1024, false), s.URL, o)
	benchmarkRequest(b, "EncodeReservedBufferE", RequestEncodeReservedBuffer(1024, true), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBuffer    ", RequestEncodeReservedBuffer(256, false), s.URL, o)
	benchmarkRequest(b, "EncodeCheatBufferE   ", RequestEncodeReservedBuffer(256, true), s.URL, o)
	benchmarkRequest(b, "EncodePipe           ", RequestEncodePipe(false), s.URL, o)
	benchmarkRequest(b, "EncodePipeE          ", RequestEncodePipe(true), s.URL, o)
}
