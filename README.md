# Push Notifications

SSH Term APP Apple Push Notifications

# Demo

## curl
```shell
curl -X POST "https://push.ssh2.app/apn_push" \
  -H "Content-Type: application/json" \
  -d '{
        "token": "token",
        "topic": "cn.sshterm.pro",
        "notification": {
          "aps": {
            "alert": {
              "title": "High server load",
              "subtitle": "OA server issues",
              "body": "cpu 500% mem 99% disk 99.9%"
            }
          }
        },
        "priority": 10
      }'
```

## wget
```shell
wget --header="Content-Type: application/json" \
     --post-data='{
        "token": "token",
        "topic": "cn.sshterm.pro",
        "notification": {
          "aps": {
            "alert": {
              "title": "High server load",
              "subtitle": "OA server issues",
              "body": "cpu 500% mem 99% disk 99.9%"
            }
          }
        },
        "priority": 10
      }' \
     -O - \
     "https://push.ssh2.app/apn_push"
```

## c
```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <curl/curl.h>

int main(void) {
    CURL *curl;
    CURLcode res;

    struct curl_slist *headers = NULL;
    headers = curl_slist_append(headers, "Content-Type: application/json");

    const char *data = "{\"token\":\"token\",\"topic\":\"cn.sshterm.pro\",\"notification\":{\"aps\":{\"alert\":{\"title\":\"High server load\",\"subtitle\":\"OA server issues\",\"body\":\"cpu 500% mem 99% disk 99.9%\"}}},\"priority\":10}";

    curl_global_init(CURL_GLOBAL_DEFAULT);
    curl = curl_easy_init();
    if(curl) {
        curl_easy_setopt(curl, CURLOPT_URL, "https://push.ssh2.app/apn_push");
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, data);

        res = curl_easy_perform(curl);
        if(res != CURLE_OK) {
            fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
        }

        curl_easy_cleanup(curl);
    }
    curl_slist_free_all(headers);
    curl_global_cleanup();

    return 0;
}
```
## C++
```cpp
#include <iostream>
#include <curl/curl.h>

int main() {
    CURL *curl;
    CURLcode res;

    curl_global_init(CURL_GLOBAL_DEFAULT);
    curl = curl_easy_init();
    if(curl) {
        const char *data = R"({"token":"token","topic":"cn.sshterm.pro","notification":{"aps":{"alert":{"title":"High server load","subtitle":"OA server issues","body":"cpu 500% mem 99% disk 99.9%"}}},"priority":10})";
        struct curl_slist *headers = NULL;
        headers = curl_slist_append(headers, "Content-Type: application/json");

        curl_easy_setopt(curl, CURLOPT_URL, "https://push.ssh2.app/apn_push");
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, data);

        res = curl_easy_perform(curl);
        if(res != CURLE_OK) {
            std::cerr << "curl_easy_perform() failed: " << curl_easy_strerror(res) << std::endl;
        }

        curl_slist_free_all(headers);
        curl_easy_cleanup(curl);
    }
    curl_global_cleanup();
    return 0;
}
```

## swift
```swift
import Foundation

let url = URL(string: "https://push.ssh2.app/apn_push")!
var request = URLRequest(url: url)
request.httpMethod = "POST"
request.addValue("application/json", forHTTPHeaderField: "Content-Type")

let data: [String: Any] = [
    "token": "token",
    "topic": "cn.sshterm.pro",
    "notification": [
        "aps": [
            "alert": [
                "title": "High server load",
                "subtitle": "OA server issues",
                "body": "cpu 500% mem 99% disk 99.9%"
            ]
        ]
    ],
    "priority": 10
]

do {
    request.httpBody = try JSONSerialization.data(withJSONObject: data, options: [])
} catch {
    print(error)
}

let task = URLSession.shared.dataTask(with: request) { data, response, error in
    if let error = error {
        print(error)
        return
    }
    if let data = data {
        print(String(data: data, encoding: .utf8)!)
    }
}
task.resume()
```

## python
```python
import requests

data = {
    "token": "token",
    "topic": "cn.sshterm.pro",
    "notification": {
        "aps": {
            "alert": {
                "title": "High server load",
                "subtitle": "OA server issues",
                "body": "cpu 500% mem 99% disk 99.9%"
            }
        }
    },
    "priority": 10
}

response = requests.post("https://push.ssh2.app/apn_push", json=data)
print(response.text)
```
## Node JS
```javascript
const https = require('https');

const data = JSON.stringify({
    token: "token",
    topic: "cn.sshterm.pro",
    notification: {
        aps: {
            alert: {
                title: "High server load",
                subtitle: "OA server issues",
                body: "cpu 500% mem 99% disk 99.9%"
            }
        }
    },
    priority: 10
});

const options = {
    hostname: 'push.ssh2.app',
    port: 443,
    path: '/apn_push',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': data.length
    }
};

const req = https.request(options, res => {
    console.log(`statusCode: ${res.statusCode}`);

    res.on('data', d => {
        process.stdout.write(d);
    });
});

req.on('error', error => {
    console.error(error);
});

req.write(data);
req.end();
```


## JS
```javascript
const url = "https://push.ssh2.app/apn_push";

const data = {
    token: "token",
    topic: "cn.sshterm.pro",
    notification: {
        aps: {
            alert: {
                title: "High server load",
                subtitle: "OA server issues",
                body: "cpu 500% mem 99% disk 99.9%"
            }
        }
    },
    priority: 10
};

fetch(url, {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
})
.then(response => response.text())
.then(result => console.log(result))
.catch(error => console.error('error:', error));
```

## php
```php
$data = [
    "token" => "token",
    "topic" => "cn.sshterm.pro",
    "notification" => [
        "aps" => [
            "alert" => [
                "title" => "High server load",
                "subtitle" => "OA server issues",
                "body" => "cpu 500% mem 99% disk 99.9%"
            ]
        ]
    ],
    "priority" => 10
];
$jsonData = json_encode($data);
$ch = curl_init('https://push.ssh2.app/apn_push');
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $jsonData);
curl_setopt($ch, CURLOPT_HTTPHEADER, ['Content-Type: application/json']);
$response = curl_exec($ch);
if (curl_errno($ch)) {
    echo 'cURL error: '. curl_error($ch);
} else {
    echo $response;
}
curl_close($ch);
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
