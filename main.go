package main

import (
	"context"
	"log"
	"os"

	diyanet "github.com/abduelhamit/DiyanetAwqatSalahAPI"
)

func main() {
	log.Println("Starting Ezanizer...")
	defer log.Println("Ezanizer stopped")

	ctx := context.TODO()

	email, ok := os.LookupEnv("DIYANET_EMAIL")
	if !ok || email == "" {
		log.Fatalln("DIYANET_EMAIL is not set")
	}

	password, ok := os.LookupEnv("DIYANET_PASSWORD")
	if !ok || password == "" {
		log.Fatalln("DIYANET_PASSWORD is not set")
	}

	countryCode, ok := os.LookupEnv("COUNTRY_CODE")
	if !ok || countryCode == "" {
		log.Fatalln("COUNTRY_CODE is not set")
	}

	stateCode, ok := os.LookupEnv("STATE_CODE")
	if !ok || stateCode == "" {
		log.Fatalln("STATE_CODE is not set")
	}

	cityCode, ok := os.LookupEnv("CITY_CODE")
	if !ok || cityCode == "" {
		log.Fatalln("CITY_CODE is not set")
	}

	config := &diyanet.Config{
		Email:    email,
		Password: password,
	}

	client := config.NewClient(ctx)
	countries, err := client.GetCountries()
	if err != nil {
		log.Fatalln("Failed to get countries:", err)
	}

	var country diyanet.Country
	found := false
	for _, c := range countries {
		if c.Code == countryCode {
			country = c
			found = true
			break
		}
	}
	if !found {
		log.Fatalf("country with code %q not found", countryCode)
	}

	states, err := client.GetStatesByCountry(country.Id)
	if err != nil {
		log.Fatalln("Failed to get states:", err)
	}

	var state diyanet.State
	found = false
	for _, s := range states {
		if s.Code == stateCode {
			state = s
			found = true
			break
		}
	}
	if !found {
		log.Fatalf("state with code %q not found", stateCode)
	}

	cities, err := client.GetCitiesByState(state.Id)
	if err != nil {
		log.Fatalln("Failed to get cities:", err)
	}

	var city diyanet.City
	found = false
	for _, c := range cities {
		if c.Code == cityCode {
			city = c
			found = true
			break
		}
	}
	if !found {
		log.Fatalf("city with code %q not found", cityCode)
	}

	cityDetail, err := client.GetCityDetail(city.Id)
	if err != nil {
		log.Fatalln("Failed to get city detail:", err)
	}

	log.Printf("City Detail: %+v\n", cityDetail)
}
