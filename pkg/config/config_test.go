package config_test

import (
	"catinator-backend/pkg/config"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	assert := require.New(t)

	os.Setenv("SERVER_PORT", ":5000")
	os.Setenv("SERVER_HOST", "localhost")

	c := &config.Config{}

	err := config.New(c, config.FromENV(""), config.FromFile("./test.yaml"))
	assert.Nil(err)

	assert.Equal("localhost", c.Server.Host)
	assert.Equal(":5000", c.Server.Port)

	assert.Equal("test", c.DB.User)
	assert.Equal("test", c.DB.Password)

	t.Logf("%+v", c)
}
