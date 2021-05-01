package util

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	pb_mani "gin_mani_center/pb"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func BuildBaseResp(err error, errMsg string) *pb_mani.BaseResp {
	if err == nil {
		return &pb_mani.BaseResp{
			State: StateSuccess,
			Msg:   MsgSuccess,
		}
	} else {
		if errMsg != "" {
			return &pb_mani.BaseResp{
				State: StateError,
				Msg:   errMsg,
			}
		} else {
			return &pb_mani.BaseResp{
				State: StateError,
				Msg:   err.Error(),
			}
		}
	}
}

func GetLimitAndOffset(pageNo int64, pageSize int64) (int64, int64) {
	if pageNo < 1 || pageSize == 0 {
		return 0, 0
	}
	offset := (pageNo - 1) * pageSize
	limit := pageSize
	return limit, offset
}

func HasMore(realCount int64, totalCount int64, offset int64) bool {
	if realCount <= 0 || totalCount <= 0 || offset < 0 {
		return false
	}
	if totalCount < offset {
		offset = totalCount
	}
	if realCount < (totalCount - offset) {
		return true
	}
	return false
}

func randString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GenUID() string {
	h := md5.New()
	h.Write([]byte(randString(16)))
	return hex.EncodeToString(h.Sum(nil))
}


func goWithRecovery(ctx context.Context, df func(ctx context.Context), f func()) {
	go func() {
		defer df(ctx)
		f()
	}()
}

func defaultF(ctx context.Context) {
	if e := recover(); e != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		log.Fatalf("goroutine panic %s", e)
	}
}

func GoParallel(ctx context.Context, f func()) {
	goWithRecovery(ctx, defaultF, f)
}
