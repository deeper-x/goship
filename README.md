# WIP - goship

Note: calls marked w/ *PC are available (because are _"Porting Complete"_)

## A - LIVE DATA SERVICES:

1. __Active trips__:

```bash
http://<REMOTE_IP>:8000/activeTripsNow?id_portinformer=<id_portinformer>
```

2. __At roadstead__ [*PC]:

```bash
http://<REMOTE_IP>:8000/anchored/<id_portinformer>
```

3. __Moored__ [*PC]:
    
```bash
http://<REMOTE_IP>:8000/moored/<id_portinformer>
```

4. __Arrivals now__ [*PC]:

```bash
http://<REMOTE_IP>:8000/arrivals/<id_portinformer>
```

5. __Departures now__ [*PC]:

```bash
http://<REMOTE_IP>:8000/departuresNow?id_portinformer=<id_portinformer>
```

6. __Arrival previsions__ [*PC]:

```bash
http://<REMOTE_IP>:8000/arrivalPrevisions/<id_portinformer>
```

7. __Shipped goods__:

```bash
http://<REMOTE_IP>:8000/shippedGoodsNow?id_portinformer=<id_portinformer>
```

8. __RO/RO + RO/PAX__:

```bash
http://<REMOTE_IP>:8000/trafficListNow?id_portinformer=<id_portinformer>
```

9. __Shifting previsions__:

```bash
http://<REMOTE_IP>:8000/shiftingPrevisions/<id_portinformer>
```

10. __Departure previsions__ [*PC]:
 
```bash
http://<REMOTE_IP>:8000/departurePrevisions/<id_portinformer>
```

## B - ARCHIVE DATA SERVICES:

1. __Trips archive [global recap, one row per trip]__:

```bash
http://<REMOTE_IP>:8000/tripsArchive?id_portinformer=<ID_PORTINFORMER>
```

2. __Trips archive [global recap, one row per commercial operation]__:

```bash
http://<REMOTE_IP>:8000/tripsArchiveMultiRows?id_portinformer=<ID_PORTINFORMER>
```

3. __Trip data archive__ [shipreport core]:

```bash
http://<REMOTE_IP>:8000/shipReportList?id_portinformer=<ID_PORTINFORMER>
```

4. __Trip data archive detailed__ [shipreport]:

```bash   
http://<REMOTE_IP>:8000/shipReportDetails?id_portinformer=<ID_PORTINFORMER>
```

5. __Arrivals archive__:

```bash
http://<REMOTE_IP>:8000/arrivalsArchive?id_portinformer=<id_portinformer>
```

6. __Departures archive__:

```bash
http://<REMOTE_IP>:8000/departuresArchive?id_portinformer=<id_portinformer>
```
7. __Shipped goods archive__:

```bash
http://<REMOTE_IP>:8000/shippedGoodsArchive?id_portinformer=<id_portinformer>
```

8. __Traffic list archive__:

```bash
http://<REMOTE_IP>:8000/trafficListArchive?id_portinformer=<id_portinformer>
```



## C - DAILY REGISTER SERVICES:

1. __Arrivals:__

```bash
http://<REMOTE_IP>:8000/registerArrivals?id_portinformer=<ID_PORTINFORMER>
```
2. __Moored:__

```bash
http://<REMOTE_IP>:8000/registerMoored?id_portinformer=<ID_PORTINFORMER>
```

3. __Roadstead:__
```bash
http://<REMOTE_IP>:8000/registerRoadstead?id_portinformer=<ID_PORTINFORMER>
```

4. __Departures:__
```bash
http://<REMOTE_IP>:8000/registerDepartures?id_portinformer=<ID_PORTINFORMER>
```

5. __Shiftings:__
```bash
http://<REMOTE_IP>:8000/registerShiftings?id_portinformer=<ID_PORTINFORMER>
```

6. __Arrival previsions:__
```bash
http://<REMOTE_IP>:8000/registerPlannedArrivals?id_portinformer=<ID_PORTINFORMER>
```

7. __Shipped goods:__
```bash
http://<REMOTE_IP>:8000/registerShippedGoods?id_portinformer=<ID_PORTINFORMER>
```

8. __RO/RO + RO/PAX:__
```bash
http://<REMOTE_IP>:8000/registerTrafficList?id_portinformer=<ID_PORTINFORMER>
```

## D - BUSINESS INTELLIGENCE SERVICES: ##

1. __Shiftings/maneuverings [per quay/berth]:__
```bash
http://<REMOTE_IP>:8000/tripsManeuverings?id_portinformer=<ID_PORTINFORMER>
```

2. __Shipped goods recap:__
```bash
http://<REMOTE_IP>:8000/shippedGoodsRecap?id_portinformer=<ID_PORTINFORMER>
```

3. __RO/RO + RO/PAX recap:__
```bash
http://<REMOTE_IP>:8000/trafficListRecap?id_portinformer=<ID_PORTINFORMER>
```

## E - METEO DATA ##
1. __Meteo data archive__:
```bash
http://<REMOTE_IP>:8000/meteoArchive?id_portinformer=<ID_PORTINFORMER>
```


# Install and (first) run 
```bash
# Install Go 1.11.9 and verify version
# ref. https://golang.org/doc/install?download=go1.11.9.linux-386.tar.gz
$ export PATH=${PATH}:/usr/local/go/bin/ # FIX w/ your installation path

# Pass vars at runtime or add to .bash_profile
$ export GOPATH=${HOME}/go
$ export GOBIN=${GOPATH}/bin
$ export PATH=${PATH}:${GOBIN}

$ go version
go version go1.11.9 linux/386

# Get goship and deploy
$ go get github.com/deeper-x/goship

$ cd ${GOPATH}/src/github.com/deeper-x/goship
$ go get -d ./...    
$ go install src/goship.go 
$ goship # RUN FOR TEST
Now listening on: http://localhost:8000
Application started. Press CTRL+C to shut down.

# Close test instance
$ <CTRL+C>  

```

# Systemd configuration

```bash
# Service install instructions: 
# This is a service file template you can start from 
$ sudo cat > /usr/lib/systemd/system/goship.service <<HEREDOC

[Unit]
Description=Shipreporting service middleware
Documentation=https://github.com/deeper-x/goship
After=network.target

[Service]
Type=simple
User=<YOUR_USER>
WorkingDirectory=<YOUR_GOPATH>/bin/
ExecStart=<YOUR_GOBIN>/goship
Restart=on-failure

[Install]
WantedBy=multi-user.target

HEREDOC
```


