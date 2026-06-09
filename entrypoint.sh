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
restcountriesurl = https://restcountries.com/v3.1
opentripmapurl = https://api.opentripmap.com/0.1/en/places

# Third-Party API Configuration Keys
opentripmapkey = ${OPENTRIPMAP_KEY}
EOF

exec ./main