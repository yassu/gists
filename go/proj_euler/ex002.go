package main

import "fmt"

func fibs(max_n int, divisable_number int) []int {
     var l []int = make([]int, max_n);
     a, b := 1, 1;
     for j:=0; j<max_n; j++ {
     	 if a % divisable_number == 0 {
     	    l = append(l, a);
	 }
     	 a, b = b, a + b;
     }
     if a % divisable_number == 0 {
     	l = append(l, a);
     }
     return l;
}

func sum(slice []int) int {
     var s int = 0;
     for j := 0; j < len(slice); j++ {
     	 s += slice[j];
     }
     return s;
}

func main() {
     var MAX_N int = 999;
     fmt.Println(fibs(MAX_N, 2));
     fmt.Println(sum(fibs(MAX_N, 2)));
}