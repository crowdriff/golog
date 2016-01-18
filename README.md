# golog

[![Build Status](https://travis-ci.org/crowdriff/golog.svg?branch=master)](https://travis-ci.org/crowdriff/golog)

golog is a logging package to write warnings, errors, requests, or any message to standard out with a set format.

## Usage

### NewLogger

`func NewLogger(app, version string) *Logger`

Create a new Logger instance with the provided app name and version.

### Log

`func (l *Logger) Log(msg string)`

Log any message to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=info app=golog msg="test message" v=v0.0.1
```

### LogError

`func (l *Logger) LogError(err error)`

Log an error to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=error app=golog msg="error message" file="github.com/crowdriff/golog/golog.go" line=24 v=v0.0.1
```

### LogWarning

`func (l *Logger) LogWarning(msg string)`

Log a warning message to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=warn app=golog msg="warning message" v=v0.0.1
```

### LogRequestMiddleware

`func LogRequestMiddleware(l *Logger) func(http.Handler) http.Handler`

Return a middleware function (`func(http.Handler) http.Handler`) that logs each request to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=info app=api code=200 dur=115745 ip="127.0.0.1:65089" method=GET size=277 uri="/test?token=token1" v=v0.0.1
```
