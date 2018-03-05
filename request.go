package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"io/ioutil"
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
