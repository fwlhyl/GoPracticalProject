package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// truckTask 模拟一辆卡车上报数据
// ctx: 用于接收取消信号
// wg: 用于告诉主线程"我跑完了"
// truckID: 卡车编号
func truckTask(ctx context.Context, wg *sync.WaitGroup, truckID int) {
	defer wg.Done()                  // 函数结束时调用，通知 WaitGroup 减少一个计数
	speed := 60 + float64(truckID)*5 // 模拟不同速度

	for {
		select {
		case <-ctx.Done(): // 监听 Context 的取消信号
			fmt.Printf("[系统] Truck-%d 收到停止指令，正在熄火...\n", truckID)
			return
		default:
			// 正常业务逻辑：上报数据
			fmt.Printf("Truck-%d 正在行驶，当前速度: %.0f km/h\n", truckID, speed)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("[系统] 车辆调度系统启动...")

	// 1. 创建 Context：设置 5 秒超时
	// context.WithTimeout 会返回两个值：
	//   - ctx: 子协程用它来监听是否该停止了
	//   - cancel: 我们手动取消用的（即使没到5秒也可以提前停）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 无论如何，main结束时调用 cancel 释放资源

	// 2. 创建 WaitGroup：用来等所有卡车协程都安全退出
	var wg sync.WaitGroup

	// 3. 启动 5 辆卡车
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 每启动一辆卡车，给 WaitGroup 加 1
		go truckTask(ctx, &wg, i)
	}

	fmt.Println("[系统] 所有卡车已启动，等待 5 秒...")

	// 4. 阻塞等待：要么等 Context 超时，要么等所有协程自己 Done()
	wg.Wait()

	fmt.Println("[系统] 所有车辆已熄火，系统关闭。")
}
