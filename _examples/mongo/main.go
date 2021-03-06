package main

import (
	"fmt"
	"github.com/gosexy/db"
	_ "github.com/gosexy/db/mongo"
	"github.com/gosexy/sugar"
)

var settings = db.DataSource{
	Host:     "debian",
	Database: "dev",
}

func main() {

	sess, err := db.Open("mongo", settings)

	if err != nil {
		panic(err)
	}

	defer sess.Close()

	animals, _ := sess.Collection("animals")

	animals.Truncate()

	animals.Append(db.Item{
		"animal": "Bird",
		"young":  "Chick",
		"female": "Hen",
		"male":   "Cock",
		"group":  "flock",
	})

	animals.Append(db.Item{
		"animal": "Bovidae",
		"young":  "Calf",
		"female": "Cow",
		"male":   "Bull",
		"group":  "Herd",
	})

	animals.Append(db.Item{
		"animal": "Canidae",
		"young":  sugar.List{"Puppy", "Pup"},
		"female": "Bitch",
		"male":   "Dog",
		"group":  "Pack",
	})

	items := animals.FindAll()

	for _, item := range items {
		fmt.Printf("animal: %s, young: %s\n", item["animal"], item["young"])
	}

}
