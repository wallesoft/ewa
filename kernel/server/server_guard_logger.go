package server

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

//logger
func (s *ServerGuard) handleAccessLog(raw string) {
	if !s.Logger.AccessLogEnabled {
		return
	}
	s.Logger.File(s.Logger.AccessLogPattern).Stdout(s.Logger.LogStdout).Printf("[Access]:Request Received-%s:\n Params:%s \n Raw:%s \n Parsed: %s \n", gtime.Datetime(), s.Request.URL.String(), gconv.String(s.bodyData.RawBody), raw)
}
