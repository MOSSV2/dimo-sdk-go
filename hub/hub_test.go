package hub

import (
	"testing"

	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code   string
	Price  uint
	Labels datatypes.JSON
}

func TestGORM(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("/tmp/test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	//datatypes.JSONSet("labels").Set("age", 20)

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	t.Log(product.ID, product.Code, product.Price)

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200).UpdateColumn("labels", datatypes.JSONSet("labels").Set("age", 20))
	db.Model(&product).UpdateColumn("labels", datatypes.JSONSet("labels").Set("sex", "male"))

	t.Log(product.ID, product.Code, product.Price, product.Labels)
	t.Log(product)

	res := datatypes.JSONQuery("labels").Extract("age")
	t.Log(res)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 300, Code: "F42"}) // 仅更新非零值字段
	t.Log(product.ID, product.Code, product.Price)
	db.Model(&product).Updates(map[string]interface{}{"Price": 400, "Code": "E42"})
	t.Log(product.ID, product.Code, product.Price)

	// Delete - 删除 product
	db.Delete(&product, 1)

	t.Log(product.ID, product.Code, product.Price)

	t.Fatal("")
}
