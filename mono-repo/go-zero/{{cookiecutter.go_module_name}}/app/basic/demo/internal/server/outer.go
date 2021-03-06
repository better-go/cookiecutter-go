package server

import (
	"context"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/router"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/service/outer"
	"{{cookiecutter.go_module_name}}/app/basic/demo/proto/config"
)

// interface HTTP/API server:
type OuterServer struct{}

func (m *OuterServer) Run(configFile string) {
	ctx := context.Background()

	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)

	// api engine:
	svc := outer.NewService(cfg, ctx)
	defer svc.Close()

	// new engine:
	engine := rest.MustNewServer(cfg.Server.Outer.RestConf)
	defer engine.Stop()

	// register routes:
	router.RegisterApiRoutes(engine, svc)

	log.Infof("starting api engine at %v\n", cfg.Server.Outer.RestConf)
	engine.Start()
}
