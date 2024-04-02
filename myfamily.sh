#! /usr/bin/bash

curl https://platform.zone01.gr/assets/superhero/all.json | jq -r -a '.[] | select (.id=='$HERO_ID') | .connections.relatives'