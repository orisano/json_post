package main

import (
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

func RequestEncodeDefaultBuffer(rawurl string, data interface{}) error {
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(data); err != nil {
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

func RequestEncodeNewBuffer(rawurl string, data interface{}) error {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(data); err != nil {
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

func RequestEncodeNewNilBuffer(rawurl string, data interface{}) error {
	b := bytes.NewBuffer(nil)
	if err := json.NewEncoder(b).Encode(data); err != nil {
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

func RequestEncodeReservedBuffer(sizeHint int) func(string, interface{}) error {
	return func(rawurl string, data interface{}) error {
		b := bytes.NewBuffer(make([]byte, 0, sizeHint))
		if err := json.NewEncoder(b).Encode(data); err != nil {
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

func RequestEncodePipe(rawurl string, data interface{}) error {
	pr, pw := io.Pipe()
	go func() {
		pw.CloseWithError(json.NewEncoder(pw).Encode(data))
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
