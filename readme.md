# Weather UMS

Used for the POC and testing purposes.

## User / Client
```json
{
  "Id":"1", 
  "name": "Joe", 
  "phone": "+64211655295",
  "emailAddress" : "joe@bloggs.com",
  "locationIds": ["1", "11", "6"],
  "subscriptions": ["warnings", "forecasts"]
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
## Catalog
```json
{
  "services" : [
    {
      "id" : 1, 
      "name" : "warnings"
    },
    {
      "id" : 2,
      "name" : "forecasts"
    }
  ]
}
```
## Subscription
```json
{
  "subscriptions" : [
    {
      "clientId" : 1,
      "services" : ["warnings", "forecasts"]
    }
  ]
}
```
