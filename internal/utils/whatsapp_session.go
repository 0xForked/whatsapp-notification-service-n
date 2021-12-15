package utils

import (
	"encoding/gob"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"os"
)

// WriteSession - Write/Store Whatsapp Session on temporary directory
func WriteSession(session whatsapp.Session) error {
	file, err := os.Create(
		os.Getenv("WAC_SESSION_PATH") +
			"/whatsapp_session.gob")
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

// ReadSession - Read Whatsapp Session from a stored file on temporary directory
func ReadSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Create(
		os.Getenv("WAC_SESSION_PATH") +
			"/whatsapp_session.gob")
	if err != nil {
		return session, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

// LogoutSession - Logout Whatsapp Session and remove the file on temporary directory
func LogoutSession(wac *whatsapp.Conn) error {
	defer func() {
		fmt.Println("Disconnecting..")
		_, _ = wac.Disconnect()
	}()

	err := wac.Logout()
	if err != nil {
		return err
	}

	_ = os.Remove(os.Getenv("WAC_SESSION_PATH") +
		"/whatsapp_session.gob")

	fmt.Println("Logout success..")

	return nil
}
