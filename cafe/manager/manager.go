package manager

import (
	"fmt"
	"net/http"
)

type ICoffee interface {
	Drink()
	Smell()
	Photo()
	TakeAway()
	Make()
}

type Latte struct {
	bean string
	milk string
}

func NewLatte() *Latte {
	return &Latte{
		bean: "云南豆子",
		milk: "光明牛奶",
	}
}

func (l *Latte) Drink() {
	fmt.Println("Drink Latte")
}

func (l *Latte) Smell() {
	fmt.Println("Smell Latte")
}

func (l *Latte) Photo() {
	fmt.Println("Photo Latte")
}

func (l *Latte) TakeAway() {
	fmt.Println("TakeAway Latte")
}

func (l *Latte) Make() {
	// ...
	// 上网查一下如何制作一杯拿铁
	http.Post("...", "...", nil)
	// ...
	fmt.Println("Make Latte")
}
