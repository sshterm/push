# Push Notifications

SSH Term APP Apple Push Notifications

## Build
```
docker run --rm -v "$PWD":/app -w /app golang:latest go build -trimpath -ldflags="-s -w" -tags 'netcgo' -o lib/server server.go && upx lib/server
```


# Demo

## Node
    https://push.sshterm.cn/apn_push
    or
    https://push.ssh2.app/apn_push

## Optional：

### token:
    Get it in the app
### priority
    The priority of the notification. If you omit this header, APNs sets the notification priority to 10.Specify 10 to send the notification immediately.
### topic
    cn.sshterm.pro (SSH Term Pro)
    or
    cn.sshterm.free (SSH Term Free)

```
curl -v \
  -X POST "https://push.sshterm.cn/apn_push" \
  -H "Content-Type: application/json" \
  -d '{
        "token": "token",
        "topic": "cn.sshterm.pro",
        "notification": {
          "aps": {
            "alert": {
              "title": "test title",
              "subtitle": "test subtitle",
              "body": "test body"
            }
          }
        },
        "priority": 10
      }'
```
