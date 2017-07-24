JWE介绍
====
JWE(Json Web Encryption),可以理解为JSON格式的网络加密,有以下几个特点:<br/>
* 整个数据分为5端，每段都用"."隔开,形如:<br />
`eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ.<br />
OKOawDo13gRp2ojaHV7LFpZcgV7T6DVZKTyKOMTYUmKoTCVJRgckCL9kiMT03JGe<br />
ipsEdY3mx_etLbbWSrFr05kLzcSr4qKAq7YN7e9jwQRb23nfa6c9d-StnImGyFDb<br />
Sv04uVuxIp5Zms1gNxKKK2Da14B8S4rzVRltdYwam_lDp5XnZAYpQdb76FdIKLaV<br />
mqgfwX7XWRxv2322i-vDxRfqNzo_tETKzpVLzfiwQyeyPGLBIO56YJ7eObdv0je8<br />
1860ppamavo35UgoRdbYaBcoh9QcfylQr66oc6vFWXRcZ_ZT2LawVCWTIy3brGPi<br />
6UklfCpIMfIjf7iGdXKHzg.<br />
48V1_ALb6US04U3b.<br />
5eym8TW_c8SuK0ltJ3rpYIzOeDQz7TALvtu6UG9oMo4vpzs9tX_EFShS8iB7j6ji<br />
SdiwkIr3ajwQzaBtQD_A.<br />
XFBoMYUZodetZdvTiFvSkQ`<br />
* 每段数据都是以Base64UrlEncode编码的数据
jwe序列化的数据包括5段,内容分别如下:
`BASE64URL(UTF8(JWE Protected Header)) || '.' ||
BASE64URL(JWE Encrypted Key) || '.' ||
BASE64URL(JWE Initialization Vector) || '.' ||
BASE64URL(JWE Ciphertext) || '.' ||
BASE64URL(JWE Authentication Tag)`
JWE工作原理
===
整个JWE数据共有3部分需要进行加密,分别为:密钥的加密,密文的加密和数字认证码的生成,在生成JWE数据之前,我们首先要对这三种算法进行指定。<br />
JWE是通过JWE Header来进行相应算法的指定说明。其格式如下:
```
{"alg":"RSA1_5","enc":"A128CBC-HS256"}
```
