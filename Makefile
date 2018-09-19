#
# @Author: sejal chougule 
# @Date: 2018-05-02 17:21:41 
# @Last Modified by:   sejal chougule 
# @Last Modified time: 2018-05-02 17:21:41 
# 

run: clean
	go run main.go

dev : clean
	go run main.go
	
clean:
	@echo "\n -> Cleaning cache and log files\n" 
	-find . -name 'nohup.out' -delete
	@echo "\n -> Cleaning done\n"

build: clean
	@echo "\n ->Fetching Dependencies \n"
	@echo "\n"	
	godep save	
	@echo "\n ->Building fresh docker image\n"
	docker build -t gpix .
	@echo "\n ->Removing intermediate containers\n"
	echo "Y" | docker image prune --filter label=stage=intermediate
	@echo "\n \n"
	@echo "\n -------- Build finished -------\n"
	
deploy: clean delete
	pm2 start index.js --name api

run-background:clean 


test:clean run-background
	@echo "Testing server api"

help:
	@echo "\nPlease call with one of these targets:\n"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F:\
        '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}'\
        | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs | tr ' ' '\n' | awk\
        '{print "    - "$$0}'
	@echo "\n"	
