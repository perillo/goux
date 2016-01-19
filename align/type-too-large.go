// Report the "type too large" error message.
//
// TODO(mperillo): The behavior depends on architecture.
// This error message is reported in two separate cases.
package main

type big struct {
	field1 [1 << 49]byte
	field2 [1 << 49]byte
}

func main() {}
