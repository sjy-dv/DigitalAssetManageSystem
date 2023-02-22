docker build --file bigchaindb.Dockerfile -t bigchaindb .
docker run -d -p 9984:9984 -p 27017:27017 bigchaindb
