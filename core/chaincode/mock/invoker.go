
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	sync "sync"

	ccprovider "github.com/hyperledger/fabric/core/common/ccprovider"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type Invoker struct {
	InvokeStub        func(*ccprovider.TransactionParams, *ccprovider.CCContext, *peer.ChaincodeInput) (*peer.ChaincodeMessage, error)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		arg1 *ccprovider.TransactionParams
		arg2 *ccprovider.CCContext
		arg3 *peer.ChaincodeInput
	}
	invokeReturns struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}
	invokeReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Invoker) Invoke(arg1 *ccprovider.TransactionParams, arg2 *ccprovider.CCContext, arg3 *peer.ChaincodeInput) (*peer.ChaincodeMessage, error) {
	fake.invokeMutex.Lock()
	ret, specificReturn := fake.invokeReturnsOnCall[len(fake.invokeArgsForCall)]
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		arg1 *ccprovider.TransactionParams
		arg2 *ccprovider.CCContext
		arg3 *peer.ChaincodeInput
	}{arg1, arg2, arg3})
	fake.recordInvocation("Invoke", []interface{}{arg1, arg2, arg3})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.invokeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Invoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *Invoker) InvokeCalls(stub func(*ccprovider.TransactionParams, *ccprovider.CCContext, *peer.ChaincodeInput) (*peer.ChaincodeMessage, error)) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = stub
}

func (fake *Invoker) InvokeArgsForCall(i int) (*ccprovider.TransactionParams, *ccprovider.CCContext, *peer.ChaincodeInput) {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	argsForCall := fake.invokeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Invoker) InvokeReturns(result1 *peer.ChaincodeMessage, result2 error) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}{result1, result2}
}

func (fake *Invoker) InvokeReturnsOnCall(i int, result1 *peer.ChaincodeMessage, result2 error) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = nil
	if fake.invokeReturnsOnCall == nil {
		fake.invokeReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeMessage
			result2 error
		})
	}
	fake.invokeReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeMessage
		result2 error
	}{result1, result2}
}

func (fake *Invoker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Invoker) recordInvocation(key string, args []interface{}) {
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
