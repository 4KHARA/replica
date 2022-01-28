package replica

/*
* Return value methods here
 */

func (m *mockData) SetReturnValues(fnName string, value ...interface{}) {
	m.MockReturnValues[fnName] = value
}

func (m *mockData) GetReturnValues(fnName string) []interface{} {
	return m.MockReturnValues[fnName]
}
