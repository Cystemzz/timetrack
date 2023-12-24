package tracker

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

const WorkspaceVersion = 1

var ErrCannotCreateWorkspace = errors.New("unable to create workspace")

func CreateWorkspace(name, path string) (*Workspace, error) {
	err := openOrCreateDatabase(path)
	if err != nil {
		return nil, err
	}

	w := &Workspace{
		Name:       name,
		Location:   path,
		LastLoaded: time.Now(),
		Version:    WorkspaceVersion,
	}

	if err = w.init(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Workspace) init() error {
	if err := db.AutoMigrate(&Job{}, &TimeCard{}, &WorkspaceMetadata{}); err != nil {
		return ErrCannotCreateWorkspace
	}
	db.Create(&WorkspaceMetadata{
		Name:    w.Name,
		Version: w.Version,
	})
	_ = w.saveWorkspace() // Not being able to save the workspace isn't a huge deal
	return nil
}

func (w *Workspace) saveWorkspace() error {
	path, err := getAppPath()
	if err != nil {
		return err
	}
	path = filepath.Join(path, "settings.json")
	var spaces []Workspace
	if _, err := os.Stat(path); err == nil {
		fileData, err := os.ReadFile(path)
		if fileData != nil {
			return err
		}
		err = json.Unmarshal(fileData, &spaces)
	}
	spaces = append(spaces, *w)
	data, _ := json.Marshal(spaces)
	err = os.WriteFile(path, data, 0600)
	return err
}

func getAppPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path = filepath.Join(path, "TimeTrack")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(path, os.ModeDir)
		if err != nil {
			return "", err
		}
	}
	return path, nil
}
