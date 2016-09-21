package ctrl

import (
	"github.com/kataras/iris"
	"github.com/jmoiron/sqlx"
	"github.com/cocotyty/summer"
)

type Index struct {
	DB *sqlx.DB `sm:"@.*"`
}

func (i *Index)Index(ctx *iris.Context) {
	ctx.JSON(200, iris.Map{"ctr", "llo"})
}
func init() {
	summer.Put(&Index{})
}