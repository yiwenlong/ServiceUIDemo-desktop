#!/usr/bin/env bash

work_home=$1

launchd_des_file=~/Library/LaunchAgents/com.1wenlong.server.plist
if [ -f "$launchd_des_file" ]; then
    launchctl unload $launchd_des_file
    echo "服务已卸载: com.1wenlong.server"
    rm -f $launchd_des_file
    echo "配置文件已移除: $launchd_des_file"
fi
echo "服务已停止！"