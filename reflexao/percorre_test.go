package reflexao

import (
	"reflect"
	"testing"
)

type Pessoa struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestPercorre(t *testing.T) {
	t.Run("Testando a função percorre com vários tipos de dados", func(t *testing.T) {
		casos := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{{
			"Struct com um campo string",
			struct{ Nome string }{"Chris"},
			[]string{"Chris"},
		}, {
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		}, {
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Chris", 33},
			[]string{"Chris"},
		}, {
			"Campos aninhados",
			Pessoa{"Chris", Profile{33, "Londres"}},
			[]string{"Chris", "Londres"},
		}, {
			"Ponteiros para coisas",
			&Pessoa{"Chris", Profile{33, "Londres"}},
			[]string{"Chris", "Londres"},
		}, {
			"Slices",
			[]Profile{{33, "Londres"}, {44, "Vancouver"}},
			[]string{"Londres", "Vancouver"},
		}, {
			"Maps",
			map[string]string{"Cow": "Moo", "Sheep": "Bee"},
			[]string{"Moo", "Bee"},
		}}

		for _, test := range casos {
			t.Run(test.Name, func(t *testing.T) {
				var output []string

				percorre(test.Input, func(input string) {
					output = append(output, input)
				})

				if !reflect.DeepEqual(output, test.ExpectedCalls) {
					t.Errorf("output %v, expected %v", output, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("With maps", func(t *testing.T) {
		mapA := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Bee",
		}

		var got []string

		percorre(mapA, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Bee")
	})
}

func assertContains(t testing.TB, total []string, needed string) {
	t.Helper()

	contains := false
	for _, x := range total {
		if x == needed {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to contain %q, but it didn't", total, needed)
	}
}
