#! /usr/bin/bash

find -type f  -exec basename {} \; -name '*.sh' |sort -r| sed 's/.sh//g' 
