// Package types @author: Violet-Eva @date  : 2025/9/3 @notes :
package types

type Permission struct {
	Users         []string     `json:"users,omitnil"`
	Accesses      []Accesses   `json:"accesses,omitnil"`
	Groups        []string     `json:"groups,omitnil"`
	Roles         []string     `json:"roles,omitnil"`
	Conditions    []Conditions `json:"conditions,omitnil"`
	DelegateAdmin bool         `json:"delegateAdmin,omitnil"`
}

func (p *Permission) SetUsers(users ...string) {
	p.Users = users
}

func (p *Permission) AddUsers(users ...string) {
	p.Users = Union(p.Users, users...)
}

func (p *Permission) DelUsers(users ...string) {
	p.Users = Difference(p.Users, users...)
}

func (p *Permission) SetAccesses(accesses ...Accesses) {
	p.Accesses = accesses
}

func (p *Permission) AddAccesses(accesses ...Accesses) {
	p.Accesses = Union(p.Accesses, accesses...)
}

func (p *Permission) SetGroups(groups ...string) {
	p.Groups = groups
}
func (p *Permission) AddGroups(groups ...string) {
	p.Groups = Union(p.Groups, groups...)
}
func (p *Permission) DelGroups(groups ...string) {
	p.Groups = Difference(p.Groups, groups...)
}
func (p *Permission) SetRoles(roles ...string) {
	p.Roles = roles
}
func (p *Permission) AddRoles(roles ...string) {
	p.Roles = Union(p.Roles, roles...)
}
func (p *Permission) DelRoles(roles ...string) {
	p.Roles = Difference(p.Roles, roles...)
}
func (p *Permission) SetConditions(conditions ...Conditions) {
	p.Conditions = conditions
}
func (p *Permission) AddConditions(conditions ...Conditions) {
	p.Conditions = append(p.Conditions, conditions...)
}
func (p *Permission) SetDelegateAdmin(delegateAdmin bool) {
	p.DelegateAdmin = delegateAdmin
}

// Accesses
// @Description: 除加密解密相关权限的其他权限
type Accesses struct {
	Type      string `json:"type,omitnil"`
	IsAllowed bool   `json:"isAllowed,omitnil"`
}

func (a *Accesses) SetType(t string) {
	a.Type = t
}

func (a *Accesses) SetIsAllowed(isAllowed bool) {
	a.IsAllowed = isAllowed
}

// Conditions
// @Description: 用户自定义限制规则
type Conditions struct {
	Values []string `json:"values,omitnil"`
	Type   string   `json:"type,omitnil"`
}

func (c *Conditions) SetValues(values ...string) {
	c.Values = values
}

func (c *Conditions) AddValues(values ...string) {
	c.Values = Union(c.Values, values...)
}

func (c *Conditions) DelValues(values ...string) {
	c.Values = Difference(c.Values, values...)
}

func (c *Conditions) SetType(t string) {
	c.Type = t
}

// ValiditySchedules
// @Description: 有效时间
type ValiditySchedules struct {
	StartTime   string       `json:"startTime,omitnil"`
	EndTime     string       `json:"endTime,omitnil"`
	TimeZone    string       `json:"timeZone,omitnil"`
	Recurrences []Recurrence `json:"recurrences,omitnil"`
}

func (v *ValiditySchedules) SetStartTime(startTime string) {
	v.StartTime = startTime
}

func (v *ValiditySchedules) SetEndTime(endTime string) {
	v.EndTime = endTime
}

func (v *ValiditySchedules) SetTimeZone(timeZone string) {
	v.TimeZone = timeZone
}

func (v *ValiditySchedules) SetRecurrences(recurrences ...Recurrence) {
	v.Recurrences = recurrences
}

func (v *ValiditySchedules) AddRecurrences(recurrences ...Recurrence) {
	v.Recurrences = append(v.Recurrences, recurrences...)
}

type Recurrence struct {
	Interval *Interval `json:"interval,omitempty,omitnil"`
	Schedule *Schedule `json:"schedule,omitempty,omitnil"`
}

type Interval struct {
}

type Schedule struct {
}

// DataMaskPolicyItems
// @Description: 加密 & 授予解密权限
type DataMaskPolicyItems struct {
	Permission
	DataMaskInfo *DataMaskInfo `json:"dataMaskInfo,omitempty,omitnil"`
}

func (d *DataMaskPolicyItems) SetDataMaskInfo(dataMaskInfo DataMaskInfo) {
	d.DataMaskInfo = &dataMaskInfo
}

// DataMaskInfo
// @Description: 加密解密相关权限
type DataMaskInfo struct {
	ConditionExpr string `json:"conditionExpr,omitnil"`
	DataMaskType  string `json:"dataMaskType,omitnil"`
	ValueExpr     string `json:"valueExpr,omitnil"`
}

func (d *DataMaskInfo) SetConditionExpr(conditionExpr string) {
	d.ConditionExpr = conditionExpr
}
func (d *DataMaskInfo) SetDataMaskType(dataMaskType string) {
	d.DataMaskType = dataMaskType
}
func (d *DataMaskInfo) SetValueExpr(valueExpr string) {
	d.ValueExpr = valueExpr
}

// RowFilterPolicyItems
// @Description: 行级过滤限制
type RowFilterPolicyItems struct {
	Permission
	RowFilterInfo *RowFilterInfo `json:"rowFilterInfo,omitempty,omitnil"`
}

func (r *RowFilterPolicyItems) SetRowFilterInfo(rowFilterInfo RowFilterInfo) {
	r.RowFilterInfo = &rowFilterInfo
}

type RowFilterInfo struct {
	FilterExpr string `json:"filterExpr,omitnil"`
}

func (r *RowFilterInfo) SetFilterExpr(filterExpr string) {
	r.FilterExpr = filterExpr
}

type Options struct {
	PolicyValiditySchedules string `json:"POLICY_VALIDITY_SCHEDULES,omitnil"` // 根据有效时间自动生成
}

func (o *Options) SetPolicyValiditySchedules(policyValiditySchedules string) {
	o.PolicyValiditySchedules = policyValiditySchedules
}
