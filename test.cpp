#include"PIMS-base.h"
#include<stdio.h>
#include<string>

using namespace std;

int main() {
    string str1 = "test1";
    string str2 = "test2";
    string str3 = "wdnmd";
    unsigned char* resultmd5 = doubleMD5((const unsigned char*)str1.c_str(),str1.length());
    unsigned char* resultsha1 = doubleSHA256((const unsigned char*)str1.c_str(),str1.length());
    unsigned char* resultsha2 = doubleSHA256((const unsigned char*)str2.c_str(),str2.length());
    encryptedData* encrypted = AES256CBCEncrypt((const unsigned char*)str3.c_str(),str2.length(),resultsha1,resultmd5);
    for(int i=0;i<encrypted->length;++i) {
        printf("%02x",encrypted->encrypted_data[i]);
    }
    printf("\n");
    unsigned char* plain = AES256CBCDecrypt(encrypted,resultsha1,resultmd5);
    if(plain == NULL) {
        printf("NULL\n");
    } else {
        printf("%s\n",plain);
    }
    return 0;
}
