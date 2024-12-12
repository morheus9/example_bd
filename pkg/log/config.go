package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	FieldKeyModule         = "module"
	FieldKeyRequestPath    = "path"
	FieldKeyRequestMethod  = "method"
	FieldKeyRequestBody    = "request"
	FieldKeyResponseBody   = "response"
	FieldKeyResponseStatus = "status"
	FieldKeySerName        = "srvname"
	FieldKeyDuration       = "duration"
	FieldKeyStatus         = "status"
	FieldKeyClientIP       = "clientIP"
	FieldKeyRequestID      = "requestId"
	FieldKeyErrorStack     = "error_stack"
)

type Format string

const (
	TextFormat Format = "text"
	JsonFormat Format = "json"
)

const (
	outputStderr = "stderr"
	outputStdout = "stdout"
	outputNull   = "null"
)

const (
	defaultFileMaxSize    = 100
	defaultFileMaxAge     = 28
	defaultFileMaxBackups = 3
)

type Env string

const (
	EnvProd Env = "prod"
	EnvDemo Env = "demo"
	EnvDev  Env = "dev"
)

type FieldValue struct {
	Version     string `env:"VERSION"`
	CommitHash  string `env:"COMMIT_HASH"`
	CommitTime  string `env:"COMMIT_TIME"`
	Environment string `env:"ENVIRONMENT"`
}

type Config struct {
	Level  string     `env:"LEVEL" envDefault:"debug"`
	Format Format     `env:"FORMAT" envDefault:"text"`
	Output string     `env:"OUTPUT" envDefault:"stdout"`
	Fields FieldValue `env:"FIELDS"`
}

func init() { //nolint:gochecknoinits
	cfg := &Config{
		Level:  "info",
		Format: TextFormat,
		Output: outputStdout,
	}

	cfg.Fields.Version = "0.0.0-dev"
	cfg.Fields.CommitHash = "last"
	cfg.Fields.CommitTime = time.Now().Format(time.RFC3339)

	_ = Init(cfg)
}

func NewLogger(cfg *Config) (*zap.Logger, error) {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, errors.Wrapf(err, "parse log level")
	}

	atomicLevel = level

	var writer zapcore.WriteSyncer

	switch cfg.Output {
	case outputStderr:
		writer = zapcore.Lock(os.Stdout)
	case outputStdout:
		writer = zapcore.Lock(os.Stderr)
	case outputNull:
		writer = zapcore.AddSync(&nullOutput{})
	default:
		writer = zapcore.AddSync(
			&lumberjack.Logger{
				Filename:   cfg.Output,
				MaxSize:    defaultFileMaxSize, // megabytes
				MaxBackups: defaultFileMaxBackups,
				MaxAge:     defaultFileMaxAge, // days
			})
	}

	var encoder zapcore.Encoder

	switch cfg.Format {
	case TextFormat:
		encoder = newTextEncoder()
	case JsonFormat:
		encoder = newJSONEncoder()
	default:
		return nil, errors.Errorf("unknown format %q", cfg.Format)
	}

	core := zapcore.NewCore(encoder, writer, level)
	logger := zap.New(core)

	logger = logger.With(
		zap.String("environment", cfg.Fields.Environment),
		zap.String("version", cfg.Fields.Version),
		zap.String("commit-hash", cfg.Fields.CommitHash),
		zap.String("commit-time", cfg.Fields.CommitTime),
	)

	return logger, nil
}

func Init(cfg *Config) error {
	logger, err := NewLogger(cfg)
	if err != nil {
		return err
	}

	SetDefaultLogger(logger)

	return nil
}

func newJSONEncoder() zapcore.Encoder {
	encConf := zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "severity",
		TimeKey:             "@timestamp",
		NameKey:             "name",
		CallerKey:           "caller",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          RFC3339TimeUTCEncoder,
		EncodeDuration:      zapcore.MillisDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}

	jsonEncoder := zapcore.NewJSONEncoder(encConf)

	return jsonEncoder
}
func newTextEncoder() zapcore.Encoder {
	encConf := zapcore.EncoderConfig{
		MessageKey:          "M",
		LevelKey:            "L",
		TimeKey:             "T",
		NameKey:             "N",
		CallerKey:           "C",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "S",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalColorLevelEncoder,
		EncodeTime:          RFC3339TimeUTCEncoder,
		EncodeDuration:      zapcore.StringDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          zapcore.FullNameEncoder,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "\t",
	}

	jsonEncoder := zapcore.NewConsoleEncoder(encConf)

	return jsonEncoder
}

type nullOutput struct {
}

func (n2 *nullOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (n2 *nullOutput) Sync() error {
	return nil
}

const timeLayout = "2006-01-02T15:04:05.999-07:00"

type appendTimeEncoder interface {
	AppendTimeLayout(time.Time, string) //nolint:inamedparam
}

func RFC3339TimeUTCEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	t = t.UTC()

	if enc, ok := enc.(appendTimeEncoder); ok {
		enc.AppendTimeLayout(t, timeLayout)

		return
	}
	// time.RFC3339
	enc.AppendString(t.Format(timeLayout))
}

func SetDefaultLogger(logger *zap.Logger) {
	var unsafePoint = (*unsafe.Pointer)(unsafe.Pointer(&defaultLogger))

	atomic.StorePointer(unsafePoint, unsafe.Pointer(logger))
}
