# Health

Library for health checks - requires gin-gonic/gin and uber

## Usage

```
import (
    "go.uber.org/zap"
    "github.com/warrenhodg/health"
)

func main() {
    logger, _ := zap.NewProduction()
    h := health.New(logger)

    svr := gin.Default()

    h.RegisterEndpoint(svr)
}

func SomethingUnstable(h health.IHealth) {
    h.SetSystemState("something", false)
}

func SomethingStable(h health.IHealth) {
    h.SetSystemState("something", true)
}
```

## Notes

Registering the endpoint registers it as /health, which returns either `200 OK`, or `500 NOT OK`.