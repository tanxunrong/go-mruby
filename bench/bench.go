package main

import (
	".."
	"fmt"
	"sync/atomic"
	"runtime/pprof"
	"os"
	"flag"
)

var total int32 = 0
var cl chan int
var num = flag.Int("n",1000,"bench func call num")

func bench() {
		mrb := mruby.NewMrb()
		clz := mrb.DefineClass("Skynet",nil)
		inst,_ := clz.New()
		inst.Call("inspect")
		mrb.Close()
		atomic.AddInt32(&total,1)
		if total >= int32(*num) {
			cl <- 1
		}
}

func main() {
	flag.Parse()
	var f *os.File
	f,_ = os.Create("pprof.log")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var f2 *os.File
	f2,_ = os.Create("mem.log")
	pprof.WriteHeapProfile(f2)
	defer f2.Close()

	cl = make(chan int)
	for i:=0;i<*num;i++ {
		go bench()
	}
	<-cl
	fmt.Println("total %d",total)
}
