#! /usr/bin/bash

INTERVIEW=$(grep "SEE INTERVIEW"  the-final-cl-test/mystery/streets/Buckingham_Place | grep -oP '\d+')

echo "$INTERVIEW"
cat the-final-cl-test/mystery/interviews/interview-$INTERVIEW
echo "$MAIN_SUSPECT"
