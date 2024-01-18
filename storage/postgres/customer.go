package postgres

import (
	"database/sql"
	"exam2/api/models"
	"exam2/storage"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type customerRepo struct {
	db *sql.DB
}

func NewCustomerRepo(db *sql.DB) storage.ICustomerRepo {
	return customerRepo{
		db,
	}
}

//customer createdstring
func (c customerRepo) Create(customer models.CreateCustomer) (string, error) {
	uid := uuid.New()
	createat := time.Now()

	if _, err := c.db.Exec(`insert into customers (id, full_name, phone, email,created_at) values ($1, $2, $3, $4, $5)`,
		uid,
		customer.FullName,
		customer.Phone,
		customer.Email,
		createat,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return "", err
	}

	return uid.String(), nil
}

//customergetbyid
func (c customerRepo) Get(id models.PrimaryKey) (models.Customer, error) {
	customers := models.Customer{}

	query := `
		select id, full_name,phone, email,  created_at from customers
`
	if err := c.db.QueryRow(query).Scan(
		&customers.ID,
		&customers.FullName,
		&customers.Phone,
		&customers.Email,
		&customers.CreatedAt,
	); err != nil {
		fmt.Println("error while scanning customers", err.Error())
		return models.Customer{}, err
	}
	return models.Customer{}, nil
}

//customergetlist
func (c customerRepo) GetList(req models.GetListRequest) (models.CustomersResponse, error) {

	var (
		costumers         = []models.Customer{}
		count             = 0
		countQuery, query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	countQuery = `
		SELECT count(1) from customers  `

	if err := c.db.QueryRow(countQuery).Scan(&count); err != nil {
		fmt.Println("error while scanning count of customers", err.Error())
		return models.CustomersResponse{}, err
	}

	query = `
		SELECT id, full_name,phone,email, created_at
			FROM customers
			   
			    `

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return models.CustomersResponse{}, err
	}

	for rows.Next() {
		costumer := models.Customer{}

		if err = rows.Scan(
			&costumer.ID,
			&costumer.FullName,
			&costumer.Phone,
			&costumer.Email,
			&costumer.CreatedAt,
		); err != nil {
			fmt.Println("error while scanning row", err.Error())
			return models.CustomersResponse{}, err
		}

		costumers = append(costumers, costumer)
	}

	return models.CustomersResponse{
		Customers: costumers,
		Count:     count,
	}, nil

}

//customer update
func (c customerRepo) Update(customerskey models.Customer) (string, error) {
	query := `
        UPDATE customers 
        SET full_name = $1, phone=$2, email=$3,created_at = $4
        WHERE id = $5
    `

	_, err := c.db.Exec(query, customerskey.FullName, customerskey.Phone, customerskey.Email, customerskey.CreatedAt, customerskey.ID)
	if err != nil {
		fmt.Println("error while updating customers data:", err.Error())
		return " ", err
	}

	return customerskey.ID, nil
}

//customerdelete
func (c customerRepo) Delete(id models.PrimaryKey) error {
	query := `
		delete from customers
			where id = $1
`
	if _, err := c.db.Exec(query, id.ID); err != nil {
		fmt.Println("error while deleting customers by id", err.Error())
		return err
	}

	return nil
}
