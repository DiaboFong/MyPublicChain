# 从0到1简易区块链开发手册V0.4

**Author: brucefeng**

**Email: brucefeng@brucefeng.com**

**编程语言: Golang**

---

## 前言

这是我这段时间学习区块链开发以来打造的第一个区块链平台，之所以叫做简易区块链，是因为它确实比较简易，仅仅是实现了底层的一些功能，不足以作为一个真正的公链使用，但通过学习，我们能够通过代码更加理解比特币白皮书中描述的各种比特币原理，区块链世界，无论是研究理论的，还是实战开发的，甚至炒币玩资本的，都离不开比特币的影子，区块链技术毕竟是从比特币中剥离抽象而来，所以，作为一个技术人员，无论是研究以太坊，超级账本，甚至是各种公链，包括某些山寨公链，都需要先去理解比特币原理，而对于开发者而言，理解原理最好的方式就是将其通过代码实现，当然，我们这里实现的一些原理只是应用实战范围之内可以实现的，比如椭圆加密算法，我们要实现的只是使用椭圆加密去实现某些加密功能，而非用代码去实现一个完整的椭圆加密代码库，这个不再本文的讨论范围内，所以本文面向的群体是:

* 对比特币原理不了解，但没时间看太多的资料文献的初学者
* 对比特币原理有所了解，但是停留在理论阶段的研究者
* 没有对比特币进行研究，想直接研究以太坊，超级账本的实战者(大神除外)
* 对Golang熟悉，但是不知道如何入手区块链的开发者或者是像我一样的运维 :-)

本文中，我们先通过命令行的方式演示区块链的工作流程以及相关原理，涉及到比较重要的内容，比如Sha256哈希，椭圆加密，Base58编码等内容，我会根据时间以及后期的工作情况进行适当调整，这注定是一个短期内没有结尾的故事。

> 为表尊敬，写在前面，建议先阅读该文档
>
> 本文的学习资料来自这位**liuxhengxu**前辈翻译的资料
>
> [1]: https://github.com/liuchengxu/blockchain-tutorial	"翻译版本"
> [2]: https://github.com/Jeiwan/blockchain_go	"原始版本"
>
> 能将资料翻译得如此完美，相比其技术功能也是相当深厚的，感谢分享

**建议大家可以先看该资料后再来看我的这系列文章**，否则可能会有一些难度，由于该资料是通过循序渐进的方式进行版本迭代，慢慢引导开发者不断对原有的代码进行优化，拓展，非常认真并细心，希望大家时间充裕的时候以及对某些本文并未写清楚的地方，强烈建议阅读该资料。

本文在此基础上进行了一些修改(谈不上改进)，我摒弃一些过于基础的以及后期需要大量重构的代码，直接通过该项目的执行流程进行代码分析，这样可以稍微节省一些大家的时间，把有限的精力放在对业务有更大提升的技术研究上。

## 一. 功能描述



```
Usage:
        createwallet
                        -- 创建钱包
        getaddresslists
                        -- 获取所有的钱包地址
        createblockchain -address address
                        -- 创建创世区块
        send -from SourceAddress -to DestAddress -amount Amount
                        -- 转账交易
        printchain
                        -- 打印区块
        getbalance -address address
                        -- 查询余额
```

本文围绕着几个功能进行讲解

* 创建钱包

  通过椭圆加密算法创建钱包地址

* 获取钱包地址

  获取区块链中所有的钱包地址

* 创建创世区块

  实现创世区块的创建，并生成区块链

* 实现转账交易

  通过转账交易，生成区块，并存入区块链

* 打印区块

  打印出所有的区块信息，实现转账交易的溯源

* 查询余额

  查询出对应钱包的余额状态

随着代码的不断完善，我们将会对以上进行改进，并提供更多的功能点进行分析探讨，我们先通过下图简单演示一下如上功能

![1.功能列表](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/动态图/1.功能列表.gif)



关于效果图，大家先大致看下即可，不需要刻意研究，在后期的课程中都会涉及。



##  二.  实现命令行功能

![1.命令执行](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/1.命令执行.png)

### 1.定义结构体

定义一个空结构体

```go
type CLI struct {
	
}
```

### 2.结构体方法

> 重要！初学者必看 
>
> 这里提到的结构体方法并不是真正实现功能的方法，而是命令行对象的方法，这些方法中会调用实际的功能对象方法进行功能实现,在本章节中，创建结构体方法即可，功能代码可以为空，如:

```
例子:
func (cli *CLI) CreateWallet() {
}
func (cli *CLI) GetAddressLists() {
}
.....
```

其他的可以在后期逐步实现，为了让有基础的同学对项目整体提前有些印象，所以，代码内容我直接复制粘贴进来，不做删减，在后期的内容中，会逐步涉及到每个调用的对象方法或者函数的作用。

#### 2.1 创建钱包

```go
func (cli *CLI) CreateWallet() {
	_, wallets := GetWallets() //获取钱包集合对象
	wallets.CreateNewWallets() //创建钱包集合
	
}
```

#### 2.2 获取钱包地址

```go
func (cli *CLI) GetAddressLists() {
	fmt.Println("钱包地址列表为:")
	    //获取钱包的集合，遍历，依次输出
	_, wallets := GetWallets() //获取钱包集合对象
	for address, _ := range wallets.WalletMap { 
		fmt.Printf("\t%s\n", address)
	}
}
```

#### 2.3 创建创世区块

```go
func (cli *CLI) CreateBlockChain(address string) {
	CreateBlockChainWithGenesisBlock(address)
	bc :=GetBlockChainObject()
	if bc == nil{
		fmt.Println("没有数据库")
		os.Exit(1)
	}
	defer bc.DB.Close()
	utxoSet:=&UTXOSet{bc}
	utxoSet.ResetUTXOSet()
}
```

#### 2.4 创建转账交易

```go
func (cli *CLI) Send(from, to, amount []string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法转账。。")
		os.Exit(1)
	}
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount)
	//添加更新
	utsoSet :=&UTXOSet{bc}
	utsoSet.Update()
}
```

#### 2.5 查询余额

```go
func (cli *CLI) GetBalance(address string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法查询。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	//total := bc.GetBalance(address,[]*Transaction{})
	utxoSet :=&UTXOSet{bc}
	total:=utxoSet.GetBalance(address)


	fmt.Printf("%s,余额是：%d\n", address, total)
}
```

#### 2.6 打印区块

```go
func (cli *CLI) PrintChains() {
	//cli.BlockChain.PrintChains()
	bc := GetBlockChainObject() //bc{Tip,DB}
	if bc == nil {
		fmt.Println("没有BlockChain，无法打印任何数据。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.PrintChains()
}
```

### 3. 相关函数

#### 3.1 判断参数是否合法

```go
func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
```

判断终端命令是否有参数输入，如果没有参数，则提示程序使用说明，并退出程序

#### 3.2 程序使用说明

```go
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatewallet\n\t\t\t-- 创建钱包")
	fmt.Println("\tgetaddresslists\n\t\t\t-- 获取所有的钱包地址")
	fmt.Println("\tcreateblockchain -address address\n\t\t\t-- 创建创世区块")
	fmt.Println("\tsend -from SourceAddress -to DestAddress -amount Amount\n\t\t\t-- 转账交易")
	fmt.Println("\tprintchain\n\t\t\t-- 打印区块")
	fmt.Println("\tgetbalance -address address\n\t\t\t-- 查询余额")
}
```

#### 3.3 JSON解析的函数

```go
func JSONToArray(jsonString string) []string {
	var arr [] string
	err := json.Unmarshal([]byte(jsonString), &arr)
	if err != nil {
		log.Panic(err)
	}
	return arr
}
```

通过该函数将JSON字符串格式转成字符串数组，用于在多笔转账交易中实现同时多个账户进行两两转账的功能。

#### 3.4 校验地址是否有效

```go
func  IsValidAddress(address []byte) bool {

	//step1：Base58解码
	//version+pubkeyHash+checksum
	full_payload := Base58Decode(address)

	//step2：获取地址中携带的checkSUm
	checkSumBytes := full_payload[len(full_payload)-addressCheckSumLen:]

	versioned_payload := full_payload[:len(full_payload)-addressCheckSumLen]

	//step3：versioned_payload，生成一次校验码
	checkSumBytes2 := CheckSum(versioned_payload)

	//step4：比较checkSumBytes，checkSumBytes2
	return bytes.Compare(checkSumBytes, checkSumBytes2) == 0

}
```

