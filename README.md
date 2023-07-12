# dpc-secretmessage
Decentrallized &amp; Privacy Communication, Protocol For Encrypt And Authenticate Messages.


dpc-secretmessage 是一个消息批量加密协议，为了保护消息不被监听和篡改，生命周期开始于节点双方秘密握手通讯链路建立后，直至有效连接结束。

模块使用XSalsa20和Poly1305对消息进行加密和身份验证密钥密码学。消息长度以明文表示。
  
调用者有责任确保：

1.例如，通过将随机数1用于第一条消息，将随机数2用于第二条消息。

2.随机数足够长，随机生成的非数值具有可忽略不计的时序风险。
  
确保消息体量很小，因为：
  
 1. The whole message needs to be held in memory to be processed.
 
 2. Using large messages pressures implementations on small machines to decrypt
 and process plaintext before authenticating it. This is very dangerous, and
 this API does not allow it, but a protocol that uses excessive message sizes
 might present some implementations with no other choice.
 
 3. Fixed overheads will be sufficiently amortised by messages as small as 8KB.
 
 4. Performance may be improved by working with messages that fit into data caches.
 
 Thus large amounts of data should be chunked so that each message is small.
 (Each message still needs a unique nonce.) If in doubt, 16KB is a reasonable
 chunk size.