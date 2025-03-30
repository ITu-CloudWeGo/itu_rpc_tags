package module

type Tag struct {
	Tid int64  `gorm:"column:pid;primaryKey;autoIncrement" json:"Pid"`
	Tag string `gorm:"column:tags" json:"tags"`
}

type Tabler interface {
	TableName() string
}

func (t *Tag) TableName() string {
	return "tag"
}
