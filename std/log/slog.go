package main

import (
	"context"
	"go-study/std/log/mylog"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

func main() {
	//https://betterstack.com/community/guides/logging/logging-in-go/
	// kv()
	// kv_mixed()
	// kvStrongTypes()
	// kvGroup()
	// childLogger()
	// childWithGroup()
	// slogHanlderOptions()
	// customLevel()
	// customLevelWithName()
	// customSource()
	// textOrJson()
	logContextKv()

	// demo1()
	//  changeLogLevel()
	// testInfof()
	// setDefault()
	// convert()
}



func logContextKv(){
	//don't work,see context_handler.go
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx := context.WithValue(context.Background(), "request_id", "req-123")
	logger.InfoContext(ctx, "image uploaded", slog.String("image_id", "img-998"))
}

func textOrJson(){
	//go run slog.go 
	//APP_ENV=production go run slog.go
	var appEnv = os.Getenv("APP_ENV")
    opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

    var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
    if appEnv == "production" {
        handler = slog.NewJSONHandler(os.Stdout, opts)
    }

    logger := slog.New(handler)

    logger.Info("Info message")
}

func customSource(){
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	ctx := context.Background()
	logger.Log(ctx, slog.LevelError, "Error message")
	logger.Log(ctx, slog.LevelInfo, "info Level")
}

func customLevelWithName(){
	const(
		LevelTrace = slog.Level(-8)
		LevelFatal = slog.Level(12)
	)
	var LevelNames = map[slog.Leveler]string{
		LevelTrace: "TRACE",
		LevelFatal: "FATAL",
	}
	opts := &slog.HandlerOptions{
		Level: LevelTrace,
		//The ReplaceAttr() function is used to customize how each key/value pair in a Record is handled by a Handler. 
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr{
			if a.Key == slog.LevelKey{
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if exists{
					a.Value = slog.StringValue(levelLabel)
				}
			}
			return a
		},
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	ctx := context.Background()
	logger.Log(ctx, slog.LevelError, "Error message")
	logger.Log(ctx, LevelTrace, "Trace message")
	logger.Log(ctx, LevelFatal, "Fatal Level")
}

func customLevel(){
	const(
		LevelTrace = slog.Level(-8)
		LevelFatal = slog.Level(12)
	)
	opts := &slog.HandlerOptions{
		Level: LevelTrace,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	ctx := context.Background()
	logger.Log(ctx, slog.LevelError, "Error message")
	logger.Log(ctx, LevelTrace, "Trace message")
	logger.Log(ctx, LevelFatal, "Fatal Level")
}

func slogHanlderOptions(){
	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warn message")
	logger.Error("Error message")
	logLevel.Set(slog.LevelWarn)
	logger.Debug("Debug message again")
	logger.Info("Info message again")
	logger.Warn("Warn message again")
	logger.Error("Error message again")

}
func childWithGroup(){
	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(handler).WithGroup("program_info")

	child := logger.With(
		slog.Int("pid", os.Getpid()),
		slog.String("go_version", buildInfo.GoVersion),
	)
	child.Warn("storage is 90% full", slog.String("avaliable_space", "900.1mb"))
}

func childLogger(){
	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(handler)
	logger.Info("partent: hello")
	child := logger.With(
		slog.Group("program_info"),
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
	)
	child.Info("child: hello")
}


func kvGroup(){
	logger := slog.Default();
	logger.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"image uploaded",
		slog.Int("id", 23123),
		slog.Group("properties", 
			slog.Int("width", 4000),
			slog.Int("height", 3000),
			slog.String("format", "jpeg"),
		),
	)
}
func kv_mixed(){
	logger := slog.Default();
	logger.Info(
		"incoming request",
		"method", "GET",
		slog.Int("time_taken_ms", 158),
		slog.String("path", "/hello/world?q=search"),
		"status", 200,
		slog.String(
		  "user_agent",
		  "Googlebot/2.1 (+http://www.google.com/bot.html)",
		),
	  )
}


func kvStrongTypes(){
	logger := slog.Default();
	logger.LogAttrs(context.Background(),
	slog.LevelInfo,
	"incoming request", slog.String("method", "GET"),
	slog.Int("time_taken_ms", 150),
	slog.String("path", "/hello/world?q=search"),
	slog.Int("status", 200),	
	slog.String("user_agenet", 
	   "Goolebot/2.1 (+http://wwww.google.com/bot.html)"),
   )	
}

func kv(){
	logger := slog.Default();
	logger.Info("incoming request",
	 slog.String("method", "GET"),
	 slog.Int("time_taken_ms", 150),
	 slog.String("path", "/hello/world?q=search"),
	 slog.Int("status", 200),	
	 slog.String("user_agenet", 
		"Goolebot/2.1 (+http://wwww.google.com/bot.html)"),
	)
}

func convert(){
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.NewLogLogger(handler, slog.LevelError)
	_ = http.Server{
		ErrorLog: logger,
	}
}

func setDefault(){
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//After this call, output from the log package's default Logger  will be logged using l's Handler, at a level controlled b
	slog.SetDefault(logger)
	log.Println("Hello form old logger")
}

func demo1(){
	//All Logger instances default to logging at the INFO level, which leads to the suppression of the DEBUG entry,
	// but you can easily update this as you see fit.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Debug msg")
	logger.Info("Info msg")
	logger.Warn("Warning msg")
	logger.Error("Error msg")
}

func testInfof(){
	mylog.Infof(slog.Default(), "hello %s", "world")
}


func basic(){
// logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	// logger.Info("hello", "count", 3)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)
}


func changeLogLevel(){
	var programLevel  = new (slog.LevelVar)
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	logger := slog.Default()
	logger.Info("hello world")
	
	programLevel.Set(slog.LevelWarn)
	logger  = slog.Default()
	logger.Info("hello world")
	logger.Warn("warning")
}