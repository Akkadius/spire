#!/usr/bin/env bash

red=$'\e[1;31m'
org=$'\e[0;33m'
grn=$'\e[1;32m'
yel=$'\e[1;33m'
blu=$'\e[1;34m'
mag=$'\e[1;35m'
cyn=$'\e[1;36m'
end=$'\e[0m'

separator='--------------------------------------------------------------------------------'

directory_spacing=$'  %-30s'
quick_command_spacing=$'  %-35s'

echo "${separator}"
printf "${yel}❯ CNF Development Console ${end}\n"
echo "${separator}"

#############################################
# environment
#############################################

printf "${org} dev-environment${end} %-10s\n"
printf "${quick_command_spacing} %-10s\n" "${grn}up${end}" "Boots containers"
printf "${quick_command_spacing} %-10s\n" "${grn}down${end}" "Halts containers"
printf "${quick_command_spacing} %-10s\n" "${grn}restart${end}" "Restarts containers"
printf "${quick_command_spacing} %-10s\n" "${grn}logs${end}" "Tail container logs "

#############################################
# quick commands
#############################################

printf "${org} shells${end}\n"
printf "${quick_command_spacing} %-10s\n" "${grn}b${end}" "Create a bash shell in the Go workspace container"
printf "${quick_command_spacing} %-10s\n" "${grn}mc${end}" "Create MySQL console shell"

printf "${org} aliases${end}\n"
printf "${quick_command_spacing} %-10s\n" "${grn}awslocal${end}" "Used for making API calls to local AWS mocked services"
printf "${quick_command_spacing} %-10s\n" "${grn}dc${end}" "Shortcut for docker-compose"
printf "${quick_command_spacing} %-10s\n" "${grn}alr${end}" "Refresh development environment aliases"

printf "${org} help${end}\n"
printf "${quick_command_spacing} %-10s\n" "${grn}help (?)${end}" "Show this menu"

SERVICES_ONLINE_COUNT=$(docker-compose ps | grep "Up" | wc -l | xargs)
echo "${separator}"
printf "${cyn}| Containers Online${end} (${yel}${SERVICES_ONLINE_COUNT}${end})\n"
echo "${separator}"

