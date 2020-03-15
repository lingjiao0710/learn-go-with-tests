package maps

type Dictionary map[string]string

/* func (d Dictionary) Search(word string) (string, error) {
	return d[word], nil
} */

const (
	ErrNotFound         = DictionaryErr("找不到你的字段值")
	ErrWordExists       = DictionaryErr("不能增加字段因为已存在")
	ErrWordDoesNotExist = DictionaryErr("由于不存在字段所以不能更新")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, def string) error {
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

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
