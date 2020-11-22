# zcp
`import "github.com/shawroger/zcp-go"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a id="pkg-overview">Overview</a>
Copyright (c) 2020 shawroger <roger@shawroger.com>




## <a id="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func CardListContain(list []CardList, p CardList) bool](#CardListContain)
* [func CardPointContain(list CardList, p CardPoint) bool](#CardPointContain)
* [func ContentEqual(p1 CardList, p2 CardList) bool](#ContentEqual)
* [func SumEqual(p1 CardList, p2 CardList) bool](#SumEqual)
* [type CardList](#CardList)
  * [func CardListCopy(p CardList) CardList](#CardListCopy)
  * [func CardListGenerate(length int) CardList](#CardListGenerate)
  * [func New(x ...interface{}) CardList](#New)
  * [func NewFromSlice(x []interface{}) CardList](#NewFromSlice)
  * [func (p CardList) ContentEqual(t CardList) bool](#CardList.ContentEqual)
  * [func (p CardList) Copy() CardList](#CardList.Copy)
  * [func (p CardList) Ergodic(width int) Result](#CardList.Ergodic)
  * [func (p CardList) Len() int](#CardList.Len)
  * [func (p CardList) Less(i, j int) bool](#CardList.Less)
  * [func (p CardList) Print()](#CardList.Print)
  * [func (p CardList) Run() Result](#CardList.Run)
  * [func (p CardList) Sort() CardList](#CardList.Sort)
  * [func (p CardList) Str() string](#CardList.Str)
  * [func (p CardList) Sum() CardPoint](#CardList.Sum)
  * [func (p CardList) SumEqual(t CardList) bool](#CardList.SumEqual)
  * [func (p CardList) Swap(i, j int)](#CardList.Swap)
* [type CardListFeature](#CardListFeature)
* [type CardListGroup](#CardListGroup)
  * [func ErgodicBase(size, width int) CardListGroup](#ErgodicBase)
  * [func ErgodicCardList(p CardList, width int) CardListGroup](#ErgodicCardList)
  * [func (c CardListGroup) Print()](#CardListGroup.Print)
* [type CardPoint](#CardPoint)
  * [func SumList(p CardList) CardPoint](#SumList)
* [type Result](#Result)
  * [func ErgodicBase2(size, width int) Result](#ErgodicBase2)
  * [func ErgodicCardList2(p CardList, width int) Result](#ErgodicCardList2)
  * [func Run(c CardList) Result](#Run)
  * [func RunWithAll(c CardList) Result](#RunWithAll)
  * [func (r Result) Print()](#Result.Print)
* [type ValidGroup](#ValidGroup)

#### <a id="pkg-examples">Examples</a>
* [CardListGenerate](#example_CardListGenerate)

#### <a id="pkg-files">Package files</a>
[calc.go](/src/github.com/shawroger/zcp-go/calc.go) [doc.go](/src/github.com/shawroger/zcp-go/doc.go) [ergodic.go](/src/github.com/shawroger/zcp-go/ergodic.go) [impl.go](/src/github.com/shawroger/zcp-go/impl.go) [type.go](/src/github.com/shawroger/zcp-go/type.go) [utils.go](/src/github.com/shawroger/zcp-go/utils.go)


## <a id="pkg-constants">Constants</a>
``` go
const Version = "0.0.2"
```
Version is the version of this package.




## <a id="CardListContain">func</a> [CardListContain](/src/target/utils.go?s=949:1003#L59)
``` go
func CardListContain(list []CardList, p CardList) bool
```
CardListContain 检查是否有重复卡组



## <a id="CardPointContain">func</a> [CardPointContain](/src/target/utils.go?s=1218:1272#L71)
``` go
func CardPointContain(list CardList, p CardPoint) bool
```
CardPointContain 检查是否有重复卡牌点数



## <a id="ContentEqual">func</a> [ContentEqual](/src/target/utils.go?s=302:350#L18)
``` go
func ContentEqual(p1 CardList, p2 CardList) bool
```
ContentEqual 判断卡牌内容相等



## <a id="SumEqual">func</a> [SumEqual](/src/target/utils.go?s=176:220#L13)
``` go
func SumEqual(p1 CardList, p2 CardList) bool
```
SumEqual 判断卡牌和相等




## <a id="CardList">type</a> [CardList](/src/target/type.go?s=106:131#L9)
``` go
type CardList []CardPoint
```
CardList 卡牌点数组







### <a id="CardListCopy">func</a> [CardListCopy](/src/target/utils.go?s=778:816#L50)
``` go
func CardListCopy(p CardList) CardList
```
CardListCopy 复制卡组


### <a id="CardListGenerate">func</a> [CardListGenerate](/src/target/utils.go?s=608:650#L40)
``` go
func CardListGenerate(length int) CardList
```
CardListGenerate 生成卡组序列


### <a id="New">func</a> [New](/src/target/type.go?s=1132:1167#L58)
``` go
func New(x ...interface{}) CardList
```
New 从空接口新建卡组


### <a id="NewFromSlice">func</a> [NewFromSlice](/src/target/type.go?s=921:964#L46)
``` go
func NewFromSlice(x []interface{}) CardList
```
NewFromSlice 从切片新建卡组





### <a id="CardList.ContentEqual">func</a> (CardList) [ContentEqual](/src/target/impl.go?s=503:550#L32)
``` go
func (p CardList) ContentEqual(t CardList) bool
```
ContentEqual 判断卡牌内容相等




### <a id="CardList.Copy">func</a> (CardList) [Copy](/src/target/impl.go?s=604:637#L37)
``` go
func (p CardList) Copy() CardList
```
Copy 复制卡组




### <a id="CardList.Ergodic">func</a> (CardList) [Ergodic](/src/target/ergodic.go?s=1990:2033#L87)
``` go
func (p CardList) Ergodic(width int) Result
```
Ergodic 直接遍历组合




### <a id="CardList.Len">func</a> (CardList) [Len](/src/target/impl.go?s=667:694#L41)
``` go
func (p CardList) Len() int
```



### <a id="CardList.Less">func</a> (CardList) [Less](/src/target/impl.go?s=715:752#L45)
``` go
func (p CardList) Less(i, j int) bool
```



### <a id="CardList.Print">func</a> (CardList) [Print](/src/target/impl.go?s=1100:1125#L69)
``` go
func (p CardList) Print()
```
Print 控制台打印




### <a id="CardList.Run">func</a> (CardList) [Run](/src/target/calc.go?s=1597:1627#L79)
``` go
func (p CardList) Run() Result
```
Run 直接获取合理结果




### <a id="CardList.Sort">func</a> (CardList) [Sort](/src/target/impl.go?s=911:944#L56)
``` go
func (p CardList) Sort() CardList
```
Sort 对自己排序

在 Copy 基础上，不修改原数据




### <a id="CardList.Str">func</a> (CardList) [Str](/src/target/impl.go?s=1010:1040#L64)
``` go
func (p CardList) Str() string
```
Str 转字符串




### <a id="CardList.Sum">func</a> (CardList) [Sum](/src/target/impl.go?s=288:321#L22)
``` go
func (p CardList) Sum() CardPoint
```
Sum 卡牌组求和




### <a id="CardList.SumEqual">func</a> (CardList) [SumEqual](/src/target/impl.go?s=380:423#L27)
``` go
func (p CardList) SumEqual(t CardList) bool
```
SumEqual 判断卡牌和相等




### <a id="CardList.Swap">func</a> (CardList) [Swap](/src/target/impl.go?s=778:810#L49)
``` go
func (p CardList) Swap(i, j int)
```



## <a id="CardListFeature">type</a> [CardListFeature](/src/target/impl.go?s=40:263#L8)
``` go
type CardListFeature interface {
    Sum() CardPoint
    SumEqual(t CardList) bool
    ContentEqual(t CardList) bool
    Copy() CardList
    Sort() CardList
    Str() string
    Print()
    Ergodic(width int) Result
    Run() Result
    sort.Interface
}
```









## <a id="CardListGroup">type</a> [CardListGroup](/src/target/type.go?s=172:201#L12)
``` go
type CardListGroup []CardList
```
CardListGroup 卡牌点数组组合







### <a id="ErgodicBase">func</a> [ErgodicBase](/src/target/ergodic.go?s=44:91#L4)
``` go
func ErgodicBase(size, width int) CardListGroup
```
ErgodicBase 遍历基函数


### <a id="ErgodicCardList">func</a> [ErgodicCardList](/src/target/ergodic.go?s=1078:1135#L49)
``` go
func ErgodicCardList(p CardList, width int) CardListGroup
```
ErgodicCardList 列出所有卡组组合





### <a id="CardListGroup.Print">func</a> (CardListGroup) [Print](/src/target/type.go?s=203:233#L14)
``` go
func (c CardListGroup) Print()
```



## <a id="CardPoint">type</a> [CardPoint](/src/target/type.go?s=57:76#L6)
``` go
type CardPoint int8
```
CardPoint 卡牌点数







### <a id="SumList">func</a> [SumList](/src/target/utils.go?s=40:74#L4)
``` go
func SumList(p CardList) CardPoint
```
SumList 卡牌组求和





## <a id="Result">type</a> [Result](/src/target/calc.go?s=69:93#L10)
``` go
type Result []ValidGroup
```






### <a id="ErgodicBase2">func</a> [ErgodicBase2](/src/target/ergodic.go?s=672:713#L30)
``` go
func ErgodicBase2(size, width int) Result
```
ErgodicBase2 遍历基函数，且包含对应剩余分组


### <a id="ErgodicCardList2">func</a> [ErgodicCardList2](/src/target/ergodic.go?s=1532:1583#L68)
``` go
func ErgodicCardList2(p CardList, width int) Result
```
ErgodicCardList2

列出所有卡组组合，且包含对应剩余分组


### <a id="Run">func</a> [Run](/src/target/calc.go?s=760:787#L41)
``` go
func Run(c CardList) Result
```
Run 获取合理结果


### <a id="RunWithAll">func</a> [RunWithAll](/src/target/calc.go?s=407:441#L24)
``` go
func RunWithAll(c CardList) Result
```
RunWithAll 获取由全部卡点组合的结果





### <a id="Result.Print">func</a> (Result) [Print](/src/target/calc.go?s=95:118#L12)
``` go
func (r Result) Print()
```



## <a id="ValidGroup">type</a> [ValidGroup](/src/target/calc.go?s=40:67#L8)
``` go
type ValidGroup [2]CardList
```


