
![Screenshot_20241030_150732](https://github.com/user-attachments/assets/55894f40-5a3c-43bd-8a38-4acd87987829)

# Proxabay 
Proxabay is a powerful reverse proxy built entirely in Golang!
This proxy is designed to supercharge the development process, providing faster response times, seamless integration, and detailed logging—all wrapped into a simple-to-use CLI tool.



## Demo

Insert gif or link to demo


## Run Locally

Clone the project

```bash
  git clone https://github.com/coderkhushal/proxabay
```

Go to the project directory

```bash
  cd proxabay
```

Install dependencies

```bash
  go mod tidy
```
Run Cli

```bash
  go run main.go --origin <Main_Server_Url> --port <Port_on_you_local_machine> 
```

Replace:
- Main_Server_Url with Primary server which you want to proxy 
- Port_on_you_local_machine with port at which you want to access the proxy


## Usage/Examples

With main.go 
- Add Proxy

```sh
    go run main.go --origin https://www.myserver.com/api/abc --port 4000
```

- Clear Cache
```sh
    go run main.go clearcache

```

With CLI installed 

- Add Proxy 
```sh
    proxabay --origin https://www.myserver.com/api/abc --port 4000
```

- Clear Cache
```sh
    proxabay clearcache
```
