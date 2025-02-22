/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package testutils

import mock "github.com/stretchr/testify/mock"

// MockConnectionConnectAwaiter is an autogenerated mock type for the ConnectionConnectAwaiter type
type MockConnectionConnectAwaiter struct {
	mock.Mock
}

type MockConnectionConnectAwaiter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectionConnectAwaiter) EXPECT() *MockConnectionConnectAwaiter_Expecter {
	return &MockConnectionConnectAwaiter_Expecter{mock: &_m.Mock}
}

// SetAwaitDisconnectComplete provides a mock function with given fields: awaitComplete
func (_m *MockConnectionConnectAwaiter) SetAwaitDisconnectComplete(awaitComplete bool) {
	_m.Called(awaitComplete)
}

// MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAwaitDisconnectComplete'
type MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call struct {
	*mock.Call
}

// SetAwaitDisconnectComplete is a helper method to define mock.On call
//   - awaitComplete bool
func (_e *MockConnectionConnectAwaiter_Expecter) SetAwaitDisconnectComplete(awaitComplete interface{}) *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call {
	return &MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call{Call: _e.mock.On("SetAwaitDisconnectComplete", awaitComplete)}
}

func (_c *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call) Run(run func(awaitComplete bool)) *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call) Return() *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call) RunAndReturn(run func(bool)) *MockConnectionConnectAwaiter_SetAwaitDisconnectComplete_Call {
	_c.Call.Return(run)
	return _c
}

// SetAwaitSetupComplete provides a mock function with given fields: awaitComplete
func (_m *MockConnectionConnectAwaiter) SetAwaitSetupComplete(awaitComplete bool) {
	_m.Called(awaitComplete)
}

// MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAwaitSetupComplete'
type MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call struct {
	*mock.Call
}

// SetAwaitSetupComplete is a helper method to define mock.On call
//   - awaitComplete bool
func (_e *MockConnectionConnectAwaiter_Expecter) SetAwaitSetupComplete(awaitComplete interface{}) *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call {
	return &MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call{Call: _e.mock.On("SetAwaitSetupComplete", awaitComplete)}
}

func (_c *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call) Run(run func(awaitComplete bool)) *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call) Return() *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call) RunAndReturn(run func(bool)) *MockConnectionConnectAwaiter_SetAwaitSetupComplete_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionConnectAwaiter creates a new instance of MockConnectionConnectAwaiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionConnectAwaiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionConnectAwaiter {
	mock := &MockConnectionConnectAwaiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
