// 2013.03.22

package main

//
// If you have trouble, see hints/chan01.txt... but try hard to get it
// first!
//

type GopherTunnel struct {
	Data chan string
}

func main() {
	// Create new GopherTunnel
	t := GopherTunnel{}

	// Channel write: sending string to t.Data
	go func() {
		t.Data <- "Can you see me?"
	}()

	// Channel read: pulling string out of t.Data
	println("GopherTunnel message:", <-t.Data)
}