package main

import (
	"io/ioutil"
	"os"
)

type SessionManager interface {
	Add(uuid string, data string) error
	Fetch(uuid string) (string, error)
	Destroy(uuid string) error
}

type FileSessionManager struct {
	dir string
}

var dirName string = "session"

func NewFileSessionManager() *FileSessionManager {
	return &FileSessionManager{
		dir: dirName,
	}
}

func (fsm *FileSessionManager) Add(uuid string, data string) error {
	d := []byte(data)
	err := ioutil.WriteFile(fsm.path(uuid), d, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (fsm *FileSessionManager) Fetch(uuid string) (string, error)  {
	content, err := ioutil.ReadFile(fsm.path(uuid))
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (fsm *FileSessionManager) Destroy(uuid string) error {
	err := os.Remove(fsm.path(uuid))
	if err != nil {
		return err
	}
	return nil
}

func (fsm *FileSessionManager) path(uuid string) string {
	filename := uuid + ".txt"
	return fsm.dir + "/" + filename
}
