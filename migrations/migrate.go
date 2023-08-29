package main

import (
	"github.com/girishyt/go-crud/initializers"
	"github.com/girishyt/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(models.Post{})
}
