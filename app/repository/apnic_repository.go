package repository

import (
	"docker_go_test/app/entity"
	"docker_go_test/app/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ApnicRepository interface {
	WhoisIp(inetNum string) (data entity.ApnicInetnum)
	InsertInetNum(inetNum entity.ApnicInetnum) error
	CreateTableMigration()
}

type ApnicRepositoryImpl struct {
	Db *gorm.DB
}

func (repo ApnicRepositoryImpl) CreateTableMigration() {
	if err := repo.Db.Set("gorm:table_options", "COLLATE=utf8_general_ci").Migrator().CreateTable(&entity.ApnicInetnum{}); err != nil {
		helper.WriteLog("migration.log", "ERROR "+err.Error())
	}
}

func (repo ApnicRepositoryImpl) InsertInetNum(inetNum entity.ApnicInetnum) error {
	err := repo.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "inetnum"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "descr"}),
	}).Create(&inetNum).Error

	return err
}

func (repo ApnicRepositoryImpl) WhoisIp(inetNum string) (data entity.ApnicInetnum) {
	repo.Db.Where("inetnum=?", inetNum).Find(&data)
	return data
}

func NewApnicRepository(db *gorm.DB) ApnicRepository {
	return ApnicRepositoryImpl{
		Db: db,
	}
}
