#ifndef PIMS_BASE_H

#define PIMS_BASE_H

#include<openssl/sha.h>
#include<openssl/aes.h>
#include<openssl/rsa.h>
#include<string.h>
#include<stdlib.h>

unsigned char* doubleSHA256(const char* input) {
    unsigned char* result1 = (unsigned char*)malloc(32);
    SHA256((const unsigned char*)input,strlen(input),result1);
    
    unsigned char* result2 = (unsigned char*)malloc(32);
    SHA256((const unsigned char*)result1,32,result2);
    
    free(result1);
    return result2;
}

unsigned char* SHA256EncryptAES128CBC(const unsigned char* plain_data,unsigned char* hash) {
    unsigned char iv[16];
    unsigned char* key = hash + 16;
    memcpy((void*)(&iv),(void*)hash,16);
    AES_KEY aes_key;
    AES_set_encrypt_key((const unsigned char*)key,128,&aes_key);
    
    unsigned int plain_length = strlen((const char*)plain_data);
    unsigned char padding = 16 - plain_length % 16;
    unsigned int all_plain_data_length = plain_length + padding;
    unsigned char* all_plain_data = (unsigned char*)malloc(all_plain_data_length);
    
    memcpy((void*)all_plain_data,(void*)plain_data,plain_length);
    for(int i=plain_length;i<all_plain_data_length;++i) {
        all_plain_data[i] = padding;
    }
    
    unsigned char data_block[16];
    unsigned char out[16];
    unsigned char* encrypted_data = (unsigned char*)malloc(all_plain_data_length);
    for(int i=0;i<all_plain_data_length;i+=16) {
        memcpy((void*)(&data_block),(void*)(all_plain_data+i),16);
        AES_cbc_encrypt(data_block,out,16,&aes_key,iv,AES_ENCRYPT);
        memcpy((void*)(encrypted_data+i),(void*)(&out),16);
    }

    free(all_plain_data);
    return encrypted_data;
}

#endif
