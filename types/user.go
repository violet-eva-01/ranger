package types

import (
	"time"
)

type XUsers struct {
	StartIndex  int      `json:"startIndex"`
	PageSize    int      `json:"pageSize"`
	TotalCount  int      `json:"totalCount"`
	ResultSize  int      `json:"resultSize"`
	SortType    string   `json:"sortType"`
	SortBy      string   `json:"sortBy"`
	QueryTimeMS int64    `json:"queryTimeMS"`
	VXUsers     []VXUser `json:"vXUsers"`
}
type VXUser struct {
	Id              int       `json:"id"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	EmailAddress    string    `json:"emailAddress"`
	Owner           string    `json:"owner"`
	UpdatedBy       string    `json:"updatedBy"`
	Name            string    `json:"name"`
	Password        string    `json:"password"`
	Description     string    `json:"description"`
	GroupIdList     []int     `json:"groupIdList"`
	GroupNameList   []string  `json:"groupNameList"`
	Status          int       `json:"status"`
	IsVisible       int       `json:"isVisible"`
	UserSource      int       `json:"userSource"`
	UserRoleList    []string  `json:"userRoleList"`
	OtherAttributes string    `json:"otherAttributes"`
	SyncSource      string    `json:"syncSource"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
}

func (u *VXUser) SetEmailAddress(email string) {
	u.EmailAddress = email
}
func (u *VXUser) SetOwner(owner string) {
	u.Owner = owner
}
func (u *VXUser) SetName(name string) {
	u.Name = name
}
func (u *VXUser) SetPassword(password string) {
	u.Password = password
}
func (u *VXUser) SetDescription(description string) {
	u.Description = description
}
func (u *VXUser) SetGroupIdList(groupIdList ...int) {
	u.GroupIdList = groupIdList
}
func (u *VXUser) AddGroupIdList(groupIdList ...int) {
	u.GroupIdList = Union(u.GroupIdList, groupIdList...)
}
func (u *VXUser) DelGroupIdList(groupIdList ...int) {
	u.GroupIdList = Difference(u.GroupIdList, groupIdList...)
}
func (u *VXUser) SetGroupNameList(groupNameList ...string) {
	u.GroupNameList = groupNameList
}
func (u *VXUser) AddGroupNameList(groupNameList ...string) {
	u.GroupNameList = Union(u.GroupNameList, groupNameList...)
}
func (u *VXUser) DelGroupNameList(groupNameList ...string) {
	u.GroupNameList = Difference(u.GroupNameList, groupNameList...)
}
func (u *VXUser) SetStatus(isEnable int) {
	u.Status = isEnable
}
func (u *VXUser) SetIsVisible(isVisible int) {
	u.IsVisible = isVisible
}
func (u *VXUser) SetUserSource(userSource int) {
	u.UserSource = userSource
}
func (u *VXUser) SetUserRoleList(roleList ...string) {
	u.UserRoleList = roleList
}
func (u *VXUser) AddUserRoleList(roleList ...string) {
	u.UserRoleList = Union(u.UserRoleList, roleList...)
}
func (u *VXUser) DelUserRoleList(roleList ...string) {
	u.UserRoleList = Difference(u.UserRoleList, roleList...)
}
func (u *VXUser) SetOtherAttributes(otherAttributes string) {
	u.OtherAttributes = otherAttributes
}
func (u *VXUser) SetSyncSource(syncSource string) {
	u.SyncSource = syncSource
}
func (u *VXUser) SetFirstName(firstName string) {
	u.FirstName = firstName
}
func (u *VXUser) SetLastName(lastName string) {
	u.LastName = lastName
}
