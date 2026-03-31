package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Client struct {
	host     string
	port     int
	path     string
	proxy    string
	userName string
	passWord string
	headers  map[string]string
}

func NewClient(host string, username, password string) *Client {
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	return &Client{
		host:     host,
		port:     6080,
		path:     "/service",
		userName: username,
		passWord: password,
		headers:  headers,
	}
}

func (c *Client) SetPort(port int) *Client {
	c.port = port
	return c
}

func (c *Client) SetPath(path string) *Client {
	c.path = path
	return c
}

func (c *Client) SetProxy(proxy string) *Client {
	c.proxy = proxy
	return c
}

func (c *Client) SetHeaders(headers map[string]string) *Client {
	c.headers = headers
	return c
}

func (c *Client) AddHeaders(headers map[string]string) *Client {
	for k, v := range headers {
		c.headers[k] = v
	}
	return c
}

func (c *Client) Request(method string, api string, body []byte) (respBody []byte, err error) {
	var (
		req     *http.Request
		urlPath strings.Builder
	)

	urlPath.WriteString("http://")
	urlPath.WriteString(c.host)
	urlPath.WriteString(":")
	urlPath.WriteString(strconv.Itoa(c.port))
	urlPath.WriteString(c.path)
	urlPath.WriteString(api)

	if req, err = http.NewRequest(method, urlPath.String(), bytes.NewBuffer(body)); err != nil {
		return
	}

	req.SetBasicAuth(c.userName, c.passWord)

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	var transport *http.Transport
	if c.proxy != "" {
		var proxyUrl *url.URL
		if proxyUrl, err = url.Parse(c.proxy); err != nil {
		}
		transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	} else {
		transport = &http.Transport{}
	}

	var resp *http.Response
	if resp, err = (&http.Client{Transport: transport}).Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	httpSuccess := map[int]bool{
		http.StatusOK:             true,
		http.StatusCreated:        true,
		http.StatusAccepted:       true,
		http.StatusNoContent:      true,
		http.StatusResetContent:   true,
		http.StatusPartialContent: true,
	}
	if !httpSuccess[resp.StatusCode] {
		err = errors.New(resp.Status)
		return
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

func (c *Client) RequestToStruct(method string, api string, body []byte, data any) (err error) {

	var respBody []byte
	if respBody, err = c.Request(method, api, body); err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, &data); err != nil {
		return err
	}

	return
}
