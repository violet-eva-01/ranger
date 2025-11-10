package types

type Resources struct {
	Database    *Resource `json:"database,omitempty,omitnil"` // hive service 相关
	Table       *Resource `json:"table,omitempty,omitnil"`
	Column      *Resource `json:"column,omitempty,omitnil"`
	Global      *Resource `json:"global,omitempty,omitnil"`
	HiveService *Resource `json:"hiveservice,omitempty,omitnil"`
	Udf         *Resource `json:"udf,omitempty,omitnil"`
	Url         *Resource `json:"url,omitempty,omitnil"`
	Bucket      *Resource `json:"bucket,omitempty,omitnil"` //tencent service 相关 cos & hdfs & chdfs
	MountPoint  *Resource `json:"mountpoint,omitempty,omitnil"`
	Path        *Resource `json:"path,omitempty,omitnil"`
	Queue       *Resource `json:"queue,omitempty,omitnil"`
	KeyName     *Resource `json:"keyname,omitempty,omitnil"` // kms service 相关
}

type Resource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

func (r *Resource) SetValues(values ...string) {
	r.Values = values
}

func (r *Resource) AddValues(values ...string) {
	r.Values = Union(r.Values, values...)
}

func (r *Resource) DelValues(values ...string) {
	r.Values = Difference(r.Values, values...)
}
func (r *Resource) SetIsExcludes(isExcludes bool) {
	r.IsExcludes = isExcludes
}
func (r *Resource) SetIsRecursive(isRecursive bool) {
	r.IsRecursive = isRecursive
}
