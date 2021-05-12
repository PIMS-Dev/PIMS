#include<openssl/sha.h>
#include<stdio.h>
#include<string>

using namespace std;

int main() {
    string str = "test1";
    unsigned char result[64];
    SHA512((const unsigned char *)str.c_str(),str.length(),result);
    for(int i=0;i<64;++i) {
        printf("%02x ",result[i]);
    }
    return 0;
}
