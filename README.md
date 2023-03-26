# go-DzMicro

计划使用 Golang 语言对现有 Python 项目 [DzMicro](https://github.com/dzming-git/DzMicro) 进行重构，以提高代码的性能、可扩展性和可维护性。

## 重构进度

| 模块名称                               | 进度   | 备注                                                         |
| -------------------------------------- | ------ | ------------------------------------------------------------ |
| api/platform_routes.go                 | 未进行 |                                                              |
| api/server_routes.go                   | 未进行 |                                                              |
| app/app.go                             | 未进行 |                                                              |
| app/message_handler/bot_commands.go    | 未进行 |                                                              |
| app/message_handler/error_handler.go   | 未进行 |                                                              |
| app/message_handler/message_handler.go | 未进行 |                                                              |
| app/services/service.go                | 未进行 |                                                              |
| conf/dzmicro.go                        | 未进行 |                                                              |
| conf/authority/authority.go            | 已完成 | 部分功能等待app/services/service.go<br />中的函数指令-函数映射表的完成 |
| conf/consul_info/consul_info.go        | 未进行 |                                                              |
| conf/route_info/route_info.go          | 未进行 |                                                              |
| utils/compare_dicts.go                 | 未进行 |                                                              |
| utils/judge_same_listener.go           | 未进行 |                                                              |
| utils/listener_manager.go              | 未进行 |                                                              |
| utils/tasks.go                         | 未进行 |                                                              |
| utils/watch_config.go                  | 未进行 |                                                              |
| utils/network/app_utils.go             | 未进行 |                                                              |
| utils/network/consul_client.go         | 未进行 |                                                              |
| utils/network/heartbeat_manager.go     | 未进行 |                                                              |
| utils/network/message_sender.go        | 未进行 |                                                              |


## 使用该引擎开发的微程序

- [go-DBot](https://github.com/dzming-git/go-DBot)  待完善

## 安装使用

### 安装

### 使用

## 授权许可

本项目使用 MIT 许可证，有关更多信息，请参阅 LICENSE 文件。

## 联系我们

如果您对本项目有任何问题或建议，请随时通过以下方式联系我们：

- Email: dzm_work@163.com
- QQ: 715558579
