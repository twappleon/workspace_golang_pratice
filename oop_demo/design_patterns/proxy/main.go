package main

import "fmt"

type Goods struct {
	Kind string
	Fact bool
}

type Shopping interface {
	Buy(goods *Goods)
}

type KoreaShopping struct {
}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国购物")
}

type AmericanShopping struct {
}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国购物")
}

type AfrikaShopping struct {
}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲购物")
}

type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) Buy(goods *Goods) {
	if op.distinguish(goods) {
		op.shopping.Buy(goods)
		op.check(goods)
	}
}

func NewProxy(shopping Shopping) *OverseasProxy {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, "， 不应该购买。")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "]进行了海关检查， 成功的带回祖国")
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	var shopping Shopping
	shopping = new(KoreaShopping)

	if g1.Fact {
		shopping = new(KoreaShopping)
	} else {
		shopping = new(AfrikaShopping)
	}

	var overseasProxy Shopping
	overseasProxy = NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)
}
