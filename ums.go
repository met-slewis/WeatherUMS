package WeatherUMS

import (
	"encoding/json"
	lib "github.com/MetServiceDev/WeatherEventLib"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	resDir            = "../commonRes/"
	clientsFile       = resDir + "clients.json"
	locationsFile     = resDir + "locations.json"
	subscriptionsFile = resDir + "subscriptions.json"
)

func DoError(level log.Level, message string, err error) {
	if err == nil {
		return
	}
	if level == log.FatalLevel {
		log.Fatalf(message+" %s", err.Error())
	} else if level == log.ErrorLevel {
		log.Errorf(message+" %s", err.Error())
	} else if level == log.WarnLevel {
		log.Warnf(message+" %s", err.Error())
	} else if level == log.InfoLevel {
		log.Infof(message+" %s", err.Error())
	} else if level == log.DebugLevel {
		log.Debugf(message+" %s", err.Error())
	}
}

// BuildUms  read the json files and construct the Subscriptions object
func CreateRuntimes() (WarningsRuntime, ForecastsRuntime) {

	clients := ReadClients()
	locations := ReadLocations()
	subscriptions := ReadSubscriptions()

	warningsRt := CreateWarningsRuntime(clients, locations, subscriptions)
	forecastsRuntime := ForecastsRuntime{}

	return warningsRt, forecastsRuntime
}

func ReadClients() ClientsMapType {
	clientData, err := os.ReadFile(clientsFile)
	DoError(log.FatalLevel, "Error reading client data. %s", err)

	var clients ClientsType
	err = json.Unmarshal(clientData, &clients)
	DoError(log.ErrorLevel, "Unable to unmarshall into ClientsType", err)

	clientMap := ClientsMapType{
		Clients: make(map[string]ClientType, 0),
	}
	for _, cl := range clients {
		clientMap.Clients[cl.ClientId] = cl
	}

	return clientMap
}

func ReadSubscriptions() SubscriptionsType {
	subscriptionsData, err := os.ReadFile(subscriptionsFile)
	DoError(log.FatalLevel, "Error reading subscriptions data. %s", err)

	var subscriptions SubscriptionsType
	err = json.Unmarshal(subscriptionsData, &subscriptions)
	DoError(log.ErrorLevel, "Unable to unmarshall into SubscriptionsType", err)

	return subscriptions
}

type TmpLocationsType struct {
	Locations []lib.LocationType `json:"locations"`
}

func ReadLocations() lib.LocationsType {
	locationsData, err := os.ReadFile(locationsFile)
	DoError(log.FatalLevel, "Error reading locations data. %s", err)

	var locations TmpLocationsType
	err = json.Unmarshal(locationsData, &locations)
	DoError(log.ErrorLevel, "Unable to unmarshall into LocationsType", err)

	locationsMap := lib.LocationsType{
		Locations: make(map[string]lib.LocationType),
	}

	for _, loc := range locations.Locations {
		locationsMap.Locations[loc.LocationId] = loc
	}
	return locationsMap
}

func GetAllLocations() lib.LocationsType {
	return ReadLocations()
}

func CreateWarningsRuntime(clients ClientsMapType, locations lib.LocationsType, subscriptions SubscriptionsType) WarningsRuntime {

	warnings := WarningsRuntime{
		Locations: make(map[string]*LocationRuntime, 0),
	}

	// iterate through the subscriptions looking for all locations of interest
	for _, client := range subscriptions.Clients {
		// look for warnings service entry
		for _, srv := range client.Services {
			if srv.Service == warningsService {
				// add all these locations to the list
				for _, loc := range srv.Locations {
					locRt, ok := warnings.Locations[loc]
					if !ok {
						locRt = &LocationRuntime{
							Location: locations.Locations[loc],
							Clients:  make([]ClientType, 0),
						}
						warnings.Locations[loc] = locRt
					}
					locRt.Clients = append(locRt.Clients, clients.Clients[client.ClientId])
				}
			}
		}
	}
	return warnings
}

func ClientsForService(service string) ClientsType {
	subscribers := make(ClientsType, 0)
	subscriptions := ReadSubscriptions()
	subs := ReadClients()

	for _, client := range subscriptions.Clients {
		for _, svc := range client.Services {
			if svc.Service == service {
				subscribers = append(subscribers, subs.Clients[client.ClientId])
			}
		}
	}

	return subscribers
}

/*
func CreateSubscriptions(clients ClientsType, locations LocationsType) RuntimeLocations {

  rtLocations := RuntimeLocations{
    Services: make(map[string]*RtServiceType),
  }

  // for each client
  for _, client := range clients {
    if len(client.Subscriptions) == 0 {
      continue
    }
    // for each of the subscriptions for this client
    for _, clientSubscription := range client.Subscriptions {
      service, srvOk := rtLocations.Services[clientSubscription]
      if !srvOk {
        service = &RtServiceType{
          Locations: make(map[string]*RtLocationType, 0),
        }
        rtLocations.Services[clientSubscription] = service
      }

      // add all locations for this client
      for _, locationId := range client.LocationIds {
        rtLoc, ok := service.Locations[locationId]
        if !ok {
          // then make a new RtLocationType
          rtLoc = &RtLocationType{
            Location: locations.Locations[locationId],
            Clients:  make(map[string]ClientType, 0),
          }
          service.Locations[locationId] = rtLoc
        }
        // now add the clientId to the (new or existing) ActiveLocation
        rtLoc.Clients[client.ClientId] = client
      }
    }
  }

  return rtLocations
}
*/
