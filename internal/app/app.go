package app

import (
	"context"
	"github.com/Azaliya1995/music_library/internal/config"
	"github.com/Azaliya1995/music_library/pkg/log"
)

type ServerApplication struct {
	conf config.Config
}

func NewServerApplication(cfg *config.Config) *ServerApplication {
	app := &ServerApplication{
		conf: *cfg,
	}
	return app
}

func (app *ServerApplication) Run(_ context.Context) error {
	log.Info("Music Library Server RUN begin ... ")

	log.Info("Music Library Server RUN finish ... ")

	return nil
}

func (app *ServerApplication) Shutdown(_ context.Context) error {
	log.Info("Music Library Server stopping...")

	//ctx := context.Background()

	//if err := app.srv.Shutdown(ctx); err != nil {
	//	if !errors.Is(err, http.ErrServerClosed) {
	//		log.Error(err.Error())
	//	}
	//}

	log.Info("Music Library Server stopped...")

	return nil
}
