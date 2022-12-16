package main

import (
	"log"
	"math"
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

func main() {
	router := gin.Default()
	router.POST("/calculate", calcMortgage)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func calcMortgage(c *gin.Context) {
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

func calcMortgageAn(inputPartData InputPart) Output {
	var monthPay = inputPartData.Amount * (inputPartData.MonthRate / (1 - math.Pow(1+inputPartData.MonthRate, -inputPartData.MonthTerm)))

	var overpaymentValue = monthPay*inputPartData.MonthTerm - inputPartData.Amount

	var repayValue = overpaymentValue + inputPartData.Amount

	var paymentTable = calcPaymentDetailAn(inputPartData.Amount, inputPartData.MonthRate, monthPay, inputPartData.MonthTerm)

	return Output{
		TakeValue:        inputPartData.Amount,
		RepayValue:       repayValue,
		OverpaymentValue: overpaymentValue,
		PaymentTable:     paymentTable,
	}
}

func calcMortgageDif(inputPartData InputPart) Output {
	var repayBody = inputPartData.Amount / inputPartData.MonthTerm

	var paymentTable = calcPaymentDetailDif(inputPartData.Amount, inputPartData.MonthRate, repayBody, inputPartData.MonthTerm)

	var repayValue float64

	for _, r := range paymentTable {
		repayValue += r.MonthPay
	}

	var overpaymentValue = repayValue - inputPartData.Amount

	return Output{
		TakeValue:        inputPartData.Amount,
		RepayValue:       repayValue,
		OverpaymentValue: overpaymentValue,
		PaymentTable:     paymentTable,
	}
}

func calcPaymentDetailAn(estMortgageBody, monthRate, monthPay, monthTerm float64) []PaymentByMonth {
	var paymentDetail = []PaymentByMonth{}
	debtEnd := estMortgageBody

	var i float64 = 1
	for ; i <= monthTerm; i++ {
		var repayPer = debtEnd * monthRate
		var repayBody = monthPay - repayPer
		debtEnd = debtEnd - repayBody

		var paymentByMonth = PaymentByMonth{
			ID:        i,
			MonthPay:  monthPay,
			RepayPer:  repayPer,
			RepayBody: repayBody,
			DebtEnd:   debtEnd,
		}

		paymentDetail = append(paymentDetail, paymentByMonth)
	}

	return paymentDetail
}

func calcPaymentDetailDif(estMortgageBody, monthRate, repayBody, monthTerm float64) []PaymentByMonth {
	var paymentDetail = []PaymentByMonth{}
	debtEnd := estMortgageBody

	var i float64 = 1
	for ; i <= monthTerm; i++ {
		var repayPer = debtEnd * monthRate
		var monthPay = repayBody + repayPer
		debtEnd = debtEnd - repayBody

		var paymentByMonth = PaymentByMonth{
			ID:        i,
			MonthPay:  monthPay,
			RepayPer:  repayPer,
			RepayBody: repayBody,
			DebtEnd:   debtEnd,
		}

		paymentDetail = append(paymentDetail, paymentByMonth)
	}

	return paymentDetail
}
