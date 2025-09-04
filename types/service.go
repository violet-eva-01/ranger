package types

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
