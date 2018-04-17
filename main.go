/*

Source of book tutorial
https://andrew-sledge.gitbooks.io/the-unofficial-pop-book/content/

*/
package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/relato/poptest2/models"
)

func main() {
	tx, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}
	if false {

		jessica := models.User{Title: "Ms.", FirstName: "Jessica", LastName: "Jones", Bio: "Private security, super hero."}
		_, err = tx.ValidateAndSave(&jessica)
		if err != nil {
			log.Panic(err)
		}
		luke := models.User{Title: "Mr.", FirstName: "Luke", LastName: "Cage", Bio: "Hero for hire."}
		_, err = tx.ValidateAndSave(&luke)
		if err != nil {
			log.Panic(err)
		}
		danny := models.User{Title: "Mr.", FirstName: "Danny", LastName: "Rand", Bio: "Martial artist."}
		_, err = tx.ValidateAndSave(&danny)
		if err != nil {
			log.Panic(err)
		}
		matt := models.User{Title: "Mr.", FirstName: "Matthew", LastName: "Murdock", Bio: "Lawyer, sees with no eyes."}
		_, err = tx.ValidateAndSave(&matt)
		if err != nil {
			log.Panic(err)
		}
		frank := models.User{Title: "Mr.", FirstName: "Frank", LastName: "Castle", Bio: "USMC, badass."}
		_, err = tx.ValidateAndSave(&frank)
		if err != nil {
			log.Panic(err)

		}

	}

	if false {
		// Case 1: good entry
		id := "c913c676-f145-40be-b911-ea01d0de745f" // frank
		jessica := models.User{}
		err = tx.Find(&jessica, id)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Success!\n")
			fmt.Printf("%v\n", jessica)
		}

		// Case 2: bad entry
		id = "00000000-0000-0000-0000-000000000000" // doesn't exist
		nuthin := models.User{}
		err = tx.Find(&nuthin, id)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Success!\n")
			fmt.Printf("%v\n", nuthin)
		}
	}

	// query all records
	if false {
		users := []models.User{}
		err = tx.All(&users)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Success!\n")
			fmt.Printf("%v\n", users)
		}
	}

	// query all with some filter
	if false {
		query := tx.Where("last_name = 'Rand' OR last_name = 'Murdock'")
		users := []models.User{}
		err = query.All(&users)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Success!\n")
			fmt.Printf("%v\n", users)
		}
	}

	// update records with some filter
	if false {
		query := tx.Where("title = 'Ms.'")
		users := []models.User{}
		err = query.All(&users)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			for i := 0; i < len(users); i++ {
				user := users[i]
				user.Title = "Mrs."
				tx.ValidateAndSave(&user)
				fmt.Print("Success!\n")
				fmt.Printf("%v\n", user)
			}
		}
	}
	// multiple updates
	if false {
		users := []models.User{}
		err = tx.All(&users)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			for i := 0; i < len(users); i++ {
				user := users[i]
				user.Location = "NYC, NY"
				tx.ValidateAndSave(&user)
				fmt.Print("Success!\n")
				fmt.Printf("%v\n", user)
			}
		}
	}

	// delete frank
	if false {
		id := "8df5d7d5-760e-4539-9f2e-1053c49228a6"
		frank := models.User{}
		err = tx.Find(&frank, id)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Success! - Now delete it.\n")
			tx.Destroy(&frank)
		}

		frank_test := models.User{}
		err = tx.Find(&frank_test, id)
		if err != nil {
			fmt.Print("Record not found!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("I shouldn't have found it.\n")
		}
	}

	// delete multiple records
	if false {
		query := tx.Where("last_name = 'Rand' OR last_name = 'Murdock'")
		users := []models.User{}
		err = query.All(&users)
		if err != nil {
			fmt.Print("ERROR!\n")
			fmt.Printf("%v\n", err)
		} else {
			fmt.Print("Found users - now delete them.\n")
			fmt.Printf("%v\n", users)
			for i := 0; i < len(users); i++ {
				user := users[i]
				tx.Destroy(&user)
			}
		}
	}

	// join with other table ( favorite_food )
	if false {
		foods := [3]string{"cake", "steak", "beer"}
		users := []models.User{}
		err = tx.All(&users)

		for i := 0; i < len(users); i++ {
			userRef := users[i]
			favoriteFood := models.FavoriteFood{User: userRef.ID, Food: foods[i]}
			fmt.Printf("%v\n", favoriteFood)
			_, err = tx.ValidateAndSave(&favoriteFood)
			if err != nil {
				log.Panic(err)
			}
		}
	}
	type Favorite struct {
		FirstName string `json:"first_name" db:"first_name"`
		LastName  string `json:"last_name" db:"last_name"`
		Food      string `json:"food" db:"food"`
	}
	if false {

		favorites := []Favorite{}
		allFoods := tx.Where("favorite_foods.food IS NOT NULL")
		query := allFoods.LeftJoin("users", "favorite_foods.user=users.id")

		sql, args := query.ToSQL(&pop.Model{Value: models.FavoriteFood{}}, "favorite_foods.food",
			"users.first_name", "users.last_name")
		err = allFoods.RawQuery(sql, args...).All(&favorites)
		for i := 0; i < len(favorites); i++ {
			fmt.Printf("%s %s really loves %s\n", favorites[i].FirstName, favorites[i].LastName, favorites[i].Food)
		}
	}

	// testando LocationValidator com ERRO
	if false {
		frank := models.User{Title: "Mr.", FirstName: "Frank", LastName: "Castle", Bio: "USMC, badass.", Location: "nowhere"}
		verrs, err := tx.ValidateAndSave(&frank)
		if verrs.Count() > 0 {
			log.Println(fmt.Sprintf("ERROR WHILE SAVING: %s\n", verrs))
		}
		if err != nil {
			log.Panic(err)
		}
	}
	// testando LocationValidator SEM ERRO
	if false {

		frank := models.User{Title: "Mr.", FirstName: "Frank", LastName: "Castle", Bio: "USMC, badass.", Location: "Hoboken, NJ"}
		verrs, err := tx.ValidateAndSave(&frank)
		if verrs.Count() > 0 {
			log.Println(fmt.Sprintf("ERROR WHILE SAVING: %s\n", verrs))
		}
		if err != nil {
			log.Panic(err)
		}
	}
	if false {

		peter := models.User{Title: "Mr.", FirstName: "Peter", LastName: "Parker", Bio: "Student", Location: "Queens, New York", Image: "https://upload.wikimedia.org/wikipedia/en/3/35/Amazing_Fantasy_15.jpg"}
		_, err = tx.ValidateAndSave(&peter)
		if err != nil {
			log.Panic(err)
		}
	}
}
