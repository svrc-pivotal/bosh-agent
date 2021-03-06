// Code generated by counterfeiter. DO NOT EDIT.
package http_blob_providerfakes

import (
	"sync"

	httpblobprovider "github.com/cloudfoundry/bosh-agent/agent/http_blob_provider"
	"github.com/cloudfoundry/bosh-utils/crypto"
)

type FakeHTTPBlobProvider struct {
	GetStub        func(string, crypto.MultipleDigest) (string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 crypto.MultipleDigest
	}
	getReturns struct {
		result1 string
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	UploadStub        func(string, string) (crypto.MultipleDigest, error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		arg1 string
		arg2 string
	}
	uploadReturns struct {
		result1 crypto.MultipleDigest
		result2 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 crypto.MultipleDigest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHTTPBlobProvider) Get(arg1 string, arg2 crypto.MultipleDigest) (string, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 crypto.MultipleDigest
	}{arg1, arg2})
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPBlobProvider) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeHTTPBlobProvider) GetCalls(stub func(string, crypto.MultipleDigest) (string, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeHTTPBlobProvider) GetArgsForCall(i int) (string, crypto.MultipleDigest) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHTTPBlobProvider) GetReturns(result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPBlobProvider) GetReturnsOnCall(i int, result1 string, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPBlobProvider) Upload(arg1 string, arg2 string) (crypto.MultipleDigest, error) {
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Upload", []interface{}{arg1, arg2})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.uploadReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHTTPBlobProvider) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeHTTPBlobProvider) UploadCalls(stub func(string, string) (crypto.MultipleDigest, error)) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = stub
}

func (fake *FakeHTTPBlobProvider) UploadArgsForCall(i int) (string, string) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	argsForCall := fake.uploadArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHTTPBlobProvider) UploadReturns(result1 crypto.MultipleDigest, result2 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 crypto.MultipleDigest
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPBlobProvider) UploadReturnsOnCall(i int, result1 crypto.MultipleDigest, result2 error) {
	fake.uploadMutex.Lock()
	defer fake.uploadMutex.Unlock()
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 crypto.MultipleDigest
			result2 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 crypto.MultipleDigest
		result2 error
	}{result1, result2}
}

func (fake *FakeHTTPBlobProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHTTPBlobProvider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ httpblobprovider.HTTPBlobProvider = new(FakeHTTPBlobProvider)
