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

~~~bash
$ git clone https://gitee.com/xyberium/kakchian.git
~~~

进入kakchain目录，并编译：

~~~bash
$ cd kakchain
$ make all
~~~

编译后的geth在目录“build/bin/”下面。geth为kakchain的主程序。

~~~bash
$ ls build/bin/
abidump  abigen  bootnode  checkpoint-admin  clef  devp2p  ethkey  evm  faucet  geth  p2psim  puppeth  rlpdump
~~~

### 2.3 运行

KAKChain节点内置了根节点信息，启动后会自动同步链上数据。根据功能不同，可以分为轻节点、RPC节点、挖矿节点等。

**轻节点：**

轻节点会自动同步数据，支持web3等相关的操作。启动轻节点比较简单，直接加上参数“--kak”就可以。

~~~bash
$ geth --kak

同步的数据会保存在用户的根目录。也可以制定目录。下面命令把链数据保存在dataxxx目录里。

~~~bash
$ geth --kak --datadir dataxxx
~~~

**RPC节点：**

RPC节点提供RPC服务，RPC服务支撑上层架构的DAPP、钱包等应用。

通过下面命令启动RPC节点。

~~~bash
$ geth --datadir node0 --syncmode=full --gcmode=archive --kak --http --http.vhosts='*' --http.addr '0.0.0.0' --http.port 8545 --http.api 'admin,debug,web3,eth,txpool,personal,miner,net' --http.corsdomain '*' --miner.gasprice 142857200000
~~~

这样就部署了一个RPC节点，通过端口8545提供web3的接口。

**挖矿节点：**

KAKChain是POS的共识机制，矿工通过质押获得出块权。

挖矿节点需要解锁账号，用于打包出块的时候签名。先用ethkey命令生成账户信息，账户信息存在“keyfile”中，并把公钥地址打印出来。

~~~bash
$ ethkey  generate keyfile
Password: 
Repeat password: 
Address: 0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e #生成账户的公钥地址
~~~

举例，链上数据存储在“node1”目录。

~~~bash
$ mkdir  -p node1/keystore #创建数据目录
$ cp keyfile node1/keystore/ #把账户文件拷贝到keystore目录
$ geth --datadir node1 --syncmode=full --gcmode=archive --kak --allow-insecure-unlock -unlock '0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e' --password password.txt --mine --miner.gasprice 142857200000 #启动节点
~~~

相应参数解释：

~~~bash
# --syncmode=full 全节点数据同步
# --allow-insecure-unlock -unlock '0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e' 解锁挖矿节点账户
# --password password.txt keyfile文件的密码
# --mine 启动挖矿
# --miner.gasprice 142857200000 最低交易手续费的price
~~~

**启动命令行界面**：执行web3的相关指令。

~~~bash
$ geth attach node1/geth.ipc
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

**通过浏览器质押：**

[通过浏览器质押地址](http://mainnet.kakscan.com/address/0xE9Da5f8dD481474b2fDCfe16b9C870d47fE4c530/write-contract)，打开这个链接后，显示质押合约的操作界面，可以通过调用第二个函数“stake”进行质押。在弹出的钱包选择界面选择矿工节点的账户钱包，然后质押合约记录矿工的地址，该地址的挖矿节点也就有了出块权。

![staking](/home/jingwei/go/src/gitee.com/xyberium/kakchian/deploy/jpg/staking.jpg)

目前质押阈值是10万个KAK，通过矿工挖矿账户的钱包做了质押后，矿工地址就有了出块权，会轮询出块。

矿工没有出块奖励，只有打包交易收取手续费的收益。



## 3. 浏览器

通过浏览器地址：http://mainnet.kakscan.com/

观察和查询链的相关信息。

