package types

type Money int64

type Category string

type Status string

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment provides information about the payment
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
