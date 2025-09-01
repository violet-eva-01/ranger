// Package client @author: Violet-Eva @date  : 2025/8/31 @notes :
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/violet-eva-01/ranger/client/functions"
	"io"
	"net/http"
	"net/url"
	"reflect"
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

func NewRangerClient(host string, userName string, passWord string) *Client {
	var (
		headers = make(map[string]string)
	)

	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	return &Client{
		host:     host,
		port:     6080,
		path:     "service",
		userName: userName,
		passWord: passWord,
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
	for k, v := range headers {
		c.headers[k] = v
	}
	return c
}

func (c *Client) Request(method string, ApiPath string, body []byte) (resp *http.Response, err error) {
	var req *http.Request

	if req, err = http.NewRequest(method,
		fmt.Sprintf("http://%s:%d/%s%s",
			c.host,
			c.port,
			c.path,
			ApiPath), bytes.NewBuffer(body)); err != nil {
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

	if resp, err = (&http.Client{Transport: transport}).Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return
}

// RequestToStruct
// @Description:
// @param method 请求方法
// @param Api ranger api
// @param body 请求体
// @param data 需要为[struct | struct slice]指针
// @return error
func (c *Client) RequestToStruct(method string, Api string, body []byte, data any) error {

	valueOf := reflect.ValueOf(data)
	if valueOf.Kind() != reflect.Ptr {
		return fmt.Errorf("data is not a pointer")
	}

	resp, respErr := c.Request(method, Api, body)
	if respErr != nil {
		return respErr
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}

	if respBody == nil {
		return errors.New("response body is nil")
	}

	juErr := json.Unmarshal(respBody, &data)
	if juErr != nil {
		return juErr
	}

	return nil
}

func (c *Client) GetServiceDefs() ([]ServiceDef, error) {
	var (
		sd         []ServiceDef
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var pd PluginsDefinitions
		if err := c.RequestToStruct("GET", fmt.Sprintf("/plugins/definitions?startIndex=%d&pageSize=%d", startIndex, pageSize), nil, &pd); err != nil {
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

func (c *Client) GetServiceDefsType() ([]ServiceDef, []ServiceTypeId, error) {
	var st []ServiceTypeId
	defs, err := c.GetServiceDefs()
	if err != nil {
		return nil, nil, err
	}
	for _, def := range defs {
		index := functions.FindIndex(strings.ToLower(def.Name), serviceTypeName)
		if index >= 0 {
			var tmpSTI ServiceTypeId
			tmpSTI.ServiceTypeId = index
			tmpSTI.ServiceType = ServiceType(index)
			st = append(st, tmpSTI)
		}
	}
	return defs, st, nil
}

func (c *Client) GetPolicyByServiceName(serviceTypes ...string) (map[string][]*PolicyBody, error) {
	servicePolicyBodies := make(map[string][]*PolicyBody)
	for _, serviceType := range serviceTypes {
		var pb []*PolicyBody
		if err := c.RequestToStruct("GET",
			fmt.Sprintf("/public/v2/api/policy?pageSize=999999&serviceType=%s",
				serviceType), nil, pb); err != nil {
			return nil, err
		}
		servicePolicyBodies[serviceType] = pb
	}

	return servicePolicyBodies, nil
}

func (c *Client) GetPolicyById(policyId int) (*PolicyBody, error) {
	pb := &PolicyBody{}
	if err := c.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy/%d",
			policyId), nil, pb); err != nil {
		return nil, err
	}
	return pb, nil
}

func (c *Client) GetPolicyByName(policyName string) (*PolicyBody, error) {
	pb := &PolicyBody{}
	if err := c.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy?policyName=%s",
			policyName), nil, pb); err != nil {
		return nil, err
	}
	return pb, nil
}

func (c *Client) GetUsers() ([]VXUser, error) {
	var (
		users      []VXUser
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var xUsers XUsers
		if err := c.RequestToStruct("GET", fmt.Sprintf("/xusers/users?startIndex=%d&pageSize=%d", startIndex, pageSize), nil, &xUsers); err != nil {
			return nil, err
		}
		users = append(users, xUsers.VXUsers...)
		startIndex += pageSize
		if xUsers.ResultSize == 0 || (xUsers.ResultSize > 0 && xUsers.ResultSize < xUsers.PageSize) {
			break
		}
	}

	return users, nil
}

func (c *Client) GetUserById(userId int) (VXUser, error) {
	var user VXUser
	if err := c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/%d", userId), nil, &user); err != nil {
		return user, err
	}
	return user, nil
}

func (c *Client) GetUserByName(userName string) (VXUser, error) {
	var user VXUser
	if err := c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, &user); err != nil {
		return user, err
	}
	return user, nil
}

func (c *Client) ChangePasswordById(userId int, newPassword string) (vxUser VXUser, err error) {

	var (
		userInformation = &VXUser{}
		reqBody         []byte
	)

	err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/%d", userId), nil, userInformation)
	if err != nil {
		return
	}

	userInformation.Password = newPassword

	reqBody, err = json.Marshal(userInformation)
	if err != nil {
		return
	}

	err = c.RequestToStruct("PUT", fmt.Sprintf("/xusers/users/%d", userId), reqBody, userInformation)
	if err != nil {
		return
	}

	vxUser = *userInformation

	return
}

func (c *Client) ChangePasswordByName(userName, newPassword string) (vxUser VXUser, err error) {

	var (
		userInformation = &VXUser{}
		reqBody         []byte
	)

	err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, userInformation)
	if err != nil {
		return
	}

	userInformation.Password = newPassword

	reqBody, err = json.Marshal(userInformation)
	if err != nil {
		return
	}

	err = c.RequestToStruct("PUT", fmt.Sprintf("/xusers/users/%d", userInformation.Id), reqBody, userInformation)
	if err != nil {
		return
	}

	vxUser = *userInformation

	return
}
