package jwt

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gutil"
	"testing"
)

func TestName(t *testing.T) {
	ctx := gctx.New()

	token, err := IssueUserIdToken(123460, g.Cfg().MustGet(ctx, "jwt.signingKey").String(), g.Cfg().MustGet(ctx, "jwt.signExpireDay").Duration())
	if err != nil {
		return
	}

	//#doctdddzyys888xzj!
	//#doctdddzyys888xzj!

	g.Log().Info(ctx, g.Cfg().MustGet(ctx, "jwt.signingKey").String(), token)
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEyMzQ2MSwiZXhwIjoyMDM4NTI4MDY2LCJuYmYiOjE3MjY2MjQwNjYsImlhdCI6MTcyNjYyNDA2Nn0.PRoFhS7umdgoxRLgf3Vxo8BXgaIVxCjKLnfpgQrL4dE"
	uid, err := ParseUserId(token, g.Cfg().MustGet(ctx, "jwt.signingKey").String())
	gutil.Dump(uid, err)
}

func TestGenToken(t *testing.T) {
	ctx := gctx.New()
	token, err := IssueUserIdToken(123461, g.Cfg().MustGet(ctx, "jwt.signingKey").String(), 10*361)
	if err != nil {
		return
	}
	g.Log().Info(ctx, token)
}
