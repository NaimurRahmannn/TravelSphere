package filters

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

// requestStartKey is where we stash the start time on the context so the
// "finish" filter can read it back to compute duration.
const requestStartKey = "request_start_time"

// LogStart records when a request began. Registered as a BeforeRouter filter
// so it runs at the very start of the request lifecycle.
func LogStart(ctx *beecontext.Context) {
	ctx.Input.SetData(requestStartKey, time.Now())
}

// LogFinish logs the method, URL, and elapsed time. Registered as a FinishRouter
// filter so it runs after the handler completes.
func LogFinish(ctx *beecontext.Context) {
	start, ok := ctx.Input.GetData(requestStartKey).(time.Time)
	if !ok {
		// No start time recorded — log without duration rather than skipping.
		logs.Info("%s %s", ctx.Input.Method(), ctx.Input.URI())
		return
	}
	duration := time.Since(start)
	logs.Info("%s %s - %v", ctx.Input.Method(), ctx.Input.URI(), duration)
}