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

func NewDatabaseResources() Resources {
	return Resources{
		Database: new(Resource),
	}
}

func NewTableResources() Resources {
	return Resources{
		Database: new(Resource),
		Table:    new(Resource),
	}
}

func NewColumnResources() Resources {
	return Resources{
		Database: new(Resource),
		Table:    new(Resource),
		Column:   new(Resource),
	}
}

func NewGlobalResources() Resources {
	return Resources{
		Global: new(Resource),
	}
}

func NewHiveServiceResources() Resources {
	return Resources{
		HiveService: new(Resource),
	}
}

func NewUdfResources() Resources {
	return Resources{
		Database: new(Resource),
		Udf:      new(Resource),
	}
}

func NewUrlResources() Resources {
	return Resources{
		Url: new(Resource),
	}
}

func NewQueueResources() Resources {
	return Resources{
		Queue: new(Resource),
	}
}

func NewCosResources() Resources {
	return Resources{
		Bucket: new(Resource),
		Path:   new(Resource),
	}
}

func NewCHDFSResources() Resources {
	return Resources{
		MountPoint: new(Resource),
		Path:       new(Resource),
	}
}

func NewKeyResources() Resources {
	return Resources{
		KeyName: new(Resource),
	}
}

func NewPathResources() Resources {
	return Resources{
		Path: new(Resource),
	}
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
