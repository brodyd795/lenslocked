package main

import (
	"fmt"

	"github.com/brodyd795/lenslocked/models"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// import (
// 	"database/sql"
// 	"fmt"

// )

// func main() {
// 	config := PostgresConfig{
// 		Host:     "localhost",
// 		Port:     "5432",
// 		User:     "baloo",
// 		Password: "junglebook",
// 		Database: "lenslocked",
// 		SSLMode:  "disable",
// 	}
// 	db, err := sql.Open("pgx", config.String())

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT,
// 		email TEXT NOT NULL
// 	  );

// 	  CREATE TABLE IF NOT EXISTS orders (
// 		id SERIAL PRIMARY KEY,
// 		user_id INT NOT NULL,
// 		amount INT,
// 		description TEXT
// 	  );`)

// 	  if err != nil {
// 		panic(err)
// 	  }

// 	  fmt.Println("tables created")

// 	//   name := "Brody Dingel"
// 	//   email := "brody@dingel.dev"

// 	// row := db.QueryRow(`
// 	// INSERT INTO users (name, email)
// 	// VALUES ($1, $2) RETURNING id;`, name, email)

// 	// var id int
// 	// err = row.Scan(&id)

// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Println("User created. id = ", id)

// 	// id := 1
// 	// id := 200 // Choose an ID NOT in your users table

// 	// row := db.QueryRow(`
// 	// SELECT name, email
// 	// FROM users
// 	// WHERE id=$1;`, id)

// 	// var name, email string

// 	// err = row.Scan(&name, &email)

// 	// if err == sql.ErrNoRows {
// 	// 	fmt.Println("Error, no rows!")
// 	//   }
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

// 	// userID := 1 // Pick an ID that exists in your DB

// 	// for i := 0; i<=5; i++ {
// 	// 	amount := i*100
// 	// 	desc := fmt.Sprintf("Fake order id: #%d", i)
// 	// 	_, err := db.Exec(`
//     // INSERT INTO orders(user_id, amount, description)
//     // VALUES($1, $2, $3)`, userID, amount, desc)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }
// 	// fmt.Println("Created fake orders.")

// 	type Order struct {
// 		ID int
// 		UserID int
// 		Amount int
// 		Description string
// 	}

// 	var orders []Order

// 	userID := 1 // Use the same ID you used in the previous lesson
// 	rows, err := db.Query(`
// 		SELECT id, amount, description
// 		FROM orders
// 		WHERE user_id=$1`, userID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var order Order
// 		order.UserID = userID
// 		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
// 		if err != nil {
// 			panic(err)
// 		}

// 		orders = append(orders, order)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("orders: ", orders)
// }

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err  := models.Open(cfg)

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("connected!")
	
	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("bob@bob.com", "bob123")

	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}