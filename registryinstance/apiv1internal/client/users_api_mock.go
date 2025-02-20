// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package registryinstanceclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that UsersApiMock does implement UsersApi.
// If this is not the case, regenerate this file with moq.
var _ UsersApi = &UsersApiMock{}

// UsersApiMock is a mock implementation of UsersApi.
//
// 	func TestSomethingThatUsesUsersApi(t *testing.T) {
//
// 		// make and configure a mocked UsersApi
// 		mockedUsersApi := &UsersApiMock{
// 			GetCurrentUserInfoFunc: func(ctx _context.Context) ApiGetCurrentUserInfoRequest {
// 				panic("mock out the GetCurrentUserInfo method")
// 			},
// 			GetCurrentUserInfoExecuteFunc: func(r ApiGetCurrentUserInfoRequest) (UserInfo, *_nethttp.Response, error) {
// 				panic("mock out the GetCurrentUserInfoExecute method")
// 			},
// 		}
//
// 		// use mockedUsersApi in code that requires UsersApi
// 		// and then make assertions.
//
// 	}
type UsersApiMock struct {
	// GetCurrentUserInfoFunc mocks the GetCurrentUserInfo method.
	GetCurrentUserInfoFunc func(ctx _context.Context) ApiGetCurrentUserInfoRequest

	// GetCurrentUserInfoExecuteFunc mocks the GetCurrentUserInfoExecute method.
	GetCurrentUserInfoExecuteFunc func(r ApiGetCurrentUserInfoRequest) (UserInfo, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetCurrentUserInfo holds details about calls to the GetCurrentUserInfo method.
		GetCurrentUserInfo []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetCurrentUserInfoExecute holds details about calls to the GetCurrentUserInfoExecute method.
		GetCurrentUserInfoExecute []struct {
			// R is the r argument value.
			R ApiGetCurrentUserInfoRequest
		}
	}
	lockGetCurrentUserInfo        sync.RWMutex
	lockGetCurrentUserInfoExecute sync.RWMutex
}

// GetCurrentUserInfo calls GetCurrentUserInfoFunc.
func (mock *UsersApiMock) GetCurrentUserInfo(ctx _context.Context) ApiGetCurrentUserInfoRequest {
	if mock.GetCurrentUserInfoFunc == nil {
		panic("UsersApiMock.GetCurrentUserInfoFunc: method is nil but UsersApi.GetCurrentUserInfo was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetCurrentUserInfo.Lock()
	mock.calls.GetCurrentUserInfo = append(mock.calls.GetCurrentUserInfo, callInfo)
	mock.lockGetCurrentUserInfo.Unlock()
	return mock.GetCurrentUserInfoFunc(ctx)
}

// GetCurrentUserInfoCalls gets all the calls that were made to GetCurrentUserInfo.
// Check the length with:
//     len(mockedUsersApi.GetCurrentUserInfoCalls())
func (mock *UsersApiMock) GetCurrentUserInfoCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetCurrentUserInfo.RLock()
	calls = mock.calls.GetCurrentUserInfo
	mock.lockGetCurrentUserInfo.RUnlock()
	return calls
}

// GetCurrentUserInfoExecute calls GetCurrentUserInfoExecuteFunc.
func (mock *UsersApiMock) GetCurrentUserInfoExecute(r ApiGetCurrentUserInfoRequest) (UserInfo, *_nethttp.Response, error) {
	if mock.GetCurrentUserInfoExecuteFunc == nil {
		panic("UsersApiMock.GetCurrentUserInfoExecuteFunc: method is nil but UsersApi.GetCurrentUserInfoExecute was just called")
	}
	callInfo := struct {
		R ApiGetCurrentUserInfoRequest
	}{
		R: r,
	}
	mock.lockGetCurrentUserInfoExecute.Lock()
	mock.calls.GetCurrentUserInfoExecute = append(mock.calls.GetCurrentUserInfoExecute, callInfo)
	mock.lockGetCurrentUserInfoExecute.Unlock()
	return mock.GetCurrentUserInfoExecuteFunc(r)
}

// GetCurrentUserInfoExecuteCalls gets all the calls that were made to GetCurrentUserInfoExecute.
// Check the length with:
//     len(mockedUsersApi.GetCurrentUserInfoExecuteCalls())
func (mock *UsersApiMock) GetCurrentUserInfoExecuteCalls() []struct {
	R ApiGetCurrentUserInfoRequest
} {
	var calls []struct {
		R ApiGetCurrentUserInfoRequest
	}
	mock.lockGetCurrentUserInfoExecute.RLock()
	calls = mock.calls.GetCurrentUserInfoExecute
	mock.lockGetCurrentUserInfoExecute.RUnlock()
	return calls
}
