package chain

type Block struct {
	Index int
	Timestamp string
	Transactions []int
	Proof int
	Previous_hash string
}
