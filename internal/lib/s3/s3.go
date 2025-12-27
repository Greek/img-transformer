package lib

import (
	"fmt"

	env "github.com/greek/img-transform/internal/lib/envloader"
)

func InitS3() {
	cfg := env.GetEnv()

	fmt.Println(cfg.S3_ACCESS_KEY)
}
