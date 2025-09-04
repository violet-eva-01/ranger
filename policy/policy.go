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

func (b *Policy) SetEnable(isEnable bool) {
	b.IsEnabled = isEnable
}

func (b *Policy) SetName(name string) {
	b.Name = name
}

func (b *Policy) SetPolicyType(policyType int) {
	b.PolicyType = policyType
}

func (b *Policy) SetPolicyPriority(policyPriority int) {
	b.PolicyPriority = policyPriority
}

func (b *Policy) SetDescription(description string) {
	b.Description = description
}

func (b *Policy) SetIsAuditEnabled(isAuditEnabled bool) {
	b.IsAuditEnabled = isAuditEnabled
}

func (b *Policy) SetResources(resources rtypes.Resources) {
	b.Resources = resources
}

func (b *Policy) SetPolicyItems(policyItems ...rtypes.Permission) {
	b.PolicyItems = policyItems
}

func (b *Policy) AddPolicyItems(policyItems ...rtypes.Permission) {
	b.PolicyItems = append(b.PolicyItems, policyItems...)
}

func (b *Policy) SetAllowExceptions(allowExceptions ...rtypes.Permission) {
	b.AllowExceptions = allowExceptions
}

func (b *Policy) AddAllowExceptions(allowExceptions ...rtypes.Permission) {
	b.AllowExceptions = append(b.AllowExceptions, allowExceptions...)
}

func (b *Policy) SetDenyPolicyItems(denyPolicyItems ...rtypes.Permission) {
	b.DenyPolicyItems = denyPolicyItems
}

func (b *Policy) AddDenyPolicyItems(policyItems ...rtypes.Permission) {
	b.DenyPolicyItems = append(b.DenyPolicyItems, policyItems...)
}

func (b *Policy) SetIsDenyAllElse(isDenyAllElse bool) {
	b.IsDenyAllElse = isDenyAllElse
}

func (b *Policy) SetDenyExceptions(denyExceptions ...rtypes.Permission) {
	b.DenyExceptions = denyExceptions
}

func (b *Policy) AddDenyExceptions(denyExceptions ...rtypes.Permission) {
	b.DenyExceptions = append(b.DenyExceptions, denyExceptions...)
}

func (b *Policy) SetDataMaskPolicyItems(dataMaskPolicyItems ...rtypes.DataMaskPolicyItems) {
	b.DataMaskPolicyItems = dataMaskPolicyItems
}

func (b *Policy) AddDataMaskPolicyItems(dataMaskPolicyItems ...rtypes.DataMaskPolicyItems) {
	b.DataMaskPolicyItems = append(b.DataMaskPolicyItems, dataMaskPolicyItems...)
}

func (b *Policy) SetRowFilterPolicyItems(rowFilterPolicyItems ...rtypes.RowFilterPolicyItems) {
	b.RowFilterPolicyItems = rowFilterPolicyItems
}

func (b *Policy) AddRowFilterPolicyItems(rowFilterPolicyItems ...rtypes.RowFilterPolicyItems) {
	b.RowFilterPolicyItems = append(b.RowFilterPolicyItems, rowFilterPolicyItems...)
}

func (b *Policy) SetOptions(options rtypes.Options) {
	b.Options = options
}

func (b *Policy) SetValiditySchedules(validitySchedules ...rtypes.ValiditySchedules) {
	b.ValiditySchedules = validitySchedules
}

func (b *Policy) AddValiditySchedules(validitySchedules ...rtypes.ValiditySchedules) {
	b.ValiditySchedules = append(b.ValiditySchedules, validitySchedules...)
}

func (b *Policy) SetPolicyLabel(policyLabel ...string) {
	b.PolicyLabels = policyLabel
}

func (b *Policy) AddPolicyLabel(policyLabel ...string) {
	b.PolicyLabels = rtypes.Union(b.PolicyLabels, policyLabel...)
}

func (b *Policy) DelPolicyLabel(policyLabel ...string) {
	b.PolicyLabels = rtypes.Difference(b.PolicyLabels, policyLabel...)
}

func (b *Policy) SetZoneName(zoneName string) {
	b.ZoneName = zoneName
}
