// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameSubjectTopic = "chii_subject_topics"

// SubjectTopic mapped from table <chii_subject_topics>
type SubjectTopic struct {
	ID        uint32 `gorm:"column:sbj_tpc_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:"sbj_tpc_id"`
	SubjectID uint32 `gorm:"column:sbj_tpc_subject_id;type:mediumint(8) unsigned;not null;index:tpc_subject_id,priority:1;index:sbj_tpc_lastpost,priority:2" json:"sbj_tpc_subject_id"`
	UID       uint32 `gorm:"column:sbj_tpc_uid;type:mediumint(8) unsigned;not null;index:sbj_tpc_uid,priority:1" json:"sbj_tpc_uid"`
	Title     string `gorm:"column:sbj_tpc_title;type:varchar(80);not null" json:"sbj_tpc_title"`
	CreatedAt uint32 `gorm:"column:sbj_tpc_dateline;type:int(10) unsigned;not null" json:"sbj_tpc_dateline"`
	UpdatedAt uint32 `gorm:"column:sbj_tpc_lastpost;type:int(10) unsigned;not null;index:sbj_tpc_lastpost,priority:1" json:"sbj_tpc_lastpost"`
	Replies   uint32 `gorm:"column:sbj_tpc_replies;type:mediumint(8) unsigned;not null" json:"sbj_tpc_replies"`
	State     uint8  `gorm:"column:sbj_tpc_state;type:tinyint(1) unsigned;not null" json:"sbj_tpc_state"`
	Status    uint8  `gorm:"column:sbj_tpc_display;type:tinyint(1) unsigned;not null;index:tpc_display,priority:1;index:sbj_tpc_lastpost,priority:3;default:1" json:"sbj_tpc_display"`
}

// TableName SubjectTopic's table name
func (*SubjectTopic) TableName() string {
	return TableNameSubjectTopic
}
