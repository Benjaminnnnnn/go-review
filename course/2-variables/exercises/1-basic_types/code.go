package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit float64 = 3.14
	var costPerSMS float64 = 0.02
	var hasPermission bool = false
	var username string = "bzz"

	fmt.Printf("%v %2.f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
