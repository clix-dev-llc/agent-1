// Package memcached_exporter embeds https://github.com/google/memcached_exporter
package memcached_exporter //nolint:golint

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/grafana/agent/pkg/integrations"
	"github.com/grafana/agent/pkg/integrations/config"
	"github.com/prometheus/memcached_exporter/pkg/exporter"
)

// DefaultConfig is the default config for memcached_exporter.
var DefaultConfig Config = Config{
	MemcachedAddress: "localhost:11211",
	Timeout:          time.Second,
}

// Config controls the memcached_exporter integration.
type Config struct {
	Common config.Common `yaml:",inline"`

	// MemcachedAddress is the address of the memcached server (host:port).
	MemcachedAddress string `yaml:"memcached_address"`

	// Timeout is the connection timeout for memcached.
	Timeout time.Duration `yaml:"timeout"`
}

// UnmarshalYAML implements yaml.Unmarshaler for Config.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig

	type plain Config
	return unmarshal((*plain)(c))
}

func (c *Config) Name() string {
	return "memcached_exporter"
}

func (c *Config) CommonConfig() config.Common {
	return c.Common
}

func (c *Config) NewIntegration(l log.Logger) (integrations.Integration, error) {
	return New(l, c)
}

func init() {
	integrations.RegisterIntegration(&Config{})
}

// New creates a new memcached_exporter integration. The integration scrapes metrics
// from a memcached server.
func New(log log.Logger, c *Config) (integrations.Integration, error) {
	return integrations.NewCollectorIntegration(
		c.Name(),
		integrations.WithCollectors(
			exporter.New(c.MemcachedAddress, c.Timeout, log),
		),
	), nil
}
