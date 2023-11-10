# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. install go:

    ```console
    sudo apt-get update  
    sudo apt-get -y upgrade  
    ```

    ```console
    wget  https://go.dev/dl/go1.21.4.linux-arm64.tar.gz
    sudo tar -xvf go1.21.4.linux-arm64.tar.gz
    sudo mv go /usr/local      
    ```

    ```console
    export GOROOT=/usr/local/go 
    export GOPATH=/usr/local/go 
    export PATH=$GOROOT/bin:$PATH 
     
    ```
    

 2. Run the server:

    ```console
    git clone https://github.com/ivanbevivino/grpc-client-server.git
    cd grpc-client-server
    go mod tidy
    
    ```

 3. Run the server:

    ```console
     go run greeter_server/main.go
    ```

3. Run the client:

   ```console
    go run greeter_client/main.go
   ```