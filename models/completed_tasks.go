package models

type CompletedTask struct {
	Charid        uint `json:"charid" gorm:"Column:charid"`
	Completedtime uint `json:"completedtime" gorm:"Column:completedtime"`
	Taskid        uint `json:"taskid" gorm:"Column:taskid"`
	Activityid    int  `json:"activityid" gorm:"Column:activityid"`
}

func (CompletedTask) TableName() string {
    return "completed_tasks"
}

func (CompletedTask) Relationships() []string {
    return []string{}
}

func (CompletedTask) Connection() string {
    return ""
}
