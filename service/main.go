package main

import (
	"context"
	"github.com/LUSHDigital/core"
	coreredis "github.com/LUSHDigital/core-redis"
	"github.com/LUSHDigital/core/env"
	"github.com/go-redis/redis"
	"github.com/LUSHDigital/core/workers/grpcsrv"
	"timestables/service/app"
)

type Dependencies struct {
	redis redis.Cmdable
}

func (d *Dependencies) Redis() redis.Cmdable {
	return d.redis
}

func main() {
	env.TryLoadDefault()
	svc := core.Service{
		Name:        "timestables",
		Type:        "service",
	}

	server := grpcsrv.New(nil)

	redis := coreredis.NewDefaultClient()


	deps := &Dependencies{redis:redis}

	ctx := context.Background()

	app.RegisterHandlers(server, deps)

	svc.MustRun(ctx, server)
}