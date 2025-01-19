package constant

type TransactionStatus string

const (
	TransactionPending TransactionStatus = "Pending"
	TransactionSuccess TransactionStatus = "Success"
	TransactionFailed  TransactionStatus = "Failed"
)

func (t TransactionStatus) GetTransactionStatus() string {
	return string(t)
}
