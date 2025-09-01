// Package client @author: Violet-Eva @date  : 2025/8/31 @notes :
package client

import (
	"errors"
	"strings"
	"time"

	"github.com/violet-eva-01/ranger/client/functions"
)

type PermissionResolution struct {
	PolicyId          int       `json:"policy_id" spark:"column:policy_id"`
	PolicyName        string    `json:"policy_name" spark:"column:policy_name"`
	PermissionType    string    `json:"permission_type" spark:"column:permission_type"`
	Permission        []string  `json:"permission" spark:"permission"`
	ObjectType        string    `json:"object_type" spark:"column:object_type"`
	ObjectName        string    `json:"object_name" spark:"column:object_name"`
	ObjectDBName      string    `json:"object_db_name" spark:"column:object_db_name"`
	ObjectTBLName     string    `json:"object_tbl_name" spark:"column:object_tbl_name"`
	ObjectColumnName  []*string `json:"object_column" spark:"column:object_column_name"`
	ObjectRestriction []*string `json:"object_restriction" spark:"column:object_restriction"`
	GranteeType       string    `json:"grantee_type" spark:"column:grantee_type"`
	Grantee           string    `json:"grantee" spark:"column:grantee"`
	IsEnable          bool      `json:"is_enable" spark:"column:is_enable"`
	IsOverride        bool      `json:"is_override" spark:"column:is_override"`
	ValiditySchedules []string  `json:"validity_schedules" spark:"column:validity_schedules"` // startTime~endTime~timeZone 2006-01-02 15:04:05~2006-01-03 15:04:05~Asia/Shanghai
	Status            bool      `json:"status" spark:"column:status"`
}

type Object struct {
	ObjectType       string    `json:"object_type"`
	ObjectName       string    `json:"object_name"`
	ObjectDBName     string    `json:"object_db_name"`
	ObjectTBLName    string    `json:"object_tbl_name"`
	ObjectColumnName []*string `json:"object_column_name"`
}

func getObjectType(policy PolicyBody) ObjectType {

	switch policy.ServiceType {
	case "hive":
		if len(*policy.DataMaskPolicyItems) > 0 {
			return Masking
		} else if len(*policy.RowFilterPolicyItems) > 0 {
			return RowFilter
		} else if len(*policy.Resources.HiveService.Values) > 0 {
			return HiveService
		} else if len(*policy.Resources.Url.Values) > 0 {
			return Url
		} else if len(*policy.Resources.Udf.Values) > 0 {
			if len(*policy.Resources.Database.Values) > 1 {
				return Udf
			} else {
				return GlobalUdf
			}
		} else if len(*policy.Resources.Column.Values) > 0 {
			return Column
		} else if len(*policy.Resources.Table.Values) > 0 {
			return Table
		} else {
			return Database
		}
	default:
		objectType := ObjectType(functions.FindIndex(strings.ToUpper(policy.ServiceType), objectTypeName))
		return objectType
	}
}

