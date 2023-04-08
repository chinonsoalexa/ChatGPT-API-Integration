package whatsapp

import (
    "github.com/skip2/go-qrcode"
    "os"
)

func qrcode() {
    // generate a QR code for the WhatsApp Web URL
    url := "https://web.whatsapp.com/"
    code, err := qrcode.Encode(url, qrcode.Medium, 256)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error generating QR code: %v\n", err)
        return
    }

    // save the QR code as a PNG file
    err = code.Save("qr-code.png")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error saving QR code: %v\n", err)
        return
    }

    fmt.Println("QR code generated successfully!")
}
