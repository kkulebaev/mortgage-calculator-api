package service

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
