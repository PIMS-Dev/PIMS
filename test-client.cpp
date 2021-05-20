#include<stdio.h>
#include"PIMS-socket.h"

int main() {
	PIMS_SOCKET* socket = createSocket();
	connectSocket(socket,"127.0.0.1",24940);
	int status = 0;
	unsigned char*data = receiveData(socket,&status);
	printf("%s",data);
	return 0;
}
