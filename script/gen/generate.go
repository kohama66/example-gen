package main

import (
	"api/infrastructure/db"

	"gorm.io/gen"
)

func main() {
	conn := db.New()
	g := gen.NewGenerator(gen.Config{
		OutPath: "infrastructure/db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(conn.DB) // reuse your gorm db

	all := g.GenerateAllTable()
	g.ApplyBasic(all...)

	g.Execute()
}
