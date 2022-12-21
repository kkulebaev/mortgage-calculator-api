package service

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
