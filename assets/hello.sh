#!/bin/bash

echo Hello is run on the \"{{ now | date "02-01-2006" }}\" day
echo Hello args: $@
echo
echo Following is the current environment:
echo
env
