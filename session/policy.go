package session

// PolicyBody
// @Description: ranger policy 和 hdfs hive yarn cos service 相关的 body
type PolicyBody struct {
	Id                   int                      `json:"id,omitnil"`
	Guid                 string                   `json:"guid,omitnil"`
	IsEnabled            bool                     `json:"isEnabled,omitnil"`
	Version              int                      `json:"version,omitnil"`
	Service              string                   `json:"service,omitnil"`
	Name                 string                   `json:"name,omitnil"`
	PolicyType           int                      `json:"policyType,omitnil"`
	PolicyPriority       int                      `json:"policyPriority,omitnil"` //0: normal 1: overrides
	Description          string                   `json:"description,omitnil"`
	IsAuditEnabled       bool                     `json:"isAuditEnabled,omitnil"`
	Resources            *Resource                `json:"resources,omitnil"`
	PolicyItems          *[]*PolicyItems          `json:"policyItems,omitempty"`          // 授权访问
	AllowExceptions      *[]*AllowExceptions      `json:"allowExceptions,omitempty"`      // 授权访问例外
	IsDenyAllElse        bool                     `json:"isDenyAllElse,omitnil"`          // 拒绝所有其他访问
	DenyPolicyItems      *[]*DenyPolicyItems      `json:"denyPolicyItems,omitempty"`      // 拒绝访问清单
	DenyExceptions       *[]*DenyExceptions       `json:"denyExceptions,omitempty"`       // 拒绝访问清单例外
	DataMaskPolicyItems  *[]*DataMaskPolicyItems  `json:"dataMaskPolicyItems,omitempty"`  // 加密解密时单独使用
	RowFilterPolicyItems *[]*RowFilterPolicyItems `json:"rowFilterPolicyItems,omitempty"` //行加密单独使用
	ServiceType          string                   `json:"serviceType,omitnil"`
	Options              *Options                 `json:"options,omitnil"`
	ValiditySchedules    *[]*ValiditySchedules    `json:"validitySchedules,omitempty"` // 有效时间
	PolicyLabels         []string                 `json:"policyLabels,omitempty"`
	ZoneName             string                   `json:"zoneName,omitnil"`
}

