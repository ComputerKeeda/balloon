# Balloon

In this project, we are learning how to upgrade a Cosmos | Ignite chain while preserving the older data. Here is what we have planned to test whether we have achieved our goal or not.

[//]: # "so first of all we will create a custom message through which we can perform some transaction and we will perform transaction so that it will be saved in the blockchain and save the transaction hashes (for the later purpose) then we will stop the blockchain and then write a new message and then again perform the transaction and try to get the latest tx hashes (which we will get easily according to me 🤔) and also try to get the old data if we get the old data "
[//]: # "Boom we have successfully achived what we wanted "

First, we will create a custom message that enables us to execute a transaction, which will then be recorded on the blockchain. We'll also save the transaction hashes for later use. After completing these transactions, we will halt the blockchain operation. Following this, we'll introduce a new message type and execute another set of transactions to generate new transaction hashes, which, in my opinion, should be easily retrievable. Additionally, we will attempt to access the previously saved data. If we can successfully retrieve the old data along with the new transaction hashes,

Boom, we will have successfully achieved our goal.

so let's create our first message

## 1. AddBalloon

```shell
ignite scaffold message addBalloon balloonName balloonHeight --response message
```

now we will run the chain and perform some transactions

```shell
ignite chain serve
```

let's check whether our message works or not

```shell
balloond tx
```
