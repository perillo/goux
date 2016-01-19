// Report the "channel element type too large" error message.
//
// NOTE(mperillo): It seems the error is reported only when using make.
package main

func main() {
	c := make(chan [1 << 32]byte)
	println(c)
}