以下三个功能实现之前需要先调用该函数进行地址校验

* 创建创世区块
* 转账交易
* 查询余额

### 4.命令行主要方法Run

```
Usage:
        createwallet
                        -- 创建钱包
        getaddresslists
                        -- 获取所有的钱包地址
        createblockchain -address address
                        -- 创建创世区块
        send -from SourceAddress -to DestAddress -amount Amount
                        -- 转账交易
        printchain
                        -- 打印区块
        getbalance -address address
                        -- 查询余额
```

我们将如上功能展示的实现功能写在Run方法中，实现命令行功能的关键是了解**os.Args**与**flag**

> 关于这两个功能，此处不再赘述，否则篇幅会无限臃肿。

代码块均在方法体Run中，下文将分步骤对代码实现进行体现

```go
func (cli *CLI) Run() {
}
```



#### 4.1 判断命令行参数是否合法

```go
isValidArgs()
```

#### 4.2 创建flagset命令对象

```go
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	getAddresslistsCmd := flag.NewFlagSet("getaddresslists", flag.ExitOnError)
	CreateBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	testMethodCmd := flag.NewFlagSet("test", flag.ExitOnError)
```

如上，通过flag.NewFlagSet方法创建命令对象，如createwallet,getaddresslists,createblockchain等命令对象

固定用法，掌握即可。

#### 4.3 设置命令后的参数对象

```go
    flagCreateBlockChainData := CreateBlockChainCmd.String("address", "GenesisBlock", "创世区块的信息")
	flagSendFromData := sendCmd.String("from", "", "转账源地址")
	flagSendToData := sendCmd.String("to", "", "转账目标地址")
	flagSendAmountData := sendCmd.String("amount", "", "转账金额")
	flagGetBalanceData := getBalanceCmd.String("address", "", "要查询余额的账户")
```

通过命令对象的String方法为命令后的参数对象

* createblockchain命令后的参数对象：address
* send命令后的参数对象: from | to |amount
* getbalance命令后的参数对象: address 

其中createwallet，getaddresslists，printchain命令没有参数对象。

#### 4.4 解析命令对象

```go
switch os.Args[1] {
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getaddresslists":
		err := getAddresslistsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := CreateBlockChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "test":
		err := testMethodCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)

	}
```

匹配对应的命令，用命令对象的Parse方法对**os.Args[2:]**进行解析。

#### 4.5 执行对应功能

```go
	//4.1 创建钱包--->交易地址
	if createWalletCmd.Parsed() {
		cli.CreateWallet()
	}
	//4.2 获取钱包地址
	if getAddresslistsCmd.Parsed() {
		cli.GetAddressLists()
	}
	//4.3 创建创世区块
	if CreateBlockChainCmd.Parsed() {
		if !IsValidAddress([]byte(*flagCreateBlockChainData)) {
			fmt.Println("地址无效，无法创建创世前区块")
			printUsage()
			os.Exit(1)
		}
		cli.CreateBlockChain(*flagCreateBlockChainData)
	}
	//4.4 转账交易
	if sendCmd.Parsed() {
		if *flagSendFromData == "" || *flagSendToData == "" || *flagSendAmountData == "" {
			fmt.Println("转账信息有误")
			printUsage()
			os.Exit(1)
		}
		//添加区块
		from := JSONToArray(*flagSendFromData)     //[]string
		to := JSONToArray(*flagSendToData)         //[]string
		amount := JSONToArray(*flagSendAmountData) //[]string
		for i := 0; i < len(from); i++ {
			if !IsValidAddress([]byte(from[i])) || !IsValidAddress([]byte(to[i])) {
				fmt.Println("地址无效，无法转账")
				printUsage()
				os.Exit(1)
			}
		}

		cli.Send(from, to, amount)
	}
	//4.5 查询余额
	if getBalanceCmd.Parsed() {
		if !IsValidAddress([]byte(*flagGetBalanceData)) {
			fmt.Println("查询地址有误")
			printUsage()
			os.Exit(1)
		}
		cli.GetBalance(*flagGetBalanceData)
	}

	//4.6 打印区块信息
	if printChainCmd.Parsed() {
		cli.PrintChains()
	}
```



### 5. 测试代码

> 在main.go中中添加测试代码

```go
package main
func main() {
	cli:=BLC.CLI{}
	cli.Run()
}
```

编译运行

```
$ go build -o mybtc main.go
```

> 测试思路
>
> 1. 查看命令行列表是否可以正常显示
> 2. 输入非法字符查看是否有错误提示
>
> 业务功能此处暂未实现，测试时忽略。



**下一篇文章**将介绍如何实现钱包/地址的生成功能，即**创建钱包**。



## 三. 创建钱包

### 1.概念

创建钱包其实就是创建比特币地址，在比特币世界中，没有账户概念，不需要也不会在任何地方存储个人数据（比如姓名，身份证件号码等）。但是，我们总要有某种途径识别出你是交易输出的所有者（也就是说，你拥有在这些输出上锁定的币），这就是比特币地址（address）需要完成的使命。

关于钱包这个概念，我个人觉得imtoken在用户引导那部分写得很清楚,此处将链接给到大家,有兴趣的可以去看看

https://www.cnblogs.com/fangbei/p/imToken-clearance.html

我们来看一下一个真实的比特币账户，1FSzfZ27CVTkfNw6TWxnHPaRLRCgpWvbFC ，比特币地址是完全公开的，如果你想要给某个人发送币，只需要知道他的地址就可以了。但是，地址（尽管地址也是独一无二的）并不是用来证明你是一个“钱包”所有者的信物。实际上，所谓的地址，只不过是将公钥表示成人类可读的形式而已，因为原生的公钥人类很难阅读。在比特币中，你的身份（identity）就是一对（或者多对）保存在你的电脑（或者你能够获取到的地方）上的公钥（public key）和私钥（private key）。比特币基于一些加密算法的组合来创建这些密钥，并且保证了在这个世界上没有其他人能够取走你的币，除非拿到你的密钥。

关于如何创建一个钱包以及钱包集合，通过下图进行简单展示

![2.0.创建钱包All](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/2.0.创建钱包All.png)

图 创建钱包与钱包集合



![2.1.创建钱包Wallet](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/2.1.创建钱包Wallet.png)

图 创建钱包wallet

![2.2.创建钱包集合Wallets](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/2.2.创建钱包集合Wallets.png)

图 创建钱包集合





### 2. 定义钱包结构体

```go
type Wallet struct {
	//1.私钥
	PrivateKey ecdsa.PrivateKey
	//2.公钥
	PublickKey []byte //原始公钥
}
```

定义钱包Wallet的属性为私钥：PrivateKey，类型为系统内置的结构体对象ecdsa.PrivateKey，公钥：PublickKey，类型为字节数组

### 3. 生成钱包地址

![address1](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/1.address1.png)



图 	从私钥到生成钱包地址的过程图		


		

####  3.1 通过椭圆曲线算法产生密钥对

```go
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	//椭圆加密
	curve := elliptic.P256() //根据椭圆加密算法，得到一个椭圆曲线值
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader) //*Private
	if err != nil {
		log.Panic(err)
	}

	//通过私钥生成原始公钥
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	return *privateKey, publicKey
}

```

椭圆曲线加密：(ECC：ellipse curve Cryptography)，非对称加密

* 根据椭圆曲线算法，产生随机私钥
* 根据私钥，产生公钥

#### 3.2 创建钱包对象

```go
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	return &Wallet{privateKey, publicKey}
}
```

通过newKeyPair函数将返回的私钥与公钥生成钱包对象Wallet



#### 3.3 定义常量值

```go
const version = byte(0x00)
const addressCheckSumLen = 4
```

* version: 版本前缀，比特币中固定为0

* addressCheckSumLen: 用于获取校验码的长度变量，取添加版本+数据进行两次SHA256之后的前4个字节

#### 3.4 根据公钥获取地址



![address2](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/2.address2.png)															



								图 从**公钥**到生成钱包地址的过程图

```go
func PubKeyHash(publickKey []byte) []byte {
	//1.sha256
	hasher := sha256.New()
	hasher.Write(publickKey)
	hash1 := hasher.Sum(nil)

	//2.ripemd160
	hasher2 := ripemd160.New()
	hasher2.Write(hash1)
	hash2 := hasher2.Sum(nil)

	//3.返回公钥哈希
	return hash2
}
```

