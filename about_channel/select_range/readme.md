### range

range 使用在channel上时，会自动等待channel动作直到channel关闭

### select

select 是一种与 switch 相似的控制结构，与 switch 不同的是，select 中虽然也有多个 case，但是这些 case 中的表达式必须都是 Channel 的收发操作。
使用 select 控制结构时，会遇到两个有趣的现象：
1. 能在 Channel 上进行非阻塞的收发操作
2. 在遇到多个 Channel 同时响应时会随机挑选 case 执行

非阻塞的收发
在通常情况下，select 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。但是如果 select 控制结构中包含 default 语句，那么这个 select 语句在执行时会遇到以下两种情况：
1. 当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
2. 当不存在可以收发的 Channel 是，执行 default 中的语句；