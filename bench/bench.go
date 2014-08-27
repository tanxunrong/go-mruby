package main

import (
	".."
	"fmt"
	"sync/atomic"
	"runtime/pprof"
	"os"
)

var total int32 = 0
var cl chan int

func bench() {
		fmt.Println("bench start")
		mrb := mruby.NewMrb()
		clz := mrb.DefineClass("Skynet",nil)
		inst,_ := clz.New()
		inst.Call("inspect")
		mrb.Close()
		atomic.AddInt32(&total,1)
		if total >= 99 {
			cl <- 1
		}
		fmt.Println("bench end")
}

func main() {
	var f *os.File
	f,_ = os.Create("pprof.log")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var f2 *os.File
	f2,_ = os.Create("mem.log")
	pprof.WriteHeapProfile(f2)
	defer f2.Close()

	cl = make(chan int)
	for i:=1;i<100;i++ {
		go bench()
	}
	<-cl
	fmt.Println("total %d",total)
}
