
# Sweeper

### Installation
```
$ conda activate <envs>
$ pip install -r requiremets.txt
$ python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. ./proto/twitter.proto
```

### Usage
```
$ python main.py 
```

```
sudo docker run -p 50081:50081 rpc-twitter:1.0 --name twitter_rpc_01
```