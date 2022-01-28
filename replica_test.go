package replica

import "testing"

func before() {
	Mocks.Clear()
}

func MockedFunction(a string) (int, []interface{}) {
	c, r := MockFn(a)

	return c, r
}

func TestFunctionCallDataAdd(t *testing.T) {
	before()

	_, _ = MockedFunction("test")

	// Check that the parameters were added
	if len(Mocks.GetCallParams("MockedFunction")[0]) != 1 {
		t.Error("Expected 1 parameter, got", len(Mocks.GetCallParams("MockedFunction")[0]))
	}
}

func TestMockFnCallCountUpdate(t *testing.T) {
	before()

	_, _ = MockedFunction("test")
	_, _ = MockedFunction("test")
	count, _ := MockedFunction("test")

	// Check that the call count is correct
	if Mocks.GetCallCount("MockedFunction") != 3 {
		t.Error("Expected 3 calls, got", Mocks.GetCallCount("MockedFunction"))
	}

	// Check that the return value is correct
	if count != 3 {
		t.Error("Expected 3 calls, got", count)
	}
}

func TestMockFnMockReturn(t *testing.T) {
	before()

	_, r := MockedFunction("test")

	// Test that first run returns nothing
	if len(r) != 0 {
		t.Error("Expected 0 return values, got", len(r))
	}

	// Check that the return value updates to match the mocked one
	Mocks.SetReturnValues("MockedFunction", "mocked")

	_, r = MockedFunction("test")

	if r[0].(string) != "mocked" {
		t.Error("Expected 'mocked', got", r[0])
	}
}
