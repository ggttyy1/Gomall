@echo off

:: 启动容器
docker-compose up -d

:: 获取 ollama 容器的容器 ID
for /f "tokens=*" %%i in ('docker ps -q --filter "ancestor=ollama/ollama"') do set ollamaContainerId=%%i

:: 执行 ollama 命令
docker exec -it %ollamaContainerId% /bin/bash -c "ollama run llama3.2"
