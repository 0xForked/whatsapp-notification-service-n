package main

import (
	"context"
	"fmt"
	"github.com/aasumitro/gowa/docs"
	"github.com/aasumitro/gowa/internal"
	"github.com/aasumitro/gowa/pkg/appconfigs"
	"github.com/aasumitro/gowa/pkg/appconsts"
	"github.com/aasumitro/gowa/resources"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
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
	appconfigs.LoadEnv()

	if !appconfigs.Instance.AppDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	appEngine = gin.Default()

	docs.SwaggerInfo.BasePath = appEngine.BasePath()
	docs.SwaggerInfo.Title = appconfigs.Instance.AppName
	docs.SwaggerInfo.Description = fmt.Sprintf("%s API Spec", appconfigs.Instance.AppName)
	docs.SwaggerInfo.Version = appconfigs.Instance.AppVersion
	docs.SwaggerInfo.Host = appconfigs.Instance.AppURL
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func main() {
	appEngine.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusPermanentRedirect, "/ns")
	})

	appEngine.StaticFS("/ns",
		http.FS(resources.Resource))

	appEngine.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(
				appconsts.GinModelsDepth)))

	internal.NewAPIProvider(ctx, appEngine)

	log.Fatal(appEngine.Run(appconfigs.Instance.AppURL))
}
