#! /usr/bin/bash

INTERVIEW=$(grep "SEE INTERVIEW" streets/Buckingham_Place | grep -oP '\d+')

echo "$INTERVIEW"
cat interviews/interview-$INTERVIEW
echo "$MAIN_SUSPECT"
