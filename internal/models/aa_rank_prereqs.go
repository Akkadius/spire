package models

type AaRankPrereq struct {
	RankId  uint `json:"rank_id" gorm:"Column:rank_id"`
	AaId    int  `json:"aa_id" gorm:"Column:aa_id"`
	Points  int  `json:"points" gorm:"Column:points"`
}

func (AaRankPrereq) TableName() string {
    return "aa_rank_prereqs"
}

func (AaRankPrereq) Relationships() []string {
    return []string{}
}

func (AaRankPrereq) Connection() string {
    return "eqemu_content"
}
