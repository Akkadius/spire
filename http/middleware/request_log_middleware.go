package middleware

import (
	"fmt"
	"github.com/Akkadius/spire/internal/influx"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type RequestLogMiddleware struct {
	influx *influx.Client
}

func NewRequestLogMiddleware(
	influx *influx.Client,
) *RequestLogMiddleware {
	return &RequestLogMiddleware{
		influx: influx,
	}
}

func (m RequestLogMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			p := req.URL.Path
			bytesIn := req.Header.Get(echo.HeaderContentLength)

			if m.influx.Alive() {
				p := influxdb2.NewPointWithMeasurement("spire_request_log").
					AddTag("ip_address", c.RealIP()).
					AddTag("host", req.Host).
					AddTag("uri", req.RequestURI).
					AddTag("method", req.Method).
					AddTag("path", p).
					AddTag("user_agent", req.UserAgent()).
					AddTag("status", fmt.Sprintf("%v", res.Status)).
					AddField("latency", int64(stop.Sub(start))).
					AddField("bytes_in", bytesIn).
					AddField("bytes_out", strconv.FormatInt(res.Size, 10)).
					AddField("response_code", c.Response().Status).
					SetTime(time.Now())
				m.influx.Writer().WritePoint(p)
				m.influx.Writer().Flush()
			}

			return nil
		}
	}
}
