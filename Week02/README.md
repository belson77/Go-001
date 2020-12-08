# 总结

## 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

1. 不应该使用 Wrap。`sql.ErrNoRows` 代表业务逻辑查询不到数据。非致命性错误，一般只有基础库、第三方库、标准库调用失败才用 Wrap。
2. dao 层出现 `sql.ErrorNoRow` 时，查询结果返回 `nil`，error 返回 `nil`。上层判断返回结果再进行下一步逻辑处理