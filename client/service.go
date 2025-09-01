// Package client @author: Violet-Eva @date  : 2025/8/31 @notes :
package client

type ServiceType int

const (
	HiveServiceType ServiceType = iota
	HdfsServiceType
	YarnServiceType
	kmsServiceType
	CosServiceType
	ChdfsServiceType
)

var serviceTypeName = []string{
	"hive",
	"hdfs",
	"cos",
	"yarn",
	"kms",
	"chdfs",
}

func (st ServiceType) String() string {
	if st >= HiveServiceType && st <= ChdfsServiceType {
		return serviceTypeName[st]
	}
	return "unknown service type"
}

type ServiceTypeId struct {
	ServiceType   ServiceType `json:"serviceType"`
	ServiceTypeId int         `json:"serviceTypeId"`
}

type ObjectType int

const (
	HiveService ObjectType = iota
	Url
	GlobalUdf
	Udf
	Database
	Table
	Column
	Masking
	RowFilter
	Hdfs
	Yarn
	Cos
	Chdfs
)

var objectTypeName = []string{
	"HIVE SERVICE",
	"URL",
	"GLOBAL UDF",
	"UDF",
	"DATABASE",
	"TABLE",
	"COLUMN",
	"MASKING",
	"ROW FILTER",
	"HDFS",
	"YARN",
	"COS",
	"CHDFS",
}

func (ot ObjectType) String() string {
	if ot >= HiveService && ot <= Chdfs {
		return objectTypeName[ot]
	}
	return "unknown service type"
}
