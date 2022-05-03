package business

import (
	"aspire/config"
	"aspire/models"
	"aspire/persistency/database"
	"testing"

	"github.com/gin-gonic/gin"
)

type getUserLoanTest struct {
	ctx      *gin.Context
	userID   int
	expected []models.LoanData
}

type payEMITest struct {
	ctx       *gin.Context
	loanID    string
	emiAmount float64
	expected  string
}

var getUserLoanTests = []getUserLoanTest{
	{&gin.Context{}, 4, []models.LoanData{}},
	{&gin.Context{}, 5, []models.LoanData{}},
}

var payEMITests = []payEMITest{
	{&gin.Context{}, "4f687486-b69c-4586-ae4f-245fb8b5d8db", 500.00, "no pending emi exists"},
	{&gin.Context{}, "4f687486-b69c-4586-ae4f-245fb8b5d8db", 0.00, "no pending emi exists"},
}

func TestGetUserLoans(t *testing.T) {
	config.InitConfigForTest()
	database.InitConnections()
	for _, test := range getUserLoanTests {
		output, _ := GetUserLoans(test.ctx, test.userID)
		for i := range test.expected {
			if test.expected[i] != output[i] {
				t.Errorf("Output %v not equal to expected %v", output, test.expected)
			}
		}
	}
}

func TestCreateUserLoan(t *testing.T) {
	config.InitConfigForTest()
	database.InitConnections()
	for _, test := range payEMITests {
		if output, _ := PayEMI(test.ctx, test.loanID, test.emiAmount); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}

}
