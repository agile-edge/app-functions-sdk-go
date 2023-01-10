// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	bootstrapinterfaces "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/interfaces"
	clientsinterfaces "github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"

	common "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	dtos "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"

	interfaces "github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"

	logger "github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"

	mock "github.com/stretchr/testify/mock"
)

// AppFunctionContext is an autogenerated mock type for the AppFunctionContext type
type AppFunctionContext struct {
	mock.Mock
}

// AddValue provides a mock function with given fields: key, value
func (_m *AppFunctionContext) AddValue(key string, value string) {
	_m.Called(key, value)
}

// ApplyValues provides a mock function with given fields: format
func (_m *AppFunctionContext) ApplyValues(format string) (string, error) {
	ret := _m.Called(format)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(format)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(format)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Clone provides a mock function with given fields:
func (_m *AppFunctionContext) Clone() interfaces.AppFunctionContext {
	ret := _m.Called()

	var r0 interfaces.AppFunctionContext
	if rf, ok := ret.Get(0).(func() interfaces.AppFunctionContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.AppFunctionContext)
		}
	}

	return r0
}

// CommandClient provides a mock function with given fields:
func (_m *AppFunctionContext) CommandClient() clientsinterfaces.CommandClient {
	ret := _m.Called()

	var r0 clientsinterfaces.CommandClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.CommandClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.CommandClient)
		}
	}

	return r0
}

// CorrelationID provides a mock function with given fields:
func (_m *AppFunctionContext) CorrelationID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DeviceClient provides a mock function with given fields:
func (_m *AppFunctionContext) DeviceClient() clientsinterfaces.DeviceClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceClient)
		}
	}

	return r0
}

// DeviceProfileClient provides a mock function with given fields:
func (_m *AppFunctionContext) DeviceProfileClient() clientsinterfaces.DeviceProfileClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceProfileClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceProfileClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceProfileClient)
		}
	}

	return r0
}

// DeviceServiceClient provides a mock function with given fields:
func (_m *AppFunctionContext) DeviceServiceClient() clientsinterfaces.DeviceServiceClient {
	ret := _m.Called()

	var r0 clientsinterfaces.DeviceServiceClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.DeviceServiceClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.DeviceServiceClient)
		}
	}

	return r0
}

// EventClient provides a mock function with given fields:
func (_m *AppFunctionContext) EventClient() clientsinterfaces.EventClient {
	ret := _m.Called()

	var r0 clientsinterfaces.EventClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.EventClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.EventClient)
		}
	}

	return r0
}

// GetAllValues provides a mock function with given fields:
func (_m *AppFunctionContext) GetAllValues() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetDeviceResource provides a mock function with given fields: profileName, resourceName
func (_m *AppFunctionContext) GetDeviceResource(profileName string, resourceName string) (dtos.DeviceResource, error) {
	ret := _m.Called(profileName, resourceName)

	var r0 dtos.DeviceResource
	if rf, ok := ret.Get(0).(func(string, string) dtos.DeviceResource); ok {
		r0 = rf(profileName, resourceName)
	} else {
		r0 = ret.Get(0).(dtos.DeviceResource)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(profileName, resourceName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValue provides a mock function with given fields: key
func (_m *AppFunctionContext) GetValue(key string) (string, bool) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// InputContentType provides a mock function with given fields:
func (_m *AppFunctionContext) InputContentType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LoggingClient provides a mock function with given fields:
func (_m *AppFunctionContext) LoggingClient() logger.LoggingClient {
	ret := _m.Called()

	var r0 logger.LoggingClient
	if rf, ok := ret.Get(0).(func() logger.LoggingClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(logger.LoggingClient)
		}
	}

	return r0
}

// MetricsManager provides a mock function with given fields:
func (_m *AppFunctionContext) MetricsManager() bootstrapinterfaces.MetricsManager {
	ret := _m.Called()

	var r0 bootstrapinterfaces.MetricsManager
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.MetricsManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.MetricsManager)
		}
	}

	return r0
}

// NotificationClient provides a mock function with given fields:
func (_m *AppFunctionContext) NotificationClient() clientsinterfaces.NotificationClient {
	ret := _m.Called()

	var r0 clientsinterfaces.NotificationClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.NotificationClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.NotificationClient)
		}
	}

	return r0
}

// PipelineId provides a mock function with given fields:
func (_m *AppFunctionContext) PipelineId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PushToCore provides a mock function with given fields: event
func (_m *AppFunctionContext) PushToCore(event dtos.Event) (common.BaseWithIdResponse, error) {
	ret := _m.Called(event)

	var r0 common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(dtos.Event) common.BaseWithIdResponse); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Get(0).(common.BaseWithIdResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dtos.Event) error); ok {
		r1 = rf(event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveValue provides a mock function with given fields: key
func (_m *AppFunctionContext) RemoveValue(key string) {
	_m.Called(key)
}

// ResponseContentType provides a mock function with given fields:
func (_m *AppFunctionContext) ResponseContentType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResponseData provides a mock function with given fields:
func (_m *AppFunctionContext) ResponseData() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// SecretProvider provides a mock function with given fields:
func (_m *AppFunctionContext) SecretProvider() bootstrapinterfaces.SecretProvider {
	ret := _m.Called()

	var r0 bootstrapinterfaces.SecretProvider
	if rf, ok := ret.Get(0).(func() bootstrapinterfaces.SecretProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bootstrapinterfaces.SecretProvider)
		}
	}

	return r0
}

// SetResponseContentType provides a mock function with given fields: _a0
func (_m *AppFunctionContext) SetResponseContentType(_a0 string) {
	_m.Called(_a0)
}

// SetResponseData provides a mock function with given fields: data
func (_m *AppFunctionContext) SetResponseData(data []byte) {
	_m.Called(data)
}

// SetRetryData provides a mock function with given fields: data
func (_m *AppFunctionContext) SetRetryData(data []byte) {
	_m.Called(data)
}

// SubscriptionClient provides a mock function with given fields:
func (_m *AppFunctionContext) SubscriptionClient() clientsinterfaces.SubscriptionClient {
	ret := _m.Called()

	var r0 clientsinterfaces.SubscriptionClient
	if rf, ok := ret.Get(0).(func() clientsinterfaces.SubscriptionClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientsinterfaces.SubscriptionClient)
		}
	}

	return r0
}

type mockConstructorTestingTNewAppFunctionContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewAppFunctionContext creates a new instance of AppFunctionContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAppFunctionContext(t mockConstructorTestingTNewAppFunctionContext) *AppFunctionContext {
	mock := &AppFunctionContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
