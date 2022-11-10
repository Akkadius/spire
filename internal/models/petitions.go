package models

import (
	"github.com/volatiletech/null/v8"
)

type Petition struct {
	Dib          uint        `json:"dib" gorm:"Column:dib"`
	Petid        uint        `json:"petid" gorm:"Column:petid"`
	Charname     string      `json:"charname" gorm:"Column:charname"`
	Accountname  string      `json:"accountname" gorm:"Column:accountname"`
	Lastgm       string      `json:"lastgm" gorm:"Column:lastgm"`
	Petitiontext string      `json:"petitiontext" gorm:"Column:petitiontext"`
	Gmtext       null.String `json:"gmtext" gorm:"Column:gmtext"`
	Zone         string      `json:"zone" gorm:"Column:zone"`
	Urgency      int         `json:"urgency" gorm:"Column:urgency"`
	Charclass    int         `json:"charclass" gorm:"Column:charclass"`
	Charrace     int         `json:"charrace" gorm:"Column:charrace"`
	Charlevel    int         `json:"charlevel" gorm:"Column:charlevel"`
	Checkouts    int         `json:"checkouts" gorm:"Column:checkouts"`
	Unavailables int         `json:"unavailables" gorm:"Column:unavailables"`
	Ischeckedout int8        `json:"ischeckedout" gorm:"Column:ischeckedout"`
	Senttime     int64       `json:"senttime" gorm:"Column:senttime"`
}

func (Petition) TableName() string {
    return "petitions"
}

func (Petition) Relationships() []string {
    return []string{}
}

func (Petition) Connection() string {
    return ""
}