通过公钥生成公钥哈希的步骤已完成。

```go
func GetAddressByPubKeyHash(pubKeyHash []byte) []byte {
	//添加版本号：
	versioned_payload := append([]byte{version}, pubKeyHash...)

	//根据versioned_payload-->两次sha256,取前4位，得到checkSum
	checkSumBytes := CheckSum(versioned_payload)

	//拼接全部数据
	full_payload := append(versioned_payload, checkSumBytes...)

	//Base58编码
	address := Base58Encode(full_payload)
	return address
}
```

> 相关函数如下

* 生成校验码

```go
func CheckSum(payload [] byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:addressCheckSumLen]
}
```

通过两次sha256哈希得到校验码，返回校验码前四位

* 字节数组转Base58加密

  ```go
  var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
  func Base58Encode(input []byte)[]byte{
  	var result [] byte
  	x := big.NewInt(0).SetBytes(input)
  
  	base :=big.NewInt(int64(len(b58Alphabet)))
  	zero:=big.NewInt(0)
  	mod:= &big.Int{}
  	for x.Cmp(zero) !=0{
  		x.DivMod(x,base,mod)
  		result = append(result,b58Alphabet[mod.Int64()])
  	}
  	ReverseBytes(result)
  	for b:=range input{
  		if b == 0x00{
  			result = append([]byte{b58Alphabet[0]},result...)
  		}else {
  			break
  		}
  	}
  
  	return result
  
  }
  ```

以上功能函数定义好之后，定义Wallet的方法GetAddress返回钱包address

```go
func (w *Wallet) GetAddress() []byte {
	pubKeyHash := PubKeyHash(w.PublickKey)
	address := GetAddressByPubKeyHash(pubKeyHash)
	return address
}

```

至此，我们已经能够生成一个比特币地址了，可以通过https://www.blockchain.com/explorer进行钱包地址查看余额。

### 4.定义钱包集合结构体

```go
type Wallets struct {
	WalletMap map[string]*Wallet
}
```

定义钱包集合结构体Wallets，属性为WalletMap,类型为Wallet集合

### 5.创建钱包集合

```go
func (ws *Wallets) CreateNewWallets() {
	wallet := NewWallet()
	var address []byte
	address = wallet.GetAddress()
	fmt.Printf("创建的钱包地址：%s\n", address)
	ws.WalletMap[string(address)] = wallet
	//将钱包集合存入到本地文件中
	ws.SaveFile()
}
```

* 创建一个钱包对象
* 通过GetAddress获取钱包对象的地址
* 将钱包地址作为钱包集合的key,钱包对象作为value存储至钱包集合中
* 通过SaveFile将钱包集合存入到本地文件中

#### 5.1 定义常量存储钱包数据

```go
const walletsFile = "Wallets.dat" //存储钱包数据的本地文件名
```

#### 5.2 本地化存储钱包对象

```go
func (ws *Wallets) SaveFile() {
	//1.将ws对象的数据--->byte[]
	var buf bytes.Buffer
	//序列化的过程中：被序列化的对象中包含了接口，那么该接口需要注册
	gob.Register(elliptic.P256()) //Curve
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}
	wsBytes := buf.Bytes()

	//2.将数据存储到文件中
	err = ioutil.WriteFile(walletsFile, wsBytes, 0644)
	if err != nil {
		log.Panic(err)
	}
}
```

### 6.获取钱包集合

此处我们提供一个函数，用户获取钱包集合

* 读取本地的钱包文件，如果文件存在，直接获取
* 如果文件不存在，创建并返回一个空的钱包对象

```go
func GetWallets()  *Wallets {
	//钱包文件不存在
	if _, err := os.Stat(walletsFile); os.IsNotExist(err) {
		fmt.Println("区块链钱包不存在")
		//创建钱包集合
		wallets := &Wallets{}
		wallets.WalletMap = make(map[string]*Wallet)
		return wallets
	}

	//钱包文件存在
	//读取本地的钱包文件中的数据
	wsBytes, err := ioutil.ReadFile(walletsFile)
	if err != nil {
		log.Panic(err)
	}
	gob.Register(elliptic.P256()) //Curve
	//将数据反序列化变成钱包集合对象
	var wallets Wallets
	reader := bytes.NewReader(wsBytes)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}
	return  &wallets
}
```

### 7.命令行中调用

#### 7.1 创建钱包

回到上一章节(二.实现命令行功能-2.1创建钱包)的命令行功能

```go
func (cli *CLI) GetAddressLists() {
	fmt.Println("钱包地址列表为:")
	    //获取钱包的集合，遍历，依次输出
	_, wallets := GetWallets() //获取钱包集合对象
	for address, _ := range wallets.WalletMap { 
		fmt.Printf("\t%s\n", address)
	}
}
```

此时进行编译运行

```
$ go build -o mybtc main.go
```

```
$ ./mybtc createwallet //创建第一个钱包
$ ./mybtc createwallet //创建第二个钱包
$ ./mybtc createwallet //创建第三个钱包
```

返回的结果:

```
创建的钱包地址：14A1b3Lp3hL5B7vZvT2UWk1W78m2Kh8MUB
创建的钱包地址：1G3SkYAJdWy5pd1hFpcciUoJi8zy8PdV11
创建的钱包地址：1AA2fyYdXCQMwLMu5NBvq7Fb9UiHqg2cQV
```

#### 7.2 获取钱包地址 

回到上一章节(二.实现命令行功能-2.2 获取钱包地址)的命令行功能

```go
func (cli *CLI) GetAddressLists() {
	fmt.Println("钱包地址列表为:")
	//获取钱包的集合，遍历，依次输出
	wallets := GetWallets()
	for address, _ := range wallets.WalletMap {

		fmt.Printf("\t%s\n", address)
	}
}
```



```
$ ./mybtc getaddresslists
```

返回的结果

```go
钱包地址列表为:
        1AA2fyYdXCQMwLMu5NBvq7Fb9UiHqg2cQV
        14A1b3Lp3hL5B7vZvT2UWk1W78m2Kh8MUB
        1G3SkYAJdWy5pd1hFpcciUoJi8zy8PdV11
```

上面我们提到生成的比特币地址可以通过https://www.blockchain.com/explorer进行钱包地址查看余额，现在我们来进行简单的查看验证，查看该地址:**1AA2fyYdXCQMwLMu5NBvq7Fb9UiHqg2cQV**

![3.在线查看比特币地址余额](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/3.在线查看比特币地址余额.png)

							图  通过搜索框进行地址搜索

![4.在线查看比特币地址余额-1](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/4.在线查看比特币地址余额-1.png)

图  钱包地址详情

> 如果修改钱包地址的某个字符，如将随后的V改为X

1AA2fyYdXCQMwLMu5NBvq7Fb9UiHqg2cQV   === > 1AA2fyYdXCQMwLMu5NBvq7Fb9UiHqg2cQ**X**

![image-20180819231300396](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/5.非法地址.png)



## 四. 持久化存储

### 1.BoltDB简介

Bolt是一个纯粹Key/Value模型的程序。该项目的目标是为不需要完整数据库服务器（如Postgres或MySQL）的项目提供一个简单，快速，可靠的数据库。

BoltDB只需要将其链接到你的应用程序代码中即可使用BoltDB提供的API来高效的存取数据。而且BoltDB支持完全可序列化的ACID事务，让应用程序可以更简单的处理复杂操作。

其源码地址为:https://github.com/boltdb/bolt



### 2.BoltDB特性

BoltDB设计源于LMDB，具有以下特点：

- 使用Go语言编写
- 不需要服务器即可运行
- 支持数据结构

- 直接使用API存取数据，没有查询语句；
- 支持完全可序列化的ACID事务，这个特性比LevelDB强；
- 数据保存在内存映射的文件里。没有wal、线程压缩和垃圾回收；
- 通过COW技术，可实现无锁的读写并发，但是无法实现无锁的写写并发，这就注定了读性能超高，但写性能一般，适合与读多写少的场景。

BoltDB是一个Key/Value（键/值）存储，这意味着没有像SQL RDBMS（MySQL，PostgreSQL等）中的表，没有行，没有列。相反，数据作为键值对存储（如在Golang Maps中）。键值对存储在Buckets中，它们旨在对相似的对进行分组（这与RDBMS中的表类似）。因此，为了获得Value(值)，需要知道该Value所在的桶和钥匙。



