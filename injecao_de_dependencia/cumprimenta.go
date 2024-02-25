package injecao

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(w io.Writer, name string) {
	fmt.Fprintf(w, "Olá, %s", name)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "Mundo")
}

func Init() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}
}
