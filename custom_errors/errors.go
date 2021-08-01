package custom_errors

import "errors"

func CouldntCompleteYourOperation() error {
	return errors.New("Couldn't complete your request, please try again later")
}

func NoSuchTag() error {
	return errors.New("No such tag..")
}

func CantGetResource() error {
	return errors.New("Can't complete your request at the moment, please try again later..")
}

func ContactUs() error {
	return errors.New("Can't complete your request at the moment, please reach out to us, we would get this resolved")
}

func GenericError() error {
	return errors.New("It's not you, it's us. Please try later.")
}