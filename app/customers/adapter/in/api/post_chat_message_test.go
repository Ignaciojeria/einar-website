package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"einar-website-api/app/shared/configuration"
	"einar-website-api/app/shared/infrastructure/serverwrapper"
	"einar-website-api/app/shared/validator"

	"github.com/labstack/echo/v4"
)

func TestPostChatMessage(t *testing.T) {
	e := echo.New()
	conf := configuration.Conf{}
	wrapper := serverwrapper.NewEchoWrapper(e,
		conf,
		validator.NewValidator())

	postChatMessage(wrapper)

	req := httptest.NewRequest(http.MethodPost, "/insert-your-custom-pattern-here", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var resp struct {
		Message string `json:"message"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedMessage := "Unimplemented"
	if resp.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, resp.Message)
	}
}
