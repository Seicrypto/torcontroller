#!/bin/bash
# Check supervisor service should be running.
supervisorResponse=$(supervisorctl status 2>&1)
if [[ $supervisorResponse == *"no such file"* ]]; then
    supervisord -c /etc/supervisor/supervisord.conf
fi
# Check and start tor / privoxy service should be running.
torsuperResponse=$(supervisorctl status tor | grep "tor")
if [[ $torsuperResponse == *"STOPPED"* ]]; then
    supervisorctl start tor
elif [[ $torsuperResponse == *"RUNNING"* ]]; then
    echo "Tor already started."
else
    echo "Unknown tor program status!"
fi
privoxysuperResponse=$(supervisorctl status privoxy | grep "privoxy")
if [[ $privoxysuperResponse == *"STOPPED"* ]]; then
    supervisorctl start privoxy
elif [[ $privoxysuperResponse == *"RUNNING"* ]]; then
    echo "privoxy already started."
else
    echo "Unknow privoxy program status!"
fi
