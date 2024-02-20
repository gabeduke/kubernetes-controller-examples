#!/bin/bash

# Polling interval in seconds
INTERVAL=5

echo "Watching for Pod changes in the default namespace. Press [CTRL+C] to stop..."
while true; do
    kubectl get pods --no-headers
    echo "Updated at: $(date)"
    sleep $INTERVAL
    clear
done
