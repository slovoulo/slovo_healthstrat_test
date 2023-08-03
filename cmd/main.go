package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Patient struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Phone        string
	County       string
	SubCounty    string
	DateOfBirth  time.Time
	Address      string
	Gender       string
	MaritalStatus string
	NextOfKin    NextOfKin
}

type NextOfKin struct {
	ID          uint   `gorm:"primaryKey"`
	PatientID   uint
	Name        string `gorm:"not null"`
	DateOfBirth time.Time
	Gender      string
	Phone       string
	IDNumber    string
	Relationship string
}

var (
	db *gorm.DB
)

func main() {
	dsn := "user=username password=password dbname=database sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Patient{}, &NextOfKin{})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/add", addPatient).Methods("POST")

	http.Handle("/", r)

	serverAddr := "localhost:8080"
	fmt.Printf("Server listening on %s...\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func addPatient(w http.ResponseWriter, r *http.Request) {
	var patient Patient
	var nextOfKin NextOfKin

	

	// Sample patient data
	patient.Name = "John Doe"
	patient.Phone = "123-456-7890"
	patient.County = "Example County"
	patient.SubCounty = "Example Sub-County"
	patient.DateOfBirth = time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	patient.Address = "123 Example St, City"
	patient.Gender = "Male"
	patient.MaritalStatus = "Single"

	// Sample next of kin data
	nextOfKin.Name = "Jane Doe"
	nextOfKin.DateOfBirth = time.Date(1988, 5, 10, 0, 0, 0, 0, time.UTC)
	nextOfKin.Gender = "Female"
	nextOfKin.Phone = "987-654-3210"
	nextOfKin.IDNumber = "ID123456"
	nextOfKin.Relationship = "Spouse"

	patient.NextOfKin = nextOfKin

	// Create patient and next of kin in the database
	if err := db.Create(&patient).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Patient added successfully")
}