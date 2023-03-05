package cmd

import (
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-template/cmd/handlers"
	"go-template/common"
	"go-template/middleware"
	appgrpc "go-template/plugin/grpc"
	"go-template/plugin/sdkgorm"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery-cart-categories"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(appgrpc.NewAuthClient(common.PluginGrpcAuthClient)),
		goservice.WithInitRunnable(appgrpc.NewUserClient(common.PluginGrpcUserClient)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			handlers.MainRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
