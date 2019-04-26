# WIP - goship

Note: calls marked w/ PC are _"Porting Complete"_

## A - LIVE DATA SERVICES:

1. __Active trips__:
```
http://<REMOTE_IP>:8000/activeTripsNow?id_portinformer=<id_portinformer>
```

2. __Roadstead now__ (PC):
```
http://<REMOTE_IP>:8000/anchored/<id_portinformer>
```

3. __Moored now__:
```
http://<REMOTE_IP>:8000/mooredNow?id_portinformer=<id_portinformer>&id_activity=5
```

4. __Arrivals now__:
```
http://<REMOTE_IP>:8000/arrivalsNow?id_portinformer=<id_portinformer>
```

5. __Departures now__:
```
http://<REMOTE_IP>:8000/departuresNow?id_portinformer=<id_portinformer>
```

6. __Arrival previsions__:
```
http://<REMOTE_IP>:8000/arrivalPrevisionsNow?id_portinformer=<id_portinformer>
```

7. __Shipped goods__:
```
http://<REMOTE_IP>:8000/shippedGoodsNow?id_portinformer=<id_portinformer>
```

8. __RO/RO + RO/PAX__:
```
http://<REMOTE_IP>:8000/trafficListNow?id_portinformer=<id_portinformer>
```

## B - ARCHIVE DATA SERVICES:

1. __Trips archive [global recap, one row per trip]__:
```
http://<REMOTE_IP>:8000/tripsArchive?id_portinformer=<ID_PORTINFORMER>
```
2. __Trips archive [global recap, one row per commercial operation]__:
```
http://<REMOTE_IP>:8000/tripsArchiveMultiRows?id_portinformer=<ID_PORTINFORMER>
```
3. __Trip data archive__ [shipreport core]:
```
http://<REMOTE_IP>:8000/shipReportList?id_portinformer=<ID_PORTINFORMER>
```

4. __Trip data archive detailed__ [shipreport]:
```   
http://<REMOTE_IP>:8000/shipReportDetails?id_portinformer=<ID_PORTINFORMER>
```

5. __Arrivals archive__:
```
http://<REMOTE_IP>:8000/arrivalsArchive?id_portinformer=<id_portinformer>
```

6. __Departures archive__:
```
http://<REMOTE_IP>:8000/departuresArchive?id_portinformer=<id_portinformer>
```
7. __Shipped goods archive__:
```
http://<REMOTE_IP>:8000/shippedGoodsArchive?id_portinformer=<id_portinformer>
```

8. __Traffic list archive__:
```
http://<REMOTE_IP>:8000/trafficListArchive?id_portinformer=<id_portinformer>
```



## C - DAILY REGISTER SERVICES:

1. __Arrivals:__
```
http://<REMOTE_IP>:8000/registerArrivals?id_portinformer=<ID_PORTINFORMER>
```
2. __Moored:__
```
http://<REMOTE_IP>:8000/registerMoored?id_portinformer=<ID_PORTINFORMER>
```
3. __Roadstead:__
```
http://<REMOTE_IP>:8000/registerRoadstead?id_portinformer=<ID_PORTINFORMER>
```

4. __Departures:__
```
http://<REMOTE_IP>:8000/registerDepartures?id_portinformer=<ID_PORTINFORMER>
```

5. __Shiftings:__
```
http://<REMOTE_IP>:8000/registerShiftings?id_portinformer=<ID_PORTINFORMER>
```

6. __Arrival previsions:__
```
http://<REMOTE_IP>:8000/registerPlannedArrivals?id_portinformer=<ID_PORTINFORMER>
```

7. __Shipped goods:__
```
http://<REMOTE_IP>:8000/registerShippedGoods?id_portinformer=<ID_PORTINFORMER>
```

8. __RO/RO + RO/PAX:__
```
http://<REMOTE_IP>:8000/registerTrafficList?id_portinformer=<ID_PORTINFORMER>
```

## D - BUSINESS INTELLIGENCE SERVICES: ##

1. __Shiftings/maneuverings [per quay/berth]:__
```
http://<REMOTE_IP>:8000/tripsManeuverings?id_portinformer=<ID_PORTINFORMER>
```

2. __Shipped goods recap:__
```
http://<REMOTE_IP>:8000/shippedGoodsRecap?id_portinformer=<ID_PORTINFORMER>
```

3. __RO/RO + RO/PAX recap:__
```
http://<REMOTE_IP>:8000/trafficListRecap?id_portinformer=<ID_PORTINFORMER>
```

## E - METEO DATA ##
1. __Meteo data archive__:
```
http://<REMOTE_IP>:8000/meteoArchive?id_portinformer=<ID_PORTINFORMER>
```



