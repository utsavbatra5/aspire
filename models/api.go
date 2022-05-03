package models

import "time"

type LoanEmiData struct {
	Userid        int
	Loanid        string
	Loanemiid     int
	Freqno        int
	Emiamount     float64
	Status        string
	Repaymentdate time.Time
	Islast        bool
}

type LoanData struct {
	Userid       int
	Loanid       string
	FreqCount    int
	Status       string
	CreationDate time.Time
	Amount       float64
}

type LoanCreateRequest struct {
	Loanid string
	Amount float64
	Freq   int
}

type ApproveLoan struct {
	Loanid string
}

type PayEMI struct {
	Loanid    string
	Emiamount float64
}
