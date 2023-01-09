package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"github.com/aasumitro/gowa/constants"
	"github.com/aasumitro/gowa/docs"
	"github.com/aasumitro/gowa/internal"
	"github.com/aasumitro/gowa/pkg/whatsapp"
	"github.com/aasumitro/gowa/web"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"log"
	"net/http"
	"os"
)

// @contact.name 	@aasumitro
// @contact.url 	https://aasumitro.id/
// @contact.email 	hello@aasumitro.id
// @license.name  	MIT
// @license.url   	https://github.com/aasumitro/whatsapp-notification-service/blob/main/LICENSE

var (
	appEngine      *gin.Engine
	whatsappClient *whatsapp.Client
	ctx            = context.Background()
)

func init() {
	configs.LoadEnv()
	configs.Instance.InitDbConn()
	initGinEngine()
	initSwaggerDocs()
	initWhatsappClient()
}

func initGinEngine() {
	if configs.Instance.AppDebug {
		accessLogFile, _ := os.Create("./storage/logs/access.log")
		gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

		errorLogFile, _ := os.Create("./storage/logs/errors.log")
		gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)
	}

	if !configs.Instance.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	appEngine = gin.Default()
}

func initSwaggerDocs() {
	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = configs.Instance.AppName
	docs.SwaggerInfo.Description = fmt.Sprintf("%s API Spec", configs.Instance.AppName)
	docs.SwaggerInfo.Version = configs.Instance.AppVersion
	docs.SwaggerInfo.Host = configs.Instance.AppURL
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func initWhatsappClient() {
	whatsappInstance := whatsapp.Client{}
	whatsappClient = whatsappInstance.MakeConnection()
}

func main() {
	defer whatsappClient.WAC.Disconnect()

	appEngine.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/home")
	})

	appEngine.StaticFS("/home",
		http.FS(web.Resource))

	appEngine.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(
				constants.GinModelsDepth)))

	internal.NewAPIProvider(ctx, appEngine, whatsappClient)

	log.Fatal(appEngine.Run(configs.Instance.AppURL))
}
