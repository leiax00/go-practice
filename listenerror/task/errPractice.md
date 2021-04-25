# Question 1
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。

# Answer
**不应该Wrap这个error并不应该抛给上层;**

`sql.ErrNoRows`: 是用于标识一个sql的scanner扫描结束, 类似于读取文件时的 `EOF` 错误;

在使用过程中, 应该作为一个正常业务标识; 应由Data provider自行处理该错误:

1. 如果sql查询出来有数据,那么在发生 `sql.ErrNoRows` 时, 应返回具体的数据(data, nil)

2. 如果sql查询出来没有数据, 那么在发生 `sqk.ErrNoRows` 时, 应该可以返回一个(nil, nil)

# code
```shell
`listenerror/task/dboperation.go`            : db操作
`listenerror/task/busnessservice.go`         : 业务处理
`listenerror/task/viewresource.go`           : 对外接口
```
