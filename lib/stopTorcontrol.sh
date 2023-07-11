#!/bin/bash
# Check supervisor service should be running.
supervisorResponse=$(supervisorctl status 2>&1)
if [[ $supervisorResponse == *"no such file"* ]]; then
    supervisord -c /etc/supervisor/supervisord.conf
fi
# Check and start tor / privoxy service should be stopped.
torsuperResponse=$(supervisorctl status tor | grep "tor")
if [[ $torsuperResponse == *"RUNNING"* ]]; then
    supervisorctl stop tor
elif [[ $torsuperResponse == *"STOPPED"* ]]; then
    echo "Tor already stopped."
else
    echo "Unknown tor program status!"
fi
privoxysuperResponse=$(supervisorctl status privoxy | grep "privoxy")
if [[ $privoxysuperResponse == *"RUNNING"* ]]; then
    supervisorctl stop privoxy
elif [[ $privoxysuperResponse == *"STOPPED"* ]]; then
    echo "privoxy already stopped."
else
    echo "Unknow privoxy program status!"
fi
