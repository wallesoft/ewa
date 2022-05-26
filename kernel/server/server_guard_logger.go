package server

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

//logger
func (s *ServerGuard) handleAccessLog(ctx context.Context, raw string) {
	if !s.Logger.AccessLogEnabled {
		return
	}
	s.Logger.File(s.Logger.AccessLogPattern).Stdout(s.Logger.LogStdout).Printf(ctx, "[Access]:Request Received-%s:\n Params:%s \n Raw:%s \n Parsed: %s \n", gtime.Datetime(), s.Request.URL.String(), gconv.String(s.bodyData.RawBody), raw)
}
