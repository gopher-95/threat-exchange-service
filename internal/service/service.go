package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

type Storage interface {
	GetInfo(ctx context.Context, md5 string) (*threat.Threat, error)
	CreateInfo(ctx context.Context, t threat.Threat) error
	UpdateInfo(ctx context.Context, t threat.Threat) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetInfoService(ctx context.Context, md5 string) (*threat.Threat, error) {
	if !validMD5(md5) {
		return nil, fmt.Errorf("ошибка в md5 хэше")
	}

	t, err := s.storage.GetInfo(ctx, md5)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Service) CreateInfoService(ctx context.Context, t *threat.ThreatCreateRequest) (*threat.Threat, error) {
	if !validMD5(t.MD5) {
		return nil, fmt.Errorf("ошибка в md5 хэше")
	}

	if !validate(t.Name, t.Severity, t.Type) {
		return nil, fmt.Errorf("неправильно указаны поля")
	}

	newThreat := threat.Threat{
		MD5:         t.MD5,
		Name:        t.Name,
		Severity:    t.Severity,
		Type:        t.Type,
		Description: t.Description,
	}

	if err := s.storage.CreateInfo(ctx, newThreat); err != nil {
		return nil, err
	}

	return s.storage.GetInfo(ctx, newThreat.MD5)
}

func (s *Service) UpdateInfoService(ctx context.Context, md5 string, t *threat.ThreatUpdateRequest) (*threat.Threat, error) {
	if !validMD5(md5) {
		return nil, fmt.Errorf("ошибка в md5 хэше")
	}

	if !validate(t.Name, t.Severity, t.Type) {
		return nil, fmt.Errorf("неправильно указаны поля")
	}

	updatedThreat := threat.Threat{
		MD5:         md5,
		Name:        t.Name,
		Severity:    t.Severity,
		Type:        t.Type,
		Description: t.Description,
	}

	if err := s.storage.UpdateInfo(ctx, updatedThreat); err != nil {
		return nil, err
	}

	return s.storage.GetInfo(ctx, md5)

}

func validMD5(md5 string) bool {
	return md5Regex.MatchString(md5)
}
