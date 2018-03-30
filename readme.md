
Chrome手件下载：https://chrome-extension-downloader.com/2ef7b976af42b7416420a756794d25f9/metamask.crx


### 疑问
区块链中每一笔交易的有效性是如何确认的？交易发送者做了些什么事？各网站节点收到这笔交易时是如何去验证这笔交易是正确的？
比特币每笔交易都有一个或多个输入TxIn,以及一个或多个输出TxOut
比特币第一笔矿工交易称为Coinbase,没有输入，它的TxIn的Hash总是被标记为00000....0000
除Coinbase外，任何一个交易的Txin都会唯一追溯到区块链上在本区块之前的某个"交易Hash,以及索引"，通过"交易Hash和索引"可唯一确定一个未花费的交易输出UTXO(Unspent Transaction Output),这样每一个TxIn都和之前的某个TxOut关联了起来
比特币没有帐户的概念，每一笔交易的TxOut并没有写上收款人的名字，也没有收款人的公钥,那么收款人想要花费BTC时，他应该如何证明自己拥有这个UTXO呢，并且其他人无法假冒收款人来花费这个UTXO?
答案是：比特币交易创建的TxOut其实并非一个简单的公钥地址，而是一个脚本，脚本意思是，谁能够提供一个签名和一个公钥，让这个脚本运行通过，谁就能花费这个UXTO.

关于 私钥 公钥 钱包地址的解释：https://zhuanlan.zhihu.com/p/28196364

```python https://github.com/snryang/pybitcointools
$ python
>>> from bitcoin import *
>>> sk = random_key()  # Generate a private key
>>> vk = privtopub(sk) # Generate a public key
>>> msg = 'hello world' # Create a simple message
>>> sig = ecdsa_sign(msg, sk) # Sign the message using your private key
>>> print sig
GxXGAt...2L/eJk=
>>> print ecdsa_verify(msg, sig, vk) # Use sig and public key to verify
True
>>> msg = 'hello mars' # Change the message
>>> print ecdsa_verify(msg, sig, vk) # Changing the msg invalidates sig
False
```

### 参考文档
- https://0dayzh.gitbooks.io/bitcoin_developer_guide/content/  比特币开发者指南
- https://yeasy.gitbooks.io/blockchain_guide/content/ 区块链技术指南
- http://dbis.rwth-aachen.de/~renzel/mobsos/lib/js/jsrsasign/sample-ecdsa.html ECDSA算法在线实验


learn blockchain 

### block data
```python
block={
    'index':1,
    'timestamp':1506057125.900785
    'transactions':[{
        'sender':"852714754545",
        'recipient':"a77f5cdfa2934df3954",
        'amount':5,
    }]
    'proof':324984774000,
    'previous_hash':"2cf24dba5fffd54f"
}
```
