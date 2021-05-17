#ifndef PIMS_BASE_H

#define PIMS_BASE_H

#include<openssl/sha.h>
#include<openssl/md5.h>
#include<openssl/aes.h>
#include<openssl/rsa.h>
#include<stdlib.h>
#include<string.h>

typedef struct {
    unsigned int length;
    unsigned char* encrypted_data;
}encryptedData;

unsigned char* doubleMD5(const unsigned char* input, unsigned int length) {
    unsigned char* result1 = (unsigned char*)malloc(16);
    MD5((const unsigned char*)input,length,result1);
    
    unsigned char* result2 = (unsigned char*)malloc(16);
    MD5((const unsigned char*)result1,16,result2);
    
    free(result1);
    return result2;
}

unsigned char* doubleSHA256(const unsigned char* input, unsigned int length) {
    unsigned char* result1 = (unsigned char*)malloc(32);
    SHA256((const unsigned char*)input,length,result1);
    
    unsigned char* result2 = (unsigned char*)malloc(32);
    SHA256((const unsigned char*)result1,32,result2);
    
    free(result1);
    return result2;
}

encryptedData* AES256CBCEncrypt(const unsigned char* plain_data, unsigned int plain_data_length, unsigned char* key, unsigned char* iv_in) {
    AES_KEY aes_key;
    AES_set_encrypt_key((const unsigned char*)key,256,&aes_key);
    unsigned char iv[16];
    memcpy((void*)(&iv),(void*)iv_in,16);
    
    unsigned char padding = 16 - plain_data_length % 16;
    unsigned int all_plain_data_length = plain_data_length + padding;
    unsigned char* all_plain_data = (unsigned char*)malloc(all_plain_data_length+16);
    
    memcpy((void*)all_plain_data,(void*)plain_data,plain_data_length);
    for(unsigned int i=plain_data_length;i<all_plain_data_length;++i) {
        all_plain_data[i] = padding;
    }
    
    for(unsigned int i=all_plain_data_length;i<all_plain_data_length+16;++i) {
        all_plain_data[i] = 0;
    }
    
    unsigned char data_block[16];
    unsigned char out[16];
    unsigned char* encrypted_data = (unsigned char*)malloc(all_plain_data_length+1);
    for(unsigned int i=0;i<all_plain_data_length+16;i+=16) {
        memcpy((void*)(&data_block),(void*)(all_plain_data+i),16);
        AES_cbc_encrypt(data_block,out,16,&aes_key,iv,AES_ENCRYPT);
        memcpy((void*)(encrypted_data+i),(void*)(&out),16);
    }
    
    free(all_plain_data);
    encryptedData* return_struct = (encryptedData*)malloc(sizeof(encryptedData));
    return_struct->length = all_plain_data_length + 16;
    return_struct->encrypted_data = encrypted_data;
    return return_struct;
}

unsigned char* AES256CBCDecrypt(encryptedData* encrypted_data, unsigned char* key, unsigned char* iv_in) {
    AES_KEY aes_key;
    AES_set_decrypt_key((const unsigned char*)key,256,&aes_key);
    unsigned char iv[16];
    memcpy((void*)(&iv),(void*)iv_in,16);
    
    unsigned char data_block[16];
    unsigned char out[16];
    unsigned char* plain_data_all = (unsigned char*)malloc(encrypted_data->length);
    for(unsigned int i=0;i<encrypted_data->length;i+=16) {
        memcpy((void*)(&data_block),(void*)(encrypted_data->encrypted_data+i),16);
        AES_cbc_encrypt(data_block,out,16,&aes_key,iv,AES_DECRYPT);
        memcpy((void*)(plain_data_all+i),(void*)(&out),16);
    }
    
    for(unsigned int i=encrypted_data->length-16;i<encrypted_data->length;++i) {
        if(*(plain_data_all+i) != 0) {
            return NULL;
        }
    }
    
    unsigned char padding = *(plain_data_all + encrypted_data->length - 17);
    unsigned int plain_data_length = encrypted_data->length - padding - 16;
    unsigned char* plain_data = (unsigned char*)malloc(plain_data_length);
    memcpy((void*)plain_data,(void*)plain_data_all,plain_data_length);
    
    free(plain_data_all);
    return plain_data;
}

#endif
