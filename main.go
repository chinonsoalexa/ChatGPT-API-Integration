package main

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func main() {


}

func init() {
		/**
	*? Set Gin-Gonic to run in release mode
	 */
	 gin.SetMode(gin.ReleaseMode)

	 /**
	 *? Initializing new route to router
	  */
	 router := gin.Default()
 
	 /**
	 *? Handle GET requests to "/"
	  */
	//  router.GET("/Test/:id", routes.User)
 
	 /**
	 *? Configure server
	  */
	 server := &http.Server{
		 Addr:    ":8080",
		 Handler: router,
		 TLSConfig: &tls.Config{
			 NextProtos:         []string{http3.NextProtoH3},
			 InsecureSkipVerify: true,
		 },
	 }
 
	 /**
	 *! Start server
	  */
	 http3.ListenAndServeQUIC(":8080", "cert.pem", "key.pem", server.Handler)
 
	 /**
	 * TODO: Server starting at [::1]:3000/login
	  */
 
	 /**
	 *! Wait for server to gracefully shut down
	  */
	 if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		 panic(err)
	 }
}