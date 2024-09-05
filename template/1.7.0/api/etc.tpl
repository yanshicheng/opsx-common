Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}

Log:
  ServiceName: {{.serviceName}}  # ServiceName 表示服务名称，可选项。
  Mode: "console"                # Mode 表示日志记录模式，默认为 "console"。# "console"：日志输出到控制台。# "file"：日志输出到文件。# "volume"：用于K8s，在日志文件名中添加主机名前缀。
  Encoding: "plain"               # Encoding 表示编码类型，默认为 "json"。# "json"：使用 JSON 编码。# "plain"：使用纯文本编码，通常用于开发环境。
  TimeFormat: "2006-01-02 15:04:05"  # TimeFormat 表示时间格式，可选项。 # 默认为 "2006-01-02T15:04:05.000Z07:00"。
  Path: "logs"                   # Path 表示日志文件路径，默认为 "logs"。
  Level: "info"                  # Level 表示日志级别，默认为 "info"。# 可选项有 "debug"、"info"、"error"、"severe"。
  MaxContentLength: 1024         # MaxContentLength 表示最大内容字节数，默认为无限制。
  Compress: false                # Compress 表示是否压缩日志文件，默认为 false。
  Stat: true                     # Stat 表示是否记录统计信息，默认为 true。
  KeepDays: 7                    # KeepDays 表示日志文件保留的天数。默认为保留所有文件。# 仅在 Mode 为 "file" 或 "volume" 时生效，并且 Rotation 为 "daily" 或 "size" 时有效。
  StackCooldownMillis: 100       # StackCooldownMillis 表示堆栈日志的冷却时间，默认为 100ms。
  MaxBackups: 0                  # MaxBackups 表示保留多少个备份日志文件。0 表示保留所有文件。# 仅在 RotationRuleType 为 "size" 时生效。# 即使 MaxBackups 设置为 0，如果达到 KeepDays 限制，日志文件仍然会被删除。
  MaxSize: 0                     # MaxSize 表示日志文件的最大大小（单位为MB），0 表示不限制。# 仅在 RotationRuleType 为 "size" 时生效。
  Rotation: "daily"              # Rotation 表示日志轮换规则类型，默认为 "daily"。# "daily"：按天轮换。  # "size"：按大小限制轮换。

Redis:
  Host: 172.16.1.61:6379
  Type: node # node or cluster
  Pass: "123456789"
  Tls: false # bool
  NonBlock: true # bool
  PingTimeout: 3s # default 1s
