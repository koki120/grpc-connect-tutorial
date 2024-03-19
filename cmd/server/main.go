package main

import (
	"context"
	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type handler struct {
	deps string
}

func NewHandler(
	deps string,
) greetv1connect.GreetServiceHandler {
	return &handler{deps: deps}
}

func (h *handler) Greet(ctx context.Context, req *connect.Request[greetv1.GreetRequest]) (*connect.Response[greetv1.GreetResponse], error) {
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("%s, %s!", h.deps, req.Msg.Name),
	})
	return res, nil
}

func main() {
	mux := chi.NewRouter()
	mux.Group(func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{}))
		path, handler := greetv1connect.NewGreetServiceHandler(NewHandler("Hello World"))
		r.Handle(path+"*", handler)
	})

	if err := http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{})); err != nil {
		log.Fatal(err)
	}

}
