// Package policy @author: Violet-Eva @date  : 2025/9/3 @notes :
package policy

import (
	"strings"

	rtypes "github.com/violet-eva-01/ranger/types"
)

func (b *Policy) GetPermissionResolution(filters ...func([]rtypes.PermissionResolution) []rtypes.PermissionResolution) ([]rtypes.PermissionResolution, error) {
	b.ParseValiditySchedules()
	parse, err := b.getPermissionResolution()
	if err != nil {
		return nil, err
	}
	for _, filter := range filters {
		parse = filter(parse)
	}
	return parse, nil
}

func (b *Policy) getPermissionResolution() ([]rtypes.PermissionResolution, error) {
	var (
		pr        []rtypes.PermissionResolution
		vss       []string
		isTimeout bool
	)

	objects := b.getObject()
	if len(b.ValiditySchedules) > 0 {
		vss = b.ParseValiditySchedules()
		timeout, err := b.JudgeTimeout()
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

	if len(b.RowFilterPolicyItems) > 0 {
		for _, rf := range b.RowFilterPolicyItems {
			permissions := getPermissions(rf.Accesses)
			restriction := rf.RowFilterInfo.FilterExpr
			tmpPR := b.getPRs(objects, rf.Users, rf.Roles, rf.Groups, permissions, "", vss, isTimeout, restriction)
			pr = append(pr, tmpPR...)
		}
	}

	if len(b.DataMaskPolicyItems) > 0 {
		for _, dmp := range b.DataMaskPolicyItems {
			permissions := getPermissions(dmp.Accesses)
			restriction := dmp.DataMaskInfo.DataMaskType
			tmpPR := b.getPRs(objects, dmp.Users, dmp.Roles, dmp.Groups, permissions, "", vss, isTimeout, restriction)
			pr = append(pr, tmpPR...)
		}
	}

	if len(b.PolicyItems) > 0 {
		permissionType := "PolicyItem"
		for _, pi := range b.PolicyItems {
			permissions := getPermissions(pi.Accesses)
			tmpPR := b.getPRs(objects, pi.Users, pi.Roles, pi.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(b.DenyPolicyItems) > 0 {
		permissionType := "DenyPolicyItem"
		for _, dpi := range b.DenyPolicyItems {
			permissions := getPermissions(dpi.Accesses)
			tmpPR := b.getPRs(objects, dpi.Users, dpi.Roles, dpi.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(b.AllowExceptions) > 0 {
		permissionType := "AllowException"
		for _, ae := range b.AllowExceptions {
			permissions := getPermissions(ae.Accesses)
			tmpPR := b.getPRs(objects, ae.Users, ae.Roles, ae.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	if len(b.DenyExceptions) > 0 {
		permissionType := "DenyException"
		for _, de := range b.DenyExceptions {
			permissions := getPermissions(de.Accesses)
			tmpPR := b.getPRs(objects, de.Users, de.Roles, de.Groups, permissions, permissionType, vss, isTimeout)
			pr = append(pr, tmpPR...)
		}
	}

	return pr, nil
}

func (b *Policy) getPRs(ojs []rtypes.Object, users []string, roles []string, groups []string, permissions []string, permissionType string, vss []string, isTimeout bool, restrictions ...string) (output []rtypes.PermissionResolution) {

	for _, oj := range ojs {
		if len(users) > 1 || (len(users) == 1 && strings.TrimSpace(users[0]) != "") {
			for _, user := range users {
				tmpPR := b.getPR(oj, permissions, permissionType, user, "USER", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
		if len(groups) > 1 || (len(groups) == 1 && strings.TrimSpace(groups[0]) != "") {
			for _, group := range groups {
				tmpPR := b.getPR(oj, permissions, permissionType, group, "GROUP", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
		if len(roles) > 1 || (len(roles) == 1 && strings.TrimSpace(roles[0]) != "") {
			for _, role := range roles {
				tmpPR := b.getPR(oj, permissions, permissionType, role, "ROLE", vss, isTimeout, restrictions...)
				output = append(output, tmpPR)
			}
		}
	}
	return
}

func (b *Policy) getPR(oj rtypes.Object, permissions []string, permissionType string, grantee string, GranteeType string, vss []string, isTimeout bool, restrictions ...string) rtypes.PermissionResolution {
	var pr rtypes.PermissionResolution
	pr.PolicyId = b.Id
	pr.PolicyName = b.Name
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
	pr.IsEnable = b.IsEnabled
	pr.IsOverride = b.PolicyPriority != 0
	pr.ValiditySchedules = vss
	if !pr.IsEnable || isTimeout || (len(b.AllowExceptions) > 0 && len(b.PolicyItems) <= 0 && permissionType == "AllowException") || (len(b.DenyExceptions) > 0 && len(b.DenyPolicyItems) <= 0 && permissionType == "DenyException") {
		pr.Status = false
	} else {
		pr.Status = true
	}
	return pr
}

func (b *Policy) getObject() (output []rtypes.Object) {

	objectType := b.getObjectType()

	switch objectType {
	case rtypes.HiveService:
		for _, hiveService := range b.Resources.HiveService.Values {
			if hiveService == "*" {
				hiveService = "ALL HIVE SERVICE"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = hiveService
			tmpObject.ObjectType = rtypes.HiveService.String()
			output = append(output, tmpObject)
		}
	case rtypes.GlobalUdf:
		for _, gu := range b.Resources.Global.Values {
			if gu == "*" {
				gu = "ALL GLOBAL UDF"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = gu
			tmpObject.ObjectType = rtypes.GlobalUdf.String()
			output = append(output, tmpObject)
		}
	case rtypes.Url:
		for _, url := range b.Resources.Url.Values {
			if url == "*" {
				url = "ALL URL"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = url
			tmpObject.ObjectType = rtypes.Url.String()
			output = append(output, tmpObject)
		}
	case rtypes.Database:
		for _, db := range b.Resources.Database.Values {
			if db == "*" {
				db = "ALL DATABASE"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectDBName = db
			tmpObject.ObjectType = rtypes.Database.String()
			output = append(output, tmpObject)
		}
	case rtypes.Hdfs:
		for _, path := range b.Resources.Path.Values {
			if path == "*" {
				path = "ALL PATH"
			}
			var tmpObject rtypes.Object
			tmpObject.ObjectName = path
			tmpObject.ObjectType = rtypes.Hdfs.String()
			output = append(output, tmpObject)
		}
	case rtypes.Yarn:
		for _, query := range b.Resources.Queue.Values {
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
		tmpObject.ObjectDBName = b.Resources.Database.Values[0]
		tmpObject.ObjectTBLName = b.Resources.Table.Values[0]
		tmpObject.ObjectType = rtypes.RowFilter.String()
		if objectType == rtypes.Masking {
			tmpObject.ObjectColumnName = b.Resources.Column.Values
			tmpObject.ObjectType = rtypes.Masking.String()
		}
		output = append(output, tmpObject)
	case rtypes.Chdfs:
		for _, mountPoint := range b.Resources.MountPoint.Values {
			if mountPoint == "*" {
				mountPoint = "ALL MOUNT POINT"
			}
			for _, path := range b.Resources.Path.Values {
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
		for _, bucket := range b.Resources.Bucket.Values {
			if bucket == "*" {
				bucket = "ALL BUCKET"
			}
			for _, path := range b.Resources.Path.Values {
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
		for _, database := range b.Resources.Database.Values {
			if database == "*" {
				database = "ALL DATABASE"
			}
			for _, table := range b.Resources.Table.Values {
				if table == "*" {
					table = "ALL TABLE"
				}
				var tmpObject rtypes.Object
				tmpObject.ObjectDBName = database
				tmpObject.ObjectTBLName = table
				tmpObject.ObjectType = rtypes.Table.String()
				if objectType == rtypes.Column {
					tmpObject.ObjectColumnName = b.Resources.Column.Values
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

func (b *Policy) getObjectType() rtypes.ObjectType {

	switch b.ServiceType {
	case "hive":
		if len(b.DataMaskPolicyItems) > 0 {
			return rtypes.Masking
		} else if len(b.RowFilterPolicyItems) > 0 {
			return rtypes.RowFilter
		} else if b.Resources.HiveService != nil && len(b.Resources.HiveService.Values) > 0 {
			return rtypes.HiveService
		} else if b.Resources.Url != nil && len(b.Resources.Url.Values) > 0 {
			return rtypes.Url
		} else if b.Resources.Udf != nil && len(b.Resources.Udf.Values) > 0 {
			if b.Resources.Database != nil && len(b.Resources.Database.Values) > 1 {
				return rtypes.Udf
			} else {
				return rtypes.GlobalUdf
			}
		} else if b.Resources.Column != nil && len(b.Resources.Column.Values) > 0 {
			return rtypes.Column
		} else if b.Resources.Table != nil && len(b.Resources.Table.Values) > 0 {
			return rtypes.Table
		} else {
			return rtypes.Database
		}
	default:
		objectType := rtypes.GetObjectType(b.ServiceType)
		return objectType
	}
}
