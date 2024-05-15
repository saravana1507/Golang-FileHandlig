package help

import "errors"

func Dberror(err error) error {
	if err != nil {
		return errors.New("database not connected")
	}
	return nil
}

func LisenError(err error) error {
	if err != nil {
		return errors.New("server not connected")
	}
	return nil
}

func WriteJsonError(err error) error {
	if err != nil {
		return errors.New("cannot write file")
	}
	return nil
}

func OpenJsonError(err error) error {
	if err != nil {
		return errors.New("cannot found file")
	}
	return nil
}

func CreateError(err error) error {
	if err != nil {
		return errors.New("cannot store data")
	}
	return nil
}
