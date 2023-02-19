# Weather UMS

Used for the POC and testing purposes.
See up-to-date json examples in the 'res' folder

Maintains some lists:  
- Has a list of 'clients' (users)  
- Has a list of 'locations' (a named lat/long)  
- Has a 'subscription' list.  For each client"
  - which services they subscribe to
  - what locations they are registered with for each of the services

## User / Client
```json
{
  "Id":"1", 
  "name": "Joe", 
  "phone": "+64211655295",
  "emailAddress" : "joe@bloggs.com",
}
```
## Location
```json
{
  "locations": [
    {
      "id": "1",
      "name": "84 Washington Ave",
      "lat": -41.3067,
      "long": 174.7665
    },
    {
      "id": "2",
      "name": "Wellington College",
      "lat": -41.3033,
      "long": 174.782
    }
  ]
}
```

## Subscription
```json
{
  "clients": [
    {
      "clientId": "1",
      "clientName": "Simon@Work",
      "services": [
        {
          "service": "warnings",
          "locations": [
            "*"
          ]
        },
        {
          "service": "forecasts",
          "locations": [
            "1",
            "10"
          ]
        }
      ]
    },
    {
      "clientId": "6",
      "clientName": "Simon@Home",
      "services": [{}]
    }
  ]
}
```
