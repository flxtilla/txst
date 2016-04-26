package txst

import (
	"os"
	"testing"

	"github.com/flxtilla/app"
	"github.com/flxtilla/cxre/log"
)

func txstLogger() app.Config {
	return app.DefaultConfig(
		func(a *app.App) error {
			a.SwapLogger(log.New(os.Stderr, log.LInfo, log.DefaultNullFormatter()))
			return nil
		},
	)
}

func txstingApp(t *testing.T, name string, conf ...app.Config) *app.App {
	conf = append(conf, app.Mode("Testing", true))
	a := app.New(name, conf...)
	err := a.Configure()
	if err != nil {
		t.Errorf("Error in app configuration: %s", err.Error())
	}
	return a
}

func TxstingApp(t *testing.T, name string, conf ...app.Config) *app.App {
	conf = append(conf, txstLogger())
	return txstingApp(t, name, conf...)
}

func VerboseTxstingApp(t *testing.T, name string, conf ...app.Config) *app.App {
	return txstingApp(t, name, conf...)
}
