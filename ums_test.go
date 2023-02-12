package WeatherUMS

import (
	"fmt"
	"testing"
)

func TestReadLocations(t *testing.T) {
	locations, err := ReadLocations()
	if err != nil {
		t.Errorf("Error reading locations. %s", err.Error())
	}
	if len(locations.Locations) == 0 {
		t.Errorf("Locations list is empty")
	}
	fmt.Printf("%d locations found.\n", len(locations.Locations))
}

func TestReadClients(t *testing.T) {
	clients, err := ReadClients()
	if err != nil {
		t.Errorf("Error reading clients. %s", err.Error())
	}
	if len(clients.Clients) == 0 {
		t.Errorf("Clients list is empty")
	}
	fmt.Printf("%d clients found.\n", len(clients.Clients))
}

func TestReadSubscriptions(t *testing.T) {
	subscriptions, err := ReadSubscriptions()
	if err != nil {
		t.Errorf("Error reading subscriptions. %s", err.Error())
	}
	if len(subscriptions.Clients) == 0 {
		t.Errorf("Subscriptions list is empty")
	}
	fmt.Printf("%d subscriptions found.\n", len(subscriptions.Clients))
}
