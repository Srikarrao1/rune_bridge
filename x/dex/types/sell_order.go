// x/dex/types/sell_order_book.go

package types

func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrder {
    book := NewOrderBook()
    return SellOrder{
        AmountDenom: AmountDenom,
        PriceDenom:  PriceDenom,
        Book:        &book,
    }
}