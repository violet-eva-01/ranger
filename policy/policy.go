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
	PolicyType           int                           `json:"policyType,omitnil"`
	PolicyPriority       int                           `json:"policyPriority,omitnil"` //0: normal 1: overrides
	Description          string                        `json:"description,omitnil"`
	IsAuditEnabled       bool                          `json:"isAuditEnabled,omitnil"`
	Resources            rtypes.Resources              `json:"resources,omitnil"`
	PolicyItems          []rtypes.Permission           `json:"policyItems,omitempty"`          // 授权访问
	AllowExceptions      []rtypes.Permission           `json:"allowExceptions,omitempty"`      // 授权访问例外
	IsDenyAllElse        bool                          `json:"isDenyAllElse,omitnil"`          // 拒绝所有其他访问
	DenyPolicyItems      []rtypes.Permission           `json:"denyPolicyItems,omitempty"`      // 拒绝访问清单
	DenyExceptions       []rtypes.Permission           `json:"denyExceptions,omitempty"`       // 拒绝访问清单例外
	DataMaskPolicyItems  []rtypes.DataMaskPolicyItems  `json:"dataMaskPolicyItems,omitempty"`  // 加密解密时单独使用
	RowFilterPolicyItems []rtypes.RowFilterPolicyItems `json:"rowFilterPolicyItems,omitempty"` //行加密单独使用
	ServiceType          string                        `json:"serviceType,omitnil"`
	Options              rtypes.Options                `json:"options,omitnil"`
	ValiditySchedules    []rtypes.ValiditySchedules    `json:"validitySchedules,omitempty"` // 有效时间
	PolicyLabels         []string                      `json:"policyLabels,omitempty"`
	ZoneName             string                        `json:"zoneName,omitnil"`
}

func (p *Policy) SetEnable(isEnable bool) {
	p.IsEnabled = isEnable
}

func (p *Policy) SetName(name string) {
	p.Name = name
}

func (p *Policy) SetPolicyType(policyType int) {
	p.PolicyType = policyType
}

func (p *Policy) SetPolicyPriority(policyPriority int) {
	p.PolicyPriority = policyPriority
}

func (p *Policy) SetDescription(description string) {
	p.Description = description
}

func (p *Policy) SetIsAuditEnabled(isAuditEnabled bool) {
	p.IsAuditEnabled = isAuditEnabled
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

func (p *Policy) SetIsDenyAllElse(isDenyAllElse bool) {
	p.IsDenyAllElse = isDenyAllElse
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
