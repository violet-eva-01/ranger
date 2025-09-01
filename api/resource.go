package api

type Resource struct {
	Database    *DatabaseResource    `json:"database,omitnil"` // hive service 相关
	Table       *TableResource       `json:"table,omitnil"`
	Column      *ColumnResource      `json:"column,omitnil"`
	Global      *GlobalResource      `json:"global,omitnil"`
	HiveService *HiveServiceResource `json:"hiveservice,omitnil"`
	Udf         *UDFResource         `json:"udf,omitnil"`
	Url         *URLResource         `json:"url,omitnil"`
	Bucket      *BucketResource      `json:"bucket,omitnil"` //tencent service 相关 cos & hdfs & chdfs
	MountPoint  *MountPointResource  `json:"mountpoint,omitnil"`
	Path        *PathResource        `json:"path,omitnil"`
	Queue       *QueueResource       `json:"queue,omitnil"`
	KeyName     *KeyNameResource     `json:"keyname,omitnil"` // kms service 相关
}

type DatabaseResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type TableResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type ColumnResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type GlobalResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type HiveServiceResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type UDFResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type URLResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type BucketResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type MountPointResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type PathResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type QueueResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}

type KeyNameResource struct {
	Values      *[]*string `json:"values,omitnil"`
	IsExcludes  *bool      `json:"isExcludes,omitnil"`
	IsRecursive *bool      `json:"isRecursive,omitnil"`
}
