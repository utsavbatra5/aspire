package business

import (
	"aspire/constants"
	"aspire/logger"
	"aspire/models"
	"aspire/persistency/database"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserLoans(ctx *gin.Context, userID int) ([]models.LoanData, error) {

	loanDataList, err := getLoans(ctx, userID)
	if err != nil {
		return []models.LoanData{}, err
	}

	return loanDataList, nil
}

func CreateUserLoan(ctx *gin.Context, req models.LoanCreateRequest, userID int) (string, error) {
	dbConnection := database.GetConnection(constants.LoanDBName)
	loanCreationTime := time.Now()

	// create entry in loan table
	_, err := dbConnection.Exec(constants.CreateLoanQuery, userID, req.Loanid, req.Freq, "PENDING", time.Now(), req.Amount)
	if err != nil {
		return "unable to create loan entry", err
	}
	emiamount := req.Amount / float64(req.Freq)
	emiamount = math.Round(emiamount*100) / 100

	var islast bool

	for i := 0; i < req.Freq; i++ {

		emiid := i + 1

		// mark last emi flag for future operations
		if emiid == req.Freq {
			islast = true
		}

		// create loan emis as per user request
		repaymentTime := loanCreationTime.AddDate(0, 0, emiid*7)
		_, err := dbConnection.Exec(constants.CreateEMIQuery, userID, req.Loanid, emiid, emiid, emiamount, "PENDING", repaymentTime, islast)
		if err != nil {
			logger.ZeroLog().Error().Err(err).Msg("failed to create loan emis")

			// in case of error, roll back database operations
			_, err := dbConnection.Exec(constants.DeleteLoanQuery, req.Loanid)
			if err != nil {
				logger.ZeroLog().Error().Err(err).Msg("failed to create loan data")
			}

			_, err = dbConnection.Exec(constants.DeleteLoanEMIQuery, req.Loanid)
			if err != nil {
				logger.ZeroLog().Error().Err(err).Msg("failed to create loan emi entry")
			}

			return "unable to create loan emi. rolling back all loan data", err
		}
	}
	return "loan created successfully", nil
}

func ApproveLoan(ctx *gin.Context, loanID string) (bool, error) {
	dbConnection := database.GetConnection(constants.LoanDBName)
	_, err := dbConnection.Exec(constants.UpdateLoanStatus, loanID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func PayEMI(ctx *gin.Context, loanID string, emiAmount float64) (string, error) {
	loanEmiData := models.LoanEmiData{}
	loanData := models.LoanData{}

	var emiPending bool

	// check for pending loan emis
	dbConnection := database.GetConnection(constants.LoanDBName)
	rows, err := dbConnection.Query(constants.GetLoanEMI, loanID)
	if err != nil {
		return "failed to fetch emi data", err
	}

	defer rows.Close()
	for rows.Next() {
		// if entry exists it means loan emis are pending
		emiPending = true
		err = rows.Scan(&loanEmiData.Userid, &loanEmiData.Loanid, &loanEmiData.Loanemiid, &loanEmiData.Freqno, &loanEmiData.Emiamount, &loanEmiData.Status, &loanEmiData.Repaymentdate, &loanEmiData.Islast)
		if err != nil {
			return "failed to read emi data", err
		}
	}

	// in case no emi is pending, return
	if !emiPending {
		return "no pending emi exists", nil
	}

	// in case last emi, return
	if loanEmiData.Islast && emiAmount != loanEmiData.Emiamount {
		return "amount more than the last payable emi", nil
	}

	if emiAmount == loanEmiData.Emiamount {
		_, err = dbConnection.Exec(constants.UpdateLoanEMI, loanID, loanEmiData.Loanemiid)
		if err != nil {
			return "unable to update loan data", err
		}
	} else {

		// if paying more than due emi, update other pending emis to balance the amount to be paid
		_, err = dbConnection.Exec(constants.UpdateLoanEMIAmount, loanID, loanEmiData.Loanemiid, emiAmount)
		if err != nil {
			return "unable to update loan emi data", err
		}
		var pendingCount int
		rows, err := dbConnection.Query(constants.CheckPendingEMI, loanID)
		if err != nil {
			return "unable to check pending emi", err
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&pendingCount)
			if err != nil {
				return "unable to get pending emi count", err
			}
		}

		// update other emis only if emi is pending
		if pendingCount != 0 {
			rows, err := dbConnection.Query(constants.LoanByLoanIdQuery, loanID)
			if err != nil {
				return "unable to fetch loan data", err
			}
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&loanData.Userid, &loanData.Loanid, &loanData.FreqCount, &loanData.Status, &loanData.CreationDate, &loanData.Amount)
				if err != nil {
					return "unable to read loan data", err
				}
			}

			rows, err = dbConnection.Query(constants.GetLoanPaidAmountQuery, loanID)
			if err != nil {
				return "unable to fetch total emi paid data", err
			}

			var emiPaid float64
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&emiPaid)
				if err != nil {
					return "unable to read total emi paid data", err
				}
			}

			pendingAmount := loanData.Amount - emiPaid
			newEmi := pendingAmount / (float64(loanData.FreqCount) - float64(loanEmiData.Loanemiid))
			newEmi = math.Round(newEmi*100) / 100

			_, err = dbConnection.Exec(constants.UpdateNewLoanEMI, loanID, newEmi)
			if err != nil {
				return "unable to update paid emi", err
			}
		} else {
			// update loan status as paid if no pending emis
			_, err := dbConnection.Exec(constants.UpdateLoanStatus, loanID)
			if err != nil {
				return "unable to update loan status", err
			}
		}
	}
	return "emi paid successfully", nil
}

func getLoans(ctx *gin.Context, userID int) ([]models.LoanData, error) {
	var loanData models.LoanData
	var loanDataList []models.LoanData

	dbConnection := database.GetConnection(constants.LoanDBName)
	rows, err := dbConnection.Query(constants.LoanDataQuery, userID)
	if err != nil {
		return []models.LoanData{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&loanData.Userid, &loanData.Loanid, &loanData.FreqCount, &loanData.Status, &loanData.CreationDate, &loanData.Amount)
		if err != nil {
			return []models.LoanData{}, err
		}
		loanDataList = append(loanDataList, loanData)
	}
	return loanDataList, nil
}
