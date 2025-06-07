#!/bin/bash


for n in $(seq 1 1 10)
do
    nohup curl -XGET curl http://liangning7.cn:7777/healthz &>/dev/null
done
