# 服务端数据储存架构
storeDir  
|-account  
| |-用户id  
| | |-publicKey.pem (用户公钥)  
| | |-iv.txt (就是用户数据的iv)  
| | |-salt.txt (用户密码生成hash加的盐)
| | |-userData.data (以用户密码作为aes密钥加密的用户私钥/群聊id列表/群里rsa私钥列表/群聊aes密钥列表)  
| | |-inviteBuffer  
| |   |-邀请id.data (用用户共钥加密的群聊id/群聊rsa私钥/群聊aes密钥)  
| |   |-...  
| |-...  
|-chat  
  |-群聊id  
  | |-publicKey.pem (群聊公钥)  
  | |-nowMaxChatNum.txt (目前最大聊天编号)  
  | |-chatData  
  |   |-按照每65536条聊天数据分的一个文件夹  
  |   | |-聊天编号.data (包含消息类型/iv/加密后的聊天数据)  
  |   | |-...  
  |   |-...  
  |-...  
