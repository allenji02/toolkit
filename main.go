package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sony/sonyflake"
	"strconv"
	"time"
)

func main() {
	t, _ := time.Parse(time.RFC3339, "2022-10-29")
	s := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: t,
	})

	h := server.Default(server.WithHostPorts("127.0.0.1:4396"))

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.GET("/uid", func(c context.Context, ctx *app.RequestContext) {
		uid, _ := s.NextID()
		ctx.String(consts.StatusOK, strconv.FormatUint(uid, 10))
	})

	h.Spin()
}
