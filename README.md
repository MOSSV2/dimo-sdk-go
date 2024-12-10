# intro
go sdk，include file, model and contract operations, hub usage

## usage 

### network

OP Sepolia: https://docs.optimism.io/chain/networks

+ explorer: https://sepolia-optimistic.etherscan.io
+ RPC URL: https://optimism-sepolia-rpc.publicnode.com
+ Faucet: https://docs.optimism.io/builders/tools/build/faucets 

note: 启动时会从服务器自动获取0.002 gas token和100 UB token

### upload file/directory/model

```shell
> git clone https://github.com/MOSSV2/dimo-sdk-go.git
> cd example/upload
> go build
# if sk not set, will generate a new key, model means upload model or regualr file/dir
> ./upload --model=false --sk=<your secret key> --path=<your local file/dir path>
# example, upload file
> ./upload --sk=4215875d8ac13ac4fb0876a0ecd0384aca0ce16b627bf975c8084915aad79470 --path=./upload
```

### download file/directory/model

```shell
> cd example/download
> go build
# if sk not set, will generate a new key, model means upload model or regualr file/dir
> ./download --model=false --sk=<your secret key>  --name=<file name> --path=<your local file/dir path to save>
# example, upload file
> ./download --sk=4215875d8ac13ac4fb0876a0ecd0384aca0ce16b627bf975c8084915aad79470 --name=4b59a3a5fa50d178dc4594c400097d497a206cff98865e815333ed7504558336 --path=./upload
```

### upload huggingface repo(todo)

```shell
> cd example/hub
> go build
# huggingface public repo name, such as: CompVis/stable-diffusion-v1-1, empty revision means latest
# hub node will sync data from hugginface and upload to dimo
> ./hub --name=<repo name> --revision=<repo revision>
# example, sync repo
> ./hub --name=CompVis/stable-diffusion-v1-1
```

## hub 

### public hub

#### download

+ web browser: http://52.76.75.134:8080/api/download?id=<your file name>&owner=<your file owner>

+ shell 

```shell
> wget http://52.76.75.134:8080/api/download?id=<your file name>\&owner=<your file owner> -O <saved name>
# or display 
> curl http://52.76.75.134:8080/api/download?id=<your file name>\&owner=<your file owner>

```

#### upload

+ upload using json 

```shell
# output: {"File":"0xabcd-0.vol","Start":0,"Size":41}
> curl -X POST http://52.76.75.134:8080/api/upload -d '{
    "id": "test1", 
    "owner":"0xabcd",
    "message":"Here is a story about llamas eating grass"
  }'
```


### private hub

```shell
> cd app/hub
> go build
> ./hub init
# run
> ./hub daemon run -b 0.0.0.0:8086
# download your file by get http://<ip>:8086/api/download?id=<your file name> 
```