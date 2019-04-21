package main

import (
	"fmt"
)

// Investor represents the Investor, the total funds avaialble, and what remains unallocated to a Loan.
type Investor struct {
	name        string
	funds       int
	unallocated int
}

// Loan represents a funds pool with the unallocated amount to fill.
type Loan struct {
	name        string
	funds       int
	unallocated int
}

// printAll prints all loans and investors.
func printAll(loans []Loan, investors []Investor) {
	fmt.Println("investors", investors)
	fmt.Println("loans", loans)
}

// printLoanInvestor prints the given loam and investor details.
func printLoanInvestor(loan *Loan, investor *Investor) {
	fmt.Println("loan", loan)
	fmt.Println("investor", investor)
}

// processLoans iterates over loans allocating funds from investors
func processLoans(loans []Loan, investors []Investor) {
	printAll(loans, investors)
	fmt.Println()

	for loanID := range loans {
		for investorID := range investors {
			fundLoan(&loans[loanID], &investors[investorID])
		}
	}

	printAll(loans, investors)
}

// fundLoan attempts to allocate investor funds to a loan.
func fundLoan(loan *Loan, investor *Investor) {
	if loan.unallocated > 0 && investor.unallocated > 0 {

		printLoanInvestor(loan, investor)

		if loan.unallocated >= investor.unallocated {
			fmt.Println("loan >= investor")
			loan.unallocated -= investor.unallocated
			investor.unallocated = 0
		} else {
			fmt.Println("loan < investor")
			investor.unallocated -= loan.unallocated
			loan.unallocated = 0
		}

		printLoanInvestor(loan, investor)
		fmt.Println()
	}
}

// main initialise and run ProcessLoans
func main() {
	investors := []Investor{
		{name: "i1", funds: 100, unallocated: 100},
		{name: "i2", funds: 200, unallocated: 200},
		{name: "i3", funds: 150, unallocated: 150},
		{name: "i4", funds: 50, unallocated: 50},
	}

	loans := []Loan{
		{name: "l1", funds: 100, unallocated: 50},
		{name: "l2", funds: 150, unallocated: 150},
		{name: "l3", funds: 100, unallocated: 50},
		{name: "l4", funds: 20, unallocated: 20},
		{name: "l5", funds: 80, unallocated: 80},
		{name: "l6", funds: 120, unallocated: 100},
	}

	processLoans(loans, investors)

}
