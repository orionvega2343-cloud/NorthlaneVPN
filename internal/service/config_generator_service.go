package service

import (
	"NorthlaneVPN/internal/models"
	"NorthlaneVPN/internal/repository"
	"encoding/base64"
	"fmt"
	"strings"
)

type ConfigGeneratorService interface {
	GenerateConfig(subscriptionId int) (string, error)
}

type ConfigGeneratorServiceImpl struct {
	subService     SubscriptionService
	subServersRepo repository.SubscriptionServersRepo
}

func NewConfigGeneratorService(subService SubscriptionService, subServersRepo repository.SubscriptionServersRepo) *ConfigGeneratorServiceImpl {
	return &ConfigGeneratorServiceImpl{subService: subService, subServersRepo: subServersRepo}
}

func (c *ConfigGeneratorServiceImpl) GenerateConfig(subscriptionId int) (string, error) {
	sub, err := c.subService.GetBySubscriptionId(subscriptionId)
	if err != nil {
		return "", err
	}

	servers, err := c.subServersRepo.GetServersForSubscription(subscriptionId)
	if err != nil {
		return "", err
	}

	links := make([]string, 0, len(servers))
	for _, srv := range servers {
		cfg := models.ServerConfig{
			Uuid:       sub.Uuid,
			Host:       srv.Host,
			Name:       srv.Name,
			RealityKey: srv.RealityKey,
			Sni:        srv.Sni,
		}
		links = append(links, buildVlessLink(cfg, srv.Port))
	}

	joined := strings.Join(links, "\n")
	return base64.StdEncoding.EncodeToString([]byte(joined)), nil
}

func buildVlessLink(cfg models.ServerConfig, port int) string {
	return fmt.Sprintf(
		"vless://%s@%s:%d?encryption=none&security=reality&sni=%s&pbk=%s&type=tcp&flow=xtls-rprx-vision#%s",
		cfg.Uuid, cfg.Host, port, cfg.Sni, cfg.RealityKey, cfg.Name,
	)
}
