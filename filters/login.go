package filters

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

const requestStartKey = "request_start_time"

//Save request start time for the finish filter.
func LogStart(ctx *beecontext.Context) {
	ctx.Input.SetData(requestStartKey, time.Now())
}

// Log request method, path and duration
func LogFinish(ctx *beecontext.Context) {
	start, ok := ctx.Input.GetData(requestStartKey).(time.Time)
	if !ok {
		logs.Info("%s %s", ctx.Input.Method(), ctx.Input.URI())
		return
	}
	duration := time.Since(start)
	logs.Info("%s %s - %v", ctx.Input.Method(), ctx.Input.URI(), duration)
}