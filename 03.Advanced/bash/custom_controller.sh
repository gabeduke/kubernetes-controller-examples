#!/bin/bash

# Interval in seconds between checks
INTERVAL=10

# Previous state of Pods
PREVIOUS_STATE=""

while true; do
    # Fetch the current state of Pods in the default namespace
    CURRENT_STATE=$(kubectl get pods --no-headers -n default)

    # Compare with previous state
    if [ "$CURRENT_STATE" != "$PREVIOUS_STATE" ]; then
        echo "Change detected in Pods. Performing actions..."
        # Insert custom logic here
        # For example, you might log the change, send a notification, or trigger some other script

        # Update the previous state
        PREVIOUS_STATE="$CURRENT_STATE"
    else
        echo "No change in Pods. Checking again in $INTERVAL seconds..."
    fi

    # Wait for the next interval
    sleep $INTERVAL
done
