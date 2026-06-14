#!/bin/sh
set -e

cat > /app/conf/app.conf <<EOF
appname = travelSphere
httpport = 8080
runmode = ${RUN_MODE:-dev}
autorender = true
ErrorsShow = true

sessionon = true
sessionname = travelspheresession

# External API Third-Party Base URLs
restcountriesurl = https://api.restcountries.com/countries/v5
opentripmapurl = https://api.opentripmap.com/0.1/en/places

# Third-Party API Configuration Keys
restcountrieskey = ${RESTCOUNTRIES_KEY}
opentripmapkey = ${OPENTRIPMAP_KEY}
EOF

exec ./main
