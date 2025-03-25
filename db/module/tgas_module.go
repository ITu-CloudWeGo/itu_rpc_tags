package module

type Tags struct {
	Pid  uint64 `gorm:"column:pid;primaryKey" json:"tags_id"`
	Tags string `gorm:"column:tags" json:"tags"`
}

type Tabler interface {
	TableName() string
}

func (t *Tags) TableName() string {
	return "tags"
}
