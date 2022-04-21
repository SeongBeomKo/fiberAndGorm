module fiber_demo

go 1.18

//importing external dependency into local folder
// 로컬에 다운받으면 아래처럼 상대경로로 의존성을 바꿔줄수잇다.
//replace github.com/CoderVlogger/go-web-frameworks/pkg => ../pkg

require (
	github.com/CoderVlogger/go-web-frameworks/pkg v0.0.0-20220316213317-1dd6ca6a3cba
	github.com/gofiber/fiber/v2 v2.31.0
	github.com/google/uuid v1.3.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.34.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)
