package models

type Modelable interface {
	Relationships() []string
	TableName() string
	Connection() string
}
