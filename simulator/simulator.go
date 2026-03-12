package simulator

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/emicklei/go-restful/v3"
	"github.com/threatmate/restfulwrapper"
)

type Simulator struct {
	server   *httptest.Server
	universe Universe
}

func New(ctx context.Context) *Simulator {
	s := &Simulator{
		universe: Universe{},
	}

	api := &API{
		universe: &s.universe,
	}

	container := restful.NewContainer()
	{
		restfulAPI := restfulwrapper.WebService("/api").
			Consumes(restful.MIME_JSON).
			Produces(restful.MIME_JSON)
		{
			ws := restfulAPI.Session().Do(handleAuthentication(&s.universe))
			ws.Register(ctx, "/", api)
		}
		container.Add(restfulAPI.WebService())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r = r.Clone(ctx)

		ctx := r.Context()
		requestContents, err := httputil.DumpRequest(r, true)
		if err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("Could not dump request: %v", err))
			return
		}
		slog.InfoContext(ctx, fmt.Sprintf("Request: %s", requestContents))

		container.ServeHTTP(w, r)
	})

	s.server = httptest.NewServer(mux)
	return s
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Close shuts down the simulator.
func (s *Simulator) Close() {
	if s.server != nil {
		s.server.Close()
	}
}

// URL returns the simulator's URL.
func (s *Simulator) URL() string {
	if s.server == nil {
		return ""
	}
	return s.server.URL
}

// Universe returns the simulator's universe.
func (s *Simulator) Universe() *Universe {
	return &s.universe
}