func (p *PolicyBody) Parse() ([]PermissionResolution, error) {
	var (
		authorizes []PermissionResolution
		vss        []string
		isTimeout  bool
	)

	objects := getObject(*p)
	if len(*p.ValiditySchedules) > 0 {
		vss = getValiditySchedules(*p.ValiditySchedules)
		timeout, err := judgeTimeout(vss)
		if err != nil {
			return nil, err
		}
		isTimeout = timeout
	}

	if len(*p.RowFilterPolicyItems) > 0 {
		for _, rf := range *p.RowFilterPolicyItems {
			permissions := getPermissions(rf.Accesses)
			restriction := rf.RowFilterInfo.FilterExpr
			authorizeSlice := p.authorizeSliceAssignment(objects, rf.Users, rf.Roles, rf.Groups, permissions, "", vss, isTimeout, restriction)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	if len(*p.DataMaskPolicyItems) > 0 {
		for _, dmp := range *p.DataMaskPolicyItems {
			permissions := getPermissions(dmp.Accesses)
			restriction := dmp.DataMaskInfo.DataMaskType
			authorizeSlice := p.authorizeSliceAssignment(objects, dmp.Users, dmp.Roles, dmp.Groups, permissions, "", vss, isTimeout, restriction)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	if len(*p.PolicyItems) > 0 {
		permissionType := "PolicyItem"
		for _, pi := range *p.PolicyItems {
			permissions := getPermissions(pi.Accesses)
			authorizeSlice := p.authorizeSliceAssignment(objects, pi.Users, pi.Roles, pi.Groups, permissions, permissionType, vss, isTimeout)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	if len(*p.DenyPolicyItems) > 0 {
		permissionType := "DenyPolicyItem"
		for _, dpi := range *p.DenyPolicyItems {
			permissions := getPermissions(dpi.Accesses)
			authorizeSlice := p.authorizeSliceAssignment(objects, dpi.Users, dpi.Roles, dpi.Groups, permissions, permissionType, vss, isTimeout)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	if len(*p.AllowExceptions) > 0 {
		permissionType := "AllowException"
		for _, ae := range *p.AllowExceptions {
			permissions := getPermissions(ae.Accesses)
			authorizeSlice := p.authorizeSliceAssignment(objects, ae.Users, ae.Roles, ae.Groups, permissions, permissionType, vss, isTimeout)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	if len(*p.DenyExceptions) > 0 {
		permissionType := "DenyException"
		for _, de := range *p.DenyExceptions {
			permissions := getPermissions(de.Accesses)
			authorizeSlice := p.authorizeSliceAssignment(objects, de.Users, de.Roles, de.Groups, permissions, permissionType, vss, isTimeout)
			authorizes = append(authorizes, authorizeSlice...)
		}
	}

	return authorizes, nil
}

// Accesses
// @Description: 除加密解密相关权限的其他权限
type Accesses struct {
	Type      *string `json:"type,omitnil"`
	IsAllowed *bool   `json:"isAllowed,omitnil"`
}

// Conditions
// @Description: 用户自定义限制规则
type Conditions struct {
	Values []*string `json:"values,omitnil"`
	Type   *string   `json:"type,omitnil"`
}

// PolicyItems
// @Description: 授权
type PolicyItems struct {
	Users         []*string     `json:"users,omitnil"`
	Accesses      []*Accesses   `json:"accesses,omitnil"`
	Groups        []*string     `json:"groups,omitnil"`
	Roles         []*string     `json:"roles,omitnil"`
	Conditions    []*Conditions `json:"conditions,omitnil"`
	DelegateAdmin *bool         `json:"delegateAdmin,omitnil"`
}

// AllowExceptions
// @Description: 除外授权
type AllowExceptions struct {
	Users         []*string     `json:"users,omitnil"`
	Accesses      []*Accesses   `json:"accesses,omitnil"`
	Groups        []*string     `json:"groups,omitnil"`
	Roles         []*string     `json:"roles,omitnil"`
	Conditions    []*Conditions `json:"conditions,omitnil"`
	DelegateAdmin *bool         `json:"delegateAdmin,omitnil"`
}

// DenyPolicyItems
// @Description: 回收权限
type DenyPolicyItems struct {
	Users         []*string     `json:"users,omitnil"`
	Accesses      []*Accesses   `json:"accesses,omitnil"`
	Groups        []*string     `json:"groups,omitnil"`
	Roles         []*string     `json:"roles,omitnil"`
	Conditions    []*Conditions `json:"conditions,omitnil"`
	DelegateAdmin *bool         `json:"delegateAdmin,omitnil"`
}

// DenyExceptions
// @Description: 除外回收权限
type DenyExceptions struct {
	Users         []*string     `json:"users,omitnil"`
	Accesses      []*Accesses   `json:"accesses,omitnil"`
	Groups        []*string     `json:"groups,omitnil"`
	Roles         []*string     `json:"roles,omitnil"`
	Conditions    []*Conditions `json:"conditions,omitnil"`
	DelegateAdmin *bool         `json:"delegateAdmin,omitnil"`
}

// ValiditySchedules
// @Description: 有效时间
type ValiditySchedules struct {
	StartTime   *string       `json:"startTime,omitnil"`
	EndTime     *string       `json:"endTime,omitnil"`
	TimeZone    *string       `json:"timeZone,omitnil"`
	Recurrences []*Recurrence `json:"recurrences,omitnil"`
}

type Recurrence struct {
	Interval *Interval `json:"interval,omitnil"`
	Schedule *Schedule `json:"schedule,omitnil"`
}

type Interval struct {
}

type Schedule struct {
}

// DataMaskInfo
// @Description: 加密解密相关权限
type DataMaskInfo struct {
	ConditionExpr *string `json:"conditionExpr,omitnil"`
	DataMaskType  *string `json:"dataMaskType,omitnil"`
	ValueExpr     *string `json:"valueExpr,omitnil"`
}

// DataMaskPolicyItems
// @Description: 加密 & 授予解密权限
type DataMaskPolicyItems struct {
	DataMaskInfo  *DataMaskInfo `json:"dataMaskInfo,omitnil"`
	Users         []*string     `json:"users,omitnil"`
	Accesses      []*Accesses   `json:"accesses,omitnil"`
	Groups        []*string     `json:"groups,omitnil"`
	Roles         []*string     `json:"roles,omitnil"`
	Conditions    []*Conditions `json:"conditions,omitnil"`
	DelegateAdmin *bool         `json:"delegateAdmin,omitnil"`
}

// RowFilterPolicyItems
// @Description: 行级过滤限制
type RowFilterPolicyItems struct {
	RowFilterInfo *RowFilterInfo `json:"rowFilterInfo,omitnil"`
	Users         []*string      `json:"users,omitnil"`
	Accesses      []*Accesses    `json:"accesses,omitnil"`
	Groups        []*string      `json:"groups,omitnil"`
	Roles         []*string      `json:"roles,omitnil"`
	Conditions    []*Conditions  `json:"conditions,omitnil"`
	DelegateAdmin *bool          `json:"delegateAdmin,omitnil"`
}

type RowFilterInfo struct {
	FilterExpr *string `json:"filterExpr,omitnil"`
}

type Options struct {
	PolicyValiditySchedules *string `json:"POLICY_VALIDITY_SCHEDULES,omitnil"` // 根据有效时间自动生成
}
