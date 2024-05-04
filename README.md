# HeliChain
This repository contains source code for Helichaind (Helichain Daemon). Helichaind binary is the official client for HeliChain. Helichaind is built using Cosmos SDK

HeliChain is a groundbreaking blockchain Layer 1 solution, driven by our exceptional team. With a decentralized approach, we strive to connect people with the web3 world, serving as the gateway to this transformative realm.

## Prerequisite
- Go 1.21

## Install Helichain daemons
Using Makefile
```
    make
```
The **helichaind** bin file is located on **${source_directory}/build/** or **GO_PATH** (default ~/go/bin/) 

## Setup a LocalNet

### Initialize the Chain
```
# <moniker> is the custom username of the node
# <chain-id> is the identity of the chain
helichaind init <moniker> --chain-id <chain-id>
```
This command will initialize the home folder containing necessary components for your chain  
(default: ~/.helichain)

### Customize the genesis file
A genesis file is a JSON file which defines the initial state of your blockchain. It can be seen as height 0 of your blockchain. The first block, at height 1, will reference the genesis file as its parent.

The docs about genesis customization: https://hub.cosmos.network/main/resources/genesis.html

### Create your validator
Create a local key pair for creating validator:
```
helichaind keys add <key_name> 
```
Add some tokens to the wallet:
```
helichaind add-genesis-account <key_name> <amount><denom>
```
Create a validtor generation transaction:
```
helichaind gentx <key_name> <amount><denom> --chain-id <chain-id>
```
Collect the gentx to genesis file:
```
helichaind collect-gentxs
```

### Run a node
```
helichaind start 
```
## Setup testnet using testnetCmd

## Contribution
The helichaind is still in development by the HeliChain team. For more information on how to contribute to this project, please contact us at support@helichain.com

## License
helichaind project source code files are made available under Apache-2.0 License, located in the LICENSE file. Basically, you can do whatever you want as long as you include the original copyright and license notice in any copy of the software/source.