### 3.BoltDB简单使用

```go
//通过go get下载并import 
import "github.com/boltdb/bolt" 
```

#### 3.1 打开或创建数据库

```go
db, err := bolt.Open("my.db", 0600, nil)
if err != nil {
   log.Fatal(err)
}
defer db.Close()
```

- 执行注意点

如果通过goland程序运行创建的my.db会保存在

```
GOPATH /src/Project目录下
如果通过go build main.go ; ./main 执行生成的my.db，会保存在当前目录GOPATH /src/Project/package下
```

#### 3.2 数据库操作

###### (1) 创建数据库表与数据写入操作

```go
//1. 调用Update方法进行数据的写入
err = db.Update(func(tx *bolt.Tx) error {
//2.通过CreateBucket()方法创建BlockBucket(表)，初次使用创建
   b, err := tx.CreateBucket([]byte("BlockBucket"))
   if err != nil {
      return fmt.Errorf("Create bucket :%s", err)
   }

//3.通过Put()方法往表里面存储一条数据(key,value)，注意类型必须为[]byte
   if b != nil {
      err := b.Put([]byte("l"), []byte("Send $100 TO Bruce"))
      if err != nil {
         log.Panic("数据存储失败..")
      }
   }

   return nil
})

//数据Update失败，退出程序
if err != nil {
   log.Panic(err)
}
```

###### (2) 数据写入

```go
//1.打开数据库
db, err := bolt.Open("my.db", 0600, nil)
if err != nil {
   log.Fatal(err)
}
defer db.Close()

err = db.Update(func(tx *bolt.Tx) error {

//2.通过Bucket()方法打开BlockBucket表
   b := tx.Bucket([]byte("BlockBucket"))
//3.通过Put()方法往表里面存储数据
   if b != nil {
      err := b.Put([]byte("l"), []byte("Send $200 TO Fengyingcong"))
      err =  b.Put([]byte("ll"), []byte("Send $100 TO Bruce"))
      if err != nil {
         log.Panic("数据存储失败..")
      }
   }

   return nil
})
//更新失败
if err != nil {
   log.Panic(err)
}
```

###### (3) 数据读取

```go
//1.打开数据库
db, err := bolt.Open("my.db", 0600, nil)
if err != nil {
   log.Fatal(err)
}
defer db.Close()

//2.通过View方法获取数据
err = db.View(func(tx *bolt.Tx) error {

//3.打开BlockBucket表，获取表对象

   b := tx.Bucket([]byte("BlockBucket"))
    
//4.Get()方法通过key读取value
   if b != nil {
      data := b.Get([]byte("l"))
      fmt.Printf("%s\n", data)
      data = b.Get([]byte("ll"))
      fmt.Printf("%s\n", data)
   }

   return nil
})

if err != nil {
   log.Panic(err)
}
```

### 4.通过BoltDB存储区块

该代码包含对BoltDB的数据库创建，表创建，区块添加，区块查询操作

```go
//1.创建一个区块对象block
block := BLC.NewBlock("Send $500 to Tom", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

//2. 打印区块对象相关信息
fmt.Printf("区块的Hash信息为:\t%x\n", block.Hash)
fmt.Printf("区块的数据信息为:\t%v\n", string(block.Data))
fmt.Printf("区块的随机数为:\t%d\n", block.Nonce)

//3. 打开数据库
db, err := bolt.Open("my.db", 0600, nil)
if err != nil {
   log.Fatal(err)
}
defer db.Close()

//4. 更新数据
err = db.Update(func(tx *bolt.Tx) error {

//4.1 打开BlockBucket表对象
   b := tx.Bucket([]byte("blocks"))
//4.2 如果表对象不存在，创建表对象
   if b == nil {
      b, err = tx.CreateBucket([]byte("blocks"))
      if err != nil {
         log.Panic("Block Table Create Failed")
      }
   }

   //4.3 往表里面存储一条数据(key,value)
   err = b.Put([]byte("l"), block.Serialize())
   if err != nil {
      log.Panic("数据存储失败..")
   }
   return nil
})

//更新失败,返回错误
if err != nil {
   log.Panic("数据更新失败")
}

//5. 查看数据
err = db.View(func(tx *bolt.Tx) error {

//5.1打开BlockBucket表对象
   b := tx.Bucket([]byte("blocks"))

   if b != nil {
       //5.2 取出key=“l”对应的value
      blockData := b.Get([]byte("l"))
		//5.3反序列化	
      block := BLC.DeserializeBlock(blockData)
 //6. 打印区块对象相关信息
	fmt.Printf("区块的Hash信息为:\t%x\n", block.Hash)
	fmt.Printf("区块的数据信息为:\t%v\n", string(block.Data))
	fmt.Printf("区块的随机数为:\t%d\n", block.Nonce)
   }

   return nil
})
//数据查看失败
if err != nil {
   log.Panic("数据更新失败")
}
```



## 五. 创建创世区块

![3.0.创建创世区块](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/3.0.创建创世区块.png)





### 1.概念

北京时间2009年1月4日2时15分5秒，比特币的第一个区块诞生了。随着时间往后推移，不断有新的区块被添加到链上，所有后续区块都可以追溯到第一个区块。第一个区块就被人们称为创世区块。

### 2. 工作量证明

在比特币世界中，获取区块记账权的过程称之为挖矿，一个矿工成功后，他会把之前打包好的网络上的交易记录到一页账本上，同步给其他人。因为这个矿工能够最先计算出超难数学题的正确答案，说明这个矿工付出了工作量，是一个有权利记账的人，因此其他人也会同意这一页账单。这种依靠工作量来证明记账权，大家来达成共识的机制叫做“工作量证明”，简而言之结果可以证明你付出了多少工作量。

![6.工作量证明](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/6.工作量证明.jpeg)

#### 2.1 定义结构体

```go
type ProofOfWork struct {
	Block  *Block   //要验证的block
	Target *big.Int //目标hash
}
```

#### 

#### 2.2 创建工作量证明对象

```go
const TargetBit = 16 //目标哈希的0个个数,16,20,24,28
func NewProofOfWork(block *Block) *ProofOfWork {
	//1.创建pow对象
	pow := &ProofOfWork{}
	//2.设置属性值
	pow.Block = block
	target := big.NewInt(1)           // 目标hash，初始值为1
	target.Lsh(target, 256-TargetBit) //左移256-16
	pow.Target = target
	return pow

}
```

我们首先设定一个难度系数值为16，即目标哈希前导0的个数，0的个数越多，挖矿难度越大，此处我们创建一个函数NewProofOfWork用于返回Pow对象。

目标Hash的长度为256**bit**，通过64个16进制byte进行展示,如下所示为前导0为16/4=4的哈希

```
0000c01d342fc51cb030f93979343de70ab771855dd8ca28e6f5888737759747
```

* 通过big.NewInt创建一个BigInt对象target
* 对target进行通过左移(256-TargetBit)位操作

#### 2.3 将int64类型转[]byte

```go
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	//将二进制数据写入w
	//
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	//转为[]byte并返回
	return buff.Bytes()
}
```

> 通过```func Write(w io.Writer, order ByteOrder, data interface{}) error```方法将一个int64的整数转为二进制后，每8bit一个byte，转为[]byte



#### 2.4 拼接区块属性数据

```go
func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join([][]byte{
		IntToHex(pow.Block.Height),
		pow.Block.PrevBlockHash,
		IntToHex(pow.Block.TimeStamp),
		pow.Block.HashTransactions(),
		IntToHex(nonce),
		IntToHex(TargetBit),
	}, []byte{})
	return data

}
```

> 通过bytes.Join方法将区块相关属性进行拼接成字节数组

#### 2.5 "挖矿"方法

```go
func (pow *ProofOfWork) Run() ([]byte, int64) {

	var nonce int64 = 0
	var hash [32]byte
	for {
		//1.根据nonce获取数据
		data := pow.prepareData(nonce)
		//2.生成hash
		hash = sha256.Sum256(data) //[32]byte
		fmt.Printf("\r%d,%x", nonce, hash)
		//3.验证：和目标hash比较
		/*
		func (x *Int) Cmp(y *Int) (r int)
		Cmp compares x and y and returns:

		   -1 if x <  y
			0 if x == y
		   +1 if x >  y
		目的：target > hashInt,成功
		 */
		hashInt := new(big.Int)
		hashInt.SetBytes(hash[:])

		if pow.Target.Cmp(hashInt) == 1 {
			break
		}

		nonce++
	}
	fmt.Println()
	return hash[:], nonce

}

```

