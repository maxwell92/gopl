package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type Car interface {
	Show() string
}

type Benz struct {
	Brand string
	*iris.Context
}

type Bmw struct {
	Price string
	*iris.Context
}

func (b Benz) Show() string {
	return b.Brand
}

func (b Bmw) Show() string {
	return b.Price
}

func (b Bmw) Get() {
	b.Write("This is bmw")
	fmt.Println("bmw get")
}

func (b *Bmw) BmwInfo(ctx *iris.Context) {
	ctx.Write(b.Show())
}

func (b Benz) Get() {
	b.Write("This is Benz")
	fmt.Println("benz get")
}

func (b Benz) BenzInfo(ctx *iris.Context) {
	ctx.Write(b.Show())
	//b.Write(b.Show())
	ctx.Next()
	//b.Next()
}

func Hi(ctx *iris.Context) {
	fmt.Println("This 2nd middleware")
	ctx.Write("hi")
}

func Display(ctx *iris.Context, c Car) {
	c.Show()
}

func main() {
	bmw := &Bmw{
		Price: "$20,000",
	}
	benz := &Benz{
		Brand: "Mercedes Benz",
	}
	//iris.API("/bmw", *bmw, bmw.BmwInfo)
	iris.API("/bmw", *bmw)
	//iris.API("/benz", *benz, Display)
	iris.API("/benz", *benz, benz.BenzInfo) // This is benz. benz get.
	//iris.API("/benz", *benz, benz.BenzInfo, Hi) // This is benz.Hi.
	iris.Listen(":8081")
}
