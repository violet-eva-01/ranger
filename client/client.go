package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/violet-eva-01/ranger/policy"
	"github.com/violet-eva-01/ranger/types"
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

func (c *Client) GetServiceDefs() ([]types.ServiceDef, error) {
	var (
		sd         []types.ServiceDef
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var pd types.PluginsDefinitions
		if err := c.RequestToStruct("GET", fmt.Sprintf("/plugins/definitions?startIndex=%d&pageSize=%d", startIndex, pageSize),
			nil, &pd); err != nil {
			return nil, err
		}
		sd = append(sd, pd.ServiceDefs...)
		startIndex += pageSize
		if pd.ResultSize == 0 || (pd.ResultSize > 0 && pd.ResultSize < pd.PageSize) {
			break
		}
	}
	return sd, nil
}

func (c *Client) GetPolicyByServiceType(serviceType string) ([]policy.Policy, error) {
	var (
		pb         []policy.Policy
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var tmpPb []policy.Policy
		if err := c.RequestToStruct("GET",
			fmt.Sprintf("/public/v2/api/policy?startIndex=%d&pageSize=%d&serviceType=%s",
				startIndex, pageSize, serviceType), nil, &tmpPb); err != nil {
			return nil, err
		}
		pb = append(pb, tmpPb...)
		if len(tmpPb) == 0 || len(tmpPb) < pageSize {
			break
		}
	}

	return pb, nil
}

func (c *Client) GetPolicyById(policyId int) (output policy.Policy, err error) {
	if err = c.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy/%d",
			policyId), nil, &output); err != nil {
		return
	}
	return
}

func (c *Client) GetPolicyByName(policyName string) (pb []policy.Policy, err error) {
	if err = c.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy?policyName=%s",
			policyName), nil, &pb); err != nil {
		return
	}
	return
}

func (c *Client) UpdatePolicy(input policy.Policy) (output policy.Policy, err error) {
	var reqBody []byte
	reqBody, err = json.Marshal(input)
	if err != nil {
		return
	}
	if err = c.RequestToStruct("PUT", fmt.Sprintf("/public/v2/api/policy/%d", input.Id), reqBody, &output); err != nil {
		return
	}
	return
}

func (c *Client) GetUsers() (users []types.VXUser, err error) {
	var (
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var xUsers types.XUsers
		if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users?startIndex=%d&pageSize=%d", startIndex, pageSize),
			nil, &xUsers); err != nil {
			return
		}
		users = append(users, xUsers.VXUsers...)
		startIndex += pageSize
		if xUsers.ResultSize == 0 || (xUsers.ResultSize > 0 && xUsers.ResultSize < xUsers.PageSize) {
			break
		}
	}

	return
}

func (c *Client) GetUserById(userId int) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/secure/users/%d", userId), nil, &u); err != nil {
		return u, err
	}
	return u, nil
}

func (c *Client) GetUserByName(userName string) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, &u); err != nil {
		return u, err
	}
	return u, nil
}

func (c *Client) UpdateUser(input types.VXUser) (output types.VXUser, err error) {

	var (
		reqBody []byte
	)
	if reqBody, err = json.Marshal(input); err != nil {
		return
	}

	if err = c.RequestToStruct("PUT", fmt.Sprintf("/xusers/secure/users/%d", input.Id), reqBody, &output); err != nil {
		return
	}
	return
}

func (c *Client) DeleteUserById(userId int, isForce bool) error {
	var sb strings.Builder
	sb.WriteString("/xusers/secure/users/id/")
	sb.WriteString(strconv.Itoa(userId))
	if isForce == true {
		sb.WriteString("?forceDelete=true")
	}
	_, err := c.Request("DELETE", sb.String(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ChangePasswordById(userId int, newPassword string) (vxUser types.VXUser, err error) {

	var ui types.VXUser
	ui, err = c.GetUserById(userId)
	if err != nil {
		return
	}

	ui.SetPassword(newPassword)

	vxUser, err = c.UpdateUser(ui)
	if err != nil {
		return
	}

	return
}

func (c *Client) ChangePasswordByName(userName, newPassword string) (vxUser types.VXUser, err error) {
	var userInformation types.VXUser
	userInformation, err = c.GetUserByName(userName)
	if err != nil {
		return
	}

	userInformation.SetPassword(newPassword)

	vxUser, err = c.UpdateUser(userInformation)

	return
}
