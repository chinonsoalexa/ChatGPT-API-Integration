package whatsapp

import (
    "fmt"
	"time"
    "github.com/Rhymen/go-whatsapp"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

func msg() {
    // create a new WhatsApp connection
    wa, err := whatsapp.NewConn(60 * time.Second)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error creating WhatsApp connection: %v\n", err)
        return
    }

    // login to WhatsApp using your phone number and QR code
    // (you will need to scan the QR code with your phone)
    login, err := wa.Login(qrCode)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error logging in to WhatsApp: %v\n", err)
        return
    }

    // create a new Gin router
    r := gin.Default()

    // define a handler function for incoming messages
    r.POST("/whatsapp", func(c *gin.Context) {
        // extract the message body from the request
        var msg whatsapp.TextMessage
        err := c.BindJSON(&msg)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
            return
        }

        // send a response message
        reply := "You said: " + msg.Text
        err = wa.Send(msg.Info.RemoteJID, reply)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "error sending response"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "response sent"})
    })

    // start the HTTP server
    err = r.Run(":8080")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error starting HTTP server: %v\n", err)
        return
    }
}
