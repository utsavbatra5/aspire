package constants

const (
	UserID                 = "Userid"
	LoanDBName             = "aspire"
	LoanDataQuery          = `SELECT * FROM loans where userid = $1`
	LoanByLoanIdQuery      = `SELECT * FROM loans where loanid = $1`
	GetLoanPaidAmountQuery = `SELECT SUM(emiamount) as emiPaid FROM loanemi where status = 'PAID' and loanid = $1`
	CreateLoanQuery        = `INSERT into loans (userid, loanid, freqcount, status, creationdate, amount) VALUES ($1, $2, $3, $4, $5, $6)`
	CreateEMIQuery         = `INSERT into loanemi (userid, loanid, loanemiid, freqno, emiamount, status, repaymentdate, islast) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	DeleteLoanQuery        = `DELETE FROM loans where loanid = $1`
	DeleteLoanEMIQuery     = `DELETE FROM loanemi where loanid = $1`
	UpdateLoanStatus       = `Update loans SET status = 'APPROVED' where loanid = $1`
	UpdateLoanEMI          = `Update loanemi SET status = 'PAID' where loanid = $1 and loanemiid = $2`
	UpdateLoanEMIAmount    = `Update loanemi SET status = 'PAID', emiamount = $3 where loanid = $1 and loanemiid = $2`
	UpdateNewLoanEMI       = `Update loanemi SET emiamount = $2 where loanid = $1 and status = 'PENDING'`
	GetLoanEMI             = `SELECT * FROM loanemi where loanid = $1 and status = 'PENDING' order by freqno LIMIT 1`
	CheckPendingEMI        = `SELECT Count(*) as pendingCount FROM loanemi where loanid = $1 and status = 'PENDING'`
)
