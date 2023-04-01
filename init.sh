#!/bin/bash

read -p "Enter module name: " VALUE
echo "replacing 'github.com/SFDC/tracks-mono/cmd/template' with:"
echo "           $VALUE"

rg github.com/SFDC/tracks-mono/cmd/template -l -g '!init.sh' | xargs sed -i -e "s+github.com/SFDC/tracks-mono/cmd/template+$VALUE+g"

rm **/**-e
