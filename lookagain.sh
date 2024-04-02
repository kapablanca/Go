#! /usr/bin/bash

find -type f -name '*.sh' |sort -r| sed 's/.sh//g' | sed 's\./\\g'
