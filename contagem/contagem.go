package contagem

import (
	"fmt"
	"io"
	"os"
	"time"
)

const lastWord = "Vai!\n"
const inicioContagem = 3

func Main() {
	sleeper := &SleeperConfigurable{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}

func Contagem(buffer io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}

	fmt.Fprint(buffer, lastWord)
}

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct{ Chamadas int }

func (s *SleeperSpy) Sleep() { s.Chamadas++ }

type SleeperStandard struct{}

func (s *SleeperStandard) Sleep() { time.Sleep(1 * time.Second) }

type SpyContagemOperations struct{ Chamadas []string }

func (s *SpyContagemOperations) Sleep() { s.Chamadas = append(s.Chamadas, pausa) }

func (s *SpyContagemOperations) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type SleeperConfigurable struct {
	duracao time.Duration
	sleep   func(time.Duration)
}

func (s *SleeperConfigurable) Sleep() {
	s.sleep(s.duracao)
}

type TempoSpy struct {
	duracaoSleep time.Duration
}

func (t *TempoSpy) Sleep(duracao time.Duration) {
	t.duracaoSleep = duracao
}

const escrita = "escrita"
const pausa = "pausa"
