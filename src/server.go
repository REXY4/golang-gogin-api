package src

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/REXY4/golang-gogin-api/src/routers"
	"github.com/REXY4/golang-gogin-api/src/setups"
	"github.com/gin-gonic/gin"
)

func SetupServer(manifest string) {
	//setup db
	db := setups.SetupConnectionSql(manifest)
	defer setups.CloseSqlConnection(db)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	var config map[string]map[string]string

	err := json.Unmarshal([]byte(manifest), &config)
	if err != nil {
		panic(err.Error())
	}
	routes := routers.Router()
	for _, route := range routes() {
		for _, routes := range route.Action {
			r.Handle(route.Method, route.Path, routes)
		}
	}
	port := ":" + config["app"]["port"]

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
