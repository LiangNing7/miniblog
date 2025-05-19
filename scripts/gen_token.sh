#!/bin/bash

# Copyright 2025 LiangNing7 <LiangNing7@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/LiangNing7/miniblog.


# 定义Header
HEADER='{"alg":"HS256","typ":"JWT"}'

# 定义Payload
# exp: 过期时间 2025-06-20 00:00:00
# iat: 签发时间 2025-05-20 00:00:00
# nbf: 生效时间 2025-05-20 00:00:00
PAYLOAD='{"exp":1750348800,"iat":1747670400,"nbf":1747670400,"x-user-id":"user-w6irkg"}'

# 定义Secret（用于签名）
SECRET="Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5"

# 1. Base64编码Header
HEADER_BASE64=$(echo -n "${HEADER}" | openssl base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')

# 2. Base64编码Payload
PAYLOAD_BASE64=$(echo -n "${PAYLOAD}" | openssl base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')

# 3. 拼接Header和Payload为签名数据
SIGNING_INPUT="${HEADER_BASE64}.${PAYLOAD_BASE64}"

# 4. 使用HMAC SHA256算法生成签名
SIGNATURE=$(echo -n "${SIGNING_INPUT}" | openssl dgst -sha256 -hmac "${SECRET}" -binary | openssl base64 | tr -d '=' | tr '/+' '_-' | tr -d '\n')

# 5. 拼接最终的JWT Token
JWT="${SIGNING_INPUT}.${SIGNATURE}"

# 输出JWT Token
echo "Generated JWT Token:"
echo "${JWT}"
