package maps

import "errors"

type Dictionary map[string]string // custom type declared here

var dict = Dictionary{} // intilaize a map dic

// can I pass a map as struct field
var (
	errorNotFound    = errors.New("No definition found!") //good practice to reduce repetations
	wordExist        = errors.New("Definition already exist!")
	ErrNoDeftoUpdate = errors.New("No definition found to update!")
)

func (d Dictionary) search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errorNotFound
	}
	return definition, nil // nil means things are going fine, with error = nil, meaning none for pointers and deroved types.
}

func (d Dictionary) add(word, definition string) error {
	_, err := d.search(word)
	switch err {
	case nil:
		return wordExist
	case errorNotFound:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, newDefinintion string) error {
	_, err := d.search(key)
	switch err {
	case errorNotFound:
		return ErrNoDeftoUpdate
	case nil:
		d[key] = newDefinintion
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(key string) {
	delete(d, key) // go's built in delete function
} // not used yet
