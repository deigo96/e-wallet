package constant

type TransactionType int

const (
	TransactionTopup TransactionType = iota + 1
	TransactionPayment
	TransactionRefund
)

var transactionType = map[TransactionType]string{
	TransactionTopup:   "Topup",
	TransactionPayment: "Payment",
	TransactionRefund:  "Refund",
}

func (t TransactionType) GetTransactionType() string {
	return transactionType[t]
}

func (t TransactionType) GetTransactionValue() int {
	return int(t)
}

func (t TransactionType) IsValidTransactionType() bool {
	if _, ok := transactionType[t]; ok {
		return true
	}
	return false
}

// func (t *TransactionType) GetTransactionType(value int) string {
// 	if name, ok := transactionType[int(t)]; ok {
// 		return name
// 	}
// 	return transactionType[TransactionUnknown]
// }

// func (t *TransactionType) GetTransactionValue(name string) int {
// 	for key, value := range transactionType {
// 		if value == name {
// 			return key
// 		}
// 	}
// 	return TransactionUnknown
// }

// func (t *TransactionType) IsValidTransactionType(value int) bool {
// 	if _, ok := transactionType[value]; ok {
// 		return true
// 	}
// 	return false
// }
