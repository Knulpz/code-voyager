package concurrency

import (
	"strconv"
	"sync"
)

//题目：
// A协程只输出0；
// B协程只输出奇数；
// C协程只输出偶数；
// 三个协程完成协程间通信：
// 例：PrintNumber(5)
// 输出：0 1 0 2 0 3 0 4 0 5

func PrintNumber(num int) string{
	number := ""

	var wg sync.WaitGroup
	wg.Add(3)

	chan1 := make(chan int, num) // 最初的0由裁判给g1
	chan1 <- 0

	chan2 := make(chan int, num) // 问题：为什么使用无缓冲通道，区别对待呢？？
	chan3 := make(chan int, num)

	// 打一场接力赛
	// g1（团队的大脑）：负责打印0，并通知g2、g3下一个要打印的数字
	go func(){
		defer wg.Done()
		for i := 1; i <= num; i++{
			select{
			case <- chan1:
				number += "0"
				// 如果是奇数
				if i % 2 == 1{
					chan2 <- i
				}else{
					chan3 <- i
				}
			}
		}
		// 接力赛结束
		//close(chan1) 细节：这里不能Close chan1，why?
		close(chan2)
		close(chan3)
	}()

	// g2（挂件1号）：眼睛盯着自己的专属通道chan2，一旦有东西，不需要思考直接打印；并通知g1要打印0
	go func(){
		defer wg.Done()
		// 细节：当通道关闭时，g2知道自己的任务完成了，否则一直不停来就dead了
		for oddNum := range chan2{
			number += strconv.Itoa(oddNum)
			chan1 <- 0
		}
	}()

	// g3（挂件2h号）：同上
	go func(){
		defer wg.Done()
		for evenNum := range chan3{
			number += strconv.Itoa(evenNum)
			chan1 <- 0
		}
	}()

	wg.Wait()
	return number
}