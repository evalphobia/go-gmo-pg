package client

const (
	StatusAuth     = "AUTH"
	StatusCapture  = "CAPTURE"
	StatusRegister = "REGISTER"
	StatusPayFail  = "PAYFAIL"
	StatusCancel   = "CANCEL"
	StatusReturn   = "RETURN"
	StatusEnd      = "END"

	AuAccountTimingSelect     = "01"
	AuAccountTimingEndOfMonth = "02"

	AuPayMethodTotal        = "01"
	AuPayMethodCreditCard   = "02"
	AuPayMethodWebMoney     = "03"
	AuPayMethodAuWallet     = "06"
	AuPaymentTypeAcceptCode = "1"
)
