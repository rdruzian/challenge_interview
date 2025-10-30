package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rdruzian/challenge_interview/inbound"
	"github.com/rdruzian/challenge_interview/model"
	"github.com/rdruzian/challenge_interview/router"
)

type mockService struct{}

func (m mockService) CreateDevice(device model.Device) error { return nil }
func (m mockService) UpdateDevice(device model.Device) (model.Device, error) { return device, nil }
func (m mockService) GetDevice(id int) (model.Device, error) {
	return model.Device{ID: int64(id), Name: "Mock", Brand: "Test", State: "available"}, nil
}
func (m mockService) GetAllDevice() ([]model.Device, error) {
	return []model.Device{{ID: 1, Name: "A", Brand: "T", State: "available"}}, nil
}
func (m mockService) GetDeviceByBrand(brand string) ([]model.Device, error) {
	return []model.Device{{ID: 2, Name: "B", Brand: brand, State: "in-use"}}, nil
}
func (m mockService) GetDeviceByState(state string) ([]model.Device, error) {
	return []model.Device{{ID: 3, Name: "C", Brand: "T", State: state}}, nil
}
func (m mockService) DeleteDevice(id int) error { return nil }

func setupRouter(s inbound.DeviceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	return router.ConfigRoutes(engine, s)
}


func TestGetAllDevices(t *testing.T) {
	r := setupRouter(mockService{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/device", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestGetDeviceById_InvalidID(t *testing.T) {
	r := setupRouter(mockService{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/device/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestGetDeviceById_OK(t *testing.T) {
	r := setupRouter(mockService{})
	req := httptest.NewRequest(http.MethodGet, "/api/v1/device/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestCreateDevice_OK(t *testing.T) {
	r := setupRouter(mockService{})
	payload := model.Device{Name: "New", Brand: "Brand", State: "available"}
	b, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/device/create", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}
}

func TestDeleteDevice_OK(t *testing.T) {
	r := setupRouter(mockService{})
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/device/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}
}

func TestNotFound(t *testing.T) {
	r := setupRouter(mockService{})
	req := httptest.NewRequest(http.MethodGet, "/nope", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}