package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ppastene/go-shortener/internal/cache"
	"github.com/ppastene/go-shortener/internal/config"
	"github.com/ppastene/go-shortener/internal/controllers"
	"github.com/ppastene/go-shortener/internal/services"
	"github.com/ppastene/go-shortener/pkg/keygen"
)

func main() {
	// Load the config file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	var cacheService cache.Cache
	switch cfg.CacheService {
	case "internal":
		cacheService = cache.NewMemoryCache(cfg)
	case "redis":
		cacheService = cache.NewRedisCache(cfg)
	default:
		log.Fatal("Error while setting the cache service")
	}
	keygen := keygen.NewKeygen()
	service := services.NewShortenerService(cacheService, *cfg)
	controller := controllers.NewShortenerController(*service, *keygen, *cfg)
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/redirect/", controller.RedirectUrl)
	http.HandleFunc("/save", controller.SaveUrl)
	http.HandleFunc("/error", controller.Error)
	http.HandleFunc("/list", controller.List)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

/*
Por hacer (PRIORIDAD):
1. Implementar redis
2. Implementar la configuracion en el resto del sistema (redis, keygen, tiempo de expiracion, intervalo limpieza, url, port)
3. Crear pruebas unitarias
4. Crear pruebas funcionales
5. Crear pruebas con selenium
6. Github actions
7. Jenkins
8. JMeter
*/

/*
Por hacer:
1. Error 404
2. Reorganizar codigo
*/
