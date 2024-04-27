package main

import (
	"api/infrastructure/db"
	"api/infrastructure/db/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	conn := db.New()
	g := gen.NewGenerator(gen.Config{
		OutPath: "infrastructure/db/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(conn.DB) // reuse your gorm db

	all := g.GenerateAllTable()
	relation := []interface{}{
		g.GenerateModel(model.TableNameUser, gen.FieldRelateModel(field.HasMany, "CreditCards", model.CreditCard{}, &field.RelateConfig{RelateSlicePointer: true})),
	}

	g.ApplyBasic(all...)
	g.ApplyBasic(relation...)

	g.Execute()
}
