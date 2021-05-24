#include<stdio.h>
#include<string>
#include"PIMS-socket.h"

using namespace std;

int main() {
	socketStartup();
	PIMS_SOCKET* socket = createSocket();
	printf("socket create\n");
	bindSocket(socket,(char*)"127.0.0.1",24940);
	PIMS_SOCKET* conn = acceptConection(socket);
	string str = "wdnmd server!";
	sendData(conn,(const unsigned char*)str.c_str(),str.length());
	closeSocket(conn);
	closeSocket(socket);
	return 0;
}
