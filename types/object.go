// Package types @author: Violet-Eva @date  : 2025/9/3 @notes :
package types

import (
	"strings"
)

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

func GetObjectType(ot string) ObjectType {
	ot = strings.ToUpper(ot)
	for index, element := range objectTypeName {
		if ot == element {
			return ObjectType(index)
		}
	}
	return ObjectType(-1)
}

type Object struct {
	ObjectType       string   `json:"object_type"`
	ObjectName       string   `json:"object_name"`
	ObjectDBName     string   `json:"object_db_name"`
	ObjectTBLName    string   `json:"object_tbl_name"`
	ObjectColumnName []string `json:"object_column_name"`
}
