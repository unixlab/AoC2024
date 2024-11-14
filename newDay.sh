#!/bin/bash -x

lastDay=$(find cmd -name "day*.go" | grep -v XX | sort -n | tail -1 | sed 's#cmd/day##' | sed 's/.go//g')
nextDay=$((10#$lastDay+1))
nextDayLeadingZero=$(printf "day%02d" $nextDay)

cp examples/dayXX examples/${nextDayLeadingZero}
cp inputs/dayXX inputs/${nextDayLeadingZero}
cp cmd/dayXX.go cmd/${nextDayLeadingZero}.go
sed -i "s/dayXX/${nextDayLeadingZero}/g" cmd/${nextDayLeadingZero}.go
cp -r internal/dayXX internal/${nextDayLeadingZero}
sed -i "s/dayXX is our template package/${nextDayLeadingZero} is our package for the ${nextDay}th AoC day/g" internal/${nextDayLeadingZero}/*.go
sed -i "s/dayXX/${nextDayLeadingZero}/g" internal/${nextDayLeadingZero}/*.go
