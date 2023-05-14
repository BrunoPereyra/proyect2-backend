package controllers

import (
	"backend/controllers"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	app.Post("/login", controllers.Login)

	testCases := []struct {
		NameUser string `json:"NameUser"`
		Password string `json:"password"`
	}{
		{
			NameUser: "bruno",
			Password: "123456789",
		},
	}

	requestBody, _ := json.Marshal(testCases[0])
	req := httptest.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, 10000)
	if err != nil {
		t.Fatalf("Error al realizar la solicitud: %s", err.Error())
	}

	defer resp.Body.Close()
	t.Logf("Estado de la reaspuesta: %s", resp.Status)
}
