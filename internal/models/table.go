package models

type TableStatus string

const (
	TableStatusAvailable TableStatus = "available"
	TableStatusReserved  TableStatus = "reserved"
)

func (ts TableStatus) String() string {
	return string(ts)
}

type Table struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Seats     int         `json:"seats"`
	Status    TableStatus `json:"status"`
	BookedBy  int         `json:"bookedBy"` // User ID of the person who booked the table
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
}
