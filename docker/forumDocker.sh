docker rm forumContainer
docker rmi forum_image
docker build -t forum_image .
docker run -it --name forumContainer -p 15040:15040 forum_image
