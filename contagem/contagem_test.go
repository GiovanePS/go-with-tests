package contagem

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Contagem(buffer, &SpyContagemOperations{})

		output := buffer.String()
		expected := "3\n2\n1\nVai!\n"

		if output != expected {
			t.Errorf("resultado %q, expected %q", output, expected)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperations{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		expected := []string{
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(spyImpressoraSleep.Chamadas, expected) {
			t.Errorf("%v function calls, expected %v", spyImpressoraSleep.Chamadas, expected)
		}
	})
}

func TestSleeperConfiguravel(t *testing.T) {
	timeSleep := 5 * time.Second

	timeSpy := &TempoSpy{}
	sleeper := SleeperConfigurable{timeSleep, timeSpy.Sleep}
	sleeper.Sleep()

	if timeSpy.duracaoSleep != timeSleep {
		t.Errorf("should paused by %v, but paused by %v", timeSleep, timeSpy.duracaoSleep)
	}
}
