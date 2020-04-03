#!/usr/bin/env bash

function load() {
  binpath=$1
  workhome=$2
  plist_template_file=$workhome/com.1wenlong.launchdui.daemon.template.plist
  plist_des_file=$workhome/com.1wenlong.launchdui.daemon.plist

  if [ ! -f "$plist_template_file" ]; then
    echo "File not found: $plist_template_file"
    exit 1
  fi

  if [ ! -f "$binpath" ]; then
    echo "File not found: $binpath"
    exit 2
  fi
  sed -e "s:launchdui.command.path:${binpath}:" "$plist_template_file" > "$plist_des_file"
  echo "Config file generated: $plist_des_file"
  launchd_des_file=~/Library/LaunchAgents/com.1wenlong.launchdui.daemon.plist
  cp "$plist_des_file" "$launchd_des_file"
  launchctl load -w $launchd_des_file
  echo "daemon loaded: $launchd_des_file"
}

function unload() {
  launchd_des_file=~/Library/LaunchAgents/com.1wenlong.launchdui.daemon.plist
  launchctl unload $launchd_des_file
  rm $launchd_des_file
}

subcommand=$1

case $subcommand in
load)
  load "$2" "$3"
  ;;
unload)
  unload
  ;;
esac