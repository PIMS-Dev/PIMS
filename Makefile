all:
	g++ test.cpp -lcrypto -o test -std=c++11
	./test

clean:
	rm test
