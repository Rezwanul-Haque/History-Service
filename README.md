# History Service
## API:
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