# PIMS
普罗米修斯实时聊天系统

## 总架构
### 用户需要知道的
1.用户名  
2.密码

### 注册流程
1.客户端获取用户名和密码  
2.客户端把用户名进行两次sha256运算后作为用户id发送至服务器进行查重，确认没有重复就进行下一步  
3.客户端把密码进行两次sha256运算，作为aes256-cbc的key，并通过目前用户id、key和一个随机数进行两次MD5后生成一个iv  
4.客户端随机生成一对rsa2048密钥，在用户数据里放入私钥  
5.客户端把统计数据用aes256-cbc加密后带上生成iv的随机数、rsa2048公钥、用户id发送至服务器，服务器储存用户id、rsa2048公钥、iv￼和加密后的用户数据，至此注册完成

### 登陆流程
1.客户端把密码进行两次sha256运算，作为用户数据aes256-cbc的key  
2.客户端把用户名进行两次sha256运算后把结果发送至服务器来获取对应的aes256-cbc加密后的用户数据和生成iv的随机数并通过用户id、key和随机数进行两次MD5后生成iv  
3.客户端用在第一步里获取的key和第二步生成的iv对用户数据进行解密，至此登陆完成

### 创建聊天
1.在登陆完成后，A客户端向获取要拉进聊天的人，并通过两次sha256后向服务器确认是否有这些人同时获取这些人的公钥，假设A拉B进来  
2.A客户端获取聊天名称，通过聊天名称、聊天内各个用户id、当前时间戳和一个随机数进行两次sha256成为聊天id并通过聊天id、当前时间戳和一个随机数进行两次sha256后作为aes256-cbc的key同时随机生成一个rsa2048密钥  
3.A客户端用所有要邀请的人的公钥加密聊天id、aes256-cbc的key、rsa2048的私钥作为聊天邀请信息，还加上rsa2048公钥、聊天id发送至服务器  
4.服务器获取到A客户端发来的数据后，把各人对应密钥加密的聊天邀请信息进行两次sha256后作为邀请信息id放入各人对应的邀请缓冲区内，并储存聊天id和对应公钥  
5.B客户端为了构造查询包向服务器请求一个随机数和时间戳  
6.B客户端用自己的用户id、经过服务器时间校准后的时间戳和从服务器获取的随机数构建一个查询包，并用自己的私钥进行数字签名后发送至服务器  
7.服务器确认查询包数字签名正确、随机数正确且时间戳与服务器当前时间戳少不超过1s后把邀请缓冲区的内容下发给B客户端    
8.B客户端解密邀请信息，把聊天id、聊天对应aes256-cbc的key和rsa2048的密钥加入自己的用户数据并重新通过目前用户id、key和一个新的随机数进行两次MD5后生成一个iv，用新的iv和key加密用户数据  
9.B客户端为了构造更改包向服务器请求一个随机数和时间戳  
10.B客户端用新加密好的用户数据、新的随机数、要删除的邀请信息id、经过服务器时间校准后的时间戳和从服务器获取的随机数构造一个更改包并用自己的私钥进行数字签名后发送至服务器  
11.服务器确认更改包数字签名正确、随机数正确且时间戳与服务器当前时间戳少不超过1s后把用户数据更新并把要求删除的邀请信息从邀请缓冲区内删除，至此创建聊天完成

### 进行通信
1.A客户端再获取要发送的消息后为了构建发送包向服务器请求一个随机数和时间戳  
2.A客户端在消息内加入消息类型、经过服务器时间校准后的时间戳并用对应群聊的key和一个随机数经过两次MD5后生成一个iv，用对应聊天id的aes256-cbc的key和生成的iv加密消息，带上群聊id、生成iv的随机数、经过服务器时间校准后的时间戳和从服务器获取的随机数并用对应聊天的rsa2048私钥进行数字签名后发送至服务器上  
3.服务器确认发送包数字签名正确、随机数正确且时间戳与服务器当前时间戳少不超过1s后把消息和生成iv的随机数放入对应群里的消息记录里并按顺序加上编号  
4.B客户端定时向服务器请求一个随机数和时间戳并构建一个消息请求包，包括聊天id、上次消息编号、经过服务器时间校准后的时间戳和从服务器获取的随机数并用对应聊天的rsa2048私钥对它进行数字签名后发至服务器  
5.服务器确认查询包数字签名正确、随机数正确且时间戳与服务器当前时间戳少不超过1s后把编号在消息请求包内消息编号以后的消息和对应消息生成iv的随机数发给B客户端  
6.B客户端收到消息后用群聊对应aes256-cbc的key和随机数进行两次MD5后生成iv进行解密，至此一次通信完成
