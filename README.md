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
    cn.sshterm.pro (SSH Term Pro iOS 7.0+)
    or
    cn.sshterm.free (SSH Term Free iOS 7.0+)

# request
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
# return

```json
{
    "StatusCode": 200,
    "Reason": "",
    "ApnsID": "B45083AA-1847-534B-528D-076BDB9C5ACF",
    "Timestamp": "0001-01-01T00:00:00Z",
    "ApnsUniqueID": "825b0498-3068-2bf3-68de-26d6355b5b3b"
}
```
## 《Message Push Service Usage Agreement》

### I. Service Description

1. This message push service is based on Apple APNS and is a decentralized service without server hosting.

2. The content of the push message only includes: title, subtitle, and body.

### II. Push Token

1. You will obtain a unique push token assigned by Apple, which must be kept confidential and not disclosed to others, otherwise it may lead to abuse of the push.

### III. Node - related

1. The push message can be sent to the official nodes push.sshterm.cn or push.ssh2.app, and you can choose at your own discretion.

2. The push node code is 100% open - source. Since the push key is not provided to you, the message needs to be forwarded to Apple APNS through the node.

3. The node may impose restrictions on requests, such as frequency limits, message length limits, etc.

### IV. Privacy Protection

1. The push message will not be stored, recorded or monitored.

If you do not agree with the contents of this agreement, you will not be able to enable this message push service. Once you start using this service, it means that you agree to all the terms of this agreement.

### Please note:

1. Although there is no emphasis on legal issues at present, in actual applications, it is still recommended to ensure that the agreement conforms to the basic legal compliance framework.

2. If there are updates or other changes in the service, it may be necessary to re - examine the terms of the agreement.