package service

import (
	"context"
	"errors"
	"testing"

	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

type mockStorage struct {
	threatToReturn *threat.Threat
	errToReturn    error
}

func (m *mockStorage) GetInfo(ctx context.Context, md5 string) (*threat.Threat, error) {
	return m.threatToReturn, m.errToReturn
}

func (m *mockStorage) CreateInfo(ctx context.Context, t threat.Threat) error {
	return m.errToReturn
}

func (m *mockStorage) UpdateInfo(ctx context.Context, t threat.Threat) error {
	return m.errToReturn
}

func TestGetInfoService_InvalidMD5(t *testing.T) {
	svc := NewService(&mockStorage{})

	_, err := svc.GetInfoService(context.Background(), "invalid-md5")

	if !errors.Is(err, threat.ErrInvalidMD5) {
		t.Errorf("expect InvalidMD5, get: %v", err)
	}
}

func TestGetInfoService_NotFound(t *testing.T) {
	svc := NewService(&mockStorage{errToReturn: threat.ErrNotFound})

	_, err := svc.GetInfoService(context.Background(), "d41d8cd98f00b204e9800998ecf8427e")

	if !errors.Is(err, threat.ErrNotFound) {
		t.Errorf("expect ErrNotFound, get: %v", err)
	}
}

func TestGetInfoService_Success(t *testing.T) {
	expected := &threat.Threat{
		MD5:  "d41d8cd98f00b204e9800998ecf8427e",
		Name: "Trojan.Test",
	}
	svc := NewService(&mockStorage{threatToReturn: expected})

	got, err := svc.GetInfoService(context.Background(), "d41d8cd98f00b204e9800998ecf8427e")

	if err != nil {
		t.Fatalf("didnt expect error, get: %v", err)
	}
	if got.MD5 != expected.MD5 {
		t.Errorf("expect MD5 %s, get %s", expected.MD5, got.MD5)
	}
}

func TestCreateInfoService_EmptyName(t *testing.T) {
	svc := NewService(&mockStorage{})

	req := &threat.ThreatCreateRequest{
		MD5:      "d41d8cd98f00b204e9800998ecf8427e",
		Name:     "",
		Severity: "high",
		Type:     "trojan",
	}

	_, err := svc.CreateInfoService(context.Background(), req)

	if !errors.Is(err, threat.ErrInvalidFields) {
		t.Errorf("expect ErrInvalidFields, get: %v", err)
	}
}
