package main

import (
	"fmt"

	"github.com/lsegal/cppgo/internal/asmcall/asmcalltest"
	"github.com/lsegal/cppgo/internal/asmcall/cdecl"
)

func main() {
	fmt.Println(cdecl.Call(asmcalltest.GetCdeclcallF6Addr(), 1, 2, 3, 4, 5, 6))
}
