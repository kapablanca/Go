#! /usr/bin/bash

curl https://platform.zone01.gr/assets/superhero/all.json | jq -r '.[] | select (.id=='$HERO_ID') | .connections.relatives'