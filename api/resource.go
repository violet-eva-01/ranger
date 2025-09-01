package api

type Resource struct {
	Database    *DatabaseResource    `json:"database,omitempty,omitnil"` // hive service 相关
	Table       *TableResource       `json:"table,omitempty,omitnil"`
	Column      *ColumnResource      `json:"column,omitempty,omitnil"`
	Global      *GlobalResource      `json:"global,omitempty,omitnil"`
	HiveService *HiveServiceResource `json:"hiveservice,omitempty,omitnil"`
	Udf         *UDFResource         `json:"udf,omitempty,omitnil"`
	Url         *URLResource         `json:"url,omitempty,omitnil"`
	Bucket      *BucketResource      `json:"bucket,omitempty,omitnil"` //tencent service 相关 cos & hdfs & chdfs
	MountPoint  *MountPointResource  `json:"mountpoint,omitempty,omitnil"`
	Path        *PathResource        `json:"path,omitempty,omitnil"`
	Queue       *QueueResource       `json:"queue,omitempty,omitnil"`
	KeyName     *KeyNameResource     `json:"keyname,omitempty,omitnil"` // kms service 相关
}

type DatabaseResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type TableResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type ColumnResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type GlobalResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type HiveServiceResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type UDFResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type URLResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type BucketResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type MountPointResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type PathResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type QueueResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}

type KeyNameResource struct {
	Values      []string `json:"values,omitnil"`
	IsExcludes  bool     `json:"isExcludes,omitnil"`
	IsRecursive bool     `json:"isRecursive,omitnil"`
}
