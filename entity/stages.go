package entity

import "time"

type Stages struct {
	StageID       int       `json:"stage_id" gorm:"type:int;primaryKey"`
	StageName     string    `json:"stage_name" gorm:"type:varchar(20);not null"`
	CompetitionID int       `json:"competition_id"`
	StageOrder    int       `json:"stage_order" gorm:"type:int;not null"`
	Deadline      time.Time `json:"deadline" gorm:"type:date;not null"`

	TeamProgresses TeamProgress `json:"team_progresses" gorm:"foreignKey:StageID;references:StageID"`
}
