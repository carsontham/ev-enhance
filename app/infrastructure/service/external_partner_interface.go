package service

type ExternalPartnerAPI interface {
	GetEVInformation(id string) (EVInformation, error)
	CreatePayment(obj ExternalPartnerObject)
}
