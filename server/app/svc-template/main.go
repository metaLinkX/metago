package main

import (
	_ "server/app/svc-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/app/svc-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
