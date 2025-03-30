package dao

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/config"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/module"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type TidPidDao struct {
	db *gorm.DB
}

type PidTidDaoImpl interface {
	InsertByPidTid(pidTid *module.PidTid) error
	GetTidByPid(pid int64) (int64, error)
}

var (
	instancePidTidDAO *TidPidDao
	onceTidPidDAO     sync.Once
)

func GetTidPidDao() *TidPidDao {
	onceTidPidDAO.Do(func() {
		instancePidTidDAO = CreatePidTidDB()
	})
	return instancePidTidDAO
}

func CreatePidTidDB() *TidPidDao {
	conf := config.Config{}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		conf.PostgresSQL.Host,
		conf.PostgresSQL.User,
		conf.PostgresSQL.Password,
		conf.PostgresSQL.DBName,
		conf.PostgresSQL.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	err = db.AutoMigrate(&module.PidTid{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	instancePidTidDAO = &TidPidDao{
		db: db,
	}
	return instancePidTidDAO
}

func (dao *TidPidDao) InsertByPidTid(pidTid *module.PidTid) error {
	return dao.db.Create(pidTid).Error
}

func (dao *TidPidDao) GetTidByPid(pid int64) ([]int64, error) {
	var pidTids []module.PidTid
	err := dao.db.Where("pid = ?", pid).Find(&pidTids).Error
	if err != nil {
		return nil, err
	}
	tids := make([]int64, 0, len(pidTids))
	for _, pidTid := range pidTids {
		tids = append(tids, pidTid.Tid)
	}
	return tids, nil
}
