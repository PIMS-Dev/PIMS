all:
	g++ test.cpp -lcrypto -o test.out -std=c++11
	./test.out

clean:
	rm test
