package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/kataras/iris/v12"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	InitLogger(LogConfig{
		Level: "debug",
		Path:  "temp/log",
		Save:  7,
	})
}

var Log *zap.SugaredLogger

type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
	Save  uint   `yaml:"save"`
}

func InitLogger(logConfig LogConfig) {
	encoder := getEncoder()

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	infoWriter := getLogWriter(logConfig.Path, "Info", logConfig.Save)

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	errorWriter := getLogWriter(logConfig.Path, "Error", logConfig.Save)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
}
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(SetColor(zapcore.InfoLevel, t.Format("2006-01-02 15:04:05")))
}

// 自定义日志级别显示
func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	levalStr := "[" + level.CapitalString() + "]"

	enc.AppendString(SetColor(level, levalStr))
}
func SetColor(level zapcore.Level, levalStr string) string {
	var str string
	switch level {
	case zapcore.DebugLevel:
		str = fmt.Sprintf("\x1b[33m%s\x1b[0m", levalStr)
	case zapcore.InfoLevel:
		str = fmt.Sprintf("\x1b[32m%s\x1b[0m", levalStr)
	case zapcore.WarnLevel:
		str = fmt.Sprintf("\x1b[33m%s\x1b[0m", levalStr)
	case zapcore.ErrorLevel:
		str = fmt.Sprintf("\x1b[31m%s\x1b[0m", levalStr)
	case zapcore.DPanicLevel:
		str = fmt.Sprintf("\x1b[35m%s\x1b[0m", levalStr)
	case zapcore.PanicLevel:
		str = fmt.Sprintf("\x1b[35m%s\x1b[0m", levalStr)
	case zapcore.FatalLevel:
		str = fmt.Sprintf("\x1b[35m%s\x1b[0m", levalStr)
	default:
		str = fmt.Sprintf("\x1b[36m%s\x1b[0m", levalStr)
	}
	return str
}

// 自定义文件：行号输出项
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString("[" + l.traceId + "]")
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.ConsoleSeparator = " "
	encoderConfig.CallerKey = "caller_line"
	encoderConfig.LevelKey = "level_name"
	encoderConfig.MessageKey = "msg"
	encoderConfig.TimeKey = "ts"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.EncodeTime = customTimeEncoder
	encoderConfig.EncodeLevel = customLevelEncoder
	encoderConfig.EncodeCaller = customCallerEncoder
	// encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath, level string, save uint) io.Writer {
	logFullPath := path.Join(logPath, level)
	hook, err := rotatelogs.New(
		logFullPath+".%Y%m%d%H",                   // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(save),        // 文件最大保存份数
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔

	)
	if err != nil {
		panic(err)
	}
	return hook
}

func LoggerHandler(ctx iris.Context) {
	p := ctx.Request().URL.Path
	method := ctx.Request().Method
	start := time.Now()
	fields := make(map[string]interface{})
	fields["title"] = "访问日志"
	fields["fun_name"] = path.Join(method, p)
	fields["ip"] = ctx.Request().RemoteAddr
	fields["method"] = method
	fields["url"] = ctx.Request().URL.String()
	fields["proto"] = ctx.Request().Proto
	// fields["header"] = ctx.Request().Header
	fields["user_agent"] = ctx.Request().UserAgent()
	fields["x_request_id"] = ctx.GetHeader("X-Request-Id")

	// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)
			fields["content_length"] = ctx.GetContentLength()
			fields["body"] = string(body)
		}
	}
	ctx.Next()

	// 下面是返回日志
	fields["res_status"] = ctx.ResponseWriter().StatusCode()
	if ctx.Values().GetString("out_err") != "" {
		fields["out_err"] = ctx.Values().GetString("out_err")
	}
	fields["res_length"] = ctx.ResponseWriter().Header().Get("size")
	if v := ctx.Values().Get("res_body"); v != nil {
		if b, ok := v.([]byte); ok {
			fields["res_body"] = string(b)
		}
	}
	token := ctx.Values().Get("jwt")
	if token != nil {
		fields["uid"] = token.(*jwt.Token).Claims
	}
	timeConsuming := time.Since(start).Nanoseconds() / 1e6
	msg := fmt.Sprintf("[http] %s-%s-%s-%d(%dms)",
		p, ctx.Request().Method, ctx.Request().RemoteAddr, ctx.ResponseWriter().StatusCode(), timeConsuming)
	fields["time_consuming"] = fmt.Sprintf("%dms", timeConsuming)

	Log.Debugf("%s\n %s", msg, prettyPrint(fields))
	// Log.Infof(msg)
}

func prettyPrint(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return ""
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return ""
	}

	return out.String()
}
