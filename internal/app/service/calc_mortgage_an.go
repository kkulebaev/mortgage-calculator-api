package service

import (
	"math"
)

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