> 代码思路

* 设置nonce值：0,1,2.......
* block-->拼接数组，产生hash
* 比较实际hash和pow的目标hash

不断更改nonce的值，计算hash，直到小于目标hash。


#### 2.6 验证区块

```go
func (pow *ProofOfWork) IsValid() bool {
   hashInt := new(big.Int)
   hashInt.SetBytes(pow.Block.Hash)
   return pow.Target.Cmp(hashInt) == 1
}
```

> 判断方式同挖矿中的策略

### 3.区块创建

#### 3.1 定义结构体

```go
type Block struct {
	//字段属性
	//1.高度：区块在区块链中的编号，第一个区块也叫创世区块，一般设定为0
	Height int64
	//2.上一个区块的Hash值
	PrevBlockHash []byte
	//3.数据：Txs，交易数据
	Txs []*Transaction
	//4.时间戳
	TimeStamp int64
	//5.自己的hash
	Hash []byte
	//6.Nonce
	Nonce int64
}
```

关于属性的定义，在代码的注释中比较清晰了，需要提一下的就是创世区块的PrevBlockHash一般设定为0 ，高度也一般设定为0

#### 3.2 创建创世区块

```go
func CreateGenesisBlock(txs []*Transaction) *Block{

	return NewBlock(txs,make([]byte,32,32),0)
}
```

设定创世区块的PrevBlockHash为0，区块高度为0

#### 3.3 序列化区块对象

```go
func (block *Block) Serialize()[]byte{
	//1.创建一个buff
	var buf bytes.Buffer

	//2.创建一个编码器
	encoder:=gob.NewEncoder(&buf)

	//3.编码
	err:=encoder.Encode(block)
	if err != nil{
		log.Panic(err)
	}

	return buf.Bytes()
}
```

> 通过gob库的Encode方法将Block对象序列化成字节数组，用于持久化存储

#### 3.4 字节数组反序列化

```go
func DeserializeBlock(blockBytes [] byte) *Block{
	var block Block
	//1.先创建一个reader
	reader:=bytes.NewReader(blockBytes)
	//2.创建解码器
	decoder:=gob.NewDecoder(reader)
	//3.解码
	err:=decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}

```

> 定义一个函数，用于将[]byte反序列化为block对象

### 4.区块链创建

#### 4.1 定义结构体

```go
type BlockChain struct {
	DB  *bolt.DB //对应的数据库对象
	Tip [] byte  //存储区块中最后一个块的hash值
}
```

定义区块链结构体属性DB用于存储对应的数据库对象，Tip用于存储区块中最后一个块的Hash值

#### 4.2 判断数据库是否存在

```go
const DBName  = "blockchain.db" //数据库的名字
const BlockBucketName = "blocks" //定义bucket
```

定义数据库名字以及定义用于存储区块数据的bucket(表)名

```go
func dbExists() bool {
	if _, err := os.Stat(DBName); os.IsNotExist(err) {  
		return false //表示文件不存在
	}
	return true //表示文件存在
}
```

> 需要注意```IsNotExist```返回```true```,则表示不存在成立，返回值为```true```，则```dbExists```函数的返回值则需要返回```false```，否则，返回```true```

#### 4.3 创建带有创世区块的区块链

```go
func CreateBlockChainWithGenesisBlock(address string) {

	/*
	1.判断数据库如果存在，直接结束方法
	2.数据库不存在，创建创世区块，并存入到数据库中
	 */
	if dbExists() {
		fmt.Println("数据库已经存在，无法创建创世区块")
		return
	}

	//数据库不存在
	fmt.Println("数据库不存在")
	fmt.Println("正在创建创世区块")
	/*
	1.创建创世区块
	2.存入到数据库中
	 */
	//创建一个txs--->CoinBase
	txCoinBase := NewCoinBaseTransaction(address)

	genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		//创世区块序列化后，存入到数据库中
		b, err := tx.CreateBucketIfNotExists([]byte(BlockBucketName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			b.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//return &BlockChain{db, genesisBlock.Hash}
}
```

> 代码分析

(1) 判断数据库是否存在，如果不存在，证明还没有创建创世区块，如果存在，则提示创世区块已存在，直接返回

```go
	if dbExists() {
		fmt.Println("数据库已经存在，无法创建创世区块")
		return
	}
```

(2) 如果数据库不存在，则提示开始调用相关函数跟方法创建创世区块

```go
	fmt.Println("数据库不存在")
	fmt.Println("正在创建创世区块")
```

(3) 创建一个交易数组Txs

> 关于交易这一部分内容，将在后面一个章节中进行详细说明，篇幅会非常长，这也是整个课程体系中最为繁琐，知识点最广的地方，届时慢慢分析

```go
txCoinBase := NewCoinBaseTransaction(address)
```

> 通过函数NewCoinBaseTransaction创建一个CoinBase交易

```go
func NewCoinBaseTransaction(address string) *Transaction {
   txInput := &TxInput{[]byte{}, -1, nil, nil}
   txOutput := NewTxOutput(10, address)
   txCoinBaseTransaction := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
   //设置交易ID
   txCoinBaseTransaction.SetID()
   return txCoinBaseTransaction
}
```

 (4) 生成创世区块

```go
genesisBlock := CreateGenesisBlock([]*Transaction{txCoinBase})
```

(5) 打开/创建数据库

```go
	db, err := bolt.Open(DBName, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
```

通过```bolt.Open```方法打开(如果不存在则创建)数据库文件，注意数据库关闭操作不能少，用defer实现延迟关闭。

(6) 将数据写入数据库

```go
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BlockBucketName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			b.Put([]byte("l"), genesisBlock.Hash)
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
```

通过```db.Upadate```方法进行数据更新操作

* 创建/打开存储区块的Bucket：BlockBucketName
* 将创世区块序列化后存入Bucket中
  * 通过Put方法更新K/V值(Key:区块哈希，Value:区块序列化后的字节数组)
  * 通过Put方法更新Key为“l”的Value为最新区块哈希值，此处即genesisBlock.Hash

### 5.命令行调用

```go
func (cli *CLI) CreateBlockChain(address string) {
	CreateBlockChainWithGenesisBlock(address)
}
```

> 测试命令

```go
$ ./mybtc  createblockchain -address 1DHPNHKfk9uUdog2f2xBvx9dq4NxpF5Q4Q
```

返回结果

```
数据库不存在
正在创建创世区块
32325,00005c7b4246aa88bd1f9664c665d6424d1522f569d981691ac2b5b5d15dd8d9
```

本章节介绍了如何创建一个带有创世区块的区块链，并持久化存储至数据库```blockchain.db```

```
$ ls
BLC             Wallets.dat     blockchain.db   main.go         mybtc
```



## 六.转账交易

创世区块创建完毕之后，按照我们的正常思路，是继续创建新的区块，并加入至区块链中，没错，这确实是学习路线，但是我们首先来了解一个区块是如何生成的，转账交易 ===>打包交易  ===>工作量证明  ===>生成区块

在上文，我们提到了钱包地址这个概念，我们一般可以简单将钱包地址理解为一个银行账户，那么交易也就可以理解为是地址与地址之间的转账过程。

