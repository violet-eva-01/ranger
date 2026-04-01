// Package client @author: Violet-Eva @date  : 2026/3/30 @notes :
package client

import (
	"encoding/json"
	"fmt"

	"github.com/violet-eva-01/ranger/policy"
)

func (c *Client) CreatePolicy(input policy.Policy) (output policy.Policy, err error) {
	var reqBody []byte
	if reqBody, err = json.Marshal(input); err != nil {
		return
	}
	err = c.RequestToStruct("POST", "/public/v2/api/policy", reqBody, output)
	return
}

func (c *Client) GetPolicyByServiceType(serviceType string) (output []policy.Policy, err error) {
	var (
		startIndex = 0
		pageSize   = 1000
	)
	for {
		var tmpPb []policy.Policy
		if err = c.RequestToStruct("GET", fmt.Sprintf("/public/v2/api/policy?startIndex=%d&pageSize=%d&serviceType=%s", startIndex, pageSize, serviceType),
			nil, &tmpPb); err != nil {
			return nil, err
		}
		output = append(output, tmpPb...)
		if len(tmpPb) == 0 || len(tmpPb) < pageSize {
			break
		} else {
			startIndex += pageSize
		}
	}
	return
}

func (c *Client) GetPolicyById(policyId int) (output policy.Policy, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/public/v2/api/policy/%d", policyId), nil, &output); err != nil {
		return
	}
	return
}

func (c *Client) GetPolicyByName(policyName string) (output []policy.Policy, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/public/v2/api/policy?policyName=%s", policyName), nil, &output); err != nil {
		return
	}
	return
}

func (c *Client) UpdatePolicy(input policy.Policy) (output policy.Policy, err error) {
	var reqBody []byte
	if reqBody, err = json.Marshal(input); err != nil {
		return
	}
	if err = c.RequestToStruct("PUT", fmt.Sprintf("/public/v2/api/policy/%d", input.Id), reqBody, &output); err != nil {
		return
	}
	return
}

func (c *Client) DisablePolicyById(policyId int) (output policy.Policy, err error) {
	var policyBody policy.Policy
	if policyBody, err = c.GetPolicyById(policyId); err != nil {
		return
	}
	policyBody.Disable()
	return c.UpdatePolicy(policyBody)
}

func (c *Client) EnablePolicyById(policyId int) (output policy.Policy, err error) {
	var policyBody policy.Policy
	if policyBody, err = c.GetPolicyById(policyId); err != nil {
		return
	}
	policyBody.Disable()
	return c.UpdatePolicy(policyBody)
}

func (c *Client) DeletePolicyById(policyId int) (respBody []byte, err error) {
	return c.Request("DELETE", fmt.Sprintf("/public/v2/api/policy/%d", policyId), nil)
}

func (c *Client) DeletePolicyByName(policyName, serviceName string) (respBody []byte, err error) {
	return c.Request("DELETE", fmt.Sprintf("/public/v2/api/policy?policyname=%sservicename=%s", policyName, serviceName), nil)
}
