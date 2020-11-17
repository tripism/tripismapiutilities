#!/bin/bash
#Git Cache clean-up script

#Set help usage text
HELPUSAGE="
Usage:
gitclean.sh [--help | --h | --?] [--aggressive | --agg]

help parameter will print this help text.

aggressive parameter will cause an aggressive git cache clean.
"
#Set aggressive cache clean OFF
AGGRESSIVE="false"

case "$1" in
        help | -help | --help | -h | --h | ? | -? | --?)
            echo "$HELPUSAGE"
            exit 1
            ;;
         
        aggressive | -aggressive | --aggressive | -agg | --agg)
            AGGRESSIVE="true"
            ;;         
        *)
    echo "$HELPUSAGE"
    exit 1
esac

cd ..
for d in */ ; do
    echo "Processing folder $d"
    cd $d
    if [ "$AGGRESSIVE" = "true" ] ; then
        git gc --aggressive
    else 
        git gc
    fi
    cd ..
done