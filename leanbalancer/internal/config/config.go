package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port        int `yaml:"port"`
		MetricsPort int `yaml:"metrics_port"`
	} `yaml:"server"`
	LoadBalancer struct {
		Algorithm string `yaml:"algorithm"`
		Timeout   string `yaml:"timeout"`
	} `yaml:"load_balancer"`
	Firewall struct {
		Enabled    bool     `yaml:"enabled"`
		BlockedIPs []string `yaml:"blocked_ips"`
	} `yaml:"firewall"`
	HealthCheck struct {
		Enabled  bool     `yaml:"enabled"`
		Interval string   `yaml:"interval"`
		Backends []string `yaml:"backends"`
	} `yaml:"health_check"`
}

func LoadConfig(filePath string) (*Config, error) {
	// Read YAML file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}
