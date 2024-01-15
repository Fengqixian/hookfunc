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
	db, _ := gorm.Open(mysql.Open("root:woxianghenichigefan@mysql@tcp(47.109.41.242:3306)/hookfun_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)
	g.GenerateModel("user_info")
	// Execute the generator
	g.Execute()
}
