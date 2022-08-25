# Mande chain core

Mande chain core built on top of cosmos-sdk and implements Proof-of-credibility for consensus

## Talk to us!
* [Discord](https://discord.gg/UdUZD9GmUq)

## Hardware Requirements

* **Minimal**
    * 1 GB RAM
    * 25 GB HDD
    * 1.4 GHz CPU
* **Recommended**
    * 2 GB RAM
    * 100 GB HDD
    * 2.0 GHz x2 CPU

> NOTE: SSDs have limited TBW before non-catastrophic data errors. Running a full node requires a TB+ writes per day,
> causing rapid deterioration of SSDs over HDDs of comparable quality.

## Operating System

* Linux/Windows/MacOS(x86)
* **Recommended**
    * Linux(x86_64)

## Installation Steps

> Prerequisite: go1.18+ required. [ref](https://golang.org/doc/install)

> Prerequisite: git. [ref](https://github.com/git/git)

> Optional requirement: GNU make. [ref](https://www.gnu.org/software/make/manual/html_node/index.html)

* Clone git repository

```shell
git clone https://github.com/mande-labs/mande-chain-core.git
```

* Build

```shell
cd mande-chain-core
make build
```

### Initialize a new chain and start node

```shell
cd build
./manded init demo --chain-id devnet
./manded keys add [key_name]
./manded add-genesis-account [key_name] 100000000mand,100000000stake
./manded gentx [key_name] 70000000stake --chain-id devnet
./manded collect-gentxs
./manded start --minimum-gas-prices=0mand
```

* To start rest server set `enable=true` in `config/app.toml` under `[api]` and restart the chain

### Reset chain

```shell
rm -rf ~/.mande
```
