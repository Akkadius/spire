package model

import (
	"strings"
)

// translateDataType translates a MySQL data type to a Go data type.
// It handles both nullable and non-nullable types.
// It also considers unsigned types for integer types.
func translateDataType(column DbSchemaRowResult) string {
	baseType := strings.ToLower(strings.Split(column.DataType, "(")[0])
	unsigned := strings.Contains(column.DataType, "unsigned")
	nullable := strings.Contains(column.IsNullable, "YES")

	var goType string

	switch baseType {
	case "tinyint":
		if unsigned {
			goType = "uint8"
		} else {
			goType = "int8"
		}
	case "smallint":
		if unsigned {
			goType = "uint16"
		} else {
			goType = "int16"
		}
	case "mediumint":
		if unsigned {
			goType = "uint32"
		} else {
			goType = "int32"
		}
	case "int", "integer":
		if unsigned {
			goType = "uint"
		} else {
			goType = "int"
		}
	case "bigint":
		if unsigned {
			goType = "uint64"
		} else {
			goType = "int64"
		}
	case "float":
		goType = "float32"
	case "double", "double precision", "real":
		goType = "float64"
	case "boolean", "bool":
		goType = "bool"
	case "date", "datetime", "timestamp":
		goType = "time.Time"
	case "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		goType = "[]byte"
	case "numeric", "decimal", "dec", "fixed":
		goType = "float32"
	case "json":
		goType = "types.JSON"
	default:
		goType = "string"
	}

	if nullable {
		switch goType {
		case "bool":
			return "null.Bool"
		case "uint8":
			return "null.Uint8"
		case "int8":
			return "null.Int8"
		case "uint16":
			return "null.Uint16"
		case "int16":
			return "null.Int16"
		case "uint32":
			return "null.Uint32"
		case "int32":
			return "null.Int32"
		case "uint":
			return "null.Uint"
		case "int":
			return "null.Int"
		case "uint64":
			return "null.Uint64"
		case "int64":
			return "null.Int64"
		case "float32":
			return "null.Float32"
		case "float64":
			return "null.Float64"
		case "time.Time":
			return "null.Time"
		case "[]byte":
			return "null.Bytes"
		case "types.JSON":
			return "null.JSON"
		default:
			return "null.String"
		}
	}

	return goType
}
