package dao

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/config"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/module"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sync"
)

type TagsDao struct {
	db *gorm.DB
}

var (
	instanceDAO *TagsDao
	onceDAO     sync.Once
)

func GetTagsDAO() *TagsDao {

	onceDAO.Do(func() {
		instanceDAO = CreateDB()
	})
	return instanceDAO
}
func CreateDB() *TagsDao {
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
	err = db.AutoMigrate(&module.Tags{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	instanceDAO = &TagsDao{
		db: db,
	}
	return instanceDAO
}
func (dao *TagsDao) Insert(tags *module.Tags) error {
	return dao.db.Create(tags).Error
}

func (dao *TagsDao) DelTags(pid uint64, tag string) error {
	tags := &module.Tags{
		Pid:  pid,
		Tags: tag,
	}
	return dao.db.Where("pid = ? AND tags =?", pid, tags).Delete(tags).Error
}

func (dao *TagsDao) CheckTags(pid uint64, tag string) (bool, error) {

	exist := dao.db.Where("pid = ? AND tags =?", pid, tag)
	if exist != nil {
		return false, fmt.Errorf("重复tags")
	}
	return true, nil
}
