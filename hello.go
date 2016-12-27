package main

import "os"

func main(){
	var arg string = "Anonymous";

	if len(os.Args) > 1 {
		arg = os.Args[1];
	}
	println("Hello " + arg);
}
