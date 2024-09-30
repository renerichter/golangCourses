package main

import "fmt"

/* func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balane - newItem.GetCost()
	if newBalance < 0.0 {
		return "insufficient funds"
	}
	oldItems = append(oldItems, newItem)
	return oldItems, newBalance
} */

type store[P product] interface {
	Sell(P)
}
type product interface {
	Price() float64
	Name() string
}
type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}
func (b book) Name() string {
	return b.title
}

type bookStore struct {
	booksSold []book
}

func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}
func main() {
	bs := bookStore{booksSold: []book{}}
	sellProducts[book](&bs, []book{{title: "The Hobbit", author: "J.R.R. Tolkien", price: 10.0}, {title: "Book2", author: "author2", price: 15.0}})
	fmt.Println(bs.booksSold)
}
