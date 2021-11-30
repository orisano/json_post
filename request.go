package json_post

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"github.com/json-iterator/go"
)

func RequestMarshal(rawurl string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, rawurl, r)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	return nil
}

func RequestEncodeDefaultBuffer(escape bool) func(rawurl string, data interface{}) error {
	return func(rawurl string, data interface{}) error {
		var b bytes.Buffer
		enc := json.NewEncoder(&b)
		enc.SetEscapeHTML(escape)
		if err := enc.Encode(data); err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, rawurl, &b)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}
}

func RequestEncodeNewBuffer(escape bool) func(rawurl string, data interface{}) error {
	return func(rawurl string, data interface{}) error {
		b := new(bytes.Buffer)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(escape)
		if err := enc.Encode(data); err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, rawurl, b)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}
}

func RequestEncodeNewNilBuffer(escape bool) func(rawurl string, data interface{}) error {
	return func(rawurl string, data interface{}) error {
		b := bytes.NewBuffer(nil)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(escape)
		if err := enc.Encode(data); err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, rawurl, b)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}
}

func RequestEncodeReservedBuffer(sizeHint int, escape bool) func(string, interface{}) error {
	return func(rawurl string, data interface{}) error {
		b := bytes.NewBuffer(make([]byte, 0, sizeHint))
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(escape)
		if err := enc.Encode(data); err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, rawurl, b)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}
}

func RequestEncodeBufferPool(escape bool) func(string, interface{}) error {
	pool := sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(nil)
		},
	}
	return func(rawurl string, data interface{}) error {
		b := pool.Get().(*bytes.Buffer)
		defer func() {
			b.Reset()
			pool.Put(b)
		}()
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(escape)
		if err := enc.Encode(data); err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, rawurl, b)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}
}

func RequestEncodePipe(escape bool) func(rawurl string, data interface{}) error {
	return func(rawurl string, data interface{}) error {
		pr, pw := io.Pipe()
		go func() {
			enc := json.NewEncoder(pw)
			enc.SetEscapeHTML(escape)
			pw.CloseWithError(enc.Encode(data))
		}()
		req, err := http.NewRequest(http.MethodPost, rawurl, pr)
		if err != nil {
			return err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		io.Copy(ioutil.Discard, resp.Body)
		return nil
	}

}

func RequestEncodeBPipe(rawurl string, data interface{}) error {
	pr, pw := io.Pipe()
	go func() {
		bw := bufio.NewWriter(pw)
		enc := json.NewEncoder(bw)
		if err := enc.Encode(data); err != nil {
			pw.CloseWithError(err)
		}
		pw.CloseWithError(bw.Flush())
	}()
	req, err := http.NewRequest(http.MethodPost, rawurl, bufio.NewReader(pr))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil

}

type SharedBufferClient struct {
	b   *bytes.Buffer
	e   *json.Encoder
	req *http.Request
}

func (s *SharedBufferClient) RequestEncode(rawurl string, data interface{}) error {
	s.b.Reset()
	if err := s.e.Encode(data); err != nil {
		return err
	}
	u, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return err
	}
	s.req.Method = http.MethodPost
	s.req.URL = u
	s.req.Host = u.Host
	for k, _ := range s.req.Header {
		s.req.Header.Del(k)
	}
	s.req.Close = false
	resp, err := http.DefaultClient.Do(s.req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func NewSharedBufferClient() *SharedBufferClient {
	b := bytes.NewBuffer(make([]byte, 0, 1*1024*1024))
	return &SharedBufferClient{
		b: b,
		e: json.NewEncoder(b),
		req: &http.Request{
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(b),
		},
	}
}

func RequestFastMarshal(rawurl string, data interface{}) error {
	jsonLib := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := jsonLib.Marshal(data)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	req, err := http.NewRequest(http.MethodPost, rawurl, r)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	return nil
}
