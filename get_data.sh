#!/bin/bash
for (( ; ; ))
do
    scp -P 1569 -i key -o ConnectTimeout=5 root@localhost:/dev/Sig_data ./
    scp -P 1569 -i key -o ConnectTimeout=5 root@localhost:/dev/Cov_data ./
    sleep 1
done
