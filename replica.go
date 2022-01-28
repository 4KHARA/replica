package replica

var Mocks = NewMocks()

type mockData struct {
	MockCallData     map[string][][]interface{}
	MockCallCount    map[string]int
	MockReturnValues map[string][]interface{}
}

func NewMocks() *mockData {
	return &mockData{
		MockCallData:     make(map[string][][]interface{}),
		MockCallCount:    make(map[string]int),
		MockReturnValues: make(map[string][]interface{}),
	}
}

func (m *mockData) Clear() {
	m.MockCallData = make(map[string][][]interface{})
	m.MockCallCount = make(map[string]int)
	m.MockReturnValues = make(map[string][]interface{})
}
