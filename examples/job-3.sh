#!/bin/bash

hnm=`hostname`

if [ true ]
then
  result=1
else
  result=0
fi

if [ $result=1 ]
then
  echo "{\"labels\": {\"host\": \"$hnm\"}, \"results\": {\"my_system_health\": 1} }"
else
  echo "{\"labels\": {\"host\": \"$hnm\"}, \"results\": {\"my_system_health\": 0} }"
fi

exit 0