func getObject(policy PolicyBody) (output []Object) {

	objectType := getObjectType(policy)

	switch objectType {
	case HiveService:
		for _, hiveService := range *policy.Resources.HiveService.Values {
			var tmpHiveService string
			if *hiveService == "*" {
				tmpHiveService = "ALL HIVE SERVICE"
			} else {
				tmpHiveService = *hiveService
			}
			var tmpObject Object
			tmpObject.ObjectName = tmpHiveService
			tmpObject.ObjectType = HiveService.String()
			output = append(output, tmpObject)
		}
	case GlobalUdf:
		for _, gu := range *policy.Resources.Global.Values {
			var tmpGU string
			if *gu == "*" {
				tmpGU = "ALL GLOBAL UDF"
			} else {
				tmpGU = *gu
			}
			var tmpObject Object
			tmpObject.ObjectName = tmpGU
			tmpObject.ObjectType = GlobalUdf.String()
			output = append(output, tmpObject)
		}
	case Url:
		for _, url := range *policy.Resources.Url.Values {
			var tmpURL string
			if *url == "*" {
				tmpURL = "ALL URL"
			} else {
				tmpURL = *url
			}
			var tmpObject Object
			tmpObject.ObjectName = tmpURL
			tmpObject.ObjectType = Url.String()
			output = append(output, tmpObject)
		}
	case Database:
		for _, db := range *policy.Resources.Database.Values {
			var tmpDB string
			if *db == "*" {
				tmpDB = "ALL DATABASE"
			} else {
				tmpDB = *db
			}
			var tmpObject Object
			tmpObject.ObjectDBName = tmpDB
			tmpObject.ObjectType = Database.String()
			output = append(output, tmpObject)
		}
	case Hdfs:
		for _, path := range *policy.Resources.Path.Values {
			var tmpPath string
			if *path == "*" {
				tmpPath = "ALL PATH"
			} else {
				tmpPath = *path
			}
			var tmpObject Object
			tmpObject.ObjectName = tmpPath
			tmpObject.ObjectType = Hdfs.String()
			output = append(output, tmpObject)
		}
	case Yarn:
		for _, query := range *policy.Resources.Queue.Values {
			var tmpQuery string
			if *query == "*" {
				tmpQuery = "ALL QUEUE"
			} else {
				tmpQuery = *query
			}
			var tmpObject Object
			tmpObject.ObjectName = tmpQuery
			tmpObject.ObjectType = Yarn.String()
		}
	// 为*规则不生效，不做特殊处理
	case Masking, RowFilter:
		var tmpObject Object
		dbValues := *policy.Resources.Database.Values
		tmpObject.ObjectDBName = *dbValues[0]
		tblValues := *policy.Resources.Table.Values
		tmpObject.ObjectTBLName = *tblValues[0]
		tmpObject.ObjectType = RowFilter.String()
		if objectType == Masking {
			colValues := *policy.Resources.Column.Values
			tmpObject.ObjectColumnName = colValues
			tmpObject.ObjectType = Masking.String()
		}
		output = append(output, tmpObject)
	case Chdfs:
		for _, mountPoint := range *policy.Resources.MountPoint.Values {
			var tmpMountPoint string
			if *mountPoint == "*" {
				tmpMountPoint = "ALL MOUNT POINT"
			} else {
				tmpMountPoint = *mountPoint
			}
			for _, path := range *policy.Resources.Path.Values {
				var tmpPath string
				if *path == "*" {
					tmpPath = "ALL PATH"
				} else {
					tmpPath = *path
				}
				var tmpObject Object
				tmpObject.ObjectName = tmpMountPoint + " AND " + tmpPath
				tmpObject.ObjectType = Chdfs.String()
				output = append(output, tmpObject)
			}
		}
	case Cos:
		for _, bucket := range *policy.Resources.Bucket.Values {
			var tmpBucket string
			if *bucket == "*" {
				tmpBucket = "ALL BUCKET"
			} else {
				tmpBucket = *bucket
			}
			for _, path := range *policy.Resources.Path.Values {
				var tmpPath string
				if *path == "*" {
					tmpPath = "ALL PATH"
				} else {
					tmpPath = *path
				}
				var tmpObject Object
				tmpObject.ObjectName = tmpBucket + " AND " + tmpPath
				tmpObject.ObjectType = Cos.String()
				output = append(output, tmpObject)
			}
		}
	case Table, Column:
		for _, database := range *policy.Resources.Database.Values {
			var tmpDatabase string
			if *database == "*" {
				tmpDatabase = "ALL DATABASE"
			} else {
				tmpDatabase = *database
			}
			for _, table := range *policy.Resources.Table.Values {
				var tmpTable string
				if *table == "*" {
					tmpTable = "ALL TABLE"
				} else {
					tmpTable = *table
				}
				var tmpObject Object
				tmpObject.ObjectDBName = tmpDatabase
				tmpObject.ObjectTBLName = tmpTable
				tmpObject.ObjectType = Table.String()
				if objectType == Column {
					tmpObject.ObjectColumnName = *policy.Resources.Column.Values
					tmpObject.ObjectType = Column.String()
				}

				output = append(output, tmpObject)
			}
		}
	default:
		panic("unhandled default case")
	}

	return
}

// validitySchedulesParse
// @Description:
// @param input 2006/1/2 15:04:05
// @return output
// @return err
func validitySchedulesParse(input string) (output string) {
	var (
		year, mount, day     string
		hour, minute, second string
	)
	timeArr := strings.Split(input, " ")
	splitYMD := strings.Split(timeArr[0], "/")
	year = splitYMD[0]
	if len(splitYMD[1]) != 2 {
		mount = "0" + splitYMD[1]
	} else {
		mount = splitYMD[1]
	}
	if len(splitYMD[2]) != 2 {
		day = "0" + splitYMD[2]
	} else {
		day = splitYMD[2]
	}

	splitHMS := strings.Split(timeArr[1], ":")
	if len(splitHMS[0]) != 2 {
		hour = "0" + hour
	} else {
		hour = splitHMS[0]
	}
	if len(splitHMS[1]) != 2 {
		minute = "0" + minute
	} else {
		minute = splitHMS[1]
	}
	if len(splitHMS[2]) != 2 {
		second = "0" + second
	} else {
		second = splitHMS[2]
	}

	output = year + "-" + mount + "-" + day + " " + hour + ":" + minute + ":" + second

	return
}

func getValiditySchedules(vss []*ValiditySchedules) (output []string) {

	for _, vs := range vss {
		startTime := validitySchedulesParse(*vs.StartTime)
		endTime := validitySchedulesParse(*vs.EndTime)
		tmpStr := startTime + "~" + endTime + "~" + *vs.TimeZone
		output = append(output, tmpStr)
	}

	return
}

