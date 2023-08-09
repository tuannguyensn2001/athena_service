package cmd

import (
	"athena_service/config"
	"athena_service/infra"
	"athena_service/middlewares"
	"athena_service/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Get()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get config")
			}
			inf, err := infra.Get(cfg)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get infra")
			}

			r := gin.Default()
			r.Use(middlewares.Cors, middlewares.Recover)

			routes.Bootstrap(r, cfg, inf)

			if err := r.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
				log.Fatal().Err(err).Msg("failed to run server")
			}
		},
	}
}
