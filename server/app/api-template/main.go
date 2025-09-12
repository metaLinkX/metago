package main

import (
	_ "server/app/api-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/app/api-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
