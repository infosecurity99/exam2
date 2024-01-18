package storage

import "exam2/api/models"

type IStorage interface {
	CloseDB()
	City() ICityRepo
	Customer() ICustomerRepo
	Driver() IDriverRepo
	Car() ICarRepo
	Trip() ITripRepo
	TripCustomer() ITripCustomerRepo
}

//cities
type ICityRepo interface {
	Create(city models.CreateCity) (string, error)
	Get(models.PrimaryKey) (models.City, error)
	GetList(req models.GetListRequest) (models.CitiesResponse, error)
	Update(city models.City) (string, error)
	Delete(models.PrimaryKey) error
}

//cistomers
type ICustomerRepo interface {
	Create(customer models.CreateCustomer) (string, error)
	Get(models.PrimaryKey) (models.Customer, error)
	GetList(req models.GetListRequest) (models.CustomersResponse, error)
	Update(customer models.Customer) (string, error)
	Delete(models.PrimaryKey) error
}

//drivers
type IDriverRepo interface {
	Create(driver models.CreateDriver) (string, error)
	Get(models.PrimaryKey) (models.Driver, error)
	GetList(req models.GetListRequest) (models.DriversResponse, error)
	Update(driver models.Driver) (string, error)
	Delete(models.PrimaryKey) error
}

//cars
type ICarRepo interface {
	Create(car models.CreateCar) (string, error)
	Get(id string) (models.Car, error)
	GetList(req models.GetListRequest) (models.CarsResponse, error)
	Update(car models.Car) (string, error)
	Delete(id string) error
	UpdateCarStatus(updateCarStatus models.UpdateCarStatus) error
	UpdateCarRoute(updateCarRoute models.UpdateCarRoute) error
}

//trips
type ITripRepo interface {
	Create(trip models.CreateTrip) (string, error)
	Get(id string) (models.Trip, error)
	GetList(req models.GetListRequest) (models.TripsResponse, error)
	Update(trip models.Trip) (string, error)
	Delete(id string) error
}

type ITripCustomerRepo interface {
	Create(tripCustomer models.CreateTripCustomer) (string, error)
	Get(models.PrimaryKey) (models.TripCustomer, error)
	GetList(req models.GetListRequest) (models.TripCustomersResponse, error)
	Update(tripCustomer models.TripCustomer) (string, error)
	Delete(models.PrimaryKey) error
}
