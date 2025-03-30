package module

type PidTid struct {
	Pid uint64 `gorm:"column:pid;primary_key" json:"Pid"`
	Tid uint64 `gorm:"column:tid;not null" json:"Tid"`
}

type PTTabler interface {
	PTTableName() string
}

func (t *Tag) PTTableName() string {
	return "pid_tid"
}
