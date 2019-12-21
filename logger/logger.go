package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error Logger
var errorLogger *zap.SugaredLogger

var levelMap  = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info": zapcore.InfoLevel,
	"warn": zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level  {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

// zap中加入Lumberjack支持，修改WriteSyncer代码
// Lumberjack Logger属性说明
//	* Filename: 日志文件的位置
//	* MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
//	* MaxBackups：保留旧文件的最大个数
//  * MaxAges：保留旧文件的最大天数
//  * Compress：是否压缩/归档旧文件
func setSyncWriter() zapcore.WriteSyncer {
	fileName := viper.GetString("log.log_file") // 日志文件
	// syncWriter : WriteSyncer 指定日志输出信息 包括文件目录等
	// zapcore.AddSync() 日志写入文件的
	return zapcore.AddSync(&lumberjack.Logger{
		Filename: fileName,
		MaxSize: 1 << 30,
		LocalTime: true,
		Compress: true,
	})
}

// EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
// EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
// EncodeDuration: zapcore.SecondsDurationEncoder,
// EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
// 编码器Encoder设置
func setEncoder() zapcore.EncoderConfig {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	return encoder
}


func InitLogger()  {

	// 设置编码器
	encoder := setEncoder()
	// 设置日志文件
	syncWriter := setSyncWriter()
	// 设置日志等级
	level := getLoggerLevel(viper.GetString("log.logger_lever"))
	atomLevel := zap.NewAtomicLevelAt(level)
	// 创建日志核心  zapcore.Core需要三个配置：Encoder，WriteSyncer，LogLevel
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, atomLevel)
	// zap.New(…) 手动传递所有配置 创建Logger
	// zap.NewProduction() 使用预置方法来创建logger
	// 创建主Logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// 获取SugaredLogger <主Logger>.Sugar()
	errorLogger = logger.Sugar()

	defer errorLogger.Sync()
	Infof("simple zap logger example")
}


func Debug(args ...interface{}) {
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	errorLogger.Infof(template, args...)
}


func Warn(args ...interface{}) {
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	errorLogger.DPanicf(template, args...)
}


func Panic(args ...interface{}) {
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	errorLogger.Fatalf(template, args...)
}
