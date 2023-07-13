#!/bin/bash
# Stop privoxy first.
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
        systemctl stop privoxy
        break
        ;;
    *"dead"*)
        echo "Privoxy stopped."
        break
        ;;
    *)
        echo "Unknown privoxy status:"
        echo "$privoxyResponse"
        break
        ;;
    esac
done
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
        systemctl stop tor
        break
        ;;
    *"dead"*)
        echo "tor stopped."
        break
        ;;
    *)
        echo "Unknown tor status:"
        echo "$torResponse"
        break
        ;;
    esac
done
unset torResponse
unset privoxyResponse
echo "Stop command succeeded."
