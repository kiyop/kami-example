package app

import (
	"core/app/config"
	"core/util"

	"net/http"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/guregu/kami"
	"golang.org/x/net/context"
)

const (
	keyConfig = "config"
)

var (
	once         = new(sync.Once)
	sharedConfig = config.DefaultConfig()
)

func init() {
	kami.Use("/", initialize)
	kami.Get("/_ah/warmup", func(ctx context.Context, w http.ResponseWriter, r *http.Request) {})
}

func initialize(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	once.Do(func() {
		ds := "/"
		configDir := "config" + ds + util.Env(r) + ds
		configExt := ".toml"
		if _, err := toml.DecodeFile(configDir+"config"+configExt, &sharedConfig); err != nil {
			panic(err)
		}
	})

	ctx = context.WithValue(ctx, keyConfig, sharedConfig)
	return ctx
}

func Config(ctx context.Context) config.Config {
	c, ok := ctx.Value(keyConfig).(config.Config)
	if !ok {
		panic("app: cannot load configure")
	}
	return c
}
