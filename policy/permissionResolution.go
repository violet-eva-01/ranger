// Package policy @author: Violet-Eva @date  : 2025/9/3 @notes :
package policy

import (
	"strings"

	rtypes "github.com/violet-eva-01/ranger/types"
)

func (p *Policy) GetPermissionResolution(filters ...func([]rtypes.PermissionResolution) []rtypes.PermissionResolution) ([]rtypes.PermissionResolution, error) {
	p.ParseValiditySchedules()
	parse, err := p.getPermissionResolution()
	if err != nil {
		return nil, err
	}
	for _, filter := range filters {
		parse = filter(parse)
	}
	return parse, nil
}

func (p *Policy) getPermissionResolution() ([]rtypes.PermissionResolution, error) {
	var (
		pr        []rtypes.PermissionResolution
		vss       []string
		isTimeout bool
	)

	objects := p.getObject()
	if len(p.ValiditySchedules) > 0 {
		vss = p.ParseValiditySchedules()
		timeout, err := p.JudgeTimeout()
		if err != nil {
			return nil, err
		}
		isTimeout = timeout
	}

	getPermissions := func(as []rtypes.Accesses) (output []string) {
		for _, i := range as {
			output = append(output, i.Type)
		}
		return
	}

	if len(p.RowFilterPolicyItems) > 0 {
		for _, rf := range p.RowFilterPolicyItems {
			permissions := getPermissions(rf.Accesses)
			restriction := rf.RowFilterInfo.FilterExpr
			tmpPR := p.getPRs(objects, rf.Users, rf.Roles, rf.Groups, permissions, "", vss, isTimeout, restriction)
			pr = append(pr, tmpPR...)
		}
	}

	if len(p.DataMaskPolicyItems) > 0 {
		for _, dmp := range p.DataMaskPolicyItems {
			permissions := getPermissions(dmp.Accesses)
			restriction := dmp.DataMaskInfo.DataMaskType
			tmpPR := p.getPRs(objects, dmp.Users, dmp.Roles, dmp.Groups, permissions, "", vss, isTimeout, restriction)
			pr = append(pr, tmpPR...)
		}
	}

	if len(p.PolicyItems) > 0 {
		permissionType := "PolicyItem"
		for _, pi := range p.PolicyItems {
			permissions := getPermissions(pi.Accesses)
			tmpPR := p.getPRs(objects, pi.Users, pi.Roles, pi.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(p.DenyPolicyItems) > 0 {
		permissionType := "DenyPolicyItem"
		for _, dpi := range p.DenyPolicyItems {
			permissions := getPermissions(dpi.Accesses)
			tmpPR := p.getPRs(objects, dpi.Users, dpi.Roles, dpi.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(p.AllowExceptions) > 0 {
		permissionType := "AllowException"
		for _, ae := range p.AllowExceptions {
			permissions := getPermissions(ae.Accesses)
			tmpPR := p.getPRs(objects, ae.Users, ae.Roles, ae.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(p.DenyExceptions) > 0 {
		permissionType := "DenyException"
		for _, de := range p.DenyExceptions {
			permissions := getPermissions(de.Accesses)
			tmpPR := p.getPRs(objects, de.Users, de.Roles, de.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	return pr, nil
}

func (p *Policy) getPRs(ojs []rtypes.Object, users []string, roles []string, groups []string, permissions []string, permissionType string, vss []string, isTimeout bool, restrictions ...string) (output []rtypes.PermissionResolution) {

	for _, oj := range ojs {
		if len(users) > 1 || (len(users) == 1 && strings.TrimSpace(users[0]) != "") {
			for _, user := range users {
				tmpPR := p.getPR(oj, permissions, permissionType, user, "USER", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
		if len(groups) > 1 || (len(groups) == 1 && strings.TrimSpace(groups[0]) != "") {
			for _, group := range groups {
				tmpPR := p.getPR(oj, permissions, permissionType, group, "GROUP", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
		if len(roles) > 1 || (len(roles) == 1 && strings.TrimSpace(roles[0]) != "") {
			for _, role := range roles {
				tmpPR := p.getPR(oj, permissions, permissionType, role, "ROLE", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
	}
	return
}

func (p *Policy) getPR(oj rtypes.Object, permissions []string, permissionType string, grantee string, GranteeType string, vss []string, isTimeout bool, restrictions ...string) rtypes.PermissionResolution {
	var pr rtypes.PermissionResolution
	pr.PolicyId = p.Id
	pr.PolicyName = p.Name
	pr.PermissionType = permissionType
	pr.Permission = permissions
	pr.ObjectType = oj.ObjectType
	pr.ObjectName = oj.ObjectName
	pr.ObjectDBName = strings.TrimSpace(oj.ObjectDBName)
	pr.ObjectTBLName = strings.TrimSpace(oj.ObjectTBLName)
	pr.ObjectColumnName = oj.ObjectColumnName
	pr.ObjectRestriction = restrictions
	pr.GranteeType = GranteeType
	pr.Grantee = strings.TrimSpace(grantee)
	pr.IsEnable = p.IsEnabled
	pr.IsOverride = p.PolicyPriority != 0
	pr.ValiditySchedules = vss
	if !pr.IsEnable || isTimeout || (len(p.AllowExceptions) > 0 && len(p.PolicyItems) <= 0 && permissionType == "AllowException") || (len(p.DenyExceptions) > 0 && len(p.DenyPolicyItems) <= 0 && permissionType == "DenyException") {
		pr.Status = false
	} else {
		pr.Status = true
	}
	return pr
}

func (p *Policy) getObject() (output []rtypes.Object) {

	objectType := p.getObjectType()

	switch objectType {
	case rtypes.HiveService:
		for _, hiveService := range p.Resources.HiveService.Values {
			if hiveService == "*" {
				hiveService = "ALL HIVE SERVICE"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = hiveService
			tmpObject.ObjectType = rtypes.HiveService.String()
			output = append(output, tmpObject)
		}
	case rtypes.GlobalUdf:
		for _, gu := range p.Resources.Global.Values {
			if gu == "*" {
				gu = "ALL GLOBAL UDF"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = gu
			tmpObject.ObjectType = rtypes.GlobalUdf.String()
			output = append(output, tmpObject)
		}
	case rtypes.Url:
		for _, url := range p.Resources.Url.Values {
			if url == "*" {
				url = "ALL URL"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = url
			tmpObject.ObjectType = rtypes.Url.String()
			output = append(output, tmpObject)
		}
	case rtypes.Database:
		for _, db := range p.Resources.Database.Values {
			if db == "*" {
				db = "ALL DATABASE"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectDBName = db
			tmpObject.ObjectType = rtypes.Database.String()
			output = append(output, tmpObject)
		}
	case rtypes.Hdfs:
		for _, path := range p.Resources.Path.Values {
			if path == "*" {
				path = "ALL PATH"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = path
			tmpObject.ObjectType = rtypes.Hdfs.String()
			output = append(output, tmpObject)
		}
	case rtypes.Yarn:
		for _, query := range p.Resources.Queue.Values {
			if query == "*" {
				query = "ALL QUEUE"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = query
			tmpObject.ObjectType = rtypes.Yarn.String()
		}
	// 为*规则不生效，不做特殊处理
	case rtypes.Masking, rtypes.RowFilter:
		var tmpObject rtypes.Object
		tmpObject.ObjectDBName = p.Resources.Database.Values[0]
		tmpObject.ObjectTBLName = p.Resources.Table.Values[0]
		tmpObject.ObjectType = rtypes.RowFilter.String()
		if objectType == rtypes.Masking {
			tmpObject.ObjectColumnName = p.Resources.Column.Values
			tmpObject.ObjectType = rtypes.Masking.String()
		}
		output = append(output, tmpObject)
	case rtypes.Chdfs:
		for _, mountPoint := range p.Resources.MountPoint.Values {
			if mountPoint == "*" {
				mountPoint = "ALL MOUNT POINT"
			}
			for _, path := range p.Resources.Path.Values {
				if path == "*" {
					path = "ALL PATH"
				}
				var tmpObject rtypes.Object
				tmpObject.ObjectName = mountPoint + " AND " + path
				tmpObject.ObjectType = rtypes.Chdfs.String()
				output = append(output, tmpObject)
			}
		}
	case rtypes.Cos:
		for _, bucket := range p.Resources.Bucket.Values {
			if bucket == "*" {
				bucket = "ALL BUCKET"
			}
			for _, path := range p.Resources.Path.Values {
				if path == "*" {
					path = "ALL PATH"
				}
				var tmpObject rtypes.Object
				tmpObject.ObjectName = bucket + " AND " + path
				tmpObject.ObjectType = rtypes.Cos.String()
				output = append(output, tmpObject)
			}
		}
	case rtypes.Table, rtypes.Column:
		for _, database := range p.Resources.Database.Values {
			if database == "*" {
				database = "ALL DATABASE"
			}
			for _, table := range p.Resources.Table.Values {
				if table == "*" {
					table = "ALL TABLE"
				}
				var tmpObject rtypes.Object
				tmpObject.ObjectDBName = database
				tmpObject.ObjectTBLName = table
				tmpObject.ObjectType = rtypes.Table.String()
				if objectType == rtypes.Column {
					tmpObject.ObjectColumnName = p.Resources.Column.Values
					tmpObject.ObjectType = rtypes.Column.String()
				}

				output = append(output, tmpObject)
			}
		}
	default:
		panic("unhandled default case")
	}
	return
}

func (p *Policy) getObjectType() rtypes.ObjectType {

	switch p.ServiceType {
	case "hive":
		if len(p.DataMaskPolicyItems) > 0 {
			return rtypes.Masking
		} else if len(p.RowFilterPolicyItems) > 0 {
			return rtypes.RowFilter
		} else if p.Resources.HiveService != nil && len(p.Resources.HiveService.Values) > 0 {
			return rtypes.HiveService
		} else if p.Resources.Url != nil && len(p.Resources.Url.Values) > 0 {
			return rtypes.Url
		} else if p.Resources.Udf != nil && len(p.Resources.Udf.Values) > 0 {
			if p.Resources.Database != nil && len(p.Resources.Database.Values) > 1 {
				return rtypes.Udf
			} else {
				return rtypes.GlobalUdf
			}
		} else if p.Resources.Column != nil && len(p.Resources.Column.Values) > 0 {
			return rtypes.Column
		} else if p.Resources.Table != nil && len(p.Resources.Table.Values) > 0 {
			return rtypes.Table
		} else {
			return rtypes.Database
		}
	default:
		objectType := rtypes.GetObjectType(p.ServiceType)
		return objectType
	}
}
