all:
	clang test.cpp -lcypto -o test -std=c++11

clean:
	rm test