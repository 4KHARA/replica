package replica

/*
* Call Data methods here
 */

func (m *mockData) AddCallData(fnName string, data ...interface{}) {
	// Check if there's an array under the key
	if _, ok := m.MockCallData[fnName]; !ok {
		m.MockCallData[fnName] = make([][]interface{}, 0)
	}

	// Add the data to the array
	m.MockCallData[fnName] = append(m.MockCallData[fnName], data)
}

func (m *mockData) GetCallParams(fnName string) [][]interface{} {
	return m.MockCallData[fnName]
}
