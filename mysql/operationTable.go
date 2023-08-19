package mysql

import "github.com/jinzhu/gorm"

// CreateTable 创建表
func CreateTable(db *gorm.DB, tableName string, model interface{}) {
	if !db.HasTable(tableName) {
		db.CreateTable(model)
	}
}

// Insert 插入一条记录
func Insert(db *gorm.DB, model interface{}) error {
	return db.Create(model).Error
}

// FindById 根据id查询，并返回结果
func FindById(db *gorm.DB, model interface{}, id int) error {
	return db.First(model, id).Error
}

// FindByWhere 根据某个条件查询多条记录
func FindByWhere(db *gorm.DB, model interface{}, where ...interface{}) error {
	return db.Where(where).Find(model).Error
}

// FindOneByWhere 根据某个条件查询一条记录
func FindOneByWhere(db *gorm.DB, model interface{}, where ...interface{}) error {
	return db.Where(where).First(model).Error
}

// FindByWhereWithOrder 根据某个条件查询多条记录，并按照某个字段排序
func FindByWhereWithOrder(db *gorm.DB, model interface{}, order string, where ...interface{}) error {
	return db.Where(where).Order(order).Find(model).Error
}

// UpdateByWhere 根据条件更新多个字段
func UpdateByWhere(db *gorm.DB, model interface{}, where ...interface{}) error {
	return db.Model(model).Where(where).Update(model).Error
}

// DeleteByWhere 删除一条记录
func DeleteByWhere(db *gorm.DB, model interface{}, where ...interface{}) error {
	return db.Where(where).Delete(model).Error
}
