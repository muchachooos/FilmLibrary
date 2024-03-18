package handler

import (
	"filmlib/model"
	"filmlib/storage"
)

type Server struct {
	Config  model.Config
	Storage *storage.Storage
}
