package WeatherUMS

import (
  "bytes"
  "context"
  "encoding/json"
  lib "github.com/MetServiceDev/WeatherEventLib"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  log "github.com/sirupsen/logrus"
  "os"
)

const (
  resDir            = "../commonRes/"
  bucketName        = "weather-event-sub"
  clientsFile       = "clients.json"
  locationsFile     = "locations.json"
  subscriptionsFile = "subscriptions.json"
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

// CreateRuntimes  read the json files and construct the Subscriptions object
func CreateRuntimes() (WarningsRuntime, ForecastsRuntime) {

  clients, err := ReadClients()
  if err != nil {
    // TODO
  }
  locations, err := ReadLocations()
  if err != nil {
    // TODO
  }
  subscriptions, err := ReadSubscriptions()
  if err != nil {
    // TODO
  }

  warningsRt := CreateWarningsRuntime(clients, locations, subscriptions)
  forecastsRuntime := ForecastsRuntime{}

  return warningsRt, forecastsRuntime
}

func ReadClients_file() ClientsMapType {

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

func ReadClients() (ClientsMapType, error) {
  log.Infof("Reading from S3, filename=%s", clientsFile)
  cf := clientsFile
  bn := bucketName
  out, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
    Bucket: &bn,
    Key:    &cf,
  })
  defer out.Body.Close()

  if err != nil {
    log.Errorf("Error reading from S3, fname=%s.  Error=%v", clientsFile, err)
    return ClientsMapType{}, err
  }

  buffer := new(bytes.Buffer)
  _, err = buffer.ReadFrom(out.Body)
  if err != nil {
    // TODO error
  }
  jsonStr := buffer.Bytes()
  var clients ClientsType
  err = json.Unmarshal(jsonStr, &clients)
  if err != nil {
    // TODO error
  }

  clientMap := ClientsMapType{
    Clients: make(map[string]ClientType, 0),
  }
  for _, cl := range clients {
    clientMap.Clients[cl.ClientId] = cl
  }

  return clientMap, nil
}

func ReadSubscriptions() (SubscriptionsType, error) {
  log.Infof("Reading from S3, filename=%s", subscriptionsFile)
  sf := subscriptionsFile
  bn := bucketName
  out, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
    Bucket: &bn,
    Key:    &sf,
  })
  defer out.Body.Close()

  if err != nil {
    log.Errorf("Error reading from S3, fname=%s.  Error=%v", subscriptionsFile, err)
    return SubscriptionsType{}, err
  }

  buffer := new(bytes.Buffer)
  _, err = buffer.ReadFrom(out.Body)
  if err != nil {
    // TODO error
  }
  jsonStr := buffer.Bytes()
  var subscriptions SubscriptionsType
  err = json.Unmarshal(jsonStr, &subscriptions)
  if err != nil {
    // TODO error
  }
  return subscriptions, nil
}

type TmpLocationsType struct {
  Locations []lib.LocationType `json:"locations"`
}

func ReadLocations_file() lib.LocationsType {
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

func ReadLocations() (lib.LocationsType, error) {
  log.Infof("Reading from S3, filename=%s", locationsFile)
  lf := locationsFile
  bn := bucketName
  out, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
    Bucket: &bn,
    Key:    &lf,
  })
  if err != nil {
    log.Errorf("Error reading from S3, fname=%s.  Error=%v", locationsFile, err)
    return lib.LocationsType{}, err
  }
  defer out.Body.Close()

  buffer := new(bytes.Buffer)
  _, err = buffer.ReadFrom(out.Body)
  if err != nil {
    // TODO error
  }
  jsonStr := buffer.Bytes()
  var tmpLocations TmpLocationsType
  err = json.Unmarshal(jsonStr, &tmpLocations)
  if err != nil {
    // TODO error
  }
  locationsMap := lib.LocationsType{
    Locations: make(map[string]lib.LocationType),
  }
  for _, loc := range tmpLocations.Locations {
    locationsMap.Locations[loc.LocationId] = loc
  }
  return locationsMap, nil
}

func GetAllLocations() (lib.LocationsType, error) {
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
  subscriptions, err := ReadSubscriptions()
  if err != nil {
    // TODO
  }
  subs, err := ReadClients()
  if err != nil {
    // TODO
  }

  for _, client := range subscriptions.Clients {
    for _, svc := range client.Services {
      if svc.Service == service {
        subscribers = append(subscribers, subs.Clients[client.ClientId])
      }
    }
  }

  return subscribers
}
