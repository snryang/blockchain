区块链中每一笔交易的有效性是如何确认的？交易发送者做了些什么事？各网站节点收到这笔交易时是如何去验证这笔交易是正确的？


关于 私钥 公钥 钱包地址的解释：https://zhuanlan.zhihu.com/p/28196364


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
