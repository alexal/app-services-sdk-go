// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package registryinstanceclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that SystemApiMock does implement SystemApi.
// If this is not the case, regenerate this file with moq.
var _ SystemApi = &SystemApiMock{}

// SystemApiMock is a mock implementation of SystemApi.
//
// 	func TestSomethingThatUsesSystemApi(t *testing.T) {
//
// 		// make and configure a mocked SystemApi
// 		mockedSystemApi := &SystemApiMock{
// 			GetSystemInfoFunc: func(ctx _context.Context) ApiGetSystemInfoRequest {
// 				panic("mock out the GetSystemInfo method")
// 			},
// 			GetSystemInfoExecuteFunc: func(r ApiGetSystemInfoRequest) (SystemInfo, *_nethttp.Response, error) {
// 				panic("mock out the GetSystemInfoExecute method")
// 			},
// 		}
//
// 		// use mockedSystemApi in code that requires SystemApi
// 		// and then make assertions.
//
// 	}
type SystemApiMock struct {
	// GetSystemInfoFunc mocks the GetSystemInfo method.
	GetSystemInfoFunc func(ctx _context.Context) ApiGetSystemInfoRequest

	// GetSystemInfoExecuteFunc mocks the GetSystemInfoExecute method.
	GetSystemInfoExecuteFunc func(r ApiGetSystemInfoRequest) (SystemInfo, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetSystemInfo holds details about calls to the GetSystemInfo method.
		GetSystemInfo []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetSystemInfoExecute holds details about calls to the GetSystemInfoExecute method.
		GetSystemInfoExecute []struct {
			// R is the r argument value.
			R ApiGetSystemInfoRequest
		}
	}
	lockGetSystemInfo        sync.RWMutex
	lockGetSystemInfoExecute sync.RWMutex
}

// GetSystemInfo calls GetSystemInfoFunc.
func (mock *SystemApiMock) GetSystemInfo(ctx _context.Context) ApiGetSystemInfoRequest {
	if mock.GetSystemInfoFunc == nil {
		panic("SystemApiMock.GetSystemInfoFunc: method is nil but SystemApi.GetSystemInfo was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetSystemInfo.Lock()
	mock.calls.GetSystemInfo = append(mock.calls.GetSystemInfo, callInfo)
	mock.lockGetSystemInfo.Unlock()
	return mock.GetSystemInfoFunc(ctx)
}

// GetSystemInfoCalls gets all the calls that were made to GetSystemInfo.
// Check the length with:
//     len(mockedSystemApi.GetSystemInfoCalls())
func (mock *SystemApiMock) GetSystemInfoCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetSystemInfo.RLock()
	calls = mock.calls.GetSystemInfo
	mock.lockGetSystemInfo.RUnlock()
	return calls
}

// GetSystemInfoExecute calls GetSystemInfoExecuteFunc.
func (mock *SystemApiMock) GetSystemInfoExecute(r ApiGetSystemInfoRequest) (SystemInfo, *_nethttp.Response, error) {
	if mock.GetSystemInfoExecuteFunc == nil {
		panic("SystemApiMock.GetSystemInfoExecuteFunc: method is nil but SystemApi.GetSystemInfoExecute was just called")
	}
	callInfo := struct {
		R ApiGetSystemInfoRequest
	}{
		R: r,
	}
	mock.lockGetSystemInfoExecute.Lock()
	mock.calls.GetSystemInfoExecute = append(mock.calls.GetSystemInfoExecute, callInfo)
	mock.lockGetSystemInfoExecute.Unlock()
	return mock.GetSystemInfoExecuteFunc(r)
}

// GetSystemInfoExecuteCalls gets all the calls that were made to GetSystemInfoExecute.
// Check the length with:
//     len(mockedSystemApi.GetSystemInfoExecuteCalls())
func (mock *SystemApiMock) GetSystemInfoExecuteCalls() []struct {
	R ApiGetSystemInfoRequest
} {
	var calls []struct {
		R ApiGetSystemInfoRequest
	}
	mock.lockGetSystemInfoExecute.RLock()
	calls = mock.calls.GetSystemInfoExecute
	mock.lockGetSystemInfoExecute.RUnlock()
	return calls
}
