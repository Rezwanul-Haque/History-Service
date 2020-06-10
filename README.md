# History Service

# Run the app
> create .env file in the **src** folder for example check .env.example
```
./run.sh
```
## If connection refused error occur then
> This error occur as mysql need some time to initialize but go request to connect 
> before mysql connection stabilize 
```
docker-compose down
# then
docker-compose up -d rabbitmq
docker-compose up -d mysql
docker-compose up -d history_service
``` 

## API:
### Health check
> # GET: /api/v1/ping
> ## **Response:**
```
pong
```

> # GET: /api/v1/history

> ## URL Query Params:
```
user_id: 1
start_date: 0
end_date: 1565260459
```
> ## Headers:
```
Content-Type: application/json
RLS-Referrer: vivacomsolutions.com
AppKey: KSAx2mit9nxoyiUW1TnuQ
```

> ## **Response:**
```
{
    "user_id": "01",
    "domain": "vivacomsolutions.com",
    "paths": [
        {
            "client_timestamp_utc": "2019-06-25 06:10:00",
            "server_timestamp_utc": "2019-07-01 12:37:59",
            "lat": 23.794459,
            "lon": 90.400998
        },
        {
            "client_timestamp_utc": "2019-06-25 06:10:00",
            "server_timestamp_utc": "2019-07-01 12:37:59",
            "lat": 23.794459,
            "lon": 90.400998
        }
    ]
}
```