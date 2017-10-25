package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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
		return nil
	}

}

func RequestEncodeBPipe(rawurl string, data interface{}) error {
	pr, pw := io.Pipe()
	go func() {
		enc := json.NewEncoder(bufio.NewWriter(pw))
		pw.CloseWithError(enc.Encode(data))
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