func judgeTimeout(vss []string) (isTimeout bool, err error) {

	for _, vs := range vss {
		timeArr := strings.Split(vs, "~")
		var location *time.Location
		var parse time.Time
		location, err = time.LoadLocation(timeArr[2])
		if err != nil {
			return
		}
		parse, err = time.ParseInLocation("2006-01-02 15:04:05", timeArr[1], location)
		if err != nil {
			return
		}
		localTime := parse.Local()
		if time.Now().Local().After(localTime) {
			isTimeout = true
		} else {
			isTimeout = false
		}
	}
	return
}

func (a *PermissionResolution) assignment(policy PolicyBody, oj Object, permissions []string, permissionType string, grantee string, GranteeType string, vss []string, isTimeout bool, restrictions ...*string) {
	a.PolicyId = policy.Id
	a.PolicyName = policy.Name
	a.PermissionType = permissionType
	a.Permission = permissions
	a.ObjectType = oj.ObjectType
	a.ObjectName = oj.ObjectName
	a.ObjectDBName = strings.TrimSpace(oj.ObjectDBName)
	a.ObjectTBLName = strings.TrimSpace(oj.ObjectTBLName)
	a.ObjectColumnName = oj.ObjectColumnName
	a.ObjectRestriction = restrictions
	a.GranteeType = GranteeType
	a.Grantee = strings.TrimSpace(grantee)
	a.IsEnable = policy.IsEnabled
	a.IsOverride = policy.PolicyPriority != 0
	a.ValiditySchedules = vss
	if !a.IsEnable || isTimeout || (len(*policy.AllowExceptions) > 0 && len(*policy.PolicyItems) <= 0 && permissionType == "AllowException") || (len(*policy.DenyExceptions) > 0 && len(*policy.DenyPolicyItems) <= 0 && permissionType == "DenyException") {
		a.Status = false
	} else {
		a.Status = true
	}
}

func (p *PolicyBody) authorizeSliceAssignment(ojs []Object, users []*string, roles []*string, groups []*string, permissions []string, permissionType string, vss []string, isTimeout bool, restrictions ...*string) (output []PermissionResolution) {

	for _, oj := range ojs {
		if len(users) > 1 || (len(users) == 1 && strings.TrimSpace(*users[0]) != "") {
			for _, user := range users {
				var tmpAuth PermissionResolution
				tmpAuth.assignment(*p, oj, permissions, permissionType, *user, "USER", vss, isTimeout, restrictions...)
				output = append(output, tmpAuth)
			}
		}
		if len(groups) > 1 || (len(groups) == 1 && strings.TrimSpace(*groups[0]) != "") {
			for _, group := range groups {
				var tmpAuth PermissionResolution
				tmpAuth.assignment(*p, oj, permissions, permissionType, *group, "GROUP", vss, isTimeout, restrictions...)
				output = append(output, tmpAuth)
			}
		}
		if len(roles) > 1 || (len(roles) == 1 && strings.TrimSpace(*roles[0]) != "") {
			for _, role := range roles {
				var tmpAuth PermissionResolution
				tmpAuth.assignment(*p, oj, permissions, permissionType, *role, "ROLE", vss, isTimeout, restrictions...)
				output = append(output, tmpAuth)
			}
		}
	}

	return
}

func getPermissions(as []*Accesses) (output []string) {
	for _, i := range as {
		output = append(output, *i.Type)
	}
	return
}

func (c *Client) AccessParse(spb map[string][]*PolicyBody, st ServiceType, filters ...func([]PermissionResolution) []PermissionResolution) ([]PermissionResolution, error) {

	var (
		prs []PermissionResolution
		err error
	)

	if spb[st.String()] == nil {
		spb, err = c.GetPolicyByServiceName(st.String())
		if err != nil {
			return nil, err
		}
	}

	for _, policy := range spb[st.String()] {
		tmpPrs, err := policy.Parse()
		if err != nil {
			return nil, err
		}
		prs = append(prs, tmpPrs...)
	}

	for _, filter := range filters {
		prs = filter(prs)
	}

	return prs, nil
}

func (c *Client) AccessParseByPolicyBody(policyBodies []PolicyBody, filters ...func([]PermissionResolution) []PermissionResolution) ([]PermissionResolution, error) {

	var (
		prs []PermissionResolution
	)

	if len(policyBodies) == 0 {
		return prs, errors.New("no policy body found")
	}

	for _, policy := range policyBodies {
		tmpPrs, err := policy.Parse()
		if err != nil {
			return nil, err
		}
		prs = append(prs, tmpPrs...)
	}

	for _, filter := range filters {
		prs = filter(prs)
	}

	return prs, nil
}
