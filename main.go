package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	goccyJson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	"github.com/fiber-go-pos-app/utils/pkg/elasticsearch"
	"github.com/fiber-go-pos-app/utils/pkg/postgres"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"

	serviceRoutes "github.com/fiber-go-pos-app/routes/services"
	webRoutes "github.com/fiber-go-pos-app/routes/web"
)

func main() {
	engine := html.New(constantsEntity.TemplateDirectory, ".html")
	engine.Reload(true)
	engine.Debug(true)

	app := fiber.New(fiber.Config{
		Views:        engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 3 * time.Second,
		JSONEncoder:  goccyJson.Marshal,
		JSONDecoder:  goccyJson.Unmarshal,
		AppName:      constantsEntity.AppName,
	})

	// Setting JWT RS256
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}

	// Setting basic configuration
	app.Use(logger.New(), recover.New(), jwtWare.New(jwtWare.Config{
		SigningMethod: constantsEntity.JWTMethod,
		SigningKey:    privateKey.Public(),
	}))

	app.Static(constantsEntity.StaticUrl, constantsEntity.StaticDirectory)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := postgres.OpenConnection(); err != nil {
		panic(err)
	}

	if err := elasticsearch.NewESClient(); err != nil {
		panic(err)
	}

	// Web Group
	webRoutes.BuildLoginRoutes(app)

	// Web handler - SIS
	sisGroup := app.Group("/sis")
	webRoutes.BuildSISRoutes(sisGroup)

	// Web handler - POS
	posGroup := app.Group("/pos")
	webRoutes.BuildPOSRoutes(posGroup)

	// Service Group
	svcGroup := app.Group("/svc")
	serviceRoutes.BuildUserRoutes(svcGroup)
	serviceRoutes.BuildLoginRoutes(svcGroup)
	serviceRoutes.BuildMemberRoutes(svcGroup)
	serviceRoutes.BuildProductRoutes(svcGroup)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
