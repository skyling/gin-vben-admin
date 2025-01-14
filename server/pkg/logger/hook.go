package logger

import "github.com/sirupsen/logrus"

type TraceIdHook struct {
	TraceId string
}

func NewTraceIdHook(traceId string) logrus.Hook {
	hook := TraceIdHook{
		TraceId: traceId,
	}
	return &hook
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["trace_id"] = hook.TraceId
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
