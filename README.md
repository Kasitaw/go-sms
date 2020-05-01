# Introduction
`go-sms` provide simple and clean API interface using Isms and sms123 provider. 

## Configurations
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

