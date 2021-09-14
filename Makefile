docker_img_name = hiroki/rand-img-srvr
docker_img_tag = 0.0.1
docker_image_name = $(docker_img_name):$(docker_img_tag)

docker_container_name = rand-img-srvr

docker_inner_port = 9811
docker_outer_port = 9811

img:
	@echo $(docker_image_name)

build:
	@docker build -t $(docker_image_name) .

rmi:
	@docker rmi $(docker_image_name)

rmc:
	@docker container rm $(docker_container_name)

up:
	@docker run -d -p $(docker_outer_port):$(docker_inner_port) --name $(docker_container_name) $(docker_image_name)

down:
	@docker stop $(docker_container_name)

autou:
	@make build
	@make up

autod:
	@make down
	@make rmc
	@make rmi