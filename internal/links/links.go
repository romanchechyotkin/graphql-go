package links

import (
	"fmt"
	database "graphql-go/internal/pkg/db/postgresql"
	"graphql-go/internal/users"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (l Link) Save() int64 {
	query := "INSERT INTO public.links (title, address) VALUES ($1, $2) RETURNING id"

	var id int64
	err := database.DB.QueryRow(query, l.Title, l.Address).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

	log.Print("Row inserted!")
	return id
}

func GetAll() []Link {
	query := "SELECT id, title, address FROM links"

	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
