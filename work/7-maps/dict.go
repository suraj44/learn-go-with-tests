package dict

type Dictionary map[string]string

const (
	ErrNotFound        = DictionaryErr("could not find what you were looking for")
	ErrWordExists      = DictionaryErr("cannot add words because it already exists")
	ErrWordDoesntExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err != nil {
		if err == ErrNotFound {
			return ErrWordDoesntExist
		} else {
			return err
		}
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	if err != nil {
		return err
	}

	delete(d, word)
	return nil

}

func (e DictionaryErr) Error() string {
	return string(e)
}
