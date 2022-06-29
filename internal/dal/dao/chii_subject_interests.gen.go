// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNameSubjectCollection = "chii_subject_interests"

// SubjectCollection mapped from table <chii_subject_interests>
type SubjectCollection struct {
	ID              uint32          `gorm:"column:interest_id;type:int(10) unsigned;primaryKey;autoIncrement:true"`
	UserID          model.UserID    `gorm:"column:interest_uid;type:mediumint(8) unsigned;not null;uniqueIndex:user_interest,priority:1;index:interest_uid,priority:1;index:tag_subject_id,priority:3;index:user_collects,priority:2;index:user_collect_type,priority:3;index:interest_id,priority:1;index:user_collect_latest,priority:3;index:interest_type_2,priority:2;index:interest_uid_2,priority:1"`
	SubjectID       model.SubjectID `gorm:"column:interest_subject_id;type:mediumint(8) unsigned;not null;uniqueIndex:user_interest,priority:2;index:interest_subject_id,priority:1;index:interest_subject_id_2,priority:1;index:subject_lasttouch,priority:1;index:subject_comment,priority:1;index:subject_collect,priority:1;index:subject_rate,priority:1;index:top_subject,priority:1"`
	SubjectType     uint8           `gorm:"column:interest_subject_type;type:smallint(6) unsigned;not null;index:interest_subject_type,priority:1;index:tag_subject_id,priority:1;index:user_collects,priority:1;index:user_collect_type,priority:1;index:top_subject,priority:2;index:user_collect_latest,priority:1"`
	Rate            uint8           `gorm:"column:interest_rate;type:tinyint(3) unsigned;not null;index:interest_rate,priority:1;index:subject_rate,priority:2"`
	Type            uint8           `gorm:"column:interest_type;type:tinyint(1) unsigned;not null;index:interest_subject_id,priority:2;index:interest_type,priority:1;index:tag_subject_id,priority:2;index:subject_collect,priority:2;index:user_collect_type,priority:2;index:user_collect_latest,priority:2;index:interest_type_2,priority:1"`
	HasComment      bool            `gorm:"column:interest_has_comment;type:tinyint(1) unsigned;not null;index:subject_comment,priority:2"`
	Comment         string          `gorm:"column:interest_comment;type:mediumtext;not null"`
	Tag             string          `gorm:"column:interest_tag;type:mediumtext;not null"`
	EpStatus        uint32          `gorm:"column:interest_ep_status;type:mediumint(8) unsigned;not null"`
	VolStatus       uint32          `gorm:"column:interest_vol_status;type:mediumint(8) unsigned;not null"` // 卷数
	WishDateline    uint32          `gorm:"column:interest_wish_dateline;type:int(10) unsigned;not null"`
	DoingDateline   uint32          `gorm:"column:interest_doing_dateline;type:int(10) unsigned;not null;index:top_subject,priority:3"`
	CollectDateline uint32          `gorm:"column:interest_collect_dateline;type:int(10) unsigned;not null;index:interest_collect_dateline,priority:1;index:subject_collect,priority:4;index:user_collect_type,priority:5"`
	OnHoldDateline  uint32          `gorm:"column:interest_on_hold_dateline;type:int(10) unsigned;not null"`
	DroppedDateline uint32          `gorm:"column:interest_dropped_dateline;type:int(10) unsigned;not null"`
	UpdatedTime     uint32          `gorm:"column:interest_lasttouch;type:int(10) unsigned;not null;index:interest_lasttouch,priority:1;index:subject_lasttouch,priority:3;index:subject_comment,priority:4;index:interest_uid_2,priority:3"`
	Private         uint8           `gorm:"column:interest_private;type:tinyint(1) unsigned;not null;index:interest_private,priority:1;index:subject_lasttouch,priority:2;index:subject_comment,priority:3;index:subject_collect,priority:3;index:user_collect_type,priority:4;index:interest_id,priority:2;index:subject_rate,priority:3;index:user_collect_latest,priority:4;index:interest_uid_2,priority:2"`
}

// TableName SubjectCollection's table name
func (*SubjectCollection) TableName() string {
	return TableNameSubjectCollection
}
