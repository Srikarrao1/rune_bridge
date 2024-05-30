

package types

func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrder {
    book := NewOrderBook()
    return SellOrder{
        AmountDenom: AmountDenom,
        PriceDenom:  PriceDenom,
        Book:        &book,
    }
}

func (s *SellOrder) AppendOrder(creator string, amount int32, price int32) (int32, error) {
    return s.Book.appendOrder(creator, amount, price, Decreasing)
}