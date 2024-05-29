// x/dex/types/buy_order.go

package types

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrder {
    book := NewOrderBook()
    return BuyOrder{
        AmountDenom: AmountDenom,
        PriceDenom:  PriceDenom,
        Book:        &book,
    }
}