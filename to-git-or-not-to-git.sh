#! /usr/bin/bash

curl https://platform.zone01.gr/assets/superhero/all.json | jq '.[170].name, .[170].powerstats.strength, .[170].appearance.gender '
