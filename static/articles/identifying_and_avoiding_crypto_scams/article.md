# Identifying and avoiding crypto scams

<div class="text-center pb-3 sm:pb-4">
<span class="font-black"><b>September 23, 2022</b></span></div>

![careful](/articles/identifying_and_avoiding_crypto_scams/images/scams_main.png)

This article is an excerpt from a book I was intending to finish writing and publishing in 2022. I ultimately decided to piece it out to release the content I wrote at the time in my blog. Note that this was all written in 2022 and some information may be outdated. I'm adding this article to my site in chronological order, so 'published' 2022-09-23, but today is 2024-07-29. I did no proofreading on this article, so if you find any mistakes, feel free to reach out.

Scams are very common in the cryptocurrency space. It is important that you are aware of all the possible ways in which you can get scammed if you're careless. Anyone can get scammed, even experienced users, so as much freedom as you have with crypto and blockchain, you also have a big responsibility of **taking care of yourself**. No one will take care of your assets except for you. Financial institutions will protect you against some scams, but this comes with the added counterparty risk of trusting said financial institution. If you learn about scams and are skeptical enough, do your own research and learn what type of scams exist, you can keep your assets safe.

***

### Article index

- [Identifying and avoiding crypto scams](#identifying-and-avoiding-crypto-scams)
    - [Article index](#article-index)
  - [Primer on scams](#primer-on-scams)
    - [Why are scams so common in crypto?](#why-are-scams-so-common-in-crypto)
    - [Stupidity vs knowledge](#stupidity-vs-knowledge)
    - [No one is safe enough from scams](#no-one-is-safe-enough-from-scams)
    - [Token approvals](#token-approvals)
      - [Tokens](#tokens)
  - [Off-chain scams](#off-chain-scams)
    - [Random direct messages to you](#random-direct-messages-to-you)
    - [Random friend requests to you](#random-friend-requests-to-you)
      - [Baiting the scammer to see where the scam leads users](#baiting-the-scammer-to-see-where-the-scam-leads-users)
    - [Doubling and giveaway scams](#doubling-and-giveaway-scams)
    - [Email spam](#email-spam)
    - [Social engineering attacks](#social-engineering-attacks)
    - [Advertisements](#advertisements)
    - [Investment courses](#investment-courses)
    - [Fake mobile apps](#fake-mobile-apps)
    - [Fake browser extensions](#fake-browser-extensions)
    - [Supply chain attacks](#supply-chain-attacks)
    - [Fake twitter pages](#fake-twitter-pages)
  - [On-chain scams](#on-chain-scams)
    - [Scam tokens sent to your address](#scam-tokens-sent-to-your-address)
      - ['From' label spoofing](#from-label-spoofing)
      - [Fake tokens](#fake-tokens)
      - [Fake tokens appearing on your wallet UI](#fake-tokens-appearing-on-your-wallet-ui)
      - [Scam NFT collections](#scam-nft-collections)
    - [Scam token trading pairs in DEXes](#scam-token-trading-pairs-in-dexes)
  - [Conclusion](#conclusion)
    - [Acknowledgements](#acknowledgements)
    - [Image sources](#image-sources)


***

## Primer on scams

First let's discuss scams as a general topic within the cryptocurrency space, which are often times similar, and often times completely different to scams in other contexts. There's a lot to unpack, but it's important that you *learn about it* because you *will encounter it*.

### Why are scams so common in crypto?

Scams are present everywhere, however, in crypto, scams are *extraordinarily common*. There's many reasons for this, but among those that pop in my head are:

* **It's easy to get away with it**: an inherent aspect of blockchain technology is that it is permissionless. If you have your assets stolen, it is difficult but definitely possible to wash out that money in ways that are much harder to track than traditional assets allow you. Data is widely available and blockchain is very transparent, but there's already pretty solid tools to wash out money.

* **It's (still) very cheap to do**: elaborate scams like fake call centers, stealing money from people's bank accounts or elaborate phishing often require rather difficult to set up systems, sites or large operations. Scammers can even get away with very low effort scams like sending private messages to people on telegram or discord which are essentially free. Setting up scam sites is also quite easy because most application frontends are open sourced and a scammer can just fork them and change the code to make it malicious.

* **There's tons of new users**: the space is young, bitcoin has only existed since 2009 and applications, wallets, defi and blockchains are still very young technology and just getting started. It will improve over time, but just like it was very easy to scam people in the early internet with a fake popup, it is now very easy to scam people because everything is new to them. Education is currently also severely lacking, which is also one of the reasons for this book to exist.

* **UX (still) sucks**: the user experience of using crypto wallets, blockchain applications, setting up your own nodes, as well as basically everything being developed in the context of smart contract blockchains that is user-facing *sucks*. I love it, but at the time of writing this book it **really sucks**. A huge part of why so many users lose funds in wrongly signed transactions or get scammed by low-effort scammers is the horrible user experience.

* **It's (still) difficult to use**: tying up with the UX, user-facing software and sites are *hard to use*. I'm certain a big part of it comes from the fact that what's being built is not really yet prepared to be used by everyone and maybe it's not even *meant to be used by everyone*. It is still important to know lots of economy concepts as well as some basic coding to be much closer to 100% safe when interfacing with blockchain applications, however, many applications and projects still attempt to cater to broader audiences which exposes poorly informed users to significant risks. It's easy to be fooled by a pretty UI and that's **bad**.

* **Advertising giants like Google do next to zero due diligence on who the advertiser is or what they're advertising**: advertising and data companies want *money* and if that money comes at the cost of some people scammed because there's a fake website set up and it shows up first like all ads do, they'll take it. I'm sure their employees and directors wish no bad upon anyone, but as long as this remains a relatively niche topic, they won't invest resources in protecting people because to them, the scam ad still means revenue.

* **It's (still) difficult to verify what a transaction does**: this is improving, but at the time of writing, it is still quite difficult to figure out what a function call on a smart contract is doing. You may press a button on a site but determining what the button triggers to be signed on your wallet is *not easy* for regular users. This is also where even being an experienced user falls short.

* **The tech stack has too many dependencies which makes apps prone to supply chain attacks**: e.g. running an `npm install` on a defi project vomits a *crap ton* of scripts into a project that uses node to build websites and whether *all of them* are safe to incorporate on a project or not is very difficult, which opens blockchain applications to supply chain based attacks. A single malicious script could inject code that modifies what the user is intended to sign.

* **DNS hijacking is a hugely problematic supply chain attack**: if you visit an application's frontend and everything *looks normal and identical as when you last used it* you can *still get scammed* if there has been a DNS hijacking event that injects malicious transaction data. High profile exploits like a recent one done on [Curve Finance's curve.fi frontend](https://www.scorechain.com/blog/curve-finance-dns-hack) managed to steal over half a million dollars in **minutes**. These types of attacks are still possible and pose significant risks to users.

* **It's too easy**: all a moderately skilled scammer has to do is build something misleading and spread it to potentially vulnerable individuals. It's still too easy to mislead people into making rushed decisions. Best case scenario for the scammer, the bait is taken, worst case, it's not.

* **It's open to the world**: location-specific scams are e.g. common in first-world countries where technology is highly integrated in basic systems and where users are targeted *personally* based on e.g. data obtained from data breaches. In crypto literally *anyone in the world* can attempt scam you regardless of whether they know *anything* personal about you or not. There's *zero* requirements to try to scam people, no barrier of entry, and everyone is a target.

### Stupidity vs knowledge

Sometimes victims are called 'stupid' for having been scammed, however, almost all scams that people fall for are a direct effect of lack of knowledge, lack of awareness, desperation or a tense moment of great concern for their assets. With this I'm not trying to absolve people's stupid decisions when they are indeed stupid, but even those that know what they're doing get scammed and we don't call them stupid.

If someone you know gets scammed, help them recover their assets if possible, but if that's not possible, don't call them stupid. The context in which a decision was made that resulted in someone being scammed is extremely important.

In the cryptocurrency space we value ethics and the philosophy of what's being built. We want to strive for openness and welcoming new users so that they, too, can avoid making decisions people may consider 'stupid'. Therefore, instead of calling people stupid, we need to show them more than anything that *there's no stupid questions* and that there's light at the end of the tunnel.

### No one is safe enough from scams

Scams are so prevalent in the crypto space that even remarkably knowledgeable people have been scammed or attempted to be scammed. Crypto company employees, developers and CEOs are regular targets for scammers. The stakes are high and therefore one must absolutely always be cautious and watch out every unusual interaction with online randoms as well as with strange websites, especially if your wallet is connected to them. Some notable examples are:

* [Hugh Karp, CEO of Nexus Mutual](https://www.coindesk.com/markets/2020/12/14/ceo-of-defi-insurer-nexus-mutual-hacked-for-8m-in-nxm-tokens/), a large decentralized DeFi insurance protocol. Karp was using a software wallet and was the victim of a remote access exploit. Theft was ~8M USD at the time. Had hugh used a hardware wallet or multisignature wallet with hardware keys for this, he would've been safe, given that remote access wouldn't have allowed the attacker to transact unless Karp confirmed the transaction for the attacker on the hardware device.

* [Wintermute, a large investment fund](https://rekt.news/wintermute-rekt-2/), given that they used a vulnerable tool called Profanity that allowed them to generate vanity addresses (semi-customized addresses on Ethereum). This theft amounted to ~160M USD at the time. Had wintermute not cared about the cost savings provided by having an address with lots of zeros (which makes transactions slightly cheaper), they wouldn't have been exploited. Saved some money transacting, but at what cost?

Obviously, you, as a new person in crypto might not be the target of such elaborate and complex attacks, however, excess caution is *always* good. This just highlights that everyone is a target and notable examples like this show that knowledgeable people and institutions are at risk of being hacked if they're careless enough about some details. Don't be lazy, protect yourself and play it safe.

### Token approvals

In this section, token approvals are mentioned as a consequence of some scams. In the previous chapters this is exemplified and covered, especially when interacting with defi applications, however, to refresh:

#### Tokens

**Tokens** (Typically ERC20 tokens) on Ethereum and EVM chains can be moved by the owner of an account/contract, but they can also be approved to be transferred by a third party contract or another account. The approval limit works in a way where you provide a number of a specified token to be moved by that third party and it's like a approval or spending limit. If you allow 1000 tokens to be spent, whenever a tokens is moved on your behalf by that third party, the spending limit goes down. Spending limits are given on a *per address* basis, so in order to allow another e.g. 2 other accounts and 3 other contract to move tokens on my behalf, I have to perform 5 transactions, one transaction per address I want to set a nonzero spending limit for.

Example addresses:

* Account A has address `0x96599F9FC94F2aB64C3823fCDDCB032B4e23764A`
* Token U has address `0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48` (The USDC contract)
* Contract B has address `0x7d2768de32b0b80b7a3454c06bdac94a69ddc7a9` (The AAVE lending pool V2)
* Account S has address `0x4809c03e3531fcc1b307DaC238e864C67B46Cd4f` (A scammer, for example)

Example actions taken:

1. Account A has 1000 units of Token U
2. Account A decides to increase the spending limit of Contract B in order to deposit 1000 units of Token U into Contract B
3. Account A approves the spending limit of Contract B up to 1000
4. Account A deposits 400 units of Token U into Contract B, thereby reducing the spending limit of Contract B by 400 (now 600)

By calling the deposit function on Contract B, Account A is telling Contract B to move tokens from Account A stash of Token U *on their behalf*.

Had Account A not allowed Contract B to move *at least* 400 units of Token U on their behalf, the transaction would have failed.

If Account A were to unintendedly approve a nonzero spending limit for Account S (a scammer), Account S could transfer tokens *out of Account A's wallet* on their behalf smaller than or equal to the amount approved by Account A for Account S.

***

## Off-chain scams

Off-chain scams are those that the victim doesn't directly find on a block explorer, on their wallet or any other piece of software that interfaces or queries blockchain data. Meaning that if I was a victim of one of these scams, I wouldn't be directly interacting with something on the blockchain, but rather I would give off some critical information, use some vulnerable piece of software, use some spoofed website or grant access to important permissions which would lead to me being scammed, sending out funds unintendedly, or granting access for someone to move funds on my behalf (a token approval). This is the most common type of scam and the one you will encounter the most in the wild or directly on your email inbox or direct messages.

### Random direct messages to you

**Note**: do not visit any of the links depicted in the message examples, **they are all scams**. In this section I intend to analyze/showcase examples of scams you might receive and how they work.

If you use platforms like Discord or Telegram which are quite common for people to use to join crypto communities, you're eventually going to be the target of a random direct message. There's lots of different things that the content of such random direct message will have, but it usually has one important characteristic and it's that *it is unexpected*. The moment you *receive something unexpected*, you should **immediately be skeptical about it**. The general rule of thumb for messages of this kind is to **ignore them**, but I want to illustrate some examples of messages you could receive and what the trick in them is:

<span style="color: red;">**Message:**</span> Invitation to a pump and dump, crypto signals or trading server

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm1.png)

<span style="color: red;">**Scam:**</span> The server may contain links to websites that suggest making malicious transactions, phishing sites among other types of harmful content. The server may recommend and instruct purchasing a token to 'pump and dump it' so that the 'exclusive server members' can take part in it, but the server members *are the ones* that will be tricked and get 'dumped on' meaning that they would purchase the token as instructed, but the server owners will sell it right after the member buys it in order to profit from their purchase.

***

<span style="color: red;">**Message:**</span> Notification that I won a prize

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm2.png)

<span style="color: red;">**Scam:**</span> There's no prize, the link directs you to a fake exchange website where you're asked to deposit some funds to be able to withdraw your prize. As a result, the funds you send to recover your prize are taken by the website owners and the number of BTC shown on the website that you won is basically just a fake number they wrote when developing the website to make people believe there's actually a prize. Also, the site has been taken down since I received the message:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm3.png)

***

<span style="color: red;">**Message:**</span> Notification that I'm whitelisted to participate in a presale

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm4.png)

<span style="color: red;">**Scam:**</span> The project (Aztec in this case but could be any) is NOT actually running a presale for any tokens. You cannot buy any *real* tokens for the project because the project *doesn't even have a token*, nothing has been announced about it and much less a presale sent to random users in the server. The user sending it is NOT a bot despite the word 'bot' being in their username, it's just something someone created to send this to randoms. When you go on the site you'll see some pretty generic interface that allows you to connect your wallet after giving up some information which you should NOT ever give and then you'll be sent some fake tokens. The tokens are worthless and whatever you give up to buy them is effectively gone, stolen from you by some internet stranger. Any sort of presale of this sort will NOT be sent to your personal messages, this is not how that works. Additionally, Aztec is a large project that has [raised millions of dollars](https://medium.com/aztec-protocol/aztec-raises-100-million-to-build-encrypted-ethereum-dd5062ba949c) to fund their operations and development, they *don't need to do a presale to raise money*.

*** 

<span style="color: red;">**Message:**</span> An offer to use a referral code on a site to mutually benefit

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm5.png)

<span style="color: red;">**Scam:**</span> Redirects to a domain that is *not the official stellar domain (stellar.org)* and redirects to a fake domain that looks quite similar with an additional 2nd level domain which probably leads to a scam website. If not, it might be a platform where you can register but requires providing a wealth of personal details that you should *never share*. The link also leads to a blog post as shown in the preview, which suggests to me that it might just be a spoofed blog post with a similar layout to that of the official stellar site announcing some fake event. In either case, the site is down, which clearly shows that it was a scam.

***

<span style="color: red;">**Message:**</span> Notification that you've won an airdrop, which is typically an event where past users of an application or other qualified people (usually contributors) are given an amount of crypto for free as a result of their participation or contributions to the community/project. These are **never** notified through a discord DM, and you would more or less know if you are going to receive it or not.

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_dm6.png)

<span style="color: red;">**Scam:**</span> Redirects to a domain that is clearly NOT the official one as shown on the image (Uniswap's real domain is Uniswap.org) where you are prompted to sign a transaction in order to claim your share of tokens. Tokens may or may not arrive to your wallet, but they will be worthless and the transaction you sign might give the contract deployer or another third party permissions to move tokens out of your wallet or a transaction that straight up sends tokens or assets from your wallet to some random address.

***

<span style="color: red;">**Message:**</span> A message telling me there's a way to look at a viral supposed sex tape of Caroline Ellison and Sam Bankman-Fried, two notorious criminals that orchestrated the Alameda + FTX fraud and former romantic partners. A large scale fraud that stole money from investors, clients, celebrities and could be the largest ever scam in the history of cryptocurrency. The sex tape is theorized to exist, but there's been zero confirmation of it and it has never actually surfaced.

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/sex_tape_dm1.png)

<span style="color: red;">**Scam:**</span> The link posted which I've intentionally hidden because *I know at least one person reading this book will try to take a look at it* leads to a fake PornHub-like website that prompts you to connect your wallet to view the video. The site prompts you to buy access to the video for 1 dollar, however, it is extremely likely that the transaction it prompts you to make does more than just send their wallet one dollar to 'unlock the video'. Signing a random transaction from a sketchy site like this could lead to complete loss of all your funds.

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/sex_tape_dm2.png)

***

<span style="color: red;">**Message:**</span> a forwarded message recommending me to join a site called 'okxloan' (OKX is a well known large cryptocurrency exchange) to 'solve my financial woes'

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm3.png)

<span style="color: red;">**Scam:**</span> The link directs you to a site where you can supposedly take loans of very large sums of money denominated in USDT (a USD-pegged stablecoin) as well as a notification to contact a person if you need money urgently. Obviously this is a scam pretending to be some sort of subsidiary of OKX. The site is all in Malay and English (having multiple languages together is a huge red flag):

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm3-1.png)

It has some menus that can be navigated and tinkered regarding the size of the loan or in how much time it can be payed:

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm3-2.png)

But it generally just leads to a dead end which is a link to chat with someone on Whatsapp:

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm3-4.png)

I'm sure you can figure out where that goes: scam.

***

<span style="color: red;">**Message:**</span> a request to swap tokens peer 2 peer with someone that 'cannot acquire them through a decentralized exchange due to restrictions in their jurisdiction'.

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm2.png)

<span style="color: red;">**Scam:**</span> This person refuses to use a VPN, refuses to use a decentralized exchange and insists on doing a peer 2 peer trade. The purpose here is to basically leave you either hanging with worthless tokens after a 'fair trade' through a 'fair platform' or convince you to send first so they run away with your assets. Another thing the person could do, which is even more convincing and what they actually ended up doing is send me some sketchy website that should function as a p2p platform, but obviously is designed for you to either sign malicious transactions or to just leave you hanging with worthless tokens:

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm2-1.png)

Peer 2 peer platforms are well known... I had never heard of this one, so it's almost certainly a scam despite the person being generally 'decent' in the way they talked to me during the chat. Obviously, I only inquired about this with the purpose of getting behind the scam. **Never engage, block + report.**

***

<span style="color: red;">**Message:**</span> just a greeting from someone that looks to be supposedly a young woman asking me about interests in cryptocurrency investments. Writing rather long messages despite my very short, clearly uninterested replies.

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm4.png)

<span style="color: red;">**Scam:**</span> Obviously, the person is fake, and it's easy to determine by reverse searching one of the many profile pictures they have (you can set multiple on telegram):

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm4-1.png)

In either case, this is essentially just an offer to supposedly invest in some sort of high yield investment platform where they're going to multiply my money by basically taking it, blocking me, deleting the conversation for both of us and disappearing. One of the great perks of telegram is that you can delete conversations fully for both parties, something that scammers love to do either after a successful scam or a failed one. Here's their message:

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm4-2.png)

The site is of course, a generic high yield investment scam website potentially packed with typos:

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/scam_dm4-3.png)

***

<span style="color: red;">**Message:**</span> an image recommending a 'web3 antivirus' with a bitly link (a commonly used URL shortener)

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/web3_antivirus_scam.png)

<span style="color: red;">**Scam:**</span> The link redirects to a 'legit looking' website with a supposed web3 antivirus software that can be downloaded to avoid scams

![](/articles/identifying_and_avoiding_crypto_scams/images/telegram/web3_antivirus_site.png)

The site looks fairly convincing, but the software is almost certainly going to be a browser extension or computer software that could do essentially anything if given the permission to. Some of these may make it so that whenever you copy a crypto address on your clipboard, the address is replaced by another (because you may unintentionally give the app clipboard access), the app may also inject malicious transaction data on websites when you make a transaction, etc. The sky is the limit. There's no real or legit web3 antivirus or anything like this. You also do NOT need any kind of software like this if you execute basic caution.

***

<span style="color: red;">**Message:**</span> a message recommending to check out a 'cool new project' that a person or group of people are building.

![](/articles/identifying_and_avoiding_crypto_scams/images/twitter/game_tester_hanni.png)

<span style="color: red;">**Scam:**</span> The message is already rather suspicious in itself, as no one will randomly recommend you to try out their game. Then the scammer tries to convince us by saying that there will be a reward of an NFT worth 0.5 ETH at the end of the beta testing process:

![](/articles/identifying_and_avoiding_crypto_scams/images/twitter/game_tester_hanni2.png)

If someone tells you that something you do will net you a large amount of crypto (0.5 ETH is about ~800 dollars at the time of writing), it's pretty clear that they're trying to incentivize you to try it. The website is also fairly good looking:

![](/articles/identifying_and_avoiding_crypto_scams/images/twitter/game_tester_hanni3.png)

This could go many ways, or they somehow attempt to trick you into sending crypto, buying some kind of worthless token or NFT, get you to sign some malicious transaction with lots of infinite approvals, among others. Generally just stay away from anything randomly sent to you no matter how legitimate it might seem.

***

### Random friend requests to you

Same situation as random messages you may receive. A friend request is usually followed by a message, especially in platforms like Discord, where if you toggle direct messages from server members off in your general settings:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_settings/privacy_and_safety.png)

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_settings/allow_direct_msg.png)

Which you can do for specific servers if you don't want to do it discord-wide:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_settings/privacy_and_safety_server_specific.png)

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_settings/privacy_and_safety_server_specific2.png)

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_settings/privacy_and_safety_server_specific3.png)

You may start to receive friend requests instead of direct message requests:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_friend_request.png)

There's **several red flags** in this friend request:

* It's random and unexpected

* It's coming from a crypto server

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_friend_request2.png)

* It's got a 'professional looking' picture

![](/articles/identifying_and_avoiding_crypto_scams/images/discord/random_friend_request3.png)

* The picture depicts a good looking woman. Women are are less commonly found in the crypto space and perhaps this is to incite men (the main demographic in crypto by far) to click on it to take a look or to make the fake friend request more attractive.

* The username says 'nft', which means non-fungible token, a type of token on the blockchain with a specific set of properties.

* The description says something that might seem like a sentence a real human wrote. Sometimes this might contain serious misspellings which would make it even sketchier.

Overall, I am 100% confident this friend request is from either a scammer (real human) or a bot intending to spread scam links. Friend requests are more often performed by humans than bots, as it's probably a more 'aggressive' activity than a message request on discord in terms of how many captchas it may prompt the attacker. I can also assume this because I receive significantly more message requests than friend requests.

Additionally, it's important to mention that it *is possible to send malware through discord*, especially if you're using Windows. Mutahar from SomeOrdinaryGamers has covered discord viruses extensively ([1](https://youtu.be/cAlP6Vt3YNo), [2](https://youtu.be/dn-B-3NPYnY), [3](https://youtu.be/smNuREXq5wc)). Accepting a DM request or accepting a friend request *already exposes you*, stay safe and do NOT engage with the scammer.

***

#### Baiting the scammer to see where the scam leads users

After accepting this scammer's request, which you **should never do**, as I mentioned before, engaging directly with a scammer exposes you to them sending malware directly to you. **Never engage**.

The only reason I did this is to illustrate another example of a common scam targeting users looking to make profits with crypto. This scam will probably end up catching lots of new crypto users, hence why it's still rather common.

First thing I did is message the scammer with a question mark:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait1.png)

Then the scammer replies and asks if I have trading experience or have ever traded crypto before:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait2.png)

Next, the scammer sends a generic message telling me that I can earn a ridiculous sum of money in an extremely short period of time and withdraw it if I want to:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait3.png)

I moderate a couple crypto discords and we get spam like this quite regularly:

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_discord_join.png)

So the scammer continues giving me explanations saying that my "profits is always 100% guaranteed", which should obviously also immediately be a red flag if there haven't been enough of those already... it is obvious the returns this scammer is talking about are basically impossible to make. The scammer also gives me figures and encourages me to invest an amount and explains they will 10x my investment in 2 days, pretty absurd but hey let's keep pressing.

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait4.png)

The scammer then asks me how much I want to trade and I say 1k, then they affirm that they will 10x my money, reaffirms that I will be able to withdraw it and asks me if I'm on Whatsapp. Messaging platforms like Whatsapp are popular for this type of scam as they can't ban you for content because messages are encrypted (or so says Meta).

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait5.png)

I continue on Whatsapp after giving the scammer a number I set up specifically for scambaiting and the scammer gives the final instructions after I tell them that I've already purchased the 1k worth of BTC (which I didn't, of course):

![](/articles/identifying_and_avoiding_crypto_scams/images/discord_scammer_bait/scammer_bait_ws.png)

The instructions are simple, send money to this address and we will 10x it in 2 days, however, when you send it, the scammer will be gone and you won't be able to contact them again. Their number is also probably set up using some service that allow you to receive SMS messages to register accounts on platforms that require it like Whatsapp.

I decided to leave the scammer on standby and they once again messaged me to 'check on me'. Of course, this means business for them, if someone reached this point, this is potential money, so they will probably continue to follow up until I tell them I was messing with them.

Once again: **never engage with scammers.**

***

### Doubling and giveaway scams

Doubling scams have exited for a long, long time. In it's more 'legitimate' form, it's a ponzi where whatever money comes in doubles whatever money goes out, meaning new money subsidizes old money. In it's least 'legitimate' form, that we see commonly in crypto, a doubling scam generally targets high profile impersonations or, at times, hacks which try to spread the 'opportunity' of having your money doubled by sending some crypto to a specific address with the promise of receiving 2x right after.

There's many variations of this, at times they can be more exaggerated (like the example in the previous section, offering a 10x in 2 days), other times it's just a simple 'giveaway' of money. In either case, it's one of the oldest scams out there and one you should *always* assume is a scam.

Around summer 2020 there was a very large, high-profile hack targeting several twitter accounts like that of the likes of Elon Musk, Barack Obama, Bill Gates, Jeff Bezos, among others. All which have a massive following. You can read more about this specific instance [in this CNBC article](https://www.cnbc.com/2020/07/15/hackers-appear-to-target-twitter-accounts-of-elon-musk-bill-gates-others-in-digital-currency-scam.html). In essence, the scam revolved around publishing tweets from these high profile accounts which mislead users into sending BTC to addresses which promised to double their money, give away money in return, among others. Of course, the result of this was that anyone that sent money to those addresses lost it.

**No one is ever going to multiply your money**. Much less a famous celebrity or billionaire that doesn't even know who you are. They won't do that *even if you know them*...

***

### Email spam

Email has always been a good source for spam. However, with the rise of DeFi and the journey of NFTs and some parts of DeFi into the mainstream, there's been more and more creative ways of crafting emails in order to mislead users. It isn't just the good ol' fake bank site that tries to phish you for your credentials, that's no longer as effective with multi-factor authentication. Now a sometimes much more profitable strategy is to just craft emails that make users believe that something happened to their digital assets.

Multitude of data breaches from crypto companies have opened the floodgates for users to receive lots of email spam:

* [Gemini data breach](https://news.bitcoin.com/report-crypto-exchange-gemini-suffers-from-data-breach-5-7-million-emails-allegedly-leaked/)

* [Coinsquare data breach](https://www.coindesk.com/tech/2022/11/26/major-canadian-crypto-exchange-coinsquare-says-client-data-breached/)

* [Customer.io data breach](https://www.coindesk.com/business/2022/06/30/opensea-reports-email-data-breach/) affecting BlockFi, Swan Bitcoin, NYDIG, Circle and Opensea

* [Robinhood data breach](https://www.coindesk.com/business/2021/11/08/robinhood-shares-fall-after-reporting-data-security-breach/)

* [BTC Markets data breach](https://www.coindesk.com/business/2020/12/02/australian-crypto-exchange-exposes-personal-data-of-270k-users/)

* [Liquid data breach](https://www.coindesk.com/markets/2020/11/18/crypto-exchange-liquid-says-user-data-possibly-exposed-in-security-breach/)

* [Ledger data breach](https://cointelegraph.com/news/ledger-data-leak-a-simple-mistake-exposed-270k-crypto-wallet-buyers)

And many more. Some of these quite notable, containing extremely personal and sensitive information like phone numbers, residential addresses, and more. The Ledger data breach in particular revealed very sensitive information of customers which are confirmed to at least hold some digital assets, given that the list of users affected by this are people who explicitly *bought a hardware wallet to keep their digital assets safe*, many receiving [death threats](https://www.reddit.com/r/ledgerwalletleak/comments/kjbrkl/recieved_the_following_mail_from_a_polish_guy/), [physical damage threats](https://www.reddit.com/r/ledgerwalletleak/comments/ki3mvc/physical_threats_begin/) or even packages containing [fake or bugged hardware wallets](https://www.coindesk.com/tech/2021/06/17/scammers-are-sending-ledger-users-fake-hardware-wallets/). A [reddit community](https://www.reddit.com/r/ledgerwalletleak/) also spun out of this breach filled with customers which intended to organize themselves with the purpose of mounting a lawsuit against Ledger.

The general rule here is that a company does NOT need to keep data for longer than what the regulation should enforce. There's not need for a crypto hardware wallet company to keep records dating back 3+ years for customers. Trezor e.g. only keeps records for 90 days. Exchanges are a more sensitive topic, because unfortunately, due to AML laws, most require identity verification, so be cautious who you entrust with this information, not every exchange deserves to have your identity information.

In either case, **never pay attention to emails telling you about any alarming situation regarding your funds, especially in crypto**. Some examples of emails you might receive (some have no images or format because my client 
blocks it) and what the scam behind them is or might be:

<span style="color: red;">**Email:**</span> An email saying that metamask (a non-custodial software wallet) has to comply with KYC regulations and needs users to verify their wallets.

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email1.png)

<span style="color: red;">**Scam:**</span> A non-custodial wallet will never require KYC because it's *non-custodial*, it's like if you needed to verify your identity with the company that sells you the wallet you use to carry your cards/cash, nonsensical. The scam here could be several things, but there's probably going to be some sort of form to capture more personal information or a requirement to connect to a site to 'verify' the wallet (something you'll never have to do, if it even means anything), where you'll probably have to sign a malicious transaction that probably steals your funds. It may also ask you for your seed phrase which you should **never ever give to anyone**. No software wallet company will ever require you to give them their seed phrase.

There's other similar scams for other similar products like this one pretending to be from WalletConnect, which isn't even a wallet but instead a middleware to connect your wallet to sites. Obviously a scam:

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email2.png)

***

<span style="color: red;">**Email:**</span> An email saying I have a pending claim for an NFT with a floor price of 3 ETH

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email3.png)

<span style="color: red;">**Scam:**</span> First of all, no NFT platform will ever email their users saying they have a pending claim for an NFT with a specific floor price, this is never going to be legit. However, there might be NFT that can be claimed on a website that's not the official NFT trading website (because the email was not sent by them). If the scammer wants to make it more believable, then there might be a few (scammer-owned) wallets that have probably wash-traded the collection to make it seem like the floor price is actually 3 ETH, however, there's several potential attack vectors here:

1. The transaction you sign to claim the NFT might be malicious and you might unintentionally send funds to an attacker or approve of them to move tokens out of your wallet.
  
2. If (1) isn't the case, then the function you call to approve of the collection to be traded in order to claim your juicy 3 ETH by putting the NFT up for sale may be the transaction that ends up stealing funds from your wallet or approving of some scammer wallet or contract to move tokens out of your wallet.

Bottom line is no one will randomly send you free NFTs, nobody will notify of it via email and nobody will conveniently include the floor price of it to make it more attractive for you to sell it.

***

<span style="color: red;">**Email:**</span> Your wallet has failed an important Ethereum upgrade (in this case the Merge, the largest in Ethereum's history)

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email4.png)

<span style="color: red;">**Scam:**</span> Wallets can't 'fail' an Ethereum upgrade. Upgrades happen without users noticing any changes to their accounts and you will never need to 'merge' your accounts or 'upgrade them'. If you were to believe this and clicked on the link which is labeled as the official metamask website (metamask.io) but redirects to a different domain which will almost certainly prompt you to sign some transaction to 'merge your wallet' which in turn will steal all your funds or approve for tokens to be moved on your behalf.

Emails for other wallets also showcase similar ridiculous claims, like this one pretending to be from TrustWallet saying that the assets will be inaccessible if the upgrade is not made:

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email6.png)

***

<span style="color: red;">**Email:**</span> an extremely extensive tutorial on how to 'upgrade' your accounts/wallet/addresses due to an incoming Ethereum upgrade

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email5.png)

<span style="color: red;">**Scam:**</span> The very 'official looking' email instructs you, through several rather complex instructions for an ordinary user, to sign a few transactions directed at some specific contracts with a very cryptic-looking transaction payload. The transaction payload will always be a hexadecimal string which gives some instructions to a contract you're interacting to, in this case, it probably calls a specific function in a smart contract deployed by some scammer which will move funds out of your wallet or set approvals for them to move tokens out of your wallet on your behalf. The Ethereum foundation or any large blockchain developing foundation or company will *never* email you instructions on how to sign some very specific transaction, this is *always* a scam.

***

<span style="color: red;">**Email:**</span> a donation request for the humanitarian disaster stemming from the war in Ukraine

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email7.png)

<span style="color: red;">**Scam:**</span> Obviously, the donations here won't go to help anyone but instead will go to the pockets of a desperate scammer using a humanitarian crisis to pose as a node to refer donations to Ukraine. The only places where you can find official addresses to donate to a cause will be the *official social media accounts* of such institutions. Do not send crypto to random addresses that you receive in an email even if you do so with the best of intention, it is **a scam**.

***

<span style="color: red;">**Email:**</span> a notification that 'your wallet is about to shut' and that it needs to be 'verified in their system'.

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email8.png)

<span style="color: red;">**Scam:**</span> wallets do not need to be verified in *any system*, this is a scam trying to bait people into signing malicious transactions. No one can shut down a non-custodial wallet like metamask, it is software that runs locally on your system. At most what could happen is that Infura, the default RPC provider of a wallet like metamask, happens to go down for maintenance, in that case, then you can just switch to another RPC provider like Alchemy or run your own node. No centralized institution can lock you out of accessing funds in self-custody on a decentralized network. In either case, the poor english on this particular email should also be a huge red flag. A self-custody wallet developing company (in this case ConsenSys) **cannot lock you out of your funds and cannot steal them** unless they backdoor software, something which *anyone can verify they don't do* because it's [open sourced](https://github.com/MetaMask), just like all major software wallets out there.

***

<span style="color: red;">**Email:**</span> a supposed TrustWallet (wrongly stylized 'Trust Wallet') notification that I have a 'pending smart contract'... whatever that means.

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email9.png)

<span style="color: red;">**Scam:**</span> you can't have 'pending smart contracts', this is just not how things work. The link will likely direct you to some malicious link that tries to prompt you to sign a malicious transaction. Obvious scam, the nonsensical language should immediately be a red flag.

***

**Emails:** Notifications of some supposed BTC transactions, from either mining, or just transfers.

![](/articles/identifying_and_avoiding_crypto_scams/images/email/email10.png)

<span style="color: red;">**Scam:**</span> all these emails contain attached PDF files. All are either malware which may attempt to execute code or lengthy documents explaining how to get some money supposedly sent to you, but instead almost certainly attempting to mislead the reader to send money to some random address. If an email is suspicious *and contains attachments*, **never** open the attachments. They might be malware. Any unexpected email containing attachments is **immediately a risk**.

***

**Always** ignore, mark as spam or immediately delete unexpected emails with explanations, fake notifications, alarming messages or just overall unusual and unexpected information that might pertain to crypto assets in self-custody. **You didn't win any prizes**, you are **not eligible for any random airdrops** that are notified through email (none are). You are not under any danger of your funds being locked out just because an email says it. **These are all scam attempts**, be very vigilant, email is like a public port for spam and scam attempts to arrive, you will receive it even if you didn't leak your email because data breaches happen and continue to happen regularly.

It's also important to note that this is obviously *not a crypto-isolated event*, there's data breaches everywhere and for all kinds of reasons. This is an unfortunately a reality of the digital age. You can mitigate it by improving your privacy in some ways I recommend on a thorough and detailed [article on my website](https://dac.ac/blog/improving_your_privacy_security/). If you'd like to check if your email address has been in any data breaches, you can use a tool like [Have I been pwned](https://haveibeenpwned.com/) to check.

***

### Social engineering attacks

A social engineering attack is a threat that replies on an interaction between an attacker and a victim with the intention of psychologically manipulating or misleading a person into skipping or breaking ordinary security procedures. What an attacker typically does is get you to give out personal/sensitive information or get you to do things that pose a risk to the safety of your assets or your online security.

These types of attacks are especially dangerous because they could be sophisticated enough to trick you. The source of the attack, whether this is a direct attacker (a person directly interacting with you) or an indirect source (like an email or text message), these attacks tend to be the most effective and easy to fall for.

In previous sections I've showed examples of emails or messages which have intended to do this exact thing, mislead you into believing that the communication is legitimate by pretending to be a specific person or entity. In some cases, if the attackers have enough information about you, they will even straight up tell you this information to attempt to convince you. Remember, anyone can know as much about you as it has been leaked in data breaches. If you happen to be the unfortunate victim of a social engineering attack or you receive emails/communications/messages which contain personal information that you *don't ever recall sharing*, do not panic, this information has almost certainly been taken out of data breaches and you should *never be alarmed*. Getting scared or alarmed is one of the first steps to get scammed. **Do not rush decisions**, especially those involving money and ask for help or other people's opinions, ask in public communities, ask online, don't get caught up in the moment, think and if you can't think, ask publicly. If you're being **threatened** online by an email or message, it's pretty much *always* a scam.

Some social engineering attacks intend to target people with lots of funds and whether that's you or not, in this case, it serves as an example of how great of an extent a scammer will go to in order to steal some money. The bigger the target, the larger the interest from scammers, but because *everyone is a target* and no scammer knows how much you may or may not have, you may as well be the victim of an attack of this sort.

The social engineering attack that [thomasg almost fell for](https://twitter.com/thomasg_eth/status/1492663192404779013?cxt=HHwWioCzifvQgLcpAAAA) is one of the most thorough attacks I have ever seen and went from a simple direct message to Thomas even signing a malicious transaction which he luckily thought to sign with a fresh address instead of his main address. Had he used his main address, it would've costed him all his funds. The story is worth reading as a thrilling and crazy story but also as a lesson to **never sign transactions that you aren't 100% sure won't steal from you**. The source of the transaction matters and the context too. The scammers that tried to attack Thomas also had *plenty of capital*, having funded their contract through a 100 ETH tornado cash deposit, which at the time was equivalent to well over 200k USD.

***

### Advertisements

The internet is *packed* with advertising. A good rule of thumb is that every application that is *free* is profiting from some type of selling or processing of user data. Data is now one of the most valuable resources a company can have. Many companies, especially those which put profits before users, which is most of them, will sell it to third parties or process it in order to somehow profile you and target you with ads.

Companies that benefit greatly from data collection are those which operate search engines, platforms where users consume content, social media networks, etc.

When searching the web or browsing social media, you may encounter ads which may link to websites that could lead you to think are legitimate but often times are attempting to scam you. This is no exception in the crypto world, which *requires* websites, applications and code to be open and public. There's no value in security by obscurity in crypto because it means you're *trusting a third party*, which goes against the purpose of blockchain technology which is to remove the middleman.

As a result of websites and applications being completely open sourced, opportunistic scammers copy the code, modify chunks of it so it still looks identical but the transaction payloads are changed so that whenever you confirm a transaction that you believe is genuine, you end up confirming a malicious one.

There's plenty of examples for this, but I'd personally like to focus on that of search engines, which we use every day to query and find information and websites when we need them. In particular, you should *never expect* companies like Google or Meta to perform any sort of meaningful due diligence when it comes to who their advertising clients are. Google, Meta or any large corporation will gladly take money from scammers unknowingly (or knowingly) just to make a profit in order to show you advertisements. Ads on Google are a particularly popular tool for scammers to publicize these modified application sites because *ads are pushed to the top of results*. All a scammer has to do is build the site, modify the transaction payloads and buy ads on google, and now whenever you search for a specific app, the scammer's site will show up *before the actual legitimate website*:

![](/articles/identifying_and_avoiding_crypto_scams/images/ads/omnibridge.png)

In this case, 3 ads are shown before the actual legitimate website link. **Do not ever click on ads**. To avoid being scammed by advertisements you should do two things:

1. **Switch to a search engine that doesn't track you and doesn't show you ads** like [those recommended by Privacyguides.org](https://www.privacyguides.org/search-engines/).

2. **Use [DefiLlama's link directory tool](https://defillama.com/directory)** in order to find an application's official link when you need it. DefiLlama operates a [great analytic dashboard](https://defillama.com) with plenty of information regarding defi protocols and in 2022 they added an amazing link directory tool through which they provide legitimate, safe links to defi applications. 

![](/articles/identifying_and_avoiding_crypto_scams/images/ads/defillama_link_directory.png)

3. **Do not blindly trust anything**, always verify. Even the DefiLlama link directory tool, which you can trust, developed by a team known to be trustworthy and reliable may contain a wrong link. It's very unlikely but possible, if something looks weird or unusual, ask publicly in communities like [r/ethfinance](https://www.reddit.com/r/ethfinance), [r/cryptocurrency](https://www.reddit.com/r/CryptoCurrency/) or on the official community for the project. Tweet at [0xngmi](https://twitter.com/0xngmi) or anyone from the DefiLlama team about it. **Always** do your own research.

4. **Ignore social media ads too**: just like with search engines, which are particularly bad in this aspect, social media will present you with ads that are also scams. Don't trust those either. If anything is marked as an ad or is sponsored, you should immediately distrust it.

***

### Investment courses

A course that charges you a lot of money to teach you how to trade cryptocurrency, invest in crypto and use defi (among other things) and have a specific website where they're advertised in way that tries too hard to be convincing are typically one of those 'looks too good to be true' type of things. They often over-promise and under-deliver, so you should *probably* not buy them. One such example could be this website:

![](/articles/identifying_and_avoiding_crypto_scams/images/get_rich_quick/cryptocourse.png)

I haven't bought or tried this specific course, there's a chance it's legitimate and contains really good content, but I would be inclined to doubt its legitimacy and quality by the look of the website. There's typical red flag statements in it like:

* "We're Determined To Help Our Students Create SERIOUS WEALTH as Crypto Investors in 2022"
* "increase your profits without wasting months trying to figure it out on your own!"
* "Sign-Up Today and Get a Bonus!"
* "[...]  access profits whenever you want to fund your lifestyle."
* "Don't Take Our Word For It"
* " [...] A buddy lost $10,000 over night because he didn't analyze the coin he bought. I made money on it!"

Among others, however, this site *isn't even that bad*. What I want to illustrate here is that most of the times, these courses will be *overpriced* and not provide valuable enough information that you can't find elsewhere.

The reality is that *most people* that trade (or attempt to) lose money trying. If you really want to learn how to trade I would recommend pursuing formal education in some kind of university coursing a degree in a field that will help you land onto real trading or learn on your own using free online resources without a specific focus on the cryptocurrency asset class. Udemy also has plenty of content for quite cheap and will almost surely beat courses like the example I showed on this section.

Youtube has lots of good and useful contents, but most crypto trading youtubers which use past data to predict the future (if at all) are usually wrong about their predictions and most of the time you **probably should not follow the investment advice or trading strategy of most of them** because even they are almost certainly not very successful at trading and instead make money through youtube ad revenue, sponsorships, running staking pools or selling courses. Some youtubers have also done some pretty scummy things like [Bitboy](https://twitter.com/zachxbt/status/1599545021887565824?s=20), [Logan Paul](https://twitter.com/coffeebreak_YT/status/1603904533691785216?s=20), among others.

***

### Fake mobile apps

Fake mobile apps pop up all the time since there's a huge incentive to steal crypto given how 'easy' it is if the victim gives you the proper information (like a seed phrase).

Metamask, and other mobile wallet apps might actually ask for a seed phrase to restore an account created through a software wallet. Because of this, you *need* to make sure that the app you're using is genuine. Many wallet app clones have popped up to scoop up the opportunity of stealing people's seed phrases under the premise that these apps _already ask you for the seed phrase_ when restoring accounts. Even the websites created to push these fake apps are extremely good looking and at times practically indistinguishable from the original site. 

In the case of the Metamask website in particular, that given Metamask's popularity must be the one that's cloned the most by attackers, there's a button to download the android app, the button, like usually, takes you to the play store. However, in the case of scams like [this one](https://blog.cyble.com/2022/04/19/fake-metamask-app-steals-cryptocurrency/), the website will directly download an APK, which of course, when you try to restore your account will send your seed phrase to some unknown third party (the attacker).

For this reason it is extremely important that you verify the source through which you download an app and **never** type out or paste your seed phrase anywhere. If you have to do this on Metamask because you're restoring an account, it means that you've created your account *on the software wallet*, which I highly recommend you never even bother doing with any wallet unless it's a requirement, and if so, *never* add funds to it.

### Fake browser extensions

Wallet browser extensions like Metamask are a [target for attackers](https://ciphertrace.com/alert-malicious-crypto-browser-extension-masked-metamask/) just like mobile apps. Given that a browser extension can be granted extensive permissions within the web browser, opportunistic attackers will code or modify extensions and then spoof the original website of that wallet extension or attempt to deliver the extension somehow else. Once you have installed one of these spoofed extensions, there's a high chance the extension will wreak havoc on your web browser if granted the permissions it asks.

Some legitimate extensions *will*, too, ask for quite aggressive permissions if they're needed for their functionality:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_software/violentmonkey_permissions.png)

Though you should be extremely cautious regarding *what you install* on the same web browser where you make transactions through whatever wallet you're using.

I highly recommend that you **use a clean web browser** to make transactions. Do NOT use an outdated, sketchy or extension-packed web browser to make transactions. A simple clean Firefox installation with your wallet software is sufficient to transact. Extra extensions in the same web browser where you transact **are a security risk** no matter how much you trust the extension developer. You can also [verify](https://medium.com/mycrypto/how-to-ensure-youre-running-the-legitimate-version-of-metamask-5fcd8ab32b96) that the extension you installed is also the correct one if you want to. If you used the right website, you should generally be fine, but it can't hurt to verify if you want to.

### Supply chain attacks

A supply chain attack is one that is made on a portion of a stack. I like to think of web applications with the analogy of a Jenga tower that has been already played with:

![jengarobot](/articles/identifying_and_avoiding_crypto_scams/images/chapter3/supply_chain/jenga.jpeg)

There's floors where you have multiple tiles, so one tile failing won't make the tower tumble, but if you happen to remove a piece that's alone in one of the floors, the tower *will* tumble and fall.

A supply chain attack in the context of a blockchain application could affect any element of the web application stack, from javascript libraries used by the application to the DNS server the domain points to.

In ordinary applications like a social media network this can lead to meaningful consequences like data loss, but in the case of a blockchain application, the effect of one of these attacks is almost always a direct monetary of user funds.

An example of this type of attack, referenced earlier in the ["Why scams are so common in crypto"](#why-are-scams-so-common-in-crypto) section was the [DNS hijacking that Curve Finance's frontend](https://www.scorechain.com/blog/curve-finance-dns-hack) was a victim of in summer 2022, which led to the loss of over half a million dollars worth of crypto at the time.

Knowing this, you eventually understand that the surface area for attacks on blockchain applications is absolutely massive. From vulnerabilities in the smart contract, to dns hijacking, compromising javascript libraries, compromising middleware that interacts with the application (wallet software), etc. Usually the effect of an attack on any link in the chain (except for an exploit to the smart contract directly) is monetary loss due to the user signing a malicious transaction unintendedly. So a big part of *knowing when to sign a transaction or not* is a responsibility that lands on the user, which makes the user experience quite poor, but it's the way to do things if you want to benefit from the amazing financial tools available in decentralized finance.

### Fake twitter pages

Twitter is rampant with scams. It is absolutely packed with fake pages, scam tweets, bots offering you help and support, etc. In fact, at times, just saying the word 'Metamask' or any other keyword like this will make it so that a massive flock of bots replies to your tweet with scam links and email addresses:

![](/articles/identifying_and_avoiding_crypto_scams/images/twitter/metamask_support.png)

Fake accounts are also incredibly common:

![](/articles/identifying_and_avoiding_crypto_scams/images/twitter/vitalik_scam.jpeg)

Make sure to *always* verify a person's username before doing anything. Also, if anything looks mildly sketchy, *always* distrust it first.

***

## On-chain scams

These are scams that happen on-chain. Either because of some backdoor, because some information was presented to the user which led to some reaction, because the user may have been mislead by some data presented by a block explorer, because the intentions of a developer were inherently evil and investors or depositors happen to have lost their money, etc...

On-chain scams differ from off-chain scams in the sense that these happen exclusively with on-chain data or through a transaction made either by a user (motivated by on-chain information obtained by the user) or by a malicious third party.

### Scam tokens sent to your address

Block explorers track transactions from all contracts and addresses on the blockchain. Some information in contract storage slots like the balance of specific tokens in your possession are data that must be read directly from the contract storage. Block explorers make this kind of information very readable and user friendly by storing and listing incoming and outgoing transactions and transaction history for every address and contract.

As the blockchain is a living decentralized network, some malicious actors may deploy token contracts with the intent of scamming users exploring their blockchain addresses and their incoming and outgoing token transactions on a block explorer.

There's plenty of examples and techniques scammers will use. I don't intend to list *all of those that exist* because this would quickly become outdated. Some of those are:

#### 'From' label spoofing

Sometimes, you'll receive tokens from addresses after you've interacted with an application or contract. For example, if you deposit assets on a Uniswap pool, as of V3, you'll receive a transaction from the Uniswap V3 Positions NFT contract that delivers an NFT to your address that represents the liquidity you've deposited to the pool. Some scammers have been successful in creating transactions where the 'From' field of the transaction shows the exact same name as the original, real contract, but instead delivers a set of tokens to your address or at times, nothing but a simple interaction with your wallet so that it shows up in your block explorer page under any of the tabs.

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/fake_from_token.png)

In this case, the sender looks to be the UNI token contract of the Uniswap Protocol, but this is NOT the case, the sender has been spoofed.

#### Fake tokens

In that same transaction shown above, the spoofed sender is transferring 7000 units of an unspecified token labeled as ERC-20 TOKEN*. Which has been labeled as a scam on Etherscan:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/fake_erc20_marked.png)

Etherscan or any other block explorers will *take a while* to mark these as scams. The general advice here is to **never** interact with these tokens, pretend they're not there.

The way token balanced work on Ethereum and other EVM chains is by creating a simple table of balances in key-value pairs. The key being any address that has previously interacted with a token and the value the amount of that token that the address has at any given moment, starting with 0.

When a random token is transferred to you, Etherscan (as a block explorer) will index and identify what any transaction is and show that you have amounts of tokens that you may or may not recognize. In this case, the individual or entity which received those 7000 tokens now has an entry in that table contained in the storage of that unspecified ERC-20 TOKEN* token contract which shows their address and the tokens, in a similar way to this:

```json
{
    "0xe357b511804f52e5ad27e8a8e09f4884e893bf99":"7000"
}
```

Block explorers will show you these amounts and historical transactions, but if the block explorer **never showed you this**, you would have **never known** about it, therefore, it's important to ignore anything like this that seems unusual, out of place, or unexpected.

Interacting with a fake token contract like this one can lead to significant consequences, like lots of unintended token approvals or even transferring funds. **Always** ignore these fake tokens.

#### Fake tokens appearing on your wallet UI

Sometimes, fake tokens will appear on your wallet UI as tokens that you 'have'. It is important to ignore these or hide them if possible, just never interact with them. Don't try to transfer them to any address, don't try to do anything with them, they're not going to do anything to your account, all the wallet software is doing is reading data from block explorers like Etherscan and displaying these.

As shown before on [Chapter 3](#multi-signature-wallets), in the section explaining multi-signature wallets, one such example of wallet software (in this case the web UI of the Gnosis Safe app), scam tokens will be shown:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/scam_tokens_gnosis_safe.png)

Ignore them and **never** visit the websites they suggest you to visit. They are scam sites that will prompt you to sign malicious transactions.

#### Scam NFT collections

If you happen to own NFTs, sometimes some addresses will be targeted by attackers by sending them scam NFTs. These scam NFTs will operate in much of the same way scam ERC20 tokens operate. You may unexpectedly receive them, see them on a site like OpenSea, LooksRare, Blur.io, or any other NFT marketplace or even on a block explorer. The NFT media content they contain *may even be nice art!*, but it is almost always guaranteed to scam you and you should **always** expect it to attempt to scam you.

The moment you sign a transaction to try to move these, they will unintendedly approve other tokens to be moved out of your wallet or perform actions that you wouldn't want to willingly perform. **Ignore them, always**, no one is going to gift you something valuable randomly.

One such example of a collection with a _ton_ of copies of it out there that scammers have made a concerted effort to target users with is The Beeings collection. Here we can see the *official* Beeings collection on Opensea:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/beeings_official_OS.png)

And the result of searching it on LooksRare with also its respective checkmark:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/beeings_official.png)

But, there's also *plenty* of other collections that are clones/fakes of this one, where interacting with any of the NFTs in the collection could lead to a significant monetary loss:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/scam_nft_beeings.png)

Your account or anyone's account could be specifically selected by the creators of one such fake collections and you could end up receiving NFTs that *may look valuable* on platforms like OpenSea or LooksRare, but are *absolutely worthless*.

### Scam token trading pairs in DEXes

Since anyone can deploy a token contract in a decentralized network, anyone can also create tokens _identical_ to those that already exist. It's very common that people who deploy scam token contracts sometimes try to emulate tokens that are just being announced, just coming out or at times that haven't even been announced.

Some emulate tokens that already exist, like in the case of SHIB (a well know dog meme token), one can search for trading pairs on a meta dex search engine like dextools and find *tens of thousands* of trading pairs across many chains:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/shib_search_pairs.png)

Many of these will be real trading pairs, some with lots of liquidity and others with not as much, but many will also be scams. Tokens created with the purpose of making believe that you are buying the real token but instead you're buying a worthless copy of it that the original pool deployer can just take the liquidity away at any time, since in order to buy it, you need to put in a valuable token (maybe ETH or BNB) to take out a portion of the worthless token, which will almost surely always have the same ticker as the original token, in this case SHIB.

Sometimes scam tokens are used to create trading pairs of tokens that just haven't even been released yet. For example, at the time of writing this section (16th January, 2023), the Arbitrum optimistic rollup on Ethereum hasn't released an official token. Since Arbitrum doesn't have an official token but lots of people expect it to have one at some point, opportunistic scammers often create trading pairs with fake ARBI or ARBITRUM tokens in order to fool real traders into believing the token has already been deployed and to show that they're buying the real deal. People may want to buy these tokens because they might have a big run up from the moment of release, as there's lots of hype surrounding token deployments like this one. It's almost certain that every single trading pair created with the name ARBI or ARBITRUM at the moment is a scam pair:

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/arbi_fake_tokens.png)

![](/articles/identifying_and_avoiding_crypto_scams/images/fake_tokens/arbitrum_fake_tokens.png)

Most as you can see also have extremely thin liquidity (which causes very high slippage) or simply no liquidity at all. Many others have been created on Arbitrum with the purpose of fooling users, since it's quite cheap to deploy there too.

***

## Conclusion

This chapter documents many of the scams I've seen while cruising through the cryptocurrency space. It is not a exhaustive list, but it is a good starting point to learn about the most common scams and how to avoid them. Many scams like those I list and explain here might change shape or form in the coming years, they'll renew themselves and they'll evolve, they'll become more sophisticated and much better at targeting victims. Therefore you should remain very aware, and always ask yourself a few questions:

- **Is this too good to be true?**: if it sounds too good to be true, it's probably a scam.

- **Is this (request, message, event, post, tweet...) unexpected, unusual or scary?**: anything that has any of these qualities is probably a scam.

- **Is this sketchy looking, do I have any reason to doubt it?**: if it does look sketchy or you do have reasons doubt it, it's probably a scam.

- **Am I being marketed this (thing) too aggressively?**: if you are, then there's probably a reason for it, and it's probably a scam.

- **Is this a scam?**: if you have to ask yourself this question, then it's probably a scam.

The most important quality of someone who self-custodies their funds and uses the blockchain as their bank, is the ability to remain vigilant and doubt everything before they act. Self-custody teaches you how to really protect yourself against scams. It's not easy, but it's worth it.

***

### Acknowledgements

Thanks to all the kind members of ethfinance which helped me collect scam examples, among which I want to highlight:

* u/alexiskef for many of the discord DM examples

* u/busterrulezzz for his writing tips and his narration of the experience he had on a crypto startup that turned out to be a scam narrated on r/CryptoCurrency

* The ethfinance mods for allowing me to ask for examples of scams publicly on the daily thread

* u/pocketwailord, u/excellentpantschoice, u/datacruncha, u/actualbadger, u/sinnU2s, u/lawfultots, u/RedPillInjester, for all the examples they sent, all quite distinct from one another which was really helpful

* u/nixorokish for offering me to collaborate on a different project on a scam-related diagram and for tagging me in a few other examples on the daily threads

***

### Image sources

Two images were sourced from external websites:

- [The main article image (finger made of bits on bear trap)](https://www.morningstar.ca/ca/news/221284/five-common-scams-(and-the-blindspots-that-make-you-vulnerable).aspx)

- [Jenga tower image](https://www.wired.com/story/a-robot-teaches-itself-to-play-jenga/)
