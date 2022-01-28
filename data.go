package replica

// Set new call data for a given mock function.
// This adds all the parameters of a function mapped to its
// function name. This makes it easier to check the parameters
// of a mocked function in your tests if you need.
//
// This is automatically handled for you through the
// MockFn function.
func (m *mockData) AddCallData(fnName string, data ...interface{}) {
	// Check if there's an array under the key
	if _, ok := m.MockCallData[fnName]; !ok {
		m.MockCallData[fnName] = make([][]interface{}, 0)
	}

	// Add the data to the array
	m.MockCallData[fnName] = append(m.MockCallData[fnName], data)
}

// Get a specific list of calls for a given function name,
// If a function is called 3 times, this will return an array of
// 3 items, each containing the exact parameters that were passed
// into the function for each of the function calls.
func (m *mockData) GetCallParams(fnName string) [][]interface{} {
	return m.MockCallData[fnName]
}
