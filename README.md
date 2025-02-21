# Push Notifications

SSH Term APP Apple Push Notifications

# Demo

## curl
```shell
TOKEN="your_unique_device_token"
TITLE="High server load"
SUBTITLE="OA server issues"
BODY="cpu 500% mem 99% disk 99.9%"

curl -X POST "https://push.ssh2.app/apn_push" \
  -H "Content-Type: application/json" \
  -d '{
        "token": "'"$TOKEN"'",
        "topic": "cn.sshterm.pro",
        "notification": {
          "aps": {
            "alert": {
              "title": "'"$TITLE"'",
              "subtitle": "'"$SUBTITLE"'",
              "body": "'"$BODY"'"
            }
          }
        },
        "priority": 10
      }'
```

# response

HTTP 200-400 status code


## Node
    https://push.sshterm.cn/apn_push (China)
    or
    https://push.ssh2.app/apn_push (Global)

## Optional:

### token:
    Get it in the app
### priority
    The priority of the notification. If you omit this header, APNs sets the notification priority to 10.Specify 10 to send the notification immediately.
### topic
    cn.sshterm.pro (SSH Term Pro (iOS macOS) 7.0.7+)
    or
    cn.sshterm.free (SSH Term Free (iOS macOS) 7.0.7+)


# Message Push Service Usage Agreement

## I. Service Overview
1. This message push service is built upon Apple's APNs (Apple Push Notification service), employing a decentralized service model without the need for server hosting.
2. The content of push messages is restricted to title, subtitle, and body.

## II. Service Purpose
1. The primary objective of this service is to facilitate server status monitoring. It is important to note that the APP does not process or retain any records of push messages. Due to iOS's notification mechanism, messages may disappear after being viewed or upon device restart.

## III. Push Token
1. You will receive a unique push token assigned by Apple. It is crucial to keep this token confidential and not share it with others to avoid the risk of push message abuse.

## IV. Node Information
1. You may choose between the official nodes: push.sshterm.cn or push.ssh2.app, based on your requirements.
2. The push node code is open-source. As the push key is not provided to you, messages need to be forwarded to Apple APNs via these nodes.
3. Nodes implement restriction measures on requests, such as frequency limits and message length limits.

## V. Privacy Protection
1. Since you are responsible for sending push messages, it is strictly forbidden to include any sensitive information.
2. Due to the decentralized nature of the service, push messages are not stored, recorded, or monitored.
3. This service does not collect any personal information from you and will not send you any advertising content.
4. If you have concerns about privacy and security, please discontinue using this service.

## VI. Important Notes
1. Although this agreement does not emphasize legal issues, it is recommended to ensure compliance with basic legal and regulatory frameworks in practical applications.
2. If there are updates or changes to the service, it may be necessary to re-evaluate the agreement terms.
3. The current service is free of charge, but there is a possibility of charging or terminating the service in the future.
4. The final interpretation right of these agreement terms belongs to the author.
