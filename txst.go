package txst

import (
	"net/http"
	"testing"

	"github.com/flxtilla/cxre/route"
	"github.com/flxtilla/cxre/state"
)

type Tanage func(*testing.T) state.Manage

type TxstApp interface {
	Manage(*route.Route)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
