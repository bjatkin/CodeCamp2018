package main

import "fmt"

func Log(i, msg string) {
	if i == CURRENT_USER {
		fmt.Println(msg)
	}
}

func Logf(i, f string, args []interface{}) {
	if i == CURRENT_USER {
		fmt.Printf(f+"\n", args...)
	}
}

func Test(i string, op func()) {
	if i == CURRENT_USER {
		op()
	}
}
