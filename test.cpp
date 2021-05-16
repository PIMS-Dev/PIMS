#include"PIMS-base.h"
#include<stdio.h>
#include<string>

using namespace std;

int main() {
    string str1 = "test1";
    string str2 = "0000000000000000111111111111111122222222222222223333333333333333";
    unsigned char* iv = (unsigned char*)malloc(16);
    memset(iv,0,16);
    unsigned char* resultmd5 = doubleMD5((const unsigned char*)str1.c_str(),str1.length());
    unsigned char* resultsha = doubleSHA256((const unsigned char*)str1.c_str(),str1.length());
    encryptedData* encrypted = AES256CBCEncrypt((const unsigned char*)str2.c_str(),str2.length(),resultsha,iv);
    for(int i=0;i<encrypted->length;++i) {
        printf("%02x",encrypted->encrypted_data[i]);
    }
    printf("\n");
    unsigned char* plain = AES256CBCDecrypt(encrypted,resultsha,iv);
    if(plain == NULL) {
        printf("NULL");
    } else {
        printf((const char*)plain);
    }
    return 0;
}
