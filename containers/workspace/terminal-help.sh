#!/usr/bin/env bash

red=$'\e[1;31m'
grn=$'\e[1;32m'
yel=$'\e[1;33m'
blu=$'\e[1;34m'
mag=$'\e[1;35m'
cyn=$'\e[1;36m'
end=$'\e[0m'

directory_spacing=$' %-30s'
quick_command_spacing=$'  %-30s'

separator='--------------------------------------------------------------------------------'

echo "${separator}"
printf "${yel} Go Workspace Console ${end}\n"
echo "${separator}"

#############################################
# navigation & shells
#############################################

printf "${directory_spacing} %-10s\n" "${grn}(c)cmd${end}" "Runs project commands (Alias of: go run main.go)"
printf "${directory_spacing} %-10s\n" "${grn}(t)tests${end}" "Runs tests"
printf "${directory_spacing} %-10s\n" "${grn}coverage${end}" "Runs tests with coverage reporting"
echo "${separator}"
printf "${directory_spacing} %-10s\n" "${grn}mc${end}" "MySQL Shell"
printf "${directory_spacing} %-10s\n" "${grn}rc${end}" "Redis Shell"
echo "${separator}"
printf "${directory_spacing} %-10s\n" "${grn}s${end}" "Regenerates Swagger docs"
echo "${separator}"
printf "${directory_spacing} %-10s\n" "${grn}help${end}" "Shows this help menu"
printf "${directory_spacing} %-10s\n" "${grn}?${end}" "Shows this help menu"


printf "\n"
