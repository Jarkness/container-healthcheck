# container-healthcheck
Microservice that returns info about docker containers on host.
RESTful JSON API.
For docker container deployment, needs the access to the docker socket via volume.
How to run:

'''bash
docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 80:80 --name container_health container_health
'''

Example response on /health: 
'''json
[
    {
        "Id": "86a62fbc15998eb756ce4b4cde71a2a9cf2740b255db0247a5fcf66d51838ce0",
        "Names": [
            "/container_info"
        ],
        "Image": "container_info",
        "ImageID": "sha256:751c1df4d9da6cddbfc5e451813f083315160fdb8b8f1d8e2f8828b34eb813ed",
        "Command": "go run main.go",
        "Created": 1677263274,
        "Ports": [
            {
                "IP": "0.0.0.0",
                "PrivatePort": 80,
                "PublicPort": 80,
                "Type": "tcp"
            }
        ],
        "Labels": {},
        "State": "running",
        "Status": "Up 10 seconds",
        "HostConfig": {
            "NetworkMode": "default"
        },
        "NetworkSettings": {
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "09cb3257034510d14da955be219ff99fe3017c143d99a1b9fb3473841fa923a5",
                    "EndpointID": "c4990eaada020c4bb9520e3d7f638a6dc10d9e114228502022be02452415cf8e",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.3",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:03",
                    "DriverOpts": null
                }
            }
        },
        "Mounts": [
            {
                "Type": "bind",
                "Source": "/var/run/docker.sock",
                "Destination": "/var/run/docker.sock",
                "Mode": "rw",
                "RW": true,
                "Propagation": "rprivate"
            }
        ]
    }
]
'''
