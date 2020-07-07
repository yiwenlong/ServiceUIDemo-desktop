#!/usr/bin/env bash

work_home=$1

echo "Work Home: $work_home"

plist_template_file=$work_home/com.1wenlong.server.template.plist
plist_des_file=$work_home/com.1wenlong.server.plist
server_binary=$work_home/server

if [ ! -f "$plist_template_file" ]; then
  echo "File not found: $plist_template_file"
  exit 1
fi

if [ ! -f "$server_binary" ]; then
  echo "File not found: $server_binary"
  exit 2
fi
sed -e "s:myhomepath:${work_home}:
s:myhomepath:${work_home}:
s:myhomepath:${work_home}:" "$plist_template_file" > "$plist_des_file"
echo "Config file generated: $plist_des_file"
launchd_des_file=~/Library/LaunchAgents/com.1wenlong.server.plist
if [ -f "$launchd_des_file" ]; then
    launchctl unload $launchd_des_file
    rm -f $launchd_des_file
fi

cp "$plist_des_file" "$launchd_des_file"

launchctl load -w $launchd_des_file
echo "server loaded: $launchd_des_file"
launchctl start com.1wenlong.server
echo "server started: http://localhost:8000"