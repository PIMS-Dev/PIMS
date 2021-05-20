#ifndef PIMS_SOCKET_H

#define PIMS_SOCKET_H

#include<stdio.h>
#include<stdlib.h>
#include<errno.h>

#ifdef WIN32

#include<winsock2.h>
#include<windows.h>

typedef struct  {
	SOCKET socket;
}PIMS_SOCKET;

WSADATA wsaData;

void socketStartup() {
	WSAStartup(MAKEWORD(2,2),&wsaData);
}

void socketCleanup() {
	WSACleanup();
}

PIMS_SOCKET* createSocket() {
	SOCKET new_socket = socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
	PIMS_SOCKET* pims_socket = (PIMS_SOCKET*)malloc(sizeof(PIMS_SOCKET));
	pims_socket->socket = new_socket;
	return pims_socket;
}

bool bindSocket(PIMS_SOCKET* socket, char* ip,int port) {
	struct sockaddr_in socket_address;
	memset((void*)(&socket_address),0,sizeof(socket_address));
	socket_address.sin_family = PF_INET;
	socket_address.sin_addr.s_addr = inet_addr(ip);
	socket_address.sin_port = htons(port);
	if(bind(socket->socket,(SOCKADDR*)(&socket_address),sizeof(SOCKADDR)) == SOCKET_ERROR) {
		return false;
	}
	listen(socket->socket,5);
	
	return true;
}

bool connectSocket(PIMS_SOCKET*socket, char* ip,int port) {
	struct sockaddr_in socket_address;
	memset((void*)(&socket_address),0,sizeof(socket_address));
	socket_address.sin_family = PF_INET;
	socket_address.sin_addr.s_addr = inet_addr(ip);
	socket_address.sin_port = htons(port);
	if(connect(socket->socket,(SOCKADDR*)(&socket_address),sizeof(SOCKADDR)) == SOCKET_ERROR) {
		return false;
	}
	return true;
}

PIMS_SOCKET* acceptConnection(PIMS_SOCKET* socket) {
	SOCKADDR client_address;
	SOCKET client_socket = accept(socket->socket,&client_address,sizeof(SOCKADDR));
	PIMS_SOCKET* connection = (PINS_SOCKET*)malloc(sizeof(PIMS_SOCKET));
	connection->socket = client_socket
	return connection;
}

void sendData(PIMS_SOCKET* socket, const unsigned char* data, unsigned int length) {
	send(socket->socket,data,length,0);
}

unsigned char* receiveData(PIMS_SOCKET* socket,int* status) {
	unsigned char recv_buffer[256];
	unsigned int ret = 0;
	
	*status = 0;
	ret = recv(socket->socket,recv_buffer,256,0);
	if(ret==0) {
		*status = 1;
		return NULL;
	} else if(ret<0) {
		if(errno==EAGAIN || errno==EWOULDBLOCK || errno==EINTR) {
			return NULL;
		} else {
			*status = -1;
			return NULL;
		}
	}

	unsigned char* recv_data = (unsigned char*)malloc(ret);
	memcpy((void*)recv_data,(void*)(&recv_buffer),ret);
	int all_length = ret;
	unsigned char* new_recv_data;
	while(ret==256) {
		ret = recv(socket->socket,recv_buffer,256,0);
		if(ret<0) {
			if(errno==EAGAIN || errno==EWOULDBLOCK || errno==EINTR) {
				break;
			}
		} else {
			*status = -1;
			return NULL;
		}
		
		all_length += ret;
		new_recv_data = (unsigned char*)realloc(recv_data,all_length);
		free(recv_data);
		recv_data = new_recv_data;
	}

	return recv_data;
}

void closeSocket(PIMS_SOCKET* socket) {
	closesocket(socket->socket);
	free(socket);
}

#endif

#endif