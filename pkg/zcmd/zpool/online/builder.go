/*
Copyright 2019 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ponline

import (
	"fmt"
	"os/exec"
	"reflect"
	"runtime"
	"strings"

	"github.com/openebs/cstor-operators/pkg/zcmd/bin"
	"github.com/pkg/errors"
)

const (
	// Operation defines type of zfs operation
	Operation = "online"
)

//PoolOnline defines structure for pool 'Online' operation
type PoolOnline struct {
	//pool name
	Pool string

	//device list
	Device []string

	//expand the device to use all available space
	ShouldExpand bool

	// command string
	Command string

	// checks is list of predicate function used for validating object
	checks []PredicateFunc

	// error
	err error
}

// NewPoolOnline returns new instance of object PoolOnline
func NewPoolOnline() *PoolOnline {
	return &PoolOnline{}
}

// WithCheck add given check to checks list
func (p *PoolOnline) WithCheck(check ...PredicateFunc) *PoolOnline {
	p.checks = append(p.checks, check...)
	return p
}

// WithPool method fills the Pool field of PoolOnline object.
func (p *PoolOnline) WithPool(Pool string) *PoolOnline {
	p.Pool = Pool
	return p
}

// WithDevice method fills the Device field of PoolOnline object.
func (p *PoolOnline) WithDevice(dev string) *PoolOnline {
	p.Device = append(p.Device, dev)
	return p
}

// WithShouldExpand method fills the ShouldExpand field of PoolOnline object.
func (p *PoolOnline) WithShouldExpand(ShouldExpand bool) *PoolOnline {
	p.ShouldExpand = ShouldExpand
	return p
}

// WithCommand method fills the Command field of PoolOnline object.
func (p *PoolOnline) WithCommand(Command string) *PoolOnline {
	p.Command = Command
	return p
}

// Validate is to validate generated PoolOnline object by builder
func (p *PoolOnline) Validate() *PoolOnline {
	for _, check := range p.checks {
		if !check(p) {
			p.err = errors.Wrapf(p.err, "validation failed {%v}", runtime.FuncForPC(reflect.ValueOf(check).Pointer()).Name())
		}
	}
	return p
}

// Execute is to execute generated PoolOnline object
func (p *PoolOnline) Execute() ([]byte, error) {
	p, err := p.Build()
	if err != nil {
		return nil, err
	}
	// execute command here
	// #nosec
	return exec.Command(bin.BASH, "-c", p.Command).CombinedOutput()
}

// Build returns the PoolOnline object generated by builder
func (p *PoolOnline) Build() (*PoolOnline, error) {
	var c strings.Builder
	p = p.Validate()
	p.appendCommand(&c, bin.ZPOOL)
	p.appendCommand(&c, fmt.Sprintf(" %s ", Operation))

	if IsShouldExpandSet()(p) {
		p.appendCommand(&c, fmt.Sprintf(" -e "))
	}

	p.appendCommand(&c, fmt.Sprintf(" %s ", p.Pool))

	if IsDeviceSet()(p) {
		for _, d := range p.Device {
			p.appendCommand(&c, fmt.Sprintf(" %s ", d))
		}
	}
	p.Command = c.String()
	return p, p.err
}

// appendCommand append string to given string builder
func (p *PoolOnline) appendCommand(c *strings.Builder, cmd string) {
	_, err := c.WriteString(cmd)
	if err != nil {
		p.err = errors.Wrapf(p.err, "Failed to append cmd{%s} : %s", cmd, err.Error())
	}
}
