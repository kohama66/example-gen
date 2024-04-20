package db

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func generate(conn *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "infrastructure/db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(conn) // reuse your gorm db

	all := g.GenerateAllTable()
	g.ApplyBasic(all...)

	g.Execute()
}
