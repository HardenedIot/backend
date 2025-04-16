package handlers

import (
	"hardenediot-client-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTechnologies(ctx *gin.Context) {
	technologies := []models.Technology{
		models.Wifi,
		models.Uart,
		models.Jtag,
		models.Bluetooth,
		models.Lte,
		models.Rfid,
		models.Nfc,
		models.Antplus,
		models.Lifi,
		models.Zigbee,
		models.Zwave,
		models.Lteadvanced,
		models.Lra,
		models.NbIot,
		models.Sigfox,
		models.NbFi,
		models.Http,
		models.Https,
		models.Coap,
		models.Mqtt,
		models.Amqp,
		models.Xmpp,
	}
	ctx.JSON(http.StatusOK, technologies)
}
