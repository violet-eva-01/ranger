package policy

import (
	rtypes "github.com/violet-eva-01/ranger/types"
)

type Policy struct {
	Id                   int                           `json:"id,omitnil"`
	Guid                 string                        `json:"guid,omitnil"`
	IsEnabled            bool                          `json:"isEnabled,omitnil"`
	Version              int                           `json:"version,omitnil"`
	Service              string                        `json:"service,omitnil"`
	Name                 string                        `json:"name,omitnil"`
	PolicyType           int                           `json:"policyType,omitnil"`     // 0: 默认 1: 脱敏 2: 过滤
	PolicyPriority       int                           `json:"policyPriority,omitnil"` // 0: normal 1: overrides
	Description          string                        `json:"description,omitnil"`
	IsAuditEnabled       bool                          `json:"isAuditEnabled,omitnil"`
	Resources            rtypes.Resources              `json:"resources,omitnil"`
	PolicyItems          []rtypes.Permission           `json:"policyItems,omitempty"`          // 授权访问
	AllowExceptions      []rtypes.Permission           `json:"allowExceptions,omitempty"`      // 授权访问例外
	IsDenyAllElse        bool                          `json:"isDenyAllElse,omitnil"`          // 拒绝所有其他访问
	DenyPolicyItems      []rtypes.Permission           `json:"denyPolicyItems,omitempty"`      // 拒绝访问清单
	DenyExceptions       []rtypes.Permission           `json:"denyExceptions,omitempty"`       // 拒绝访问清单例外
	DataMaskPolicyItems  []rtypes.DataMaskPolicyItems  `json:"dataMaskPolicyItems,omitempty"`  // 加密解密时单独使用
	RowFilterPolicyItems []rtypes.RowFilterPolicyItems `json:"rowFilterPolicyItems,omitempty"` // 行加密单独使用
	ServiceType          string                        `json:"serviceType,omitnil"`
	Options              rtypes.Options                `json:"options,omitnil"`
	ValiditySchedules    []rtypes.ValiditySchedules    `json:"validitySchedules,omitempty"` // 有效时间
	PolicyLabels         []string                      `json:"policyLabels,omitempty"`
	ZoneName             string                        `json:"zoneName,omitnil"`
}

func (p *Policy) Enable() {
	p.IsEnabled = true
}

func (p *Policy) Disable() {
	p.IsEnabled = false
}

func (p *Policy) SetName(name string) {
	p.Name = name
}

func (p *Policy) SetService(service string) {
	p.Service = service
}

func (p *Policy) SetNormal() {
	p.PolicyPriority = 0
}

func (p *Policy) SetOverrides() {
	p.PolicyPriority = 1
}

type Type int

const (
	Default Type = iota
	Mask
	Filter
)

func (p *Policy) SetPolicyType(policyType Type) {
	p.PolicyType = int(policyType)
}

func (p *Policy) SetDescription(description string) {
	p.Description = description
}

func (p *Policy) EnabledAudit() {
	p.IsAuditEnabled = true
}

func (p *Policy) DisableAudit() {
	p.IsAuditEnabled = false
}

func (p *Policy) SetResources(resources rtypes.Resources) {
	p.Resources = resources
}

func (p *Policy) SetPolicyItems(policyItems ...rtypes.Permission) {
	p.PolicyItems = policyItems
}

func (p *Policy) AddPolicyItems(policyItems ...rtypes.Permission) {
	p.PolicyItems = append(p.PolicyItems, policyItems...)
}

func (p *Policy) SetAllowExceptions(allowExceptions ...rtypes.Permission) {
	p.AllowExceptions = allowExceptions
}

func (p *Policy) AddAllowExceptions(allowExceptions ...rtypes.Permission) {
	p.AllowExceptions = append(p.AllowExceptions, allowExceptions...)
}

func (p *Policy) SetDenyPolicyItems(denyPolicyItems ...rtypes.Permission) {
	p.DenyPolicyItems = denyPolicyItems
}

func (p *Policy) AddDenyPolicyItems(policyItems ...rtypes.Permission) {
	p.DenyPolicyItems = append(p.DenyPolicyItems, policyItems...)
}

func (p *Policy) EnableDenyAllElse() {
	p.IsDenyAllElse = true
}

