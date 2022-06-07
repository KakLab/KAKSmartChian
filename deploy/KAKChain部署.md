# KAKChain部署

## 1. 简介

KAKChain是一个POS共识的区块链，基于官方以太坊开发。为了保持最大的兼容性，EVM、合约、Web3、RPC等接口和以太坊完全一致。KAKChain为了降低能源功耗，摒弃了以太坊的POW共识机制，实现了质押出块的POS共识机制。

## 2. 部署

### 2.1 硬件要求：

最小:

- 双核CPU
- 4GB 内存
- 1TB 空余硬盘同步数据
- 8 MBit/sec 网速带宽

推荐:

- 高速四核以上CPU
- 16GB以上内存
- 至少1T的高速固态硬盘
- 25+ MBit/sec 网络速度

### 2.2 编译

KAKChain依赖golang，需要先部署golang的运行环境，请参照golang官方说明。

下载KAKChain源码：

~~~
git clone https://github.com/KakLab/KAKSmartChian.git
~~~

进入KAKSmartChian目录，并编译：

~~~
cd KAKSmartChian
make all
~~~

编译后的geth在目录“build/bin/”下面。geth为kakchain的主程序。

### 2.3 运行

**初始化节点：**

创世配置文件的路径：“deploy/genesis.json”。

~~~
geth init genesis.json
~~~

默认在用户的根目录下生成初始化文件，也可以指定目录：

~~~
mkdir node1
geth --datadir node1  init genesis.json
~~~

**启动节点：**

启动节点需要访问kakchain的根节点，通过bootnode来进行路由控制。bootnode的地址为：

~~~
enode://895ee59590233648b19f2e111784cbf23e7842c364d68c9434282491ba84c5f60dcd42028bbf05d3ea77b8a991f65e2d6c0f835465ce90ae02611cd5aee1ab05@103.50.206.103:0?discport=30310
~~~

kakchain的chainid和networkid：5198。

启动全节点：全节点会同步所有的区块，并把区块保存在本地。并且启动web3接口，可以通过rpc调用。

~~~
geth --datadir node1 --syncmode=full --gcmode=archive --networkid 5198 --port 30311 --http --http.vhosts='*' --http.addr '0.0.0.0' --http.port 8545 --http.api 'admin,debug,web3,eth,txpool,personal,miner,net' --http.corsdomain '*' --allow-insecure-unlock --bootnodes 'enode://895ee59590233648b19f2e111784cbf23e7842c364d68c9434282491ba84c5f60dcd42028bbf05d3ea77b8a991f65e2d6c0f835465ce90ae02611cd5aee1ab05@103.50.206.103:0?discport=30310' 
~~~

启动默认挖矿节点：

~~~
geth --datadir node1 --syncmode=full --gcmode=archive --networkid 5198 --port 30311 --http --http.vhosts='*' --http.addr '0.0.0.0' --http.port 8545 --http.api 'admin,debug,web3,eth,txpool,personal,miner,net' --http.corsdomain '*' --allow-insecure-unlock --bootnodes 'enode://895ee59590233648b19f2e111784cbf23e7842c364d68c9434282491ba84c5f60dcd42028bbf05d3ea77b8a991f65e2d6c0f835465ce90ae02611cd5aee1ab05@103.50.206.103:0?discport=30310'  -unlock '0x6e9dee4f886a7bb1ee824700b4f7302388b00510' --password password.txt --mine
~~~

相应参数解释：

~~~
# 0x6e9dee4f886a7bb1ee824700b4f7302388b00510 为节点的公钥地址，需要替换为自己的公钥地址
# --password 指定私钥的解密秘钥
# --mine 开启挖矿流程
~~~

启动命令行界面：执行web3的相关指令。

~~~
geth attach node1/geth.ipc
~~~

### 2.4 质押

KAKChain通过POS质押合约，自动获得出块权。

目前的质押阈值为10万个KAK，质押超过阈值后，进行投票，收到1/2的矿工节点投票后，新矿工地址写入区块，全网广播。新矿工进入出块流程。

POS质押合约地址：0xE9Da5f8dD481474b2fDCfe16b9C870d47fE4c530

合约abi：

~~~
[
	{
		"inputs": [],
		"name": "stake",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "unstake",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "validators",
		"outputs": [
			{
				"internalType": "address[]",
				"name": "",
				"type": "address[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
~~~

通过接口stake和unstake进行质押和赎回操作。

合约的详情请参见合约源码：deploy/stake/stake.sol。

## 3. 浏览器

通过浏览器地址：http://103.50.206.103:4000/

观察和查询链的相关信息。

