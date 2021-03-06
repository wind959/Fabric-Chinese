
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/orderer/consensus/etcdraft"
	"github.com/hyperledger/fabric/protos/orderer"
)

type FakeRPC struct {
	StepStub        func(dest uint64, msg *orderer.StepRequest) (*orderer.StepResponse, error)
	stepMutex       sync.RWMutex
	stepArgsForCall []struct {
		dest uint64
		msg  *orderer.StepRequest
	}
	stepReturns struct {
		result1 *orderer.StepResponse
		result2 error
	}
	stepReturnsOnCall map[int]struct {
		result1 *orderer.StepResponse
		result2 error
	}
	SendSubmitStub        func(dest uint64, request *orderer.SubmitRequest) error
	sendSubmitMutex       sync.RWMutex
	sendSubmitArgsForCall []struct {
		dest    uint64
		request *orderer.SubmitRequest
	}
	sendSubmitReturns struct {
		result1 error
	}
	sendSubmitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRPC) Step(dest uint64, msg *orderer.StepRequest) (*orderer.StepResponse, error) {
	fake.stepMutex.Lock()
	ret, specificReturn := fake.stepReturnsOnCall[len(fake.stepArgsForCall)]
	fake.stepArgsForCall = append(fake.stepArgsForCall, struct {
		dest uint64
		msg  *orderer.StepRequest
	}{dest, msg})
	fake.recordInvocation("Step", []interface{}{dest, msg})
	fake.stepMutex.Unlock()
	if fake.StepStub != nil {
		return fake.StepStub(dest, msg)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.stepReturns.result1, fake.stepReturns.result2
}

func (fake *FakeRPC) StepCallCount() int {
	fake.stepMutex.RLock()
	defer fake.stepMutex.RUnlock()
	return len(fake.stepArgsForCall)
}

func (fake *FakeRPC) StepArgsForCall(i int) (uint64, *orderer.StepRequest) {
	fake.stepMutex.RLock()
	defer fake.stepMutex.RUnlock()
	return fake.stepArgsForCall[i].dest, fake.stepArgsForCall[i].msg
}

func (fake *FakeRPC) StepReturns(result1 *orderer.StepResponse, result2 error) {
	fake.StepStub = nil
	fake.stepReturns = struct {
		result1 *orderer.StepResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRPC) StepReturnsOnCall(i int, result1 *orderer.StepResponse, result2 error) {
	fake.StepStub = nil
	if fake.stepReturnsOnCall == nil {
		fake.stepReturnsOnCall = make(map[int]struct {
			result1 *orderer.StepResponse
			result2 error
		})
	}
	fake.stepReturnsOnCall[i] = struct {
		result1 *orderer.StepResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeRPC) SendSubmit(dest uint64, request *orderer.SubmitRequest) error {
	fake.sendSubmitMutex.Lock()
	ret, specificReturn := fake.sendSubmitReturnsOnCall[len(fake.sendSubmitArgsForCall)]
	fake.sendSubmitArgsForCall = append(fake.sendSubmitArgsForCall, struct {
		dest    uint64
		request *orderer.SubmitRequest
	}{dest, request})
	fake.recordInvocation("SendSubmit", []interface{}{dest, request})
	fake.sendSubmitMutex.Unlock()
	if fake.SendSubmitStub != nil {
		return fake.SendSubmitStub(dest, request)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sendSubmitReturns.result1
}

func (fake *FakeRPC) SendSubmitCallCount() int {
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	return len(fake.sendSubmitArgsForCall)
}

func (fake *FakeRPC) SendSubmitArgsForCall(i int) (uint64, *orderer.SubmitRequest) {
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	return fake.sendSubmitArgsForCall[i].dest, fake.sendSubmitArgsForCall[i].request
}

func (fake *FakeRPC) SendSubmitReturns(result1 error) {
	fake.SendSubmitStub = nil
	fake.sendSubmitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) SendSubmitReturnsOnCall(i int, result1 error) {
	fake.SendSubmitStub = nil
	if fake.sendSubmitReturnsOnCall == nil {
		fake.sendSubmitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendSubmitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stepMutex.RLock()
	defer fake.stepMutex.RUnlock()
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRPC) recordInvocation(key string, args []interface{}) {
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

var _ etcdraft.RPC = new(FakeRPC)
