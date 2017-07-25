# JWE介绍

JWE(Json Web Encryption),可以理解为JSON格式的网络加密,有以下几个特点:<br/>
* 整个数据分为5端，每段都用"."隔开,形如:<br/>
>   eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ.<br/>
    OKOawDo13gRp2ojaHV7LFpZcgV7T6DVZKTyKOMTYUmKoTCVJRgckCL9kiMT03JGe<br/>
    ipsEdY3mx_etLbbWSrFr05kLzcSr4qKAq7YN7e9jwQRb23nfa6c9d-StnImGyFDb<br/>
    Sv04uVuxIp5Zms1gNxKKK2Da14B8S4rzVRltdYwam_lDp5XnZAYpQdb76FdIKLaV<br/>
    mqgfwX7XWRxv2322i-vDxRfqNzo_tETKzpVLzfiwQyeyPGLBIO56YJ7eObdv0je8<br/>
    1860ppamavo35UgoRdbYaBcoh9QcfylQr66oc6vFWXRcZ_ZT2LawVCWTIy3brGPi<br/>
    6UklfCpIMfIjf7iGdXKHzg.<br/>
    48V1_ALb6US04U3b.<br/>
    5eym8TW_c8SuK0ltJ3rpYIzOeDQz7TALvtu6UG9oMo4vpzs9tX_EFShS8iB7j6ji<br/>
    SdiwkIr3ajwQzaBtQD_A.<br/>
    XFBoMYUZodetZdvTiFvSkQ<br/>
* 每段数据都是以Base64UrlEncode编码的数据
jwe序列化的数据包括5段,内容分别如下:
>   BASE64URL(UTF8(JWE Protected Header)) || '.' || <br/>
    BASE64URL(JWE Encrypted Key) || '.' || <br/>
    BASE64URL(JWE Initialization Vector) || '.' || <br/>
    BASE64URL(JWE Ciphertext) || '.' || <br/>
    BASE64URL(JWE Authentication Tag)` <br/>
# JWE的工作原理

## 1.加密算法的选择
整个JWE数据共有3部分需要进行加密,分别为:密钥的加密,密文的加密和数字认证码的生成,在生成JWE数据之前,我们首先要对这三种算法进行指定。<br />
JWE是通过JWE Header来进行相应算法的指定说明。其格式如下:
```
    {"alg":"RSA1_5","enc":"A128CBC-HS256"}
```
## 2.密的过程
>   Step1.选择算法，生成JWE Header<br/>
    Step2.生成密钥并加密密钥，得到Encrypted Key<br/>
    Step3.生成向量数据，得到Initialization Vector<br/>
    Step4.加密原始报文，得到Cipher text<br/>
    Step5.生成认证码，得到Authentication Tag<br/>
    Step6.拼接数据，得到JWE Object<br/>
# 怎么使用
## JWE加密实例
根据上文的加密过程，我们来具体介绍每一步的使用：
### Step1.选择算法，生成JWE Header<br/>
>   此版本密钥加密我们选择RSA,padding方式为：pkcs1-v1_5的加密方法,密文我们选择AES128CBC加密，而认证码的生成则选择HMAC With SH256<br/>
   JWEHeader为：{"alg":"RSA1_5","enc":"A128CBC-HS256"}<br/>
    header := NewHeader(ALG_RSA1_5, ENC_A128CBC_HS256)  <br/>
    jsonHeader, err := utils.JsonEncode(header)  <br/> 
### Step2.生成密钥并加密密钥，得到Encrypted Key
>   随机生成一组AES的Key,然后用RSA进行加密<br/>
    rsa := rsa.EncryptionMethodRSA{} <br/>
    publickey, _ := rsa.GetPublicKey() <br/>
    key, RasKey := jwe.GetEncryptedKey(header, 16, publickey) <br/>           
#### func GetEncryptedKey(header Header, size int, key interface{}) (aesKey, RsaKey []byte)函数说明
>   参数<br/>
    header: Step1生成的JWE Header<br/>
    size:AES加密的密钥长度<br/>
    key:对密钥进行RSA加密的publicKey<br/>
    return：aesKey和RsaKey
### Step3.生成向量数据,得到Initialization Vector
随机生成一组向量数据（这里为了方便，与Step4合并了）
### Step4.加密原始报文，得到Cipher text
利用Step2的密钥和Step3的向量数据，将原始报文进行AES加密
```
cipher, Iv := jwe.GetCipherText(header, []byte(plant), key)
```
#### func GetCipherText(header Header, plant []byte, key interface{}) (ciphertext, IV []byte)函数说明
>   参数<br/>
    header: Step1生成的JWE Header<br/>
    plant:原始报文<br/>
    key:step2得到的aesKey<br/>
    return：密文和Iv数据向量<br/>
### Step5.生成认证码，得到Authentication Tag
>   用"."的方式将Step2的加密密钥，Step3的向量数据，Step4的密文进行拼接，然后得到认证码<br/>
    Atag := jwe.GetAuthenticationTag(header, []string{string(RasKey), string(Iv), string(cipher)}, key)<br/>
#### GetAuthenticationTag(header Header, args []string, key interface{}) (Atag []byte)函数说明
>   参数<br/>
    header: Step1生成的JWE Header<br/>
    args:Step2的加密密钥，Step3的向量数据，Step4的密文组成的数组<br/>
    key:进行Hmac加密的Key<br/>
    return:生成的认证码<br/>
### Step6.拼接以上数据，得到JWE Object
>   把以上5个步骤的数据进行Base64UrlEncode，然后按照顺序拼接，用"."分割，得到最后的数据。<br/>
    jw := jwe.GetJWE(args)<br/>
#### func GetJWE(args []string) string 函数说明
>   参数<br/>
    args:以上5的步骤的数据<br/>
    return：以Base64Base64UrlEncode编码的JWE Object<br/>
 # JWE的本质
 * JWE是一个数据加解密协议(protocol)，是一种流程和数据结构的规范<br/>
 * JWE有着非常强的定制性和可扩展性。<br/>
 开发者可以根据自己的需要，选择适合的算法对数据进行JWE加解密
 # 未完待续。。。。。
 后续还会增加其他的算法。
