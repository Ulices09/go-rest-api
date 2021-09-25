package logger

import "go.uber.org/zap/zaptest"

func NewMockLogger(t zaptest.TestingT, opts ...zaptest.LoggerOption) (logger Logger) {
	testLogger := zaptest.NewLogger(t, opts...).Sugar()
	logger.SugaredLogger = testLogger
	return
}
