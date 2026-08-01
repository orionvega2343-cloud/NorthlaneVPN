package service

import (
	"NorthlaneVPN/internal/client"
	"context"
	"log"
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
	//Реализация status mapping для перевода кода провайдера в статус приложения
	m := make(map[int]string)
	m[0] = "success"
	m[1] = "pending"
	m[3] = "failed"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	//Таймер для периодического опроса - short polling
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
			//На каждой итерации,
			//делаем запрос к GetPaymentStatus
		case <-ticker.C:
			status, err := p.client.GetPaymentStatus(transactionId)
			if err != nil {
				log.Println("GetPaymentStatus", err)
				return
			}

			//переводим код провайдера в статус приложения через status-mapping
			val, ok := m[status]
			if !ok {
				log.Println(err)
				return
			}
			if val != "pending" {
				err = p.svc.UpdatePayment(paymentId, val)
				if err != nil {
					log.Println("UpdatePayment", err)
					return
				}
				return
			}
		}
	}

}
