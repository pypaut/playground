package transaction

type Transaction struct {
	fromAddress string
	toAddress   string
	amount      int64
}

func (t *Transaction) GetFromAddress() string {
	return t.fromAddress
}

func (t *Transaction) GetToAddress() string {
	return t.toAddress
}

func (t *Transaction) GetAmount() int64 {
	return t.amount
}
