package main

// #include <stdio.h>
// #include <errno.h>
// #include "sort.h"
// int f() {
//  int v[3] = {3, 2 ,5 };
// 	for (int i = 0; i <3; i++) {
// 		printf("%d \n", v[i]);
// 	}
//  return v[2];
// }
import "C"
import "fmt"

func main() {
	v := C.f()
	fmt.Println(v)
	C.cpp()
}
