package logclient

import (
	"context"
	"log"
	"octcarp/sustech/cs328/a2/api/config"
	logpb "octcarp/sustech/cs328/a2/gogrpc/glog/pb"
	"sync"
	"time"
)

type LogSender struct {
	client      logpb.LoggingServiceClient
	stream      logpb.LoggingService_StreamLogsClient
	serviceName string
	mu          sync.Mutex
}

var (
	globalSender *LogSender
	senderOnce   sync.Once
)

func initGlobalSender() *LogSender {
	senderOnce.Do(func() {
		globalSender = &LogSender{
			client:      getLogClient(),
			serviceName: config.GetConfig().ServiceName,
		}
		if err := globalSender.initStream(); err != nil {
			log.Printf("Failed to init stream: %v", err)
		}
	})
	return globalSender
}

func getGlobalSender() *LogSender {
	if globalSender == nil {
		initGlobalSender()
	}
	return globalSender
}

func (s *LogSender) initStream() error {
	stream, err := s.client.StreamLogs(context.Background())
	if err != nil {
		return err
	}
	s.stream = stream
	return nil
}

func (s *LogSender) sendLog(level logpb.LogLevel, message string, traceID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.stream == nil {
		if err := s.initStream(); err != nil {
			return err
		}
	}

	logMsg := &logpb.LogMessage{
		Level:       level,
		ServiceName: s.serviceName,
		Message:     message,
		Timestamp:   time.Now().Unix(),
		TraceId:     traceID,
	}

	err := s.stream.Send(logMsg)
	if err != nil {
		if initErr := s.initStream(); initErr != nil {
			return initErr
		}
		return s.stream.Send(logMsg)
	}

	return nil
}

func Info(message string, traceID string) {
	go func() {
		sender := getGlobalSender()
		if err := sender.sendLog(logpb.LogLevel_INFO, message, traceID); err != nil {
			log.Printf("Failed to send log async: %v", err)
		}
	}()
}

func Error(message string, traceID string) {
	go func() {
		sender := getGlobalSender()
		if err := sender.sendLog(logpb.LogLevel_INFO, message, traceID); err != nil {
			log.Printf("Failed to send log async: %v", err)
		}
	}()
}

func Warning(message string, traceID string) {
	go func() {
		sender := getGlobalSender()
		if err := sender.sendLog(logpb.LogLevel_INFO, message, traceID); err != nil {
			log.Printf("Failed to send log async: %v", err)
		}
	}()
}

func Debug(message string, traceID string) {
	go func() {
		sender := getGlobalSender()
		if err := sender.sendLog(logpb.LogLevel_INFO, message, traceID); err != nil {
			log.Printf("Failed to send log async: %v", err)
		}
	}()
}

func CloseStream() error {
	if globalSender == nil {
		return nil
	}

	globalSender.mu.Lock()
	defer globalSender.mu.Unlock()

	if globalSender.stream != nil {
		resp, err := globalSender.stream.CloseAndRecv()
		if err != nil {
			return err
		}
		_ = resp
	}
	return nil
}
