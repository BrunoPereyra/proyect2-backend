package controllers

import (
	"backend/controllers"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(t *testing.T) {

	ctx := &fiber.Ctx{}
	ctx.Request().SetBody([]byte(`{"username": "admin", "password": "123456"}`))
	err := controllers.Login(ctx)

	if err != nil {
		t.Fatalf("Se esperaba que no ocurriera un error, pero se obtuvo: %v", err)
	}

	if ctx.Response().StatusCode() != fiber.StatusOK {
		t.Errorf("Se esperaba un código de estado 200, pero se obtuvo: %d", ctx.Response().StatusCode())
	}

	expectedResponse := `{"message": "Token created", "token": "..."}` // Ajusta el contenido esperado según tu implementación
	if string(ctx.Response().Body()) != expectedResponse {
		t.Errorf("La respuesta no coincide con el resultado esperado. Se esperaba: %s, pero se obtuvo: %s", expectedResponse, ctx.Response().Body())
	}
}
