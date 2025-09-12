package main

import (
	_ "server/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
