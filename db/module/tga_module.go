package module

type Tag struct {
	Tid int64  `gorm:"column:pid;primaryKey;autoIncrement" json:"Pid"`
	Tag string `gorm:"column:tags" json:"tags"`
}

type TagTable interface {
	TagTableName() string
}

func (t *Tag) TagTableName() string {
	return "tag"
}
