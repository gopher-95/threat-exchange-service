package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

// md5Regex — предкомпилированное регулярное выражение для валидации MD5-хеша.
var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

// Storage определяет контракт для хранилища угроз.
type Storage interface {
	GetInfo(ctx context.Context, md5 string) (*threat.Threat, error)
	CreateInfo(ctx context.Context, t threat.Threat) error
	UpdateInfo(ctx context.Context, t threat.Threat) error
}

// Service реализует бизнес-логику для работы с угрозами.
type Service struct {
	storage Storage
}

// NewService создает новый экземпляр Service с указанным хранилищем.
func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

// GetInfoService возвращает информацию об угрозе по её MD5-хешу.
// Возвращает ErrInvalidMD5, если формат хеша неверен, и ErrNotFound, если угроза не найдена.
func (s *Service) GetInfoService(ctx context.Context, md5 string) (*threat.Threat, error) {
	if !validMD5(md5) {
		return nil, threat.ErrInvalidMD5
	}

	t, err := s.storage.GetInfo(ctx, md5)
	if err != nil {
		if errors.Is(err, threat.ErrNotFound) {
			return nil, err
		}

		return nil, fmt.Errorf("get info: %w", err)
	}

	return t, nil
}

// CreateInfoService создает новую запись об угрозе.
// Валидирует MD5, имя, уровень опасности и тип.
// Возвращает threat.ErrInvalidMD5 или threat.ErrInvalidFields при ошибках валидации.
// Возвращает threat.ErrConflict, если угроза с таким MD5 уже существует.
func (s *Service) CreateInfoService(ctx context.Context, t *threat.ThreatCreateRequest) (*threat.Threat, error) {
	if !validMD5(t.MD5) {
		return nil, threat.ErrInvalidMD5
	}

	if !validate(t.Name, t.Severity, t.Type) {
		return nil, threat.ErrInvalidFields
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

// UpdateInfoService обновляет существующую запись об угрозе.
// Валидирует MD5, имя, уровень опасности и тип.
// Возвращает threat.ErrNotFound, если угроза с таким MD5 не существует.
func (s *Service) UpdateInfoService(ctx context.Context, md5 string, t *threat.ThreatUpdateRequest) (*threat.Threat, error) {
	if !validMD5(md5) {
		return nil, threat.ErrInvalidMD5
	}

	if !validate(t.Name, t.Severity, t.Type) {
		return nil, threat.ErrInvalidFields
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

// validMD5 проверяет, соответствует ли строка формату MD5-хеша (32 hex-символа).
func validMD5(md5 string) bool {
	return md5Regex.MatchString(md5)
}
