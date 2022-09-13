# KAKSmart Chain Deployment

## 1. Introduction

KAKSmartChain(KSC) is a POS consensus blockchain developed based on the Ethereum. In order to maintain maximum compatibility, EVM, contract, Web3, RPC and other interfaces are exactly the same as those of Ethereum. In order to reduce energy consumption, KSC abandoned the POW consensus mechanism of Ethereum and realized the POS consensus mechanism to generate blocks.

## 2. Deployment

### 2.1 Hardware Requirement：

Minimum:

- Dual core CPU
- 4GB Memory
- 1TB Free hard disk sync data
- 8 MBit/sec Bandwidth

Recommendation:

- High-speed quad-core CPU or above
- 16GB Memory or above
- 1T high-speed SSD or above
- 25+ MBit/sec Bandwidth

### 2.2 Compile

KAKSmartChain relys on golang, so you need to deploy golang operating environment first. Please refer to golang official document.
Download KAKSmartChain source code：

~~~bash
$ git clone https://github.com/KakLab/KAKSmartChian.git
~~~

Enter KAKSmartChian directory, and compile:

~~~bash
$ cd KAKSmartChian
$ make all
~~~

The compiled geth is under the directory "build/bin/". geth is the main program of KAKSmartChian.

~~~bash
$ ls build/bin/
abidump  abigen  bootnode  checkpoint-admin  clef  devp2p  ethkey  evm  faucet  geth  p2psim  puppeth  rlpdump
~~~

### 2.3 Operating

The KAKSmartChain node has built-in root node information, whose data will be automatically synchronized on the chain after startup. According to different functions, it can be divided into light nodes, RPC nodes, mining nodes, etc.

**light nodes：**

Light nodes will automatically synchronize data and support web3 and other related operations. It is relatively simple to start a light node, just add the parameter "--kak" directly.

~~~bash
$ geth --kak

# Synchronized data will be saved in the user's root directory, or a custom one. The following command saves the chain data in the dataxxx directory.。
$ geth --kak --datadir dataxxx
~~~

**RPC node：**

RPC nodes provide RPC services, which support applications such as DAPPs and wallets in the upper-layer architecture.
Start RPC nodes with the following command.

~~~bash
$ geth --datadir node0 --syncmode=full --gcmode=archive --kak --http --http.vhosts='*' --http.addr '0.0.0.0' --http.port 8545 --http.api 'admin,debug,web3,eth,txpool,personal,miner,net' --http.corsdomain '*' --miner.gasprice 142857200000
~~~

An RPC node is deployed to provide the web3 interface through port 8545.

**Mining node:：**

KAKSmartChain applies the POS consensus mechanism, and miners obtain block generation right through staking.
Mining nodes need to unlock their account before signing when a block is packaged. First use the command ethkey to generate account information, which is stored in the "keyfile", and the public key address is printed out.

~~~bash
$ ethkey  generate keyfile
Password: 
Repeat password: 
Address: 0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e #Generate the public key address of the account
~~~

For example, on-chain data is stored in the "node1" directory.

~~~bash
$ mkdir  -p node1/keystore #Create data directory
$ cp keyfile node1/keystore/ #Copy account file tokeystore directory
$ geth --datadir node1 --syncmode=full --gcmode=archive --kak --allow-insecure-unlock -unlock '0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e' --password password.txt --mine --miner.gasprice 142857200000 --http --http.vhosts='*' --http.addr 'localhost' --http.port 8545 --http.api 'admin,web3,eth,txpool,personal,miner,net' --http.corsdomain '*' #Start node
~~~

Corresponding parameter explanation：

~~~bash
--syncmode=full #Data sync of full node
--allow-insecure-unlock -unlock '0xE061Eeb8E33CFaBb1C8Eb4A8302c5616aFc3E50e' # Unlock mining node account
--password password.txt # keyfile File password
--mine # Start mining
--miner.gasprice 142857200000 # minimum transaction fee price,As the network changes
--http #start http server for Staking contract
~~~

:exclamation:**security warning** :exclamation:Enable the firewall and prohibit external access to port 8545.

**Start the command line**：Execute the relevant instructions of web3.

~~~bash
$ geth attach node1/geth.ipc
~~~

### 2.4 Staking

KAKSmartChain automatically obtains the right to generate blocks through the POS staking contract.
The current staking limit is 100,000 KAK. After the staking exceeds the limit, a vote will be held. After receiving the votes of 1/2 of the miner nodes, the new miner address will be written into the block and broadcast on the entire network. The new miners enter the block generation process.

POS staking contract address：0xE9Da5f8dD481474b2fDCfe16b9C870d47fE4c530

Contract abi：

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

Staking and redemption operations are performed through the interface stake and unstake.

For details of the contract, please refer to the contract source code：deploy/stake/stake.sol。

**Staking via Browser：**

Through the [Browser Staking Address](http://mainnet.kakscan.com/address/0xE9Da5f8dD481474b2fDCfe16b9C870d47fE4c530/write-contract). After opening this link, the operation interface of the staking contract is displayed, and the staking can be made by calling the second function "stake". In the pop-up wallet selection interface, select the account wallet of the miner node, and then record the address of the miner in thestaking contract, and the mining node of this address will have the right to generate blocks.

![staking](/home/jingwei/go/src/gitee.com/xyberium/kakchian/deploy/jpg/staking.jpg)

The current staking limit is 100,000 KAK. After staking through the wallet of the miner's account, the miner's address has the right to generate blocks in queue.

There is no block generation reward for miners, only the income from packaging and transaction fees.。



## 3. **Browser**

Browse related information through：http://mainnet.kakscan.com/



