# Introduction
`go-sms` provide simple and clean API interface using Isms and sms123 provider. 

### Installation
Get the package
```
go get github.com/Kasitaw/go-sms
```

CD into the package config's folder
```
cd ~/go/src/github.com/Kasitaw/go-sms/configs
```

Build main.go file
```
go build main.go
```

Follow `configurations` section. 

### Configurations
Create `credential.yaml` file inside `configs` folder and put below code:

```yaml
---
# supported drivers
drivers:
  - name: "isms"
    username: "your-isms-username"
    password: "your-isms-password"
    url: "https://www.isms.com.my/isms_send.php"

  - name: "sms123"
    username: ""
    password: "your-sms123-apiKey"
    url: "https://www.sms123.net/api/send.php"

# default provider that will be used
default: "isms"
```

### Run
`go run main.go`

Server will be run on port `:2346`.

### Create supervisord config to ensure the server keeps restarting if down (Adjust to your environment's setup)
```
# /etc/supervisor/conf.d/gosms.conf

[program:go-sms]
process_name=%(program_name)s_%(process_num)02d
directory=/home/<user-directory>/go/src/github.com/Kasitaw/go-sms
command=/home/<user-directory>/go/src/github.com/Kasitaw/go-sms/main
autostart=true
autorestart=true
user=your-user
numprocs=1
redirect_stderr=true
```

Once the configuration file has been created, you may update the Supervisor configuration and start the processes using the following commands

```
sudo supervisorctl reread

sudo supervisorctl update

sudo supervisorctl start
```

