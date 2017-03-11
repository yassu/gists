package main

import "fmt"

func main() {
    var s int = 0;
    for j:=0; j<1000; j++ {
    	if j % 3 == 0 || j % 5 == 0 {
	   s += j;
	}
    }
    fmt.Println(s);
}