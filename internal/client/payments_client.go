package client

import "NorthlaneVPN/internal/models"

type PaymentClient interface {
	CreatePayment(m models.Payment) (models.Payment, error)
	GetPaymentStatus(transactionId string) (models.Payment, error)
}

type PaymentClientImpl struct {
	Url    string
	ApiKey string
}

func NewPaymentClient(url string, apikey string) *PaymentClientImpl {
	return &PaymentClientImpl{Url: url, ApiKey: apikey}
}

func (p *PaymentClientImpl) CreatePayment(m models.Payment) (models.Payment, error) {
	//TODO: описать подключение платежных провайдеров: СБП, Т-Банк, Ю-Касса
	return m, nil
}

func (p *PaymentClientImpl) GetPaymentStatus(transactionId string) (models.Payment, error) {
	//TODO: получить статусы оплаты, передать в сервисы, для автоматической выдачи подписки
	return models.Payment{}, nil
}
