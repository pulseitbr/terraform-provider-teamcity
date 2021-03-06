package teamcity

import (
	api "github.com/cvbarros/go-teamcity-sdk/pkg/teamcity"
)

// Config Used to configure an api client for TeamCity
type Config struct {
	Address  string
	Username string
	Password string
}

// Client Returns a new TeamCity api client configured with this instance parameters
func (c *Config) Client() (*api.Client, error) {
	return api.New(c.Username, c.Password)
}
