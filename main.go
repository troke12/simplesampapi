package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Southclaws/go-samp-query"
	"time"
	"fmt"
	"context"
	"net/http"
	"os"
	"encoding/json"
)

func main() {
	router := gin.Default()

	router.GET("/info", serverinfo)
	router.Run(":7000")
}

func serverinfo(c *gin.Context) {

	ipaddress := "MASUKIN_IP_WITH_PORT"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	server, err := sampquery.GetServerInfo(ctx, ipaddress, true)
	if err != nil {
		fmt.Println(err)
	}

	if err = json.NewEncoder(os.Stdout).Encode(server); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, server)
}
