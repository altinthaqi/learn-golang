package main

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

var (
	ErrNotFound         = DictionaryErr("cannot find the word you are looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (d Dictionary) Search(k string) (string, error) {
	def, ok := d[k]

	if !ok {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dictionary) Add(k, v string) error {

	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
