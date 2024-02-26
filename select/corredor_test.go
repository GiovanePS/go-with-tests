package corredor

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCorredor(t *testing.T) {
	t.Run("Verifica se o servidor mais rápido é retornado corretamente", func(t *testing.T) {
		servidorLento := createServerWithDelay(20 * time.Millisecond)
		servidorRapido := createServerWithDelay(0 * time.Millisecond)

		defer servidorLento.Close()
		defer servidorRapido.Close()

		URLLenta := servidorLento.URL
		URLRapida := servidorRapido.URL

		output, err := Corredor(URLLenta, URLRapida)
		expected := URLRapida

		if err != nil {
			t.Fatalf("Unexpected error ocurred: %v", err.Error())
		}

		if output != expected {
			t.Errorf("output: %q, expected %q", output, expected)
		}
	})

	t.Run("Retorna um erro se o servidor não responder dentro de 10 segundos", func(t *testing.T) {
		servidorA := createServerWithDelay(10 * time.Millisecond)

		defer servidorA.Close()

		_, err := Configurable(servidorA.URL, servidorA.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("Expected error, but none ocurred")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
