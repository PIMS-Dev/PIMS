#include"PIMS-base.h"
#include<stdio.h>
#include<string>

using namespace std;

int main() {
    string str1 = "test1";
    string str2 = "test2";
    unsigned char* result = doubleSHA256(str1.c_str());
    for(int i=0;i<32;++i) {
        printf("%02x",result[i]);
    }
    printf("\n");
    unsigned char* encrypted = SHA256EncryptAES128CBC((const unsigned char*)str2.c_str(),result);
    for(int i=0;i<16;++i) {
        printf("%02x",encrypted[i]);
    }
    printf("\n");
    return 0;
}
