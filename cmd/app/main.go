package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/docs"
	"github.com/aasumitro/gowa/internal"
	"github.com/aasumitro/gowa/pkg/appconfig"
	"github.com/aasumitro/gowa/pkg/appconstant"
	"github.com/aasumitro/gowa/resources"
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
// @license.url   	https://github.com/aasumitro/pokewar/blob/main/LICENSE

var (
	appEngine *gin.Engine
	ctx       = context.Background()
)

func init() {
	appconfig.LoadEnv()

	if appconfig.Instance.AppDebug {
		accessLogFile, _ := os.Create("./storage/logs/access.log")
		gin.DefaultWriter = io.MultiWriter(accessLogFile, os.Stdout)

		errorLogFile, _ := os.Create("./storage/logs/errors.log")
		gin.DefaultErrorWriter = io.MultiWriter(errorLogFile, os.Stdout)
	}

	if !appconfig.Instance.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	appEngine = gin.Default()

	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = appconfig.Instance.AppName
	docs.SwaggerInfo.Description = fmt.Sprintf("%s API Spec", appconfig.Instance.AppName)
	docs.SwaggerInfo.Version = appconfig.Instance.AppVersion
	docs.SwaggerInfo.Host = appconfig.Instance.AppURL
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	appEngine.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/home")
	})

	appEngine.StaticFS("/home",
		http.FS(resources.Resource))

	appEngine.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(
				appconstant.GinModelsDepth)))

	internal.NewAPIProvider(ctx, appEngine)

	log.Fatal(appEngine.Run(appconfig.Instance.AppURL))
}
