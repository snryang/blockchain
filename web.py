import hashlib
import json
import Blockchain
from textwrap import dedent
from time import time
from uuid import uuid4
from flask import Flask

app = Flask(__name__)
node_identifier = str(uuid4()).replace('-','')

blockchain = Blockchain.Blockchain()
@app.route('/main',methods=['GET'])
def main():
    return "we'll mine a new Block"

@app.route('/transactions/new',methods=['POST'])
def new_transaction():
    return "we'll add a new transaction"

@app.route('/chain',methods=['GET'])
def full_chain():
    response = {
        'chain':blockchain.chain,
        'length':len(blockchain.chain),
    }
    return jsonify(response),200

if __name__ == '__main__':
    app.run(host="0.0.0.0",port=5000)