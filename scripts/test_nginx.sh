#!/bin/bash

# Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/LiangNing7/miniblog.



for n in $(seq 1 1 10)
do
    nohup curl -XGET curl http://liangning7.cn:7777/healthz &>/dev/null
done
