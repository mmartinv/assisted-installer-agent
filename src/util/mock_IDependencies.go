// Code generated by mockery v1.0.0. DO NOT EDIT.

package util

import (
	fs "io/fs"

	block "github.com/jaypipes/ghw/pkg/block"

	gpu "github.com/jaypipes/ghw/pkg/gpu"

	memory "github.com/jaypipes/ghw/pkg/memory"

	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"

	option "github.com/jaypipes/ghw/pkg/option"

	product "github.com/jaypipes/ghw/pkg/product"
)

// MockIDependencies is an autogenerated mock type for the IDependencies type
type MockIDependencies struct {
	mock.Mock
}

// Abs provides a mock function with given fields: path
func (_m *MockIDependencies) Abs(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Block provides a mock function with given fields: opts
func (_m *MockIDependencies) Block(opts ...*option.Option) (*block.Info, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *block.Info
	if rf, ok := ret.Get(0).(func(...*option.Option) *block.Info); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*block.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*option.Option) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvalSymlinks provides a mock function with given fields: path
func (_m *MockIDependencies) EvalSymlinks(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: command, args
func (_m *MockIDependencies) Execute(command string, args ...string) (string, string, int) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, ...string) string); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string, ...string) int); ok {
		r2 = rf(command, args...)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// GPU provides a mock function with given fields: opts
func (_m *MockIDependencies) GPU(opts ...*option.Option) (*gpu.Info, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gpu.Info
	if rf, ok := ret.Get(0).(func(...*option.Option) *gpu.Info); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gpu.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*option.Option) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGhwChrootRoot provides a mock function with given fields:
func (_m *MockIDependencies) GetGhwChrootRoot() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Hostname provides a mock function with given fields:
func (_m *MockIDependencies) Hostname() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interfaces provides a mock function with given fields:
func (_m *MockIDependencies) Interfaces() ([]Interface, error) {
	ret := _m.Called()

	var r0 []Interface
	if rf, ok := ret.Get(0).(func() []Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkByName provides a mock function with given fields: name
func (_m *MockIDependencies) LinkByName(name string) (netlink.Link, error) {
	ret := _m.Called(name)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Memory provides a mock function with given fields: opts
func (_m *MockIDependencies) Memory(opts ...*option.Option) (*memory.Info, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *memory.Info
	if rf, ok := ret.Get(0).(func(...*option.Option) *memory.Info); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*memory.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*option.Option) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Product provides a mock function with given fields: opts
func (_m *MockIDependencies) Product(opts ...*option.Option) (*product.Info, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *product.Info
	if rf, ok := ret.Get(0).(func(...*option.Option) *product.Info); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Info)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*option.Option) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDir provides a mock function with given fields: dirname
func (_m *MockIDependencies) ReadDir(dirname string) ([]fs.FileInfo, error) {
	ret := _m.Called(dirname)

	var r0 []fs.FileInfo
	if rf, ok := ret.Get(0).(func(string) []fs.FileInfo); ok {
		r0 = rf(dirname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]fs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dirname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFile provides a mock function with given fields: fname
func (_m *MockIDependencies) ReadFile(fname string) ([]byte, error) {
	ret := _m.Called(fname)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(fname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RouteList provides a mock function with given fields: link, family
func (_m *MockIDependencies) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.Called(link, family)

	var r0 []netlink.Route
	if rf, ok := ret.Get(0).(func(netlink.Link, int) []netlink.Route); ok {
		r0 = rf(link, family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(netlink.Link, int) error); ok {
		r1 = rf(link, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields: fname
func (_m *MockIDependencies) Stat(fname string) (fs.FileInfo, error) {
	ret := _m.Called(fname)

	var r0 fs.FileInfo
	if rf, ok := ret.Get(0).(func(string) fs.FileInfo); ok {
		r0 = rf(fname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fs.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
