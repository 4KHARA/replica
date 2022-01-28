package replica

import (
	"runtime"
	"strings"
)

/*
* Call Count methods here
 */

func (m *mockData) Call(fnName string) {
	m.MockCallCount[fnName]++
}

func (m *mockData) GetCallCount(fnName string) int {
	return m.MockCallCount[fnName]
}

// Every mocked function should call this function with all its
// parameters and the function name so that metrics can be
// gathered when the mocked function is called.
func MockFn(params ...interface{}) (int, []interface{}) {
	// Get the function name using reflection
	counter, _, _, success := runtime.Caller(1)
	if !success {
		panic("Couldn't get caller information")
	}

	fnFullName := runtime.FuncForPC(counter).Name()
	fnName := fnFullName[strings.LastIndex(fnFullName, ".")+1:]

	Mocks.AddCallData(fnName, params...)
	Mocks.Call(fnName)

	// Return the call count so that the mocked function can
	// do different things based on the call count.
	return Mocks.GetCallCount(fnName), Mocks.GetReturnValues(fnName)
}
