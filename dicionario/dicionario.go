package dicionario

type Dicionario map[string]string
type ErrDicionario string

const (
	ErrWordNotFound              = ErrDicionario("word not found on dictionary")
	ErrWordExistsOnDictionary    = ErrDicionario("word already exists on dictionary")
	ErrWordDontExistOnDictionary = ErrDicionario("word don't can be updated because it don't exist on dictionary")
)

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(word string) (string, error) {
	output, exist := d[word]

	if !exist {
		return "", ErrWordNotFound
	}

	return output, nil
}

func (d Dicionario) Adicionar(key string, value string) error {
	_, err := d.Busca(key)

	switch err {
	case ErrWordNotFound:
		d[key] = value

	case nil:
		return ErrWordExistsOnDictionary

	default:
		return err
	}

	return nil
}

func (d Dicionario) Atualizar(key string, newValue string) error {
	_, err := d.Busca(key)

	switch err {
	case ErrWordNotFound:
		return ErrWordDontExistOnDictionary

	case nil:
		d[key] = newValue

	default:
		return err
	}

	return nil
}

func (d Dicionario) Remover(key string) error {
	delete(d, key)
	return nil
}
