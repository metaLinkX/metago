package main

import (
	_ "server/app/backend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/app/backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