func (p *Policy) DisableDenyAllElse() {
	p.IsDenyAllElse = false
}

func (p *Policy) SetDenyExceptions(denyExceptions ...rtypes.Permission) {
	p.DenyExceptions = denyExceptions
}

func (p *Policy) AddDenyExceptions(denyExceptions ...rtypes.Permission) {
	p.DenyExceptions = append(p.DenyExceptions, denyExceptions...)
}

func (p *Policy) SetDataMaskPolicyItems(dataMaskPolicyItems ...rtypes.DataMaskPolicyItems) {
	p.DataMaskPolicyItems = dataMaskPolicyItems
}

func (p *Policy) AddDataMaskPolicyItems(dataMaskPolicyItems ...rtypes.DataMaskPolicyItems) {
	p.DataMaskPolicyItems = append(p.DataMaskPolicyItems, dataMaskPolicyItems...)
}

func (p *Policy) SetRowFilterPolicyItems(rowFilterPolicyItems ...rtypes.RowFilterPolicyItems) {
	p.RowFilterPolicyItems = rowFilterPolicyItems
}

func (p *Policy) AddRowFilterPolicyItems(rowFilterPolicyItems ...rtypes.RowFilterPolicyItems) {
	p.RowFilterPolicyItems = append(p.RowFilterPolicyItems, rowFilterPolicyItems...)
}

func (p *Policy) SetOptions(options rtypes.Options) {
	p.Options = options
}

func (p *Policy) SetValiditySchedules(validitySchedules ...rtypes.ValiditySchedules) {
	p.ValiditySchedules = validitySchedules
}

func (p *Policy) AddValiditySchedules(validitySchedules ...rtypes.ValiditySchedules) {
	p.ValiditySchedules = append(p.ValiditySchedules, validitySchedules...)
}

func (p *Policy) SetPolicyLabel(policyLabel ...string) {
	p.PolicyLabels = policyLabel
}

func (p *Policy) AddPolicyLabel(policyLabel ...string) {
	p.PolicyLabels = rtypes.Union(p.PolicyLabels, policyLabel...)
}

func (p *Policy) DelPolicyLabel(policyLabel ...string) {
	p.PolicyLabels = rtypes.Difference(p.PolicyLabels, policyLabel...)
}

func (p *Policy) SetZoneName(zoneName string) {
	p.ZoneName = zoneName
}

func NewBasicPolicy(service, name, desc string, label []string) (policy Policy) {

	policy.SetName(name)
	policy.Enable()
	policy.SetService(service)
	policy.SetDescription(desc)
	policy.SetPolicyLabel(label...)
	policy.EnabledAudit()
	policy.SetPolicyType(Default)

	return
}

func NewBasicPermission(accesses, users, roles, groups []string) rtypes.Permission {

	var policyAccesses []rtypes.Accesses
	for _, access := range accesses {
		var tmpPolicyAccess rtypes.Accesses
		tmpPolicyAccess.SetType(access)
		tmpPolicyAccess.SetIsAllowed(true)
		policyAccesses = append(policyAccesses, tmpPolicyAccess)
	}

	var permission rtypes.Permission
	permission.SetAccesses(policyAccesses...)
	permission.SetUsers(users...)
	permission.SetRoles(roles...)
	permission.SetGroups(groups...)

	return permission
}

func NewHiveBasicResources(dbName, tblName, colName []string) rtypes.Resources {

	var resources rtypes.Resources
	switch {
	case len(colName) > 0:
		resources = rtypes.NewColumnResources()
		resources.Column.SetValues(colName...)
		resources.Table.SetValues(tblName...)
	case len(tblName) > 0:
		resources = rtypes.NewTableResources()
		resources.Table.SetValues(tblName...)
	case len(dbName) > 0:
		resources = rtypes.NewDatabaseResources()
	}
	resources.Database.SetValues(dbName...)
	return resources
}

func NewHiveItem(service, name, desc string, label, dbName, tblName, colName, accesses, users, roles, groups []string) (policy Policy) {

	resources := NewHiveBasicResources(dbName, tblName, colName)

	permission := NewBasicPermission(accesses, users, roles, groups)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetResources(resources)
	policy.SetPolicyItems(permission)

	return
}

