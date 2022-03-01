package maps

const (
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	if def, ok := d[key]; ok {
		return def, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrNotFound:
		d[k] = v
		return nil
	case nil:
		return ErrKeyExists
	default:
		return err
	}
}
