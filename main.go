package main

import(
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Book struct{
	Author string `json:"author"`
	Title string `json:"title"`
	Publisher string `json:"publisher"`
}


type Repository struct{
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book  := Book()
	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
err = r.DB.Create(&book).error
	if err!= nil {
	context.Status(http.StatusBadRequest).JSON(
		&fiber.Map{"message":"could not create book"})
	return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book has been added"
	})
	return nil

}

func (r*Repository) GetBooks(context *fiber.Ctx)error{
	bookModels := &[]mogels.Books{}

	err := r.DB.Find(bookModels).Error
	if error != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"messsage":"could not get books"})
			return err
	}

	context.Satus(http.StatusOK).JSON(
		&fiber.Map{"message":"book fetched successfully",
		"data": bookModels,	
	})
	return nil
}


func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.deleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks) 

}

func main(){
	err != godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)
	if err != nil{
		log.fatal("could not load the database")
	}

	r:= Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}