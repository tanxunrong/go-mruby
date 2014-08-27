

all :
	if [ ! -f libmruby.a ]
	then
		cp ../libmruby.a ./
	fi
		
	go build ./bench.go

clean :
	rm libmruby.a
	rm bench
	rm *.log

test : all
	./bench

.PHONY : all test clean
