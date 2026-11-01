package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"github.com/gopher-95/threat-exchange-service/internal/threat"
)

type mockService struct {
	threatToReturn *threat.Threat
	errToReturn    error
}

func (m *mockService) GetInfoService(ctx context.Context, md5 string) (*threat.Threat, error) {
	return m.threatToReturn, m.errToReturn
}
func (m *mockService) CreateInfoService(ctx context.Context, t *threat.ThreatCreateRequest) (*threat.Threat, error) {
	return m.threatToReturn, m.errToReturn
}
func (m *mockService) UpdateInfoService(ctx context.Context, md5 string, t *threat.ThreatUpdateRequest) (*threat.Threat, error) {
	return m.threatToReturn, m.errToReturn
}

func addChiURLParam(req *http.Request, key, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
}

func TestGetInfoHandle_NotFound(t *testing.T) {
	h := NewHandler(&mockService{errToReturn: threat.ErrNotFound})

	req := httptest.NewRequest(http.MethodGet, "/info/d41d8cd98f00b204e9800998ecf8427e", nil)
	req = addChiURLParam(req, "md5", "d41d8cd98f00b204e9800998ecf8427e")
	rr := httptest.NewRecorder()

	h.GetInfoHandle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expect 404, get %d", rr.Code)
	}
}

func TestGetInfoHandle_Success(t *testing.T) {
	expected := &threat.Threat{MD5: "d41d8cd98f00b204e9800998ecf8427e", Name: "Test"}
	h := NewHandler(&mockService{threatToReturn: expected})

	req := httptest.NewRequest(http.MethodGet, "/info/d41d8cd98f00b204e9800998ecf8427e", nil)
	req = addChiURLParam(req, "md5", "d41d8cd98f00b204e9800998ecf8427e")
	rr := httptest.NewRecorder()

	h.GetInfoHandle(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expect 200, get %d", rr.Code)
	}
}

func TestCreateInfoHandle_InvalidJSON(t *testing.T) {
	h := NewHandler(&mockService{})

	req := httptest.NewRequest(http.MethodPost, "/info", strings.NewReader(`{invalid json`))
	rr := httptest.NewRecorder()

	h.CreateInfoHandle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expect 400, get %d", rr.Code)
	}
}

func TestUpdateInfoHandle_NotFound(t *testing.T) {
	h := NewHandler(&mockService{errToReturn: threat.ErrNotFound})

	body := `{"name":"Updated","severity":"high","type":"trojan"}`
	req := httptest.NewRequest(http.MethodPut, "/info/d41d8cd98f00b204e9800998ecf8427e", strings.NewReader(body))

	req = addChiURLParam(req, "md5", "d41d8cd98f00b204e9800998ecf8427e")
	rr := httptest.NewRecorder()

	h.UpdateInfoHandle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("ожидали 404, получили %d", rr.Code)
	}
}
