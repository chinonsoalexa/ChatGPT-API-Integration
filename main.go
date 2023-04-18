package main

import (
    // "crypto/tls"
    // "net/http"

    "chat/route"

    // "github.com/quic-go/quic-go/http3"
)

func main() {
    // Set up the router
    router := route.SetupRouter()
	router.Run(":8080")

    // // Configure the server
    // server := &http.Server{
    //     Addr: ":8080",
    //     Handler: router,
    //     TLSConfig: &tls.Config{
    //         NextProtos:         []string{http3.NextProtoH3},
    //         InsecureSkipVerify: true,
    //     },
    // }

    // Start the server using HTTP/3 with QUIC
    // err := http3.ListenAndServeQUIC(":8080", "server.crt", "server.key", server.Handler)
    // if err != nil {
    //     panic(err)
    // }
}