因为这部分内容非常重要，设置可以说交易就是比特币原理的核心，所以，为了保证大家对概念有充分的了解，本章节的理论描述部分此处摘录[liuchengxu](https://liuchengxu.gitbooks.io/blockchain-tutorial/content/part-4/transactions-1.html)中关于对交易的翻译。

### 1.概念

交易（transaction）是比特币的核心所在，而区块链唯一的目的，也正是为了能够安全可靠地存储交易。在区块链中，交易一旦被创建，就没有任何人能够再去修改或是删除它。今天，我们将会开始实现交易。不过，由于交易是很大的话题，我会把它分为两部分来讲：在今天这个部分，我们会实现交易的基本框架。在第二部分，我们会继续讨论它的一些细节。

由于比特币采用的是 UTXO 模型，并非账户模型，并不直接存在“余额”这个概念，余额需要通过遍历整个交易历史得来。

> 关于UTXO模型，这在比特币中也是非常重要的概念模型，务必数量掌握。

点击[此处](https://www.blockchain.com/zh-cn/btc/tx/b6f6b339b546a13822192b06ccbdd817afea5311845f769702ae2912f7d94ab5)查看相关的交易信息

![image-20180821104144181](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/7.查看交易信息.png)

图 交易记录

![image-20180821110337138](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/7.1.交易输入脚本.png)

图 输入脚本

![image-20180821110412726](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/其他图片/7.2.交易输出脚本.png)

关于转账交易涉及到的内容非常多，由于时间原因，目前可能无法做到非常全面的讲解，姑且将自己梳理好能够解释清楚的地方分享出来，由于比特币世界中的交易规则会更加复杂化，所以，希望大家能够通过本章节的阅读，在一定程度上对某些概念有一些初步或者稍微深刻的理解，那么本章节的目的也就达到了，更深的分析笔者将会在后期的工作中根据实际的工作场景进行优化并做相关记录。

### 2.结构体定义

其实再转账交易这个功能里面，涉及了本文所有的结构体对象，由于区块与区块链对象等在上文已经有所提及，这里先列出跟转账交易关系最为密切的一些结构体。



![image-20180822133236933](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/4.0.转账交易-结构体定义.png)

#### 2.1 交易Transaction

```go
type Transaction struct {
	//1.交易ID
	TxID []byte
	//2.输入
	Vins []*TxInput
	//3.输出
	Vouts []*TxOutput
}
```

* **TxID** : 交易ID，一般通过对交易进行哈希后得到
* **Vins**:  交易输入数组
* **Vouts**:交易输出数组

#### 2.2 交易输入TxInput

```
type TxInput struct {
	//1.交易ID：
	TxID []byte
	//2.下标
	Vout int
	//3.数字签名
	Signature []byte 
	//4.原始公钥，钱包里的公钥
	PublicKey []byte 

}
```

* **TxID**: 交易ID，表示该TxInput引用的TxOutput所在的交易ID
* **Vout**:下标，表示该TxInput引用的TxOutput在交易中的位置
* **Signature**：数字签名，用于对引用的TxOutput交易的解锁
* **PublicKey**: 钱包的公钥，原始公钥

#### 2.3 交易输出TxOutput

```go
type TxOutput struct {
	//金额
	Value int64  //金额
	//锁定脚本，也叫输出脚本，公钥，目前先理解为用户名，钥花费这笔前，必须钥先解锁脚本
	//ScriptPubKey string
	PubKeyHash [] byte//公钥哈希
}
```

* **Value** : 金额，转账/找零金额

* **PubKeyHash**:输出脚本，此处为公钥哈希，用于锁定该笔交易输出



#### 2.4 未花费交易输出UTXO

```go
type UTXO struct {
	//1.该output所在的交易id
	TxID []byte
	//2.该output 的下标
	Index int
	//3.output
	Output *TxOutput
}
```

UTXO:Unspent Transaction output

* **TxID**: 该TxOutput所在的交易id
* **Index**:该TxOutput 的下标
* **Output**:TxOutput对象



#### 2.5 未花费交易输出集合 UTXOSet

```go
type UTXOSet struct {
	BlockChian *BlockChain
}
```



```go
const utxosettable = "utxoset"
```

定义一个常量用于标识存入数据库中的Bucket表名



### 3.转账交易流程

单笔转账

```
$ ./mybtc send -from 源地址 -to 目标地址 -amount 转账金额
```

多笔转账

```
$ ./mybtc send \
  -from '["源地址1","源地址2","源地址N"]' \
  -to   '["目标地址1","目标地址2","目标地址3"]' \
  -amount '["转账金额1","转账金额2","转账金额3"]'
```

这部分内容理解起来有些难度，所以我做了一张图，希望能够帮助大家能够理顺思路，这样在后面的学习以及代码理解上面会稍微容易一些。



![2.转账交易.001](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/Keynote源文件/2.转账交易.jpeg)



本图介绍了从创世区块后的三次转账过程，分别产生了三个区块，为了让读者有更直观的了解，我又将该图做成了动态图的方式供大家参考，通过该图，希望大家能够大致对转账交易有个印象。

![2.转账交易](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/动态图/2.转账交易.gif)

动态图演示了新区块中的输入交易引用的是哪个区块中的交易输出，从而实现了区块链每次转账的金额都有据可依，也从另外一个角度展示了比特币中UTXO的概念模型。

### 4.代码分析

由于代码量巨大，为了让整个过程的理解更加流程，我改变前面几篇文章的思路，从执行命令的代码块进行一步一步的代码分析，希望能将自己的思路理顺，从而可以更好得引导读者朋友。

同样，因为很多概念性的东西，我不准备在文章里面啰嗦，如果感觉阅读难度比较大，建议先仔细阅读这篇文章

https://github.com/liuchengxu/blockchain-tutorial

然后再回头来看我的这篇文章，会事半功倍

```go
func (cli *CLI) Send(from, to, amount []string) {
	bc := GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法转账。。")
		os.Exit(1)
	}
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount)
	utsoSet :=&UTXOSet{bc}
	utsoSet.Update()
}

```



#### 4.1 获取blockchain对象

```go
func GetBlockChainObject() *BlockChain {
	/*
		1.数据库存在，读取数据库，返回blockchain即可
		2.数据库 不存在，返回nil
	 */

	if dbExists() {
		//fmt.Println("数据库已经存在。。。")
		//打开数据库
		db, err := bolt.Open(DBName, 0600, nil)
		if err != nil {
			log.Panic(err)
		}

		var blockchain *BlockChain

		err = db.View(func(tx *bolt.Tx) error {
			//打开bucket，读取l对应的最新的hash
			b := tx.Bucket([]byte(BlockBucketName))
			if b != nil {
				//读取最新hash
				hash := b.Get([]byte("l"))
				blockchain = &BlockChain{db, hash}
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		return blockchain
	} else {
		fmt.Println("数据库不存在，无法获取BlockChain对象。。。")
		return nil
	}
}
```

> 判断存储区块的DB文件是否存在，如果存在，直接从数据库Bucket中读取"l"对应Hash值，将db对象与获取到hash值赋值给需要返回的区块链对象，如果DB文件不存在，说明创世区块并未创建，没有区块链对象，直接退出程序。



#### 4.2 生成区块

获取到区块链对象之后，我们调用MineNewBlock方法进行区块的创建

![image-20180822192727438](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/5.0.转账交易-生成区块-1.png)

```go
func (bc *BlockChain) MineNewBlock(from, to, amount []string) {
	/*
	1.新建交易
	2.新建区块：
		读取数据库，获取最后一块block
	3.存入到数据库中
	 */

	//1.新建交易集合
	var txs [] *Transaction

	utxoSet := &UTXOSet{bc}

	for i := 0; i < len(from); i++ {
		//amount[0]-->int
		amountInt, _ := strconv.ParseInt(amount[i], 10, 64)
		tx := NewSimpleTransaction(from[i], to[i], amountInt, utxoSet, txs)
		txs = append(txs, tx)

	}
	/*
	分析：循环第一次：i=0
		txs[transaction1, ]
		循环第二次：i=1
		txs [transaction1, transaction2]
	 */

	//交易的验证：
	for _, tx := range txs {
		if bc.VerifityTransaction(tx, txs) == false {
			log.Panic("数字签名验证失败。。。")
		}

	}

	/*
	奖励：reward：
	创建一个CoinBase交易--->Tx
	 */
	coinBaseTransaction := NewCoinBaseTransaction(from[0])
	txs = append(txs, coinBaseTransaction)

	//2.新建区块
	newBlock := new(Block)
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//读取数据库
			blockBytes := b.Get(bc.Tip)
			lastBlock := DeserializeBlock(blockBytes)

			newBlock = NewBlock(txs, lastBlock.Hash, lastBlock.Height+1)

		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//3.存入到数据库中
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//将新block存入到数据库中
			b.Put(newBlock.Hash, newBlock.Serialize())
			//更新l
			b.Put([]byte("l"), newBlock.Hash)
			//tip
			bc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}
```

##### 4.2.1 新建交易

在以上的代码中，涉及到几个比较重要的方法，其中一个NewSimpleTransaction用于创建交易并打包，这里对代码进行了简单梳理,由于内容实在太多，在**文章末尾**我会将github的源代码地址贴出，供大家查看。

###### (1) 创建该交易的TxInput

![image-20180822193658819](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/5.1.转账交易-创建交易的TxInput.png)



###### (2) 创建该交易的TxOutput

![image-20180822193902120](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/5.1.转账交易-创建交易的TxOutput.png)



###### (3) 创建交易

![image-20180822194002399](/Users/brucefeng/Desktop/公众号文章/公链开发/从0到1简易区块链开发手册/脑图截图/5.1.转账交易-创建交易.png)



**NewSimpleTransaction**

```go
func NewSimpleTransaction(from, to string, amount int64, utxoSet *UTXOSet, txs []*Transaction) *Transaction {
	//1.定义Input和Output的数组
	var txInputs []*TxInput
	var txOuputs [] *TxOutput

	//2.创建Input
	/*
	创世区块中交易ID：c16d3ad93450cd532dcd7ef53d8f396e46b2e59aa853ad44c284314c7b9db1b4
	 */

	//获取本次转账要使用output
	//total, spentableUTXO := bc.FindSpentableUTXOs(from, amount, txs) //map[txID]-->[]int{index}
	total, spentableUTXO := utxoSet.FindSpentableUTXOs(from, amount, txs) //map[txID]-->[]int{index}


	//获取钱包的集合：
	wallets := GetWallets()
	wallet := wallets.WalletMap[from]

	for txID, indexArray := range spentableUTXO {
		txIDBytes, _ := hex.DecodeString(txID)
		for _, index := range indexArray {
			txInput := &TxInput{txIDBytes, index, nil, wallet.PublickKey}
			txInputs = append(txInputs, txInput)
		}
	}


	//3.创建Output

	//转账
	txOutput := NewTxOutput(amount, to)
	txOuputs = append(txOuputs, txOutput)

	//找零
	//txOutput2 := &TxOutput{total - amount, from}
	txOutput2 := NewTxOutput(total-amount, from)
	txOuputs = append(txOuputs, txOutput2)

	//4.创建交易
	tx := &Transaction{[]byte{}, txInputs, txOuputs}

	//设置交易的ID
	tx.SetID()


	//设置签名
	utxoSet.BlockChian.SignTransaction(tx,wallet.PrivateKey,txs)


	return tx

}
```

**FindSpentableUTXOs**

```go
func (utxoSet *UTXOSet) FindSpentableUTXOs(from string, amount int64, txs []*Transaction) (int64, map[string][]int) {
	var total int64
	//用于存储转账所使用utxo
	spentableUTXOMap := make(map[string][]int)
	//1.查询未打包可以使用的utxo：txs
	unPackageSpentableUTXOs := utxoSet.FindUnpackeSpentableUTXO(from, txs)

	for _, utxo := range unPackageSpentableUTXOs {
		total += utxo.Output.Value
		txIDStr := hex.EncodeToString(utxo.TxID)
		spentableUTXOMap[txIDStr] = append(spentableUTXOMap[txIDStr], utxo.Index)
		if total >= amount {
			return total, spentableUTXOMap
		}
	}



	//2.查询utxotable，查询utxo
	//已经存储的但是未花费的utxo
	err := utxoSet.BlockChian.DB.View(func(tx *bolt.Tx) error {
		//查询utxotable中，未花费的utxo
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			//查询
			c := b.Cursor()
		dbLoop:
			for k, v := c.First(); k != nil; k, v = c.Next() {
				txOutputs := DeserializeTxOutputs(v)
				for _, utxo := range txOutputs.UTXOs {
					if utxo.Output.UnlockWithAddress(from) {
						total += utxo.Output.Value
						txIDStr := hex.EncodeToString(utxo.TxID)
						spentableUTXOMap[txIDStr] = append(spentableUTXOMap[txIDStr], utxo.Index)
						if total >= amount {
							break dbLoop
							//return nil
						}
					}
				}

			}

		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}

	return total, spentableUTXOMap
}

```

**FindUnpackeSpentableUTXO**

```go
func (utxoSet *UTXOSet) FindUnpackeSpentableUTXO(from string, txs []*Transaction) []*UTXO {
	//存储可以使用的未花费utxo
	var unUTXOs []*UTXO

	//存储已经花费的input
	spentedMap := make(map[string][]int)

	for i := len(txs) - 1; i >= 0; i-- {
		//func caculate(tx *Transaction, address string, spentTxOutputMap map[string][]int, unSpentUTXOs []*UTXO) []*UTXO {
		unUTXOs = caculate(txs[i], from, spentedMap, unUTXOs)
	}

	return unUTXOs
}
```

**caculate**

```go
func caculate(tx *Transaction, address string, spentTxOutputMap map[string][]int, unSpentUTXOs []*UTXO) []*UTXO {
   //遍历每个tx：txID，Vins，Vouts

   //遍历所有的TxInput
   if !tx.IsCoinBaseTransaction() { //tx不是CoinBase交易，遍历TxInput
      for _, txInput := range tx.Vins {
         //txInput-->TxInput
         full_payload := Base58Decode([]byte(address))

         pubKeyHash := full_payload[1 : len(full_payload)-addressCheckSumLen]

         if txInput.UnlockWithAddress(pubKeyHash) {
            //txInput的解锁脚本(用户名) 如果和要查询的余额的用户名相同，
            key := hex.EncodeToString(txInput.TxID)
            spentTxOutputMap[key] = append(spentTxOutputMap[key], txInput.Vout)
            /*
            map[key]-->value
            map[key] -->[]int
             */
         }
      }
   }

   //遍历所有的TxOutput
outputs:
   for index, txOutput := range tx.Vouts { //index= 0,txoutput.锁定脚本
      if txOutput.UnlockWithAddress(address) {
         if len(spentTxOutputMap) != 0 {
            var isSpentOutput bool //false
            //遍历map
            for txID, indexArray := range spentTxOutputMap { //143d,[]int{1}
               //遍历 记录已经花费的下标的数组
               for _, i := range indexArray {
                  if i == index && hex.EncodeToString(tx.TxID) == txID {
                     isSpentOutput = true //标记当前的txOutput是已经花费
                     continue outputs
                  }
               }
            }

            if !isSpentOutput {
               //unSpentTxOutput = append(unSpentTxOutput, txOutput)
               //根据未花费的output，创建utxo对象--->数组
               utxo := &UTXO{tx.TxID, index, txOutput}
               unSpentUTXOs = append(unSpentUTXOs, utxo)
            }

         } else {
            //如果map长度为0,证明还没有花费记录，output无需判断
            //unSpentTxOutput = append(unSpentTxOutput, txOutput)
            utxo := &UTXO{tx.TxID, index, txOutput}
            unSpentUTXOs = append(unSpentUTXOs, utxo)
         }
      }
   }
   return unSpentUTXOs

}
```



###### (4) 设置签名

![image-20180822194100407](/Users/brucefeng/Desktop/%E5%85%AC%E4%BC%97%E5%8F%B7%E6%96%87%E7%AB%A0/%E5%85%AC%E9%93%BE%E5%BC%80%E5%8F%91/%E4%BB%8E0%E5%88%B01%E7%AE%80%E6%98%93%E5%8C%BA%E5%9D%97%E9%93%BE%E5%BC%80%E5%8F%91%E6%89%8B%E5%86%8C/%E8%84%91%E5%9B%BE%E6%88%AA%E5%9B%BE/5.1.%E8%BD%AC%E8%B4%A6%E4%BA%A4%E6%98%93-%E8%AE%BE%E7%BD%AE%E7%AD%BE%E5%90%8D.png)

**SignTransaction**

```go
func (bc *BlockChain) SignTransaction(tx *Transaction, privateKey ecdsa.PrivateKey, txs []*Transaction) {
   //1.判断要签名的tx，如果是coninbase交易直接返回
   if tx.IsCoinBaseTransaction() {
      return
   }

   //2.获取该tx中的Input，引用之前的transaction中的未花费的output
   prevTxs := make(map[string]*Transaction)
   for _, input := range tx.Vins {
      txIDStr := hex.EncodeToString(input.TxID)
      prevTxs[txIDStr] = bc.FindTransactionByTxID(input.TxID, txs)
   }

   //3.签名
   tx.Sign(privateKey, prevTxs)

}
```

**Sign**

```go
func (tx *Transaction) Sign(privateKey ecdsa.PrivateKey, prevTxsmap map[string]*Transaction) {
	//1.判断当前tx是否是coinbase交易
	if tx.IsCoinBaseTransaction() {
		return
	}

	//2.获取input对应的output所在的tx，如果不存在，无法进行签名
	for _, input := range tx.Vins {
		if prevTxsmap[hex.EncodeToString(input.TxID)] == nil {
			log.Panic("当前的Input，没有找到对应的output所在的Transaction，无法签名。。")
		}
	}

	//即将进行签名:私钥，要签名的数据
	txCopy := tx.TrimmedCopy()

	for index, input := range txCopy.Vins {

		prevTx := prevTxsmap[hex.EncodeToString(input.TxID)]

		txCopy.Vins[index].Signature = nil    
		txCopy.Vins[index].PublicKey = prevTx.Vouts[input.Vout].PubKeyHash //设置input中的publickey为对应的output的公钥哈希


		txCopy.TxID = txCopy.NewTxID()//产生要签名的交易的TxID

		//为了方便下一个input，将数据再置为空
		txCopy.Vins[index].PublicKey = nil


		/*
		第一个参数
		第二个参数：私钥
		第三个参数：要签名的数据


		func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)
		r + s--->sign
		input.Signatrue = sign
	 */
		r,s,err:=ecdsa.Sign(rand.Reader, &privateKey, txCopy.TxID )
		if err != nil{
			log.Panic(err)
		}

		sign:=append(r.Bytes(),s.Bytes()...)
		tx.Vins[index].Signature = sign
	}

}

```

**TrimmedCopy**

```go
func (tx *Transaction) TrimmedCopy() *Transaction {
	var inputs [] *TxInput
	var outputs [] *TxOutput
	for _, in := range tx.Vins {
		inputs = append(inputs, &TxInput{in.TxID, in.Vout, nil, nil})
	}

	for _, out := range tx.Vouts {
		outputs = append(outputs, &TxOutput{out.Value, out.PubKeyHash})
	}

	txCopy := &Transaction{tx.TxID, inputs, outputs}
	return txCopy

}
```



##### 4.2.2 验证交易

```go
func (tx *Transaction) Verifity(prevTxs map[string]*Transaction)bool{
	//1.如果是coinbase交易，不需要验证
	if tx.IsCoinBaseTransaction(){
		return true
	}

	//prevTxs
	for _,input:=range prevTxs{
		if prevTxs[hex.EncodeToString(input.TxID)] == nil{
			log.Panic("当前的input没有找到对应的Transaction，无法验证。。")
		}
	}

	//验证
	txCopy:= tx.TrimmedCopy()

	curev:= elliptic.P256() //曲线

	for index,input:=range tx.Vins{
		//原理：再次获取 要签名的数据  + 公钥哈希 + 签名
		/*
		验证签名的有效性：
		第一个参数：公钥
		第二个参数：签名的数据
		第三、四个参数：签名：r，s
		func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
		 */
		//ecdsa.Verify()

	//获取要签名的数据
		prevTx:=prevTxs[hex.EncodeToString(input.TxID)]

		txCopy.Vins[index].Signature = nil
		txCopy.Vins[index].PublicKey = prevTx.Vouts[input.Vout].PubKeyHash
		txCopy.TxID = txCopy.NewTxID() //要签名的数据

		txCopy.Vins[index].PublicKey = nil

		//获取公钥
		/*
		type PublicKey struct {
			elliptic.Curve
			X, Y *big.Int
		}
		 */

		x:=big.Int{}
		y:=big.Int{}
		keyLen:=len(input.PublicKey)
		x.SetBytes(input.PublicKey[:keyLen/2])
		y.SetBytes(input.PublicKey[keyLen/2:])



		rawPublicKey:=ecdsa.PublicKey{curev,&x,&y}


		//获取签名：

		r :=big.Int{}
		s :=big.Int{}

		signLen:=len(input.Signature)
		r.SetBytes(input.Signature[:signLen/2])
		s.SetBytes(input.Signature[signLen/2:])

		if ecdsa.Verify(&rawPublicKey,txCopy.TxID,&r,&s) == false{
			return false
		}

	}
	return true
}
```



##### 4.2.3 创建CoinBase交易

```go
func NewCoinBaseTransaction(address string) *Transaction {
	txInput := &TxInput{[]byte{}, -1, nil, nil}
	//txOutput := &TxOutput{10, address}
	txOutput := NewTxOutput(10, address)
	txCoinBaseTransaction := &Transaction{[]byte{}, []*TxInput{txInput}, []*TxOutput{txOutput}}
	//设置交易ID
	txCoinBaseTransaction.SetID()
	return txCoinBaseTransaction
}

```

在每个区块中创建一个CoinBase交易作为奖励机制。

##### 4.2.4 新建区块

```go
func NewBlock(txs []*Transaction, prevBlockHash [] byte, height int64) *Block {
	//创建区块
	block := &Block{height, prevBlockHash, txs, time.Now().Unix(), nil,0}
	//设置hash
	//block.SetHash()
	pow:=NewProofOfWork(block)
	hash,nonce:=pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}
```

##### 4.2.5 持久化存储

```go
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucketName))
		if b != nil {
			//将新block存入到数据库中
			b.Put(newBlock.Hash, newBlock.Serialize())
			//更新l
			b.Put([]byte("l"), newBlock.Hash)
			//tip
			bc.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
```





#### 4.3 更新UTXO集

```
func (utxoSet *UTXOSet) Update() {
	/*
	表：key：txID
		value：TxOutputs
			UTXOs []UTXO
	 */

	//1.获取最后(从后超前遍历)一个区块,遍历该区块中的所有tx
	newBlock := utxoSet.BlockChian.Iterator().Next()
	//2.获取所有的input
	inputs := [] *TxInput{}
	//遍历交易，获取所有的input
	for _, tx := range newBlock.Txs {
		if !tx.IsCoinBaseTransaction() {
			for _, in := range tx.Vins {
				inputs = append(inputs, in)
			}
		}
	}


	//存储该区块中的，tx中的未花费
	outsMap := make(map[string]*TxOutputs)

	//3.获取所有的output
	for _, tx := range newBlock.Txs {
		utxos := []*UTXO{}
		//找出交易中的未花费
		for index, output := range tx.Vouts {
			isSpent := false
			//遍历inputs的数组，比较是否有intput和该output对应，如果满足，表示花费了
			for _, input := range inputs {
				if bytes.Compare(tx.TxID, input.TxID) == 0 && index == input.Vout {
					if bytes.Compare(output.PubKeyHash, PubKeyHash(input.PublicKey)) == 0 {
						isSpent = true
					}
				}
			}
			if isSpent == false {
				//output未花
				utxo := &UTXO{tx.TxID, index, output}
				utxos = append(utxos, utxo)
			}
		}

		//utxos,
		if len(utxos) > 0 {
			txIDStr := hex.EncodeToString(tx.TxID)
			outsMap[txIDStr] = &TxOutputs{utxos}
		}

	}


	//删除花费了数据,添加未花费
	err := utxoSet.BlockChian.DB.Update(func(tx *bolt.Tx) error {
	
		b := tx.Bucket([]byte(utxosettable))
		if b != nil {
			//遍历inputs，删除
			for _, input := range inputs {
				txOutputsBytes := b.Get(input.TxID)
				if len(txOutputsBytes) == 0 {
					continue
				}

				//反序列化
				txOutputs := DeserializeTxOutputs(txOutputsBytes)
				//是否需要被删除
				isNeedDelete := false

				//存储该txoutout中未花费utxo
				utxos := []*UTXO{}

				for _, utxo := range txOutputs.UTXOs {
					if bytes.Compare(utxo.Output.PubKeyHash, PubKeyHash(input.PublicKey)) == 0 && input.Vout == utxo.Index {
						isNeedDelete = true
					} else {
						utxos = append(utxos, utxo)
					}
				}

				if isNeedDelete == true {
					b.Delete(input.TxID)
					if len(utxos) > 0 {
						txOutputs := &TxOutputs{utxos}
						b.Put(input.TxID, txOutputs.Serialize())
					}
				}
			}

			//遍历map，添加
			for txIDStr, txOutputs := range outsMap {
				txID, _ := hex.DecodeString(txIDStr)
				b.Put(txID, txOutputs.Serialize())

			}

		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
```

* 删除本次交易产生的input对应的utxo
* 添加本次交易产生的新utxo



### 5.代码共享

由于转账交易这一块的内容代码量特别大，脑图跟交易流程图我也是花费了大量的时间进行整理，但是要一项一项进行代码分析，时间成本还是太大了，所以，将github的代码共享给大家，可以照着文章思路与思维导图中的路线进行适当的分析。https://github.com/DiaboFong/MyPublicChain

