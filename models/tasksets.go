package models

type Taskset struct {
	ID     uint `json:"id" gorm:"Column:id"`
	Taskid uint `json:"taskid" gorm:"Column:taskid"`
}

func (Taskset) TableName() string {
    return "tasksets"
}

func (Taskset) Relationships() []string {
    return []string{}
}

func (Taskset) Connection() string {
    return "eqemu_content"
}
