// Report the "type larger than address space" error message.
//
// TODO(mperillo): The behavior depends on the architecture.
package main

type big [1 << 50]byte

func main() {}
