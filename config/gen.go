package config

import (
	"ST/dal"
	"ST/dal/query"
	"gorm.io/gen"
)

func InitGen() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./biz/dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})
	g.UseDB(dal.DB) //DB的初始化需要在UserDB前，不然会报空指针
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()

	query.SetDefault(dal.DB)
}
