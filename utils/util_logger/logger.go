package utillogger

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

// type for log severity level
type LogLevel string
type color string

const (
	INFO     LogLevel = "INFO"     // info level log
	WARNING  LogLevel = "WARNING"  // warning level log
	ERROR    LogLevel = "ERROR"    // error level log
	REQUEST  LogLevel = "REQUEST"  // request level log
	RESPONSE LogLevel = "RESPONSE" // response level log
	GREEN    color    = "\033[32m"
	YELLOW   color    = "\033[33m"
	RED      color    = "\033[31m"
	CYAN     color    = "\033[36m"
	RESET    color    = "\033[0m"
)

type logData struct {
	Level          LogLevel    `json:"level,omitempty"`
	Timestamp      time.Time   `json:"timestamp,omitempty"`
	Caller         string      `json:"caller,omitempty"`
	Message        string      `json:"message,omitempty"`
	AdditionalInfo interface{} `json:"additionalInfo,omitempty"`
}

type logColor struct {
	Start color
	Reset color
}

var levelColor = map[LogLevel]logColor{
	INFO: {
		Start: GREEN,
		Reset: RESET,
	},
	WARNING: {
		Start: YELLOW,
		Reset: RESET,
	},
	ERROR: {
		Start: RED,
		Reset: RESET,
	},
	REQUEST: {
		Start: CYAN,
		Reset: RESET,
	},
	RESPONSE: {
		Start: CYAN,
		Reset: RESET,
	},
}

func setColor(level LogLevel) string {
	var (
		color    = levelColor[level]
		levelStr = fmt.Sprintf("[%v]", level)
	)

	return fmt.Sprint(color.Start, levelStr, color.Reset)
}

func logger(level LogLevel, message string, additionalInfo map[string]interface{}) {
	header := setColor(level)
	_, file, line, _ := runtime.Caller(2)
	funcName := fmt.Sprintf("%s:%d", file, line)

	currentTime := time.Now()
	newLog := logData{
		Level:          level,
		Timestamp:      currentTime,
		Caller:         funcName,
		Message:        message,
		AdditionalInfo: additionalInfo,
	}

	logByte, err := json.MarshalIndent(newLog, "", "	")
	if err != nil {
		log.Println(header)
		fmt.Println(err.Error())
	}

	logString := string(logByte)
	fmt.Println()
	log.Printf("%s\n%s", header, logString)
}

// Generate error logger on terminal
func Error(err error, additionalInfo map[string]interface{}) {
	level := ERROR
	logger(level, err.Error(), additionalInfo)
}

// Generate warning logger on terminal
func Warning(message string, additionalInfo map[string]interface{}) {
	level := WARNING
	logger(level, message, additionalInfo)
}

// Generate info logger on terminal
func Info(message string, additionalInfo map[string]interface{}) {
	level := INFO
	logger(level, message, additionalInfo)
}

func Request(message string, c *fiber.Ctx) {
	level := REQUEST
	additionalInfo := RequestLogger(c)
	logger(level, message, additionalInfo)
}

func Response(message string, c *fiber.Ctx) {
	level := RESPONSE
	additionalInfo := RequestLogger(c)
	logger(level, message, additionalInfo)
}

func RequestLogger(c *fiber.Ctx) map[string]interface{} {
	body := map[string]interface{}{}
	params := c.AllParams()

	fmt.Printf("%+v", params["noteId"])
	json.Unmarshal(c.Body(), &body)

	mapformat := map[string]interface{}{
		"time":       c.Context().Time().Format(time.RFC3339Nano),
		"httpMethod": c.Method(),
		"path":       c.Path(),
		"params":     params,
		"body":       body,
	}

	return mapformat
}
