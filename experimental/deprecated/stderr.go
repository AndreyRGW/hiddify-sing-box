package deprecated

import "github.com/sagernet/sing/common/logger"

type stderrManager struct {
	logger   logger.Logger
	reported map[string]bool
}

func NewStderrManager(logger logger.Logger) Manager {
	return &stderrManager{
		logger:   logger,
		reported: make(map[string]bool),
	}
}

func (f *stderrManager) ReportDeprecated(feature Note) {
	// no-op: skip deprecated warnings to avoid concurrent map writes
}
