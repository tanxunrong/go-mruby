

all :
	cp ../libmruby.a ./
	go build ./bench.go

clean :
	rm libmruby.a
	rm bench
	rm *.log

test : all
	./bench

.PHONY : all test clean
