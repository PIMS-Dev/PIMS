# 服务端数据储存文件结构
storeDir  
|-account  
| |-用户id  
| | |-publiUserData.data (包含用户生成iv的随机byte/用户密码生成hash加的盐/随机byte/用用户aes密钥加密后的随机byte/用户公钥)  
| | |-userData.data (以用户密码作为aes密钥加密的用户rsa私钥/群聊id列表/群聊rsa私钥列表/群聊aes密钥列表)  
| | |-inviteBuffer  
| |   |-邀请id.data (用用户公钥加密的群聊id/群聊aes密钥/群聊rsa私钥)  
| |   |-...  
| |-...  
|-chat  
  |-群聊id  
  | |-chatPublicData.data (包括目前最大聊天编号和群聊公钥)  
  | |-chatData  
  |   |-按照每65536条聊天数据分的一个文件夹  
  |   | |-聊天编号.data (包含消息类型/生成当前消息iv的随机byte/加密后的聊天数据)  
  |   | |-...  
  |   |-...  
  |-...  

# 服务端文件结构
## storeDir/account/用户id/publiUserData.data
16byte 生成iv的随机byte
32byte 用户密码生成hash加的盐(随机byte)
64byte 随机byte
64byte 用用户aes密钥加密后的随机byte
未定 用户公钥

## storeDir/account/用户id/userData.data(解密后)
4byte rsa私钥长度(uint32)
未定 rsa私钥数据
4byte 群聊数量
32byte 群聊1id
4byte 群聊1rsa私钥长度
未定 群聊1rsa私钥数据
32byte 群聊1aes密钥
...

## storeDir/account/用户id/inviteBuffer/邀请id.data(解密后)
32byte 群聊id
32byte 群聊aes密钥
未定 群聊rsa私钥

## storeDir/chat/chatPublicData.data
8byte 目前最大聊天编号(uint64)
未定 群聊公钥

## storeDir/chat/chatData/按照每65536条聊天数据分的一个文件夹/聊天编号.data
1byte 消息类型
16byte 生成当前消息iv的随机byte
未定 消息正文

# 数据传输包结构
## 总包头
1byte 包类型识别码
8byte 余下包的长度(uint64)

## 用户查询(查看用户是否存在/获取用户公钥)
### 客户端上传
识别码: 0x00
32byte 用户id
1byte 是否获取用户公钥

### 服务端下发
识别码: 0x01
1byte 存在/不存在(bool)
以下内容仅限于获取用户公钥为true且存在用户
4byte 用户公钥长度
未定 用户公钥数据

## 注册
### 客户端上传
识别码: 0x02
16byte 生成iv的随机byte
32byte 用户密码生成hash加的盐(随机byte)
32byte 用户id
64byte 用户随机byte
64byte 加密后的随机byte
4byte 下面用户公钥的长度(uint32)
未定 用户公钥
4byte 下面用户数据的长度(uint32)
未定 用户数据

### 服务端下发
识别码: 0x03
1byte 状态码(0x00为注册成功，0x01为用户id重复，0x03为不支持注册)

## 登陆(获取用户数据)
### 客户端上传1
识别码: 0x04
32byte 用户id

### 服务端下发1
识别码: 0x05
16byte 生成iv的随机byte
32byte 盐(随机byte)
64byte 加密后的用户随机byte
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码: 0x06
32byte 用用户随机byte加密后的时间戳和临时随机byte

### 服务端下发2
识别码: 0x07
1byte 状态码(0x00为登陆成功，0x01为超时/临时随机byte不对，0x02为密钥校验出错)
登陆成功才有下面的部分
4byte 下面用户数据的长度(uint32)
未定 用户数据

## 创建聊天
### 客户端上传
识别码: 0x08
32byte 聊天id
4byte 要拉入的人数(uint32)
32byte 用户1的id
4byte 下面邀请信息数据的长度(uint32)
未定 给用户1的邀请信息
...

