package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentByMonth struct {
	ID        float64 `json:"id"`
	MonthPay  float64 `json:"month_pay"`
	RepayPer  float64 `json:"repay_per"`
	RepayBody float64 `json:"repay_body"`
	DebtEnd   float64 `json:"debt_end"`
}

type Input struct {
	Amount       float64 `json:"amount"`
	Term         float64 `json:"term"`
	Period       string  `json:"period"`
	Rate         float64 `json:"rate"`
	MortgageType string  `json:"mortgage_type"`
}

type Output struct {
	TakeValue        float64          `json:"take_value"`
	RepayValue       float64          `json:"repay_value"`
	OverpaymentValue float64          `json:"overpayment_value"`
	PaymentTable     []PaymentByMonth `json:"payment_table"`
}

type InputPart struct {
	Amount    float64 `json:"amount"`
	MonthTerm float64 `json:"month_term"`
	MonthRate float64 `json:"month_rate"`
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
