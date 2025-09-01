package client

import "time"

type VXUser struct {
	Id              int       `json:"id"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	EmailAddress    string    `json:"emailAddress,omitempty"`
	Owner           string    `json:"owner,omitempty"`
	UpdatedBy       string    `json:"updatedBy,omitempty"`
	Name            string    `json:"name"`
	Password        string    `json:"password,omitempty"`
	Description     string    `json:"description"`
	GroupIdList     []int     `json:"groupIdList"`
	GroupNameList   []string  `json:"groupNameList"`
	Status          int       `json:"status"`
	IsVisible       int       `json:"isVisible"`
	UserSource      int       `json:"userSource"`
	UserRoleList    []string  `json:"userRoleList"`
	OtherAttributes string    `json:"otherAttributes,omitempty"`
	SyncSource      string    `json:"syncSource,omitempty"`
	FirstName       string    `json:"firstName,omitempty"`
	LastName        string    `json:"lastName,omitempty"`
}

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
