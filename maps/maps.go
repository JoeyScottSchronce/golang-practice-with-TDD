package main

type Dict map[string]string
type DictErr string

const (
	ErrNotFound         = DictErr("can't find the word you're searching for")
	ErrWordExists       = DictErr("word already exists")
	ErrWordDoesNotExist = DictErr("word does not exist in dict")
)

func (e DictErr) Error() string {
	return string(e)
}

func (d Dict) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dict) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dict) Update(word, newdef string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newdef
	default:
		return err
	}

	return nil
}

func (d Dict) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
