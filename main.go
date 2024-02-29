package main

import (
	"context"
	"log"

	"github.com/richardktran/KMY-Drug-Store/bootstrap"
	"github.com/richardktran/KMY-Drug-Store/pkg/env"
	"go.uber.org/fx"
)

func init() {
	env.Setup()
}

func main() {
	app := fx.New(
		bootstrap.Module,
	)
	ctx := context.Background()
	err := app.Start(ctx)
	defer app.Stop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
