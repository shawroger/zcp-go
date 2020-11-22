 # zcp

`import "github.com/shawroger/zcp-go"`

张昌蒲计算器 golang 版本

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Copyright (c) 2020 shawroger <roger@shawroger.com>

## <a name="pkg-index">Index</a>
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


## <a name="CardListContain">func</a> CardListContain
``` go
func CardListContain(list []CardList, p CardList) bool
```
CardListContain 检查是否有重复卡组



## <a name="CardPointContain">func</a> CardPointContain
``` go
func CardPointContain(list CardList, p CardPoint) bool
```
CardPointContain 检查是否有重复卡牌点数



## <a name="ContentEqual">func</a> ContentEqual
``` go
func ContentEqual(p1 CardList, p2 CardList) bool
```
ContentEqual 判断卡牌内容相等



## <a name="SumEqual">func</a> SumEqual
``` go
func SumEqual(p1 CardList, p2 CardList) bool
```
SumEqual 判断卡牌和相等




## <a name="CardList">type</a> CardList
``` go
type CardList []CardPoint
```
CardList 卡牌点数组







### <a name="CardListCopy">func</a> CardListCopy
``` go
func CardListCopy(p CardList) CardList
```
CardListCopy 复制卡组


### <a name="CardListGenerate">func</a> CardListGenerate
``` go
func CardListGenerate(length int) CardList
```
CardListGenerate 生成卡组序列


### <a name="New">func</a> New
``` go
func New(x ...interface{}) CardList
```
New 从空接口新建卡组


### <a name="NewFromSlice">func</a> NewFromSlice
``` go
func NewFromSlice(x []interface{}) CardList
```
NewFromSlice 从切片新建卡组





### <a name="CardList.ContentEqual">func</a> (CardList) ContentEqual
``` go
func (p CardList) ContentEqual(t CardList) bool
```
ContentEqual 判断卡牌内容相等




### <a name="CardList.Copy">func</a> (CardList) Copy
``` go
func (p CardList) Copy() CardList
```
Copy 复制卡组




### <a name="CardList.Ergodic">func</a> (CardList) Ergodic
``` go
func (p CardList) Ergodic(width int) Result
```
Ergodic 直接遍历组合




### <a name="CardList.Len">func</a> (CardList) Len
``` go
func (p CardList) Len() int
```



### <a name="CardList.Less">func</a> (CardList) Less
``` go
func (p CardList) Less(i, j int) bool
```



### <a name="CardList.Print">func</a> (CardList) Print
``` go
func (p CardList) Print()
```
Print 控制台打印




### <a name="CardList.Run">func</a> (CardList) Run
``` go
func (p CardList) Run() Result
```
Run 直接获取合理结果




### <a name="CardList.Sort">func</a> (CardList) Sort
``` go
func (p CardList) Sort() CardList
```
Sort 对自己排序

在 Copy 基础上，不修改原数据




### <a name="CardList.Str">func</a> (CardList) Str
``` go
func (p CardList) Str() string
```
Str 转字符串




### <a name="CardList.Sum">func</a> (CardList) Sum
``` go
func (p CardList) Sum() CardPoint
```
Sum 卡牌组求和




### <a name="CardList.SumEqual">func</a> (CardList) SumEqual
``` go
func (p CardList) SumEqual(t CardList) bool
```
SumEqual 判断卡牌和相等




### <a name="CardList.Swap">func</a> (CardList) Swap
``` go
func (p CardList) Swap(i, j int)
```



## <a name="CardListFeature">type</a> CardListFeature
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









## <a name="CardListGroup">type</a> CardListGroup
``` go
type CardListGroup []CardList
```
CardListGroup 卡牌点数组组合







### <a name="ErgodicBase">func</a> ErgodicBase
``` go
func ErgodicBase(size, width int) CardListGroup
```
ErgodicBase 遍历基函数


### <a name="ErgodicCardList">func</a> ErgodicCardList
``` go
func ErgodicCardList(p CardList, width int) CardListGroup
```
ErgodicCardList 列出所有卡组组合





### <a name="CardListGroup.Print">func</a> (CardListGroup) Print
``` go
func (c CardListGroup) Print()
```



## <a name="CardPoint">type</a> CardPoint
``` go
type CardPoint int8
```
CardPoint 卡牌点数







### <a name="SumList">func</a> SumList
``` go
func SumList(p CardList) CardPoint
```
SumList 卡牌组求和





## <a name="Result">type</a> Result
``` go
type Result []ValidGroup
```






### <a name="ErgodicBase2">func</a> ErgodicBase2
``` go
func ErgodicBase2(size, width int) Result
```
ErgodicBase2 遍历基函数，且包含对应剩余分组


### <a name="ErgodicCardList2">func</a> ErgodicCardList2
``` go
func ErgodicCardList2(p CardList, width int) Result
```
ErgodicCardList2

列出所有卡组组合，且包含对应剩余分组


### <a name="Run">func</a> Run
``` go
func Run(c CardList) Result
```
Run 获取合理结果


### <a name="RunWithAll">func</a> RunWithAll
``` go
func RunWithAll(c CardList) Result
```
RunWithAll 获取由全部卡点组合的结果





### <a name="Result.Print">func</a> (Result) Print
``` go
func (r Result) Print()
```



## <a name="ValidGroup">type</a> ValidGroup
``` go
type ValidGroup [2]CardList
```
