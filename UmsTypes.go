package WeatherUMS

import "github.com/MetServiceDev/WeatherEventLib"

const (
	warningsService = "warnings"
	forecastService = "forecast"
)

// SubscriptionsType  who is subscribed to what services and
// for which locations
type SubscriptionsType struct {
	Clients []ClientSubType `json:"clients"`
}

type ClientSubType struct {
	ClientId   string                `json:"clientId"`
	ClientName string                `json:"clientName"`
	Services   []ServiceLocationType `json:"services"`
}

type ServiceLocationType struct {
	Service   string   `json:"service"`
	Locations []string `json:"locations"`
}

type ServiceType struct {
	Locations map[string]*ActiveLocationType `json:"activeLocations"`
}

type ActiveLocationType struct {
	LocationId *string                `json:"locationId"`
	Clients    map[string]*ClientType `json:"clients"`
}

// ClientsType  The clients (users)
//
//	type ClientsType struct {
//	 Clients []ClientType `json:"clients"`
//	}
type ClientsType []ClientType

type ClientsMapType struct {
	Clients map[string]ClientType
}

type ClientType struct {
	ClientId     string `json:"clientId"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	EmailAddress string `json:"emailAddress"`
	//LocationIds   []string `json:"locationIds"`
	//Subscriptions []string `json:"subscriptions"`
}

// RuntimeLocations  Generated object with
// an entry for each of the locations we are interested in
// for a service.  Can save as json if that turns out to be
// useful
type RuntimeLocations struct {
	Services map[string]*RtServiceType `json:"services"`
}

type RtServiceType struct {
	Locations map[string]*RtLocationType `json:"locations"`
}

type RtLocationType struct {
	Location WeatherEventLib.LocationType `json:"location"`
	Clients  map[string]ClientType        `json:"clients"`
}

type WarningsRuntime struct {
	Locations map[string]*LocationRuntime `json:"locations"`
}

type LocationRuntime struct {
	Location WeatherEventLib.LocationType `json:"location"`
	Clients  []ClientType                 `json:"clients"`
}

type ForecastsRuntime WarningsRuntime