func NewHiveDenyItem(service, name, desc string, label, dbName, tblName, colName, accesses, users, roles, group []string) (policy Policy) {

	resources := NewHiveBasicResources(dbName, tblName, colName)

	permission := NewBasicPermission(accesses, users, roles, group)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetOverrides()
	policy.SetResources(resources)
	policy.SetDenyPolicyItems(permission)

	return
}

func NewMask(service, name, desc, dbName, tblName, colName string, label, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewColumnResources()
	resources.Database.SetValues(dbName)
	resources.Table.SetValues(tblName)
	resources.Column.SetValues(colName)

	accesses := rtypes.Accesses{
		Type:      "select",
		IsAllowed: true,
	}

	var dataMaskPolicyItems []rtypes.DataMaskPolicyItems
	if len(users) > 0 || len(roles) > 0 || len(groups) > 0 {
		var unHashDMI rtypes.DataMaskInfo
		unHashDMI.SetDataMaskType("MASK_NONE")
		var unHashDMPI rtypes.DataMaskPolicyItems
		unHashDMPI.SetUsers(users...)
		unHashDMPI.SetRoles(roles...)
		unHashDMPI.SetGroups(groups...)
		unHashDMPI.SetAccesses(accesses)
		unHashDMPI.SetDataMaskInfo(unHashDMI)
		dataMaskPolicyItems = append(dataMaskPolicyItems, unHashDMPI)
	}

	var hashDMI rtypes.DataMaskInfo
	hashDMI.SetDataMaskType("MASK_HASH")

	var hashDMPI rtypes.DataMaskPolicyItems
	hashDMPI.SetUsers("{USER}")
	hashDMPI.SetAccesses(accesses)
	hashDMPI.SetDataMaskInfo(hashDMI)
	dataMaskPolicyItems = append(dataMaskPolicyItems, hashDMPI)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetPolicyType(Mask)
	policy.SetResources(resources)
	policy.SetDataMaskPolicyItems(dataMaskPolicyItems...)

	return
}

func NewFilter(service, name, desc, dbName, tblName, filterExpr string, label, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewTableResources()
	resources.Database.SetValues(dbName)
	resources.Table.SetValues(tblName)

	accesses := rtypes.Accesses{
		Type:      "select",
		IsAllowed: true,
	}

	var rowFilterInfo rtypes.RowFilterInfo
	rowFilterInfo.SetFilterExpr(filterExpr)

	var filterPolicyItem rtypes.RowFilterPolicyItems
	filterPolicyItem.SetUsers(users...)
	filterPolicyItem.SetRoles(roles...)
	filterPolicyItem.SetGroups(groups...)
	filterPolicyItem.SetAccesses(accesses)
	filterPolicyItem.SetRowFilterInfo(rowFilterInfo)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetPolicyType(Filter)
	policy.SetResources(resources)
	policy.SetRowFilterPolicyItems(filterPolicyItem)

	return
}

func NewKMS(service, name, desc string, label, keyNames, accesses, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewKeyResources()
	resources.KeyName.SetValues(keyNames...)

	permission := NewBasicPermission(accesses, users, roles, groups)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetResources(resources)
	policy.SetPolicyItems(permission)

	return
}

func NewHDFS(service, name, desc string, label, path, accesses, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewPathResources()
	resources.Path.SetValues(path...)

	permission := NewBasicPermission(accesses, users, roles, groups)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetResources(resources)
	policy.SetPolicyItems(permission)

	return
}

func NewCos(service, name, desc string, label, buckets, paths, accesses, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewCosResources()
	resources.Bucket.SetValues(buckets...)
	resources.Path.SetValues(paths...)

	permission := NewBasicPermission(accesses, users, roles, groups)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetResources(resources)
	policy.SetPolicyItems(permission)

	return
}

func NewCHDFS(service, name, desc string, label, mountPoint, path, accesses, users, roles, groups []string) (policy Policy) {

	resources := rtypes.NewCHDFSResources()
	resources.MountPoint.SetValues(mountPoint...)
	resources.Path.SetValues(path...)

	permission := NewBasicPermission(accesses, users, roles, groups)

	policy = NewBasicPolicy(service, name, desc, label)
	policy.SetResources(resources)
	policy.SetPolicyItems(permission)

	return
}
