package models

type Modelable interface {
	Relationships() []string
	TableName() string
	Connection() string
}

type Indexable interface {
	Indexes() map[string][]string
}
