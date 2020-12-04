// SPDX-License-Identifier: MIT-0

package vci_template_go

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Config struct {
	writeMu       sync.Mutex
	currentConfig atomic.Value
}

func NewConfig() *Config {
	cfg := &Config{}
	return cfg
}

func (c *Config) Get() map[string]interface{} {
	// Return what your current configuration when requested.
	return c.currentConfig.Load().(map[string]interface{})
}

func (c *Config) Set(new map[string]interface{}) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	// Do some thing with the configuration
	fmt.Println("setting the config to", new)
	c.currentConfig.Store(new)
	return nil
}

func (c *Config) Check(new map[string]interface{}) error {
	// Implement any additional configuration validation you wish here.
	return nil
}

type State struct{}

func NewState() *State {
	return &State{}
}

func (s *State) Get() map[string]interface{} {
	//Do something to retrieve the state here
	return map[string]interface{}{
		"vci-template-go-v1:vci-template-go": map[string]interface{}{
			"state": map[string]interface{}{
				"foo": "bar",
			},
		},
	}
}

type RPC struct{}

func NewRPC() *RPC {
	return &RPC{}
}
func (r *RPC) RPC1(in map[string]interface{}) (map[string]interface{}, error) {
	// Implement your RPC logic here.
	fmt.Println("RPC1 called", in)
	return in, nil
}
