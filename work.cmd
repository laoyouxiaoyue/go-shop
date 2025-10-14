@echo off
title Start Consul, Nacos and Redis
echo Starting Consul, Nacos and Redis services...

REM Start Redis
echo Starting Redis...
start "Redis" cmd /k "D:\Redis\redis-server.exe"

REM Wait for Redis to start
timeout /t 3

REM Start Consul
echo Starting Consul...
start "Consul" cmd /k "D:\nacos\bin\consul.exe agent -dev"

REM Wait for Consul to start
timeout /t 5

REM Start Nacos
echo Starting Nacos...
start "Nacos" cmd /k "D:\nacos\bin\startup.cmd -m standalone"

echo All services started successfully!