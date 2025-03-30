package dao

import (
	"errors"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/config"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/db/module"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"sync"
)

type TagDao struct {
	db *gorm.DB
}

type TagDaoImpl interface {
	InsertByTag(tag *module.Tag) error
	GetTidByTag(tag string) (uint64, error)
	GetTagByTid(tid uint64) (string, error)
}

var (
	instanceTagDAO *TagDao
	onceTagDAO     sync.Once
)

func GetTagDao() *TagDao {

	onceTagDAO.Do(func() {
		instanceTagDAO = CreateTagDB()
	})
	return instanceTagDAO
}
func CreateTagDB() *TagDao {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		conf.PostgresSQL.Host,
		conf.PostgresSQL.User,
		conf.PostgresSQL.Password,
		conf.PostgresSQL.DBName,
		conf.PostgresSQL.Port,
		conf.PostgresSQL.Sslmode,
	)
	fmt.Printf(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	err = db.AutoMigrate(&module.Tag{})
	if err != nil {
		panic(fmt.Sprintf("failed to migrate database: %v", err))
	}
	instanceTagDAO = &TagDao{
		db: db,
	}
	return instanceTagDAO
}
func (dao *TagDao) InsertByTag(tag *module.Tag) error {
	return dao.db.Create(tag).Error
}

func (dao *TagDao) GetTidByTag(tag string) (uint64, error) {
	var existingTag module.Tag
	err := dao.db.Where("tag = ?", tag).First(&existingTag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return 0, nil
		}
		return -1, err
	}
	return existingTag.Tid, nil
}

func (dao *TagDao) GetTagByTid(tid uint64) (string, error) {
	var existingTag module.Tag
	err := dao.db.Where("tid = ?", tid).First(&existingTag).Error
	if err != nil {
		return "", err
	}
	return existingTag.Tag, nil
}
