package service

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
