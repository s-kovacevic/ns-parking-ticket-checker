# Parking Ticket Checker
Simple utility that checks if you have a parking ticket in Novi Sad, Serbia.

## Building
1. Replace `PLATE_NUMBER` constant with your plate number
2. ```go build -o parkingcheck```

## Running
1. Move parkingcheck to /usr/local/bin/ and fix permissions if necessary
2. Add crontab job with `crontab -e` : `0 18 * * * /usr/local/bin/parkingcheck > /dev/null 2>&1`
