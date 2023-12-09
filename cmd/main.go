package main

import (
	"fmt"

	"github.com/br4tech/video-aulas-go/internal/infra/database"
	"github.com/br4tech/video-aulas-go/internal/infra/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(localhost:3306)/video_aula_app?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha de conexao!")
	}

	adapter := database.NewGormAdapter(db)
	db = adapter.GetDB()

	db.AutoMigrate(
		&model.User{},
		&model.Post{},
	)

	user := model.User{
		Name:  "Guilherme",
		Email: "Guilherme.silva@gmail.com",
		Post: []model.Post{
			{Title: "Titulo A", Content: "Conteudo do titulo A"},
			{Title: "Titulo B", Content: "Conteudo do titulo B"},
			{Title: "Titulo C", Content: "Conteudo do titulo C"},
		},
	}

	db.Create(&user)

	var userCreted model.User

	db.Preload("Post").First(&userCreted, user.ID)

	fmt.Printf("Usuario: %s \n", userCreted.Name)
	fmt.Printf("Posts: \n")
	for _, post := range userCreted.Post {
		fmt.Printf("Title: %s \n", post.Title)
		fmt.Printf("Conteudo: %s \n", post.Content)
	}
}
