package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type cityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) storage.ICityRepo {
	return &cityRepo{
		db: db,
	}
}

//  create city
func (c *cityRepo) Create(city models.CreateCity) (string, error) {
	id := uuid.New()
	createdAt := time.Now()
	if _, err := c.db.Exec(`INSERT INTO cities VALUES ($1, $2, $3)`, id, city.Name, createdAt); err != nil {
		log.Println("error while inserting data", err.Error())
		return "", err
	}
	return id.String(), nil
}

//getbyidcity
func (c *cityRepo) Get(pKey models.PrimaryKey) (models.City, error) {
	user := models.City{}

	query := `
		select id, name, created_at from cities  where id=$1
`
	if err := c.db.QueryRow(query, pKey.ID).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	); err != nil {
		fmt.Println("error while scanning user", err.Error())
		return models.City{}, err
	}

	return models.City{}, nil
}

//getlistcity
func (c *cityRepo) GetList(req models.GetListRequest) (models.CitiesResponse, error) {
	var (
		cities            = []models.City{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from cities  `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of cities", err.Error())
		return models.CitiesResponse{}, err
	}

	query = `
		SELECT id, name, created_at
			FROM cities
			   
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.CitiesResponse{}, err
	}

	for rows.Next() {
		city := models.City{}

		if err = rows.Scan(
			&city.ID,
			&city.Name,
			&city.CreatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.CitiesResponse{}, err
		}

		cities = append(cities, city)
	}

	return models.CitiesResponse{
		Cities: cities,
		Count:  count,
	}, nil

}

//updatecity
func (c *cityRepo) Update(citykey models.City) (string, error) {
	query := `
        UPDATE cities 
        SET name = $1, created_at = $2
        WHERE id = $3
    `

	_, err := c.db.Exec(query, citykey.Name, citykey.CreatedAt, citykey.ID)
	if err != nil {
		fmt.Println("error while updating cities data:", err.Error())
		return " ", err
	}

	return citykey.ID, nil
}

//delete   city
func (c *cityRepo) Delete(pKey models.PrimaryKey) error {
	query := `
		delete from cities
			where id = $1
`
	if _, err := c.db.Exec(query, pKey.ID); err != nil {
		fmt.Println("error while deleting cities by id", err.Error())
		return err
	}

	return nil

}