### 服务端下发
识别码: 0x09
1byte 状态码(0x00为创建成功, 0x01为聊天id重复, 0x02为不支持创建聊天)

## 获取聊天邀请信息
### 客户端上传1
识别码: 0x0A
没了

### 服务端下发1
识别码: 0x0B
8byte 到毫秒的时间戳(uint64)
16byte 随机byte

### 客户端上传2
识别码: 0x0C
32byte 用户id
8byte 经过校准后得时间戳(uint64)
16byte 下发的随机byte
4byte 下面数字签名的长度(uint32)
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码: 0x0D
1byte 状态码(0x00为验证成功，0x01为超时/随机byte不对，0x02为签名校验失败)
验证成功才有下面的部分
4byte 邀请信息的条数(uint32)
16byte 第二次下发随机byte
4byte 下面邀请信息1数据的长度(uint32)
未定 邀请信息1
...

## 更新用户数据
### 客户端上传1
识别码:0x0E
没了

### 服务端下发1
识别码: 0x0F
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码: 0x10
32byte 用户id
8byte 经过校准后的时间戳(uint64)
16byte 临时随机byte
16byte 用户随机byte
16byte 加密后的用户随机byte
16byte 生成iv的随机byte
32byte 盐(随机byte)
32byte 用户id
4byte 下面用户公钥的长度(uint32)
未定 用户公钥
4byte 下面用户数据的长度(uint32)
未定 用户数据
4byte 下面数字签名的长度(uint32)
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码: 0x11
1byte 是否成功(bool)

## 发送信息
### 客户端上传1
识别码: 0x12
没了

### 服务端下发2
识别码:0x13
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码:0x12
8byte 经过校准后的时间戳(uint64)
16byte 临时随机byte
32byte 聊天id
1byte 消息类型
16byte 生成当前消息iv的随机byte
8byte 消息正文长度(uint64)
未定 消息正文
4byte 签名长度
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码:0x13
1byte 状态码(0x0为发送成功，0x1为正文容量过大，0x2为超时/随机数不正确，0x3为签名校验失败)

## 获取聊天当前最大序号
### 客户端上传1
识别码: 0x14
没了

### 服务端下发1
识别码:0x15
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码: 0x16
8byte 经过校准后的时间戳(uint64)
16byte 临时随机byte
32byte 聊天id
4byte 签名长度
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码: 0x17
1byte 状态码(0x00为验证成功，0x01为超时/随机byte不对，0x02为签名校验失败)
验证成功才有下面的部分
8byte 当前聊天最大序号(uint64)

## 获取聊天消息类型列表
### 客户端上传1
识别码: 0x18
没了

### 服务端下发1
识别码:0x19
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码: 0x1A
8byte 经过校准后的时间戳(uint64)
16byte 临时随机byte
32byte 聊天id
8byte 起始序号(较小)(uint64)
8byte 终结序号(较大)(uint64)
4byte 签名长度
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码: 0x1B
1byte 状态码(0x00为验证成功，0x01为超时/随机byte不对，0x02为签名校验失败)
验证成功才有下面的部分
4byte 有几条消息(uint32)
8byte 消息1序号(uint64)
1byte 消息1类型
...

## 获取聊天消息正文
### 客户端上传1
识别码: 0x18
没了

### 服务端下发1
识别码:0x19
8byte 到毫秒的时间戳(uint64)
16byte 临时随机byte

### 客户端上传2
识别码: 0x1A
8byte 经过校准后的时间戳(uint64)
16byte 临时随机byte
32byte 聊天id
8byte 起始序号(较小)(uint64)
8byte 终结序号(较大)(uint64)
1byte 消息类型
4byte 签名长度
未定 给上面所有数据除数字签名长度以外的数字签名

### 服务端下发2
识别码: 0x1B
1byte 状态码(0x00为验证成功，0x01为超时/随机byte不对，0x02为签名校验失败)
验证成功才有下面的部分
8byte 消息1序号(uint64)
16byte 生成消息1iv的随机byte
8byte 消息1正文长度(uint64)
未定 消息1正文
...