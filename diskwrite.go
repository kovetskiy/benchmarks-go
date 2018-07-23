package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/kovetskiy/goa/uuid"
)

type Packet struct {
	ID int64 `json:"id"`

	AccountDebit  uuid.UUID `json:"account_debit" binding:"required"`
	AccountCredit uuid.UUID `json:"account_credit" binding:"required"`

	Status int64    `json:"status,omitempty"`
	Side   int64    `json:"side"`
	Kind   int64    `json:"kind"`
	Market [10]byte `json:"market"`
	Amount int64    `json:"amount"`
	Price  int64    `json:"price"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

func main() {
	file, err := os.OpenFile("diskwrite.test", os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		panic(err)
	}

	packet := Packet{}

	buffer := bytes.NewBuffer(nil)
	err = binary.Write(buffer, binary.BigEndian, packet)
	if err != nil {
		panic(err)
	}

	amount := 10000

	started := time.Now()
	for i := 0; i < amount; i++ {
		_, err := file.Write(buffer.Bytes())
		if err != nil {
			panic(err)
		}
	}
	stopped := time.Now()

	duration := stopped.Sub(started)

	fmt.Printf("%d - %s (%.2f)\n", amount, duration, float64(amount)/duration.Seconds())
}
