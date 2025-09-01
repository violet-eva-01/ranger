package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/violet-eva-01/ranger/api/funcs"
)

type Session struct {
	host     string
	port     int
	path     string
	proxy    string
	userName string
	passWord string
	headers  map[string]string
}

func (s *Session) Request(method string, ApiPath string, body []byte) (respBody []byte, err error) {
	var req *http.Request

	if req, err = http.NewRequest(method,
		fmt.Sprintf("http://%s:%d/%s%s",
			s.host,
			s.port,
			s.path,
			ApiPath), bytes.NewBuffer(body)); err != nil {
		return
	}

	req.SetBasicAuth(s.userName, s.passWord)

	for k, v := range s.headers {
		req.Header.Set(k, v)
	}

	var transport *http.Transport
	if s.proxy != "" {
		var proxyUrl *url.URL
		if proxyUrl, err = url.Parse(s.proxy); err != nil {
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

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// RequestToStruct
// @Description:
// @param method 请求方法
// @param Api api api
// @param body 请求体
// @param data 需要为[struct | struct slice]指针
// @return error
func (s *Session) RequestToStruct(method string, Api string, body []byte, data any) (err error) {

	valueOf := reflect.ValueOf(data)
	if valueOf.Kind() != reflect.Ptr {
		return fmt.Errorf("data is not a pointer")
	}

	var respBody []byte
	if respBody, err = s.Request(method, Api, body); err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, &data); err != nil {
		return err
	}

	return
}

func (s *Session) GetServiceDefs() ([]ServiceDef, error) {
	var (
		sd         []ServiceDef
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var pd PluginsDefinitions
		if err := s.RequestToStruct("GET", fmt.Sprintf("/plugins/definitions?startIndex=%d&pageSize=%d", startIndex, pageSize),
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

func (s *Session) GetServiceDefsType() ([]ServiceDef, []ServiceTypeId, error) {
	var st []ServiceTypeId
	defs, err := s.GetServiceDefs()
	if err != nil {
		return nil, nil, err
	}
	for _, def := range defs {
		index := funcs.FindIndex(strings.ToLower(def.Name), serviceTypeName)
		if index >= 0 {
			var tmpSTI ServiceTypeId
			tmpSTI.ServiceTypeId = index
			tmpSTI.ServiceType = ServiceType(index)
			st = append(st, tmpSTI)
		}
	}
	return defs, st, nil
}

func (s *Session) GetPolicyByServiceName(serviceTypes ...string) (map[string][]*PolicyBody, error) {
	servicePolicyBodies := make(map[string][]*PolicyBody)
	for _, serviceType := range serviceTypes {
		var pb []*PolicyBody
		if err := s.RequestToStruct("GET",
			fmt.Sprintf("/public/v2/api/policy?pageSize=999999&serviceType=%s",
				serviceType), nil, pb); err != nil {
			return nil, err
		}
		servicePolicyBodies[serviceType] = pb
	}

	return servicePolicyBodies, nil
}

func (s *Session) GetPolicyById(policyId int) (*PolicyBody, error) {
	pb := &PolicyBody{}
	if err := s.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy/%d",
			policyId), nil, pb); err != nil {
		return nil, err
	}
	return pb, nil
}

func (s *Session) GetPolicyByName(policyName string) (*PolicyBody, error) {
	pb := &PolicyBody{}
	if err := s.RequestToStruct("GET",
		fmt.Sprintf("/public/v2/api/policy?policyName=%s",
			policyName), nil, pb); err != nil {
		return nil, err
	}
	return pb, nil
}

func (s *Session) UpdatePolicy(policy *PolicyBody) (p PolicyBody, err error) {
	pb := &PolicyBody{}
	if err = s.RequestToStruct("PUT", fmt.Sprintf("/public/v2/api/policy/%d", policy.Id), nil, pb); err != nil {
		return
	}
	p = *pb
	return
}

func (s *Session) GetUsers() ([]VXUser, error) {
	var (
		users      []VXUser
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var xUsers XUsers
		if err := s.RequestToStruct("GET", fmt.Sprintf("/xusers/users?startIndex=%d&pageSize=%d", startIndex, pageSize),
			nil, &xUsers); err != nil {
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

func (s *Session) GetUserById(userId int) (*VXUser, error) {
	var user = &VXUser{}
	if err := s.RequestToStruct("GET", fmt.Sprintf("/xusers/users/%d", userId), nil, user); err != nil {
		return user, err
	}
	return user, nil
}

func (s *Session) GetUserByName(userName string) (*VXUser, error) {
	var user = &VXUser{}
	if err := s.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, user); err != nil {
		return user, err
	}
	return user, nil
}

func (s *Session) UpdateUser(user *VXUser) (vxUser VXUser, err error) {

	var (
		reqBody []byte
	)
	if reqBody, err = json.Marshal(user); err != nil {
		return
	}

	u := &VXUser{}
	if err = s.RequestToStruct("PUT", fmt.Sprintf("/xusers/users/%d", user.Id), reqBody, u); err != nil {
		return
	}
	vxUser = *u
	return
}

func (s *Session) ChangePasswordById(userId int, newPassword string) (vxUser VXUser, err error) {

	var (
		userInformation = &VXUser{}
	)

	userInformation, err = s.GetUserById(userId)
	if err != nil {
		return
	}

	userInformation.Password = newPassword

	vxUser, err = s.UpdateUser(userInformation)
	if err != nil {
		return
	}

	return
}

func (s *Session) ChangePasswordByName(userName, newPassword string) (vxUser VXUser, err error) {

	var (
		userInformation = &VXUser{}
	)

	userInformation, err = s.GetUserByName(userName)
	if err != nil {
		return
	}

	userInformation.Password = newPassword

	vxUser, err = s.UpdateUser(userInformation)

	return
}
