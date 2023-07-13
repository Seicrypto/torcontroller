#!/bin/bash
# Restart tor
# Reset counter
stateCounter=0
while true; do
    torResponse=$(systemctl status tor)
    case "$torResponse" in
    *"waiting"*)
        if [ $stateCounter -ge 60 ]; then
            echo "tor failed to start within 60 seconds, exiting."
            exit 1
        fi
        echo "Waiting for tor..."
        sleep 1
        stateCounter=$(($stateCounter + 1))
        continue
        ;;
    *"running"*)
        systemctl restart tor
        break
        ;;
    *"dead"*)
        systemctl start tor
        break
        ;;
    *)
        echo "Unknown tor status:"
        echo "$torResponse"
        break
        ;;
    esac
done
# Check privoxy status.
# Reset counter
stateCounter=0
while true; do
    privoxyResponse=$(systemctl status privoxy)
    case "$privoxyResponse" in
    *"waiting"*)
        if [ $stateCounter -ge 60 ]; then
            echo "Privoxy failed to start within 60 seconds, exiting."
            exit 1
        fi
        echo "Waiting for privoxy..."
        sleep 1
        stateCounter=$(($stateCounter + 1))
        continue
        ;;
    *"running"*)
        echo "Checked privoxy."
        break
        ;;
    *"dead"*)
        systemctl start privoxy
        break
        ;;
    *)
        echo "Unknown privoxy status:"
        echo "$privoxyResponse"
        break
        ;;
    esac
done
echo "New using IP:"
# Now curl package is neccesary for this script.
curl -x 127.0.0.1:8118 http://icanhazip.com/
