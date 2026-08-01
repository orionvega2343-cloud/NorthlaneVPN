package service

import (
	"NorthlaneVPN/internal/client"
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
)

type PaymentService interface {
	CreatePayment(m models.Payment) (models.Payment, error)
	GetPaymentById(id int) (models.Payment, error)
	UpdatePayment(id int, status string) error
}

type PaymentServiceImpl struct {
	repo    repository.PaymentRepo
	client  client.PaymentClient
	polling PaymentPolling
}

func NewPaymentService(repo repository.PaymentRepo, client client.PaymentClient, polling PaymentPolling) *PaymentServiceImpl {
	return &PaymentServiceImpl{repo: repo, client: client, polling: polling}
}

func (s *PaymentServiceImpl) CreatePayment(m models.Payment) (models.Payment, error) {
	//Подключаем оплату подписок
	cl, err := s.client.CreatePayment(m)
	if err != nil {
		return models.Payment{}, err
	}

	//Сохраняем оплату в БД
	payment, err := s.repo.CreatePayment(cl)
	if err != nil {
		return models.Payment{}, err
	}

	go s.polling.Polling(payment.TransactionId, payment.Id)
	return payment, nil
}

func (s *PaymentServiceImpl) GetPaymentById(id int) (models.Payment, error) {
	return s.repo.GetPaymentById(id)
}

func (s *PaymentServiceImpl) UpdatePayment(id int, status string) error {
	return s.repo.UpdatePayment(id, status)
}
