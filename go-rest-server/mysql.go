package main

import "github.com/jinzhu/gorm"

func testMySQL(){
	db, err := gorm.Open("mysql", "root:p$%57p@tcp(localhost:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}
