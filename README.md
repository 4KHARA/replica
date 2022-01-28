# replica
Simple struct for fetching data related to your mocked data structures inside tests and setting expectations for those mocked structures.


<br/>

Helper functions for wrapping everything up nicely. Usually it gets pretty messy when there are
a lot of small projects accessing third party services that need to be mocked away, all of them need some type of 
solution but it might be overkill sometimes.

This is for those times

<br/>

### Use Case

Imagine you have the following interface:
```go
type Thing interface {
    Run(message string) string
}
```

You implement this interface somehow and start using it within your application
and everything is nice and testable. 

Then it comes to testing and you need to create a mock of that interface so you can 
have full control over the functions.
```go
type Test struct {}

func (t *Test) Run(message string) string {
    MockFn(message)
}
```
Now that this function is mocked away, within your tests, you can observe basic information about
it. This information includes `call count` and `parameter values`. This is usually all you need
to make sure everything is running as you expect it to.

Notice that this happens because we call `MockFn` in the mocked function.

Since we're now "hooked" into the function, we can either leave it as it is
and check the call counts and parameter values in our tests, or we can decide what 
it will return.

If we want to override the return values we need to update it a bit so that
it automatically converts them to the right return type:
```go
// Updated Run function
func (t *Test) Run(message string) string {
    callCount, returnValues := MockFn(message)

    // Check if we have a mocked return value before returning
    r := ""

    if returnValues[0] != nil {
        r = returnValues[0].(string)
    }

    return r
}
```
With that handled, we can now pass custom values into it from our tests:
```go
func TestRun(t *testing.T) {
    // Make sure to clear all the mocks data at the start of each test
    // so tests are independent of each other. 'Mocks' is a globally provided
    // object.
    Mocks.Clear() 

    // Mock the return value
    Mocks.SetReturnValues("Run", "this will be returned")

    // Check that the value matches the mocked value
    test := Test{}
    res := test.Run("aaa")

    if res != "this will be returned" {
        panic("Things don't add up!")
    }
}
```
Everything gets accomplished with the `Mocks` object and the `MockFn` function.
You can access all methods to interact with the data from the `MockFn` function using the 
`Mocks` object in your tests. Each function is fairly well documented.
