// Copyright (c) 2015 Cameron King. All rights reserved.
// License: BSD 2-clause.
// Website: https://github.com/ckxng/wakeup

package config

import (
	"fmt"
)

type Config struct {
	Host			string
	Port			int
	EnableServer	bool
	EnableWindow	bool
	Path			string
	Protocol		string
	CachePath		string
	Title			string
}

func NewConfig() (*Config) {
	return &Config{
		Host:			"127.0.0.1",
		Port:			3000,
		EnableServer:	true,
		EnableWindow:	true,
		Path:			"/",
		Protocol:		"http",
		CachePath:		"webcache",
		Title:			"Application",
	}
}

func (cfg *Config) URL() (string) {
	return fmt.Sprintf(
		"%s://%s:%d%s", 
		cfg.Protocol, 
		cfg.Host, 
		cfg.Port, 
		cfg.Path)
}