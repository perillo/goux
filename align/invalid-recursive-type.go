// Report the "invalid recursive type" error message.
//
// TODO(mperillo): This error message is reported in two separate cases.
package main

type invalid struct {
	parent invalid
}

func main() {}
