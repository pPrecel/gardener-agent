// Code generated by mockery v2.33.3. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/pPrecel/cloudagent/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0, _a1
func (_m *Client) List(_a0 context.Context, _a1 v1.ListOptions) (*types.ShootList, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.ShootList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) (*types.ShootList, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, v1.ListOptions) *types.ShootList); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.ShootList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, v1.ListOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
