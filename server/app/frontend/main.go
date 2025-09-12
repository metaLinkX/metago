package main

import (
	_ "server/app/frontend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/app/frontend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
