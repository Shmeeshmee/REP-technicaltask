package router

import (
	"context"
	"fmt"
	"net/http"
	"re-parteners-tech-task/config"

	"github.com/gorilla/mux"
)

func Router(ctx context.Context, cfg *config.Config) {
	r := mux.NewRouter()

	routes(cfg, r)

	port := fmt.Sprintf(":%s", cfg.Port)
	gopher()
	fmt.Println("app listening at port", port)

	fmt.Println(http.ListenAndServe(port, r).Error())
}

func gopher() {
	fmt.Println("   __             __")
	fmt.Println("  ╱▲╲╲__________╱▲╲╲")
	fmt.Println(" |   ____    ____  |||")
	fmt.Println(" |  /    \\  /    \\ |||")
	fmt.Println(" | |⚫   | |⚫    ||||")
	fmt.Println(" |  \\____/__\\____/ |||")
	fmt.Println(" |     [_ ▲ _]     |||")
	fmt.Println(" |      |_|_|      |||")
	fmt.Println(" |                 |||")
}
