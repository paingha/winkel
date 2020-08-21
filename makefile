.PHONY: start-api
start-api:
	if [ -a ./api/api.exe ]; then rm -rf ./api/api.exe; fi;
	@echo "Starting API Service"
ifeq ($(OS),Windows_NT)
	cd ./api && ./start.cmd
else
	cd ./api && ./start.sh
endif