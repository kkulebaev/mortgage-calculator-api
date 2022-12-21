package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentByMonth struct {
	ID        float64 `json:"id"`
	MonthPay  float64 `json:"monthPay"`
	RepayPer  float64 `json:"repayPer"`
	RepayBody float64 `json:"repayBody"`
	DebtEnd   float64 `json:"debtEnd"`
}

type Input struct {
	Amount       float64 `json:"amount"`
	Term         float64 `json:"term"`
	Period       string  `json:"period"`
	Rate         float64 `json:"rate"`
	MortgageType string  `json:"mortgageType"`
}

type Output struct {
	TakeValue        float64          `json:"takeValue"`
	RepayValue       float64          `json:"repayValue"`
	OverpaymentValue float64          `json:"overpaymentValue"`
	PaymentTable     []PaymentByMonth `json:"paymentTable"`
}

type InputPart struct {
	Amount    float64 `json:"amount"`
	MonthTerm float64 `json:"monthTerm"`
	MonthRate float64 `json:"monthRate"`
}

func Ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "All is OK")
}

func CalcMortgage(c *gin.Context) {
	var input Input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var monthTerm = input.Term
	if input.Period == "years" {
		monthTerm = monthTerm * 12
	}

	var monthRate = input.Rate / 100 / 12

	var mortgageOutput Output

	inputPartData := InputPart{Amount: input.Amount, MonthTerm: monthTerm, MonthRate: monthRate}
	if input.MortgageType == "an" {
		mortgageOutput = calcMortgageAn(inputPartData)
	} else {
		mortgageOutput = calcMortgageDif(inputPartData)
	}

	c.IndentedJSON(http.StatusOK, mortgageOutput)
}
