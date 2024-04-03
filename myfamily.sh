#! /usr/bin/bash

<<<<<<< HEAD
curl  -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select (.id=='$HERO_ID') | .connections.relatives' | sed 's/"//g'
=======
curl  -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select (.id=='$HERO_ID') | .connections.relatives' | sed 's/"//g'
>>>>>>> 491e381c428a836cd05e2de3ddf31e645745aa21
