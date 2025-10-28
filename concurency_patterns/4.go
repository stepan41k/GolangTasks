package main

import (
	"fmt"
	"math/rand"
)

type Transaction struct {
	ID     int64
	Amount float64
}

func generateTransactions(count int) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for i := 0; i < count; i++ {
			out <- Transaction{
				ID: int64(i),
				Amount: rand.Float64() * 200 - 100, // -100, 100
			}
		}

		close(out)
	}()

	return out
}

func filterTransactions(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			

			if tr.Amount >= 0 {
				out <- tr
			} else {
				fmt.Println("Negative", tr.Amount)
			}
		}

		close(out)
	}()

	return out
}

func convertTransactions(in <-chan Transaction) <-chan Transaction {
	out := make(chan Transaction)

	go func() {
		for tr := range in {
			tr.Amount *= 0.8
			out <- tr
		}

		close(out)
	}()

	return out
}

func saveTransactions(in <-chan Transaction)  {

	for tr := range in {
		fmt.Printf("%d, %.2f\n", tr.ID, tr.Amount)
	}
}

func main() {
	transaction := generateTransactions(10)
	filtered := filterTransactions(transaction)
	converted := convertTransactions(filtered)
	saveTransactions(converted)
}
