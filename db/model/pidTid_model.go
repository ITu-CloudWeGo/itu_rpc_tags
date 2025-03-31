package model

type PidTid struct {
	Pid int64 `gorm:"column:pid;primary_key" json:"Pid"`
	Tid int64 `gorm:"column:tid;not null" json:"Tid"`
}

type PTTable interface {
	PTTableName() string
}

func (t *Tag) PTTableName() string {
	return "pid_tid"
}
