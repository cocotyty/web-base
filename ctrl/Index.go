package ctrl

import (
	"github.com/cocotyty/summer"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/context"
)

type Index struct {
	DB *sqlx.DB `sm:"@.*"`
}

func (i *Index) Index(ctx context.Context) {
	ctx.JSON(context.Map{"ctr", "llo"})
}

func init() {
	summer.Put(&Index{})
}
