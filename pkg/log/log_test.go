package log_test

import (
	"catinator-backend/pkg/log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLog(t *testing.T) {
	assert := require.New(t)
	logger := log.New(true)
	assert.NotNil(logger)
}
