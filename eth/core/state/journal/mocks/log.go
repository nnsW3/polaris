// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	journal "github.com/berachain/polaris/eth/core/state/journal"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Log is an autogenerated mock type for the Log type
type Log struct {
	mock.Mock
}

type Log_Expecter struct {
	mock *mock.Mock
}

func (_m *Log) EXPECT() *Log_Expecter {
	return &Log_Expecter{mock: &_m.Mock}
}

// AddLog provides a mock function with given fields: _a0
func (_m *Log) AddLog(_a0 *types.Log) {
	_m.Called(_a0)
}

// Log_AddLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddLog'
type Log_AddLog_Call struct {
	*mock.Call
}

// AddLog is a helper method to define mock.On call
//   - _a0 *types.Log
func (_e *Log_Expecter) AddLog(_a0 interface{}) *Log_AddLog_Call {
	return &Log_AddLog_Call{Call: _e.mock.On("AddLog", _a0)}
}

func (_c *Log_AddLog_Call) Run(run func(_a0 *types.Log)) *Log_AddLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Log))
	})
	return _c
}

func (_c *Log_AddLog_Call) Return() *Log_AddLog_Call {
	_c.Call.Return()
	return _c
}

func (_c *Log_AddLog_Call) RunAndReturn(run func(*types.Log)) *Log_AddLog_Call {
	_c.Call.Return(run)
	return _c
}

// Clone provides a mock function with given fields:
func (_m *Log) Clone() journal.Log {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clone")
	}

	var r0 journal.Log
	if rf, ok := ret.Get(0).(func() journal.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(journal.Log)
		}
	}

	return r0
}

// Log_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type Log_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *Log_Expecter) Clone() *Log_Clone_Call {
	return &Log_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *Log_Clone_Call) Run(run func()) *Log_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_Clone_Call) Return(_a0 journal.Log) *Log_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_Clone_Call) RunAndReturn(run func() journal.Log) *Log_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *Log) Finalize() {
	_m.Called()
}

// Log_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type Log_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *Log_Expecter) Finalize() *Log_Finalize_Call {
	return &Log_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *Log_Finalize_Call) Run(run func()) *Log_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_Finalize_Call) Return() *Log_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *Log_Finalize_Call) RunAndReturn(run func()) *Log_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// GetLogs provides a mock function with given fields: hash, blockNumber, blockHash
func (_m *Log) GetLogs(hash common.Hash, blockNumber uint64, blockHash common.Hash) []*types.Log {
	ret := _m.Called(hash, blockNumber, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetLogs")
	}

	var r0 []*types.Log
	if rf, ok := ret.Get(0).(func(common.Hash, uint64, common.Hash) []*types.Log); ok {
		r0 = rf(hash, blockNumber, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Log)
		}
	}

	return r0
}

// Log_GetLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogs'
type Log_GetLogs_Call struct {
	*mock.Call
}

// GetLogs is a helper method to define mock.On call
//   - hash common.Hash
//   - blockNumber uint64
//   - blockHash common.Hash
func (_e *Log_Expecter) GetLogs(hash interface{}, blockNumber interface{}, blockHash interface{}) *Log_GetLogs_Call {
	return &Log_GetLogs_Call{Call: _e.mock.On("GetLogs", hash, blockNumber, blockHash)}
}

func (_c *Log_GetLogs_Call) Run(run func(hash common.Hash, blockNumber uint64, blockHash common.Hash)) *Log_GetLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Hash), args[1].(uint64), args[2].(common.Hash))
	})
	return _c
}

func (_c *Log_GetLogs_Call) Return(_a0 []*types.Log) *Log_GetLogs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_GetLogs_Call) RunAndReturn(run func(common.Hash, uint64, common.Hash) []*types.Log) *Log_GetLogs_Call {
	_c.Call.Return(run)
	return _c
}

// Logs provides a mock function with given fields:
func (_m *Log) Logs() []*types.Log {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Logs")
	}

	var r0 []*types.Log
	if rf, ok := ret.Get(0).(func() []*types.Log); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Log)
		}
	}

	return r0
}

// Log_Logs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Logs'
type Log_Logs_Call struct {
	*mock.Call
}

// Logs is a helper method to define mock.On call
func (_e *Log_Expecter) Logs() *Log_Logs_Call {
	return &Log_Logs_Call{Call: _e.mock.On("Logs")}
}

func (_c *Log_Logs_Call) Run(run func()) *Log_Logs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_Logs_Call) Return(_a0 []*types.Log) *Log_Logs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_Logs_Call) RunAndReturn(run func() []*types.Log) *Log_Logs_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *Log) RegistryKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RegistryKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Log_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type Log_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *Log_Expecter) RegistryKey() *Log_RegistryKey_Call {
	return &Log_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *Log_RegistryKey_Call) Run(run func()) *Log_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_RegistryKey_Call) Return(_a0 string) *Log_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_RegistryKey_Call) RunAndReturn(run func() string) *Log_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *Log) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// Log_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type Log_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *Log_Expecter) RevertToSnapshot(_a0 interface{}) *Log_RevertToSnapshot_Call {
	return &Log_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *Log_RevertToSnapshot_Call) Run(run func(_a0 int)) *Log_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Log_RevertToSnapshot_Call) Return() *Log_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *Log_RevertToSnapshot_Call) RunAndReturn(run func(int)) *Log_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SetTxContext provides a mock function with given fields: thash, ti
func (_m *Log) SetTxContext(thash common.Hash, ti int) {
	_m.Called(thash, ti)
}

// Log_SetTxContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTxContext'
type Log_SetTxContext_Call struct {
	*mock.Call
}

// SetTxContext is a helper method to define mock.On call
//   - thash common.Hash
//   - ti int
func (_e *Log_Expecter) SetTxContext(thash interface{}, ti interface{}) *Log_SetTxContext_Call {
	return &Log_SetTxContext_Call{Call: _e.mock.On("SetTxContext", thash, ti)}
}

func (_c *Log_SetTxContext_Call) Run(run func(thash common.Hash, ti int)) *Log_SetTxContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Hash), args[1].(int))
	})
	return _c
}

func (_c *Log_SetTxContext_Call) Return() *Log_SetTxContext_Call {
	_c.Call.Return()
	return _c
}

func (_c *Log_SetTxContext_Call) RunAndReturn(run func(common.Hash, int)) *Log_SetTxContext_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *Log) Snapshot() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Snapshot")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Log_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type Log_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *Log_Expecter) Snapshot() *Log_Snapshot_Call {
	return &Log_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *Log_Snapshot_Call) Run(run func()) *Log_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_Snapshot_Call) Return(_a0 int) *Log_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_Snapshot_Call) RunAndReturn(run func() int) *Log_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// TxIndex provides a mock function with given fields:
func (_m *Log) TxIndex() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxIndex")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Log_TxIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxIndex'
type Log_TxIndex_Call struct {
	*mock.Call
}

// TxIndex is a helper method to define mock.On call
func (_e *Log_Expecter) TxIndex() *Log_TxIndex_Call {
	return &Log_TxIndex_Call{Call: _e.mock.On("TxIndex")}
}

func (_c *Log_TxIndex_Call) Run(run func()) *Log_TxIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Log_TxIndex_Call) Return(_a0 int) *Log_TxIndex_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Log_TxIndex_Call) RunAndReturn(run func() int) *Log_TxIndex_Call {
	_c.Call.Return(run)
	return _c
}

// NewLog creates a new instance of Log. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLog(t interface {
	mock.TestingT
	Cleanup(func())
}) *Log {
	mock := &Log{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
