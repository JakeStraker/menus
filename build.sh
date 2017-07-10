cp Templates/index.html Dockerfile/index.html
go build -o Dockerfile/bettermenus *.go
cd Dockerfile
docker build -t artifacts.ath.bskyb.com:5001/maa62/bettermenus . && docker push artifacts.ath.bskyb.com:5001/maa62/bettermenus
cd ../
rm Dockerfile/index.html Dockerfile/bettermenus
