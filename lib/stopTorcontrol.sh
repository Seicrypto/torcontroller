#!/bin/bash
service privoxy stop
kill $(pidof tor)
service tor stop
