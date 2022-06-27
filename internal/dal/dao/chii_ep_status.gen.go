// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNameEpCollection = "chii_ep_status"

// EpCollection mapped from table <chii_ep_status>
type EpCollection struct {
	ID        uint32          `gorm:"column:ep_stt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	UserID    model.UserID    `gorm:"column:ep_stt_uid;type:mediumint(8) unsigned;not null;uniqueIndex:ep_stt_uniq,priority:1"`
	SubjectID model.SubjectID `gorm:"column:ep_stt_sid;type:mediumint(8) unsigned;not null;uniqueIndex:ep_stt_uniq,priority:2"`
	OnPrg     bool            `gorm:"column:ep_stt_on_prg;type:tinyint(1) unsigned;not null"`
	Status    bytes           `gorm:"column:ep_stt_status;type:mediumtext;not null"`
	UpdatedAt uint32          `gorm:"column:ep_stt_lasttouch;type:int(10) unsigned;not null"`
}

// TableName EpCollection's table name
func (*EpCollection) TableName() string {
	return TableNameEpCollection
}
