package payment

import (
	"backend/config"

	"github.com/fabianMendez/mercadopago"
	"github.com/gofiber/fiber/v2"
)

func TicketPaymentConfimation(c *fiber.Ctx) error {
	// crear en el front para la solicitud
	TEST_ACCESS_TOKEN := config.TEST_ACCESS_TOKEN()
	PUBLICKEY := config.PUBLICKEY()
	client := mercadopago.NewClient("https://api.mercadopago.com/v1", PUBLICKEY, TEST_ACCESS_TOKEN)

	buyer, err := client.NewTestUser(mercadopago.TestUserParams{SiteID: "MCO"})
	if err != nil {
		panic(err)
	}

	identification := mercadopago.Identification{Type: "CC", Number: "19119119100"}

	cardToken, err := client.NewCardToken(mercadopago.CardTokenParams{
		ExpirationMonth: 11,
		ExpirationYear:  2025,
		Cardholder:      mercadopago.Cardholder{Name: "APRO", Identification: identification},
		SecurityCode:    "123",
		CardNumber:      "4013540682746260",
	})
	if err != nil {
		panic(err)
	}

	payment, err := client.NewPayment(mercadopago.PaymentParams{
		PaymentMethodID:   "visa",
		TransactionAmount: 1234.5,
		Payer: mercadopago.Payer{
			Email:          buyer.Email,
			Identification: identification,
		},
		Token:        cardToken.ID,
		Description:  "Test Payment",
		Installments: 1,
	})
	if err != nil {
		panic(err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": payment.Status,
	})
	// db, err := database.NewMongoDB(10)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Pool.Disconnect(context.Background())
	// databaseGoMongodb := db.Pool.Database("goMoongodb")

	// ver si el usuario fue aceptado
	// for _, v := range collectionChampionship.AcceptedApplicants {
	// if v == idUser{

	// }
	// }

	// update := bson.M{
	// 	"$pull": bson.M{
	// 		"applicants": UserIDReq,
	// 	},
	// 	"$set": bson.M{
	// 		"participantesquepagaronlaentrada": Championship.AcceptedApplicants,
	// 	},
	// }

	//errUpdateOne := collection.UpdateOne(context.TODO(), find, update)
	// if errUpdateOne != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Internal Server Error",
	// 	})
	// }
}
