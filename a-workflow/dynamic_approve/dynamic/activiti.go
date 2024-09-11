package dynamic

import (
	"context"
	"log"
	"time"
)

// ApproveRequest 活动：模拟审核请求
func ApproveRequest(ctx context.Context, requestID int) (bool, error) {
	log.Printf("Approving request %d", requestID)
	time.Sleep(time.Second * 2) // 模拟耗时操作
	return true, nil
}

// RejectRequest 活动：模拟拒绝请求
func RejectRequest(ctx context.Context, requestID int) (bool, error) {
	log.Printf("Rejecting request %d", requestID)
	time.Sleep(time.Second * 2) // 模拟耗时操作
	return false, nil
}
