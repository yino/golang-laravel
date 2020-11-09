#!/bin/sh
// -c 并发个数
// -n 请求总数

ab -c 500 -n 1200 -p 'postdata.txt' -T 'application/x-www-form-urlencoded' 'http://127.0.0.1:8000/corp/login'