package main

import (
	_ "server/internal/packed"
	"server/library/hggen"
	"server/library/system"

	_ "server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"server/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	system.InitServerRoot(ctx)
	system.InitGFMode(ctx)
	hggen.InIt(ctx)

	cmd.Main.Run(ctx)
}
