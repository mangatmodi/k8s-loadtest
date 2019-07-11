#!/bin/bash
LOCUST_CMD="/usr/local/bin/locust"
LOCUST_DUMMY="/locustfile.py"
echo -n "=> Starting locust "
CMD="$LOCUST_CMD --master -f $LOCUST_DUMMY"
echo "$CMD"
$CMD

