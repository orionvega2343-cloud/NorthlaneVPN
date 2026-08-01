package service

import (
	"NorthlaneVPN/internal/client"
	"context"
	"time"
)

type PaymentPolling struct {
	client client.PaymentClient
	svc    PaymentService
}

func NewPaymentPolling(client client.PaymentClient, svc PaymentService) *PaymentPolling {
	return &PaymentPolling{client: client, svc: svc}
}

func (p *PaymentPolling) Polling(transactionId string, paymentId int) {
	//Реализация status mapping для валидации статусов
	m := make(map[int]string)
	m[0] = "success"
	m[1] = "pending"
	m[3] = "failed"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	//Создание счетчика, для short polling паттерна
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			status, err := p.client.GetPaymentStatus(transactionId)
			if err != nil {
				return
			}

			val, ok := m[status]
			if !ok {
				return
			}
			if val != "pending" {
				err = p.svc.UpdatePayment(paymentId, val)
				if err != nil {
					return
				}
				return
			}
		}
	}

}
