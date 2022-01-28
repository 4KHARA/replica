package replica

// Global object to access all the mocked
// data information.
var Mocks = NewMocks()

// Structure containing everything
type mockData struct {
	// Data for each Mocked function call grouped by function name
	MockCallData map[string][][]interface{}

	// Call count for each Mocked function, grouped by function name
	MockCallCount map[string]int

	// Specific return values that a mocked function of the given name
	// should return.
	MockReturnValues map[string][]interface{}
}

// Create a fresh Mock object with zero values.
// This is usually called automatically and all you need
// to worry about is the Mocks object but if you ever need this
// it's here.
func NewMocks() *mockData {
	return &mockData{
		MockCallData:     make(map[string][][]interface{}),
		MockCallCount:    make(map[string]int),
		MockReturnValues: make(map[string][]interface{}),
	}
}

// Reset all the data in the Mock object.
// This should usually be executed at the start of
// every test if you want no persistence between test cases.
func (m *mockData) Clear() {
	m.MockCallData = make(map[string][][]interface{})
	m.MockCallCount = make(map[string]int)
	m.MockReturnValues = make(map[string][]interface{})
}
