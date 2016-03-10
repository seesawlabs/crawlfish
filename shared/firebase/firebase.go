package firebase

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/cosn/firebase"
	"github.com/dropbox/godropbox/errors"
	"github.com/seesawlabs/crawlfish/shared/config"
)

var (
	ErrFirebaseChildIsNil = errors.New("Firebase child is nil")
)

type firebaseProvider struct {
	client *firebase.Client
	config config.Firebase
}

type IFirebase interface {
	Push(value interface{}) (*string, error)
	Update(path string, value interface{}) error
}

func (f *firebaseProvider) Push(value interface{}) (*string, error) {
	path := fmt.Sprintf("crawls")

	childValue := f.client.Child(path, nil, nil)
	if childValue == nil {
		logrus.Error(fmt.Sprintf("Path: %s, Value: %s", path, value))
		return nil, ErrFirebaseChildIsNil
	}

	result, err := childValue.Push(value, nil)
	if err != nil {
		logrus.Error("Push error", err)
		return nil, err
	}

	return &result.Url, nil
}

func (f *firebaseProvider) Update(path string, value interface{}) error {
	err := f.client.Update(path, value, nil)
	if err != nil {
		logrus.Error(fmt.Sprintf("Firebase Update Error. Path: %s, Value: %s", path, value))
	}

	return err
}

func NewFirebase(c config.Firebase) IFirebase {
	f := &firebaseProvider{
		client: &firebase.Client{},
		config: c,
	}

	f.client.Init(c.Url, c.Auth, nil)

	return f
}
