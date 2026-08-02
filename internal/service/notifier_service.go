package service

import (
	tele "gopkg.in/telebot.v3"
)

type NotifierService struct {
	bot         *tele.Bot
	subService  SubscriptionService
	userService UserService
}

func NewNotifierService(bot *tele.Bot, subService SubscriptionService, userService UserService) *NotifierService {
	return &NotifierService{bot: bot, subService: subService, userService: userService}
}

func (s *NotifierService) SendNotify() error {
	subs, err := s.subService.GetActiveExpiring()
	if err != nil {
		return err
	}

	const text = "Упс.. Спешим сообщить, что ваша подписка вот-вот закончится (через 24 часа), успейте продлить подписку!"

	for _, sub := range subs {
		user, err := s.userService.GetUserById(sub.UserId)
		if err != nil {
			return err
		}

		if _, err := s.bot.Send(&tele.User{ID: user.TgId}, text); err != nil {
			return err
		}
	}

	return nil
}
