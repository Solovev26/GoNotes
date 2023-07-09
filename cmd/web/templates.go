package main

import (
	"awesomeProject/pkg/models"
)

type templateData struct {
	Note  *models.Note
	Notes []*models.Note
}
