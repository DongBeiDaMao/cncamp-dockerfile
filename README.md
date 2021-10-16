# cncamp-dockerfile

main.go是上个作业的http服务器代码，bin/amd64/my_httpserver是编译出来的可执行文件，
Dockerfile中使用这个可执行文件创建一个docker镜像

将服务器的可执行文件打包成docker镜像：
docker build -t wentian2021/my_httpserver:v1.0.0 .

将镜像推送到仓库：
docker push wentian2021/my_httpserver:v1.0.0
