#!/bin/bash
unset oldTorPWD
unset newTorPWD
unset checknewTorPWD
# Function trim delete the space front and end.
trim() {
    local var=$1
    var="${var#"${var%%[![:space:]]*}"}"
    var="${var%"${var##*[![:space:]]}"}"
    echo -n "$var"
}
# Check TOR service is running.
torServiceStatus=$(service tor status | head -n 1)
torServiceStatus=$(trim "$torServiceStatus")
if [[ $torServiceStatus != "* tor is running" ]]; then
    echo "Please command 'torcontroller --start' to start torserver first"
    echo "Because need TOR server to verify."
    exit 1
fi
# Info about --setpassword option.
echo "Important!"
echo "To protect your persinal torcontroller use."
echo "torcontroller:"
echo "1. Would NOT storage your password in any file."
echo "2. Hash your password and record to tor file."
echo "3. Verify though tor server."
echo "4. Turn off your tor during setting new password and restart after it."
echo "------Accourding tor service, not able to change password during connecting."
echo "If you haven't set up your password."
echo "'torcontroller' is the default password."
echo "Please change it to your own."
echo "              -- torcontroller Dev"
echo -n "Enter old TOR password:"
while IFS= read -r -s -n 1 char; do
    if [[ $char == $'\0' ]]; then
        break
    elif [[ $char == $'\177' ]]; then
        if [[ -n $oldTorPWD ]]; then
            oldTorPWD=${oldTorPWD%?}
            echo -en "\b \b"
        fi
    else
        echo -n '*'
        oldTorPWD+="$char"
    fi
done
echo ""
# Verify that the old password hash is correct.
torVerifyResponse=$(echo -e "AUTHENTICATE \"$oldTorPWD\"\r\nQUIT" | nc -q 1 127.0.0.1 9051 | head -n 1)
torVerifyResponse=$(trim "$torVerifyResponse")
if [[ $torVerifyResponse != "250 OK" ]]; then
    echo -e "\nIncorrect old password. Exiting..."
    exit 1
fi
# User set up new_password.
echo -n "Enter new password:"
while IFS= read -r -s -n 1 char; do
    if [[ $char == $'\0' ]]; then
        break
    elif [[ $char == $'\177' ]]; then
        if [[ -n $newTorPWD ]]; then
            newTorPWD=${newTorPWD%?}
            echo -en "\b \b"
        fi
    else
        echo -n '*'
        newTorPWD+="$char"
    fi
done
echo ""
# read -s -p "Enter new password again: " checknewTorPWD
echo -n "Enter new password again:"
while IFS= read -r -s -n 1 char; do
    if [[ $char == $'\0' ]]; then
        break
    elif [[ $char == $'\177' ]]; then
        if [[ -n $checknewTorPWD ]]; then
            checknewTorPWD=${checknewTorPWD%?}
            echo -en "\b \b"
        fi
    else
        echo -n '*'
        checknewTorPWD+="$char"
    fi
done
echo ""
if [[ "$newTorPWD" != "$checknewTorPWD" ]]; then
    echo "New password input inconsistency."
    exit 1
fi
# Before user set up new password,
# we need to stop TOR server,
# accourding TOR rules.
## Maybe just call stopTorServer.sh script in the future,
## if the script getting better.
kill $(pidof tor)
service tor stop
# Hash the new password and hash it,
# then record the hash code to tor setting.
inbashTorHashPWD=$(tor --hash-password "$newTorPWD" | tail -n 1)
sed -i "/HashedControlPassword/s/.*/HashedControlPassword $inbashTorHashPWD/" /etc/tor/torrc
# Restart tor server
service tor start
# Check new password legal or not.
## torVerifyResponse again. for Dev
torVerifyResponse=$(echo -e "AUTHENTICATE \"$newTorPWD\"\r\nQUIT" | nc -q 1 127.0.0.1 9051 | head -n 1)
torVerifyResponse=$(trim "$torVerifyResponse")
if [[ $torVerifyResponse != "250 OK" ]]; then
    echo "New password legal."
    exit 1
fi
echo "TOR Hashed password updated successfully."
