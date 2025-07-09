#!/bin/bash

NATS_SERVER="nats://localhost:4222"
NATS_SUBJECT="enshrouded-logins-dev"
SLEEP_SECONDS=5

echo "Sending Logon Event to $NATS_SERVER on subject $NATS_SUBJECT"
cat ./login_event.json | nats pub $NATS_SUBJECT --server $NATS_SERVER

sleep $SLEEP_SECONDS
echo "Sending Logoff Event to $NATS_SERVER on subject $NATS_SUBJECT"
cat ./logoff_event.json | nats pub $NATS_SUBJECT --server $NATS_SERVER

echo "Finished."
