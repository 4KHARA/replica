package replica

import (
	"runtime"
	"strings"
)

// Update the call count for a specific function.
// You don't need to use this since its used automatically
// for every function that calls the MockFn function.
func (m *mockData) Call(fnName string) {
	m.MockCallCount[fnName]++
}

// Get the total calls for a specific mocked function.
func (m *mockData) GetCallCount(fnName string) int {
	return m.MockCallCount[fnName]
}

// Every mocked function should call this function with all its
// parameters and the function name so that metrics can be
// gathered when the mocked function is called.
//
// This function will automatically return the mocked return values
// as interface{} types if they are set, and the total call count
// for the mocked function so far.
//
// This function will automatically map the caller function name to
// the Mocks object so you just need to tell it what parameters
// are coming in if you wish to check the contents of said parameters
// in your tests.
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
