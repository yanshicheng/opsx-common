Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc

Mysql:
  DataSource: root:123456@tcp(172.16.1.61:3307)/zero_demo?charset=utf8mb4&parseTime=True&loc=Local
Cache:
  - Host: 172.16.1.61:6379
    Type: node # node or cluster
    Pass: "123456"
    Tls: false # bool
    NonBlock: true # bool
    PingTimeout: 3s # default 1s