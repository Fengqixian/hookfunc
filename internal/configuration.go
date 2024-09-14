package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/model", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	db, _ := gorm.Open(mysql.Open("root:WZkYigvhY99QRFGB6@tcp(117.72.44.89:13306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)
	g.GenerateModel("order_goods")
	// Execute the generator
	g.Execute()
}
