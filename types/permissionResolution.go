package types

type PermissionResolution struct {
	PolicyId          int      `json:"policy_id" spark:"column:policy_id"`
	PolicyName        string   `json:"policy_name" spark:"column:policy_name"`
	PermissionType    string   `json:"permission_type" spark:"column:permission_type"`
	Permission        []string `json:"permission" spark:"permission"`
	ObjectType        string   `json:"object_type" spark:"column:object_type"`
	ObjectName        string   `json:"object_name" spark:"column:object_name"`
	ObjectDBName      string   `json:"object_db_name" spark:"column:object_db_name"`
	ObjectTBLName     string   `json:"object_tbl_name" spark:"column:object_tbl_name"`
	ObjectColumnName  []string `json:"object_column" spark:"column:object_column_name"`
	ObjectRestriction []string `json:"object_restriction" spark:"column:object_restriction"`
	GranteeType       string   `json:"grantee_type" spark:"column:grantee_type"`
	Grantee           string   `json:"grantee" spark:"column:grantee"`
	IsEnable          bool     `json:"is_enable" spark:"column:is_enable"`
	IsOverride        bool     `json:"is_override" spark:"column:is_override"`
	ValiditySchedules []string `json:"validity_schedules" spark:"column:validity_schedules"` // startTime~endTime~timeZone 2006-01-02 15:04:05~2006-01-03 15:04:05~Asia/Shanghai
	Status            bool     `json:"status" spark:"column:status"`
}
