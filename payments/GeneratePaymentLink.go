package payments

import (
	"github.com/mercadopago/mercadopago-sdk-go/payments"
)

func GeneratePaymentLink(amount float64, description string) (string, error) {
	// Crear una instancia del cliente de Mercado Pago
	mp := payments.NewClient()

	// Configurar las credenciales
	mp.SetAccessToken("ACCESS_TOKEN_HERE")

	// Crear un pago
	payment := &payments.Payment{
		TransactionAmount: amount,
		Description:       description,
		PaymentMethodId:   "rapipago",                                  // Método de pago a utilizar
		NotificationUrl:   "https://yourapp.com/payments/notification", // URL de notificación de pagos
	}

	// Crear el pago
	response, err := mp.CreatePayment(payment)
	if err != nil {
		return "", err
	}

	// Devolver el link de pago generado
	return response.InitPoint, nil
}
