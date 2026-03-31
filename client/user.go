// Package client @author: Violet-Eva @date  : 2026/3/30 @notes :
package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/violet-eva-01/ranger/types"
)

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
		if xUsers.ResultSize == 0 || (xUsers.ResultSize > 0 && xUsers.ResultSize < xUsers.PageSize) {
			break
		} else {
			startIndex += pageSize
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

func (c *Client) DisableUserById(userId int) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/secure/users/%d", userId), nil, &u); err != nil {
		return u, err
	}
	u.Disable()
	return c.UpdateUser(u)
}

func (c *Client) DisableUserByName(userName string) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, &u); err != nil {
		return u, err
	}
	u.Disable()
	return c.UpdateUser(u)
}

func (c *Client) EnableUserById(userId int) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/secure/users/%d", userId), nil, &u); err != nil {
		return u, err
	}
	u.Enable()
	return c.UpdateUser(u)
}

func (c *Client) EnableUserByName(userName string) (u types.VXUser, err error) {
	if err = c.RequestToStruct("GET", fmt.Sprintf("/xusers/users/userName/%s", userName), nil, &u); err != nil {
		return u, err
	}
	u.Enable()
	return c.UpdateUser(u)
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
