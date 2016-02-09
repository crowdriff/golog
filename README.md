# golog

[![Build Status](https://travis-ci.org/crowdriff/golog.svg?branch=master)](https://travis-ci.org/crowdriff/golog)

golog is a logging package to write warnings, errors, requests, or any message to standard out with a set format.

## Usage

### NewLogger

`func Init(app, version string)`

Initialize the global logger with the provided app name and version.

### Log

`func Log(msg string)`

Log any message to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=info app=golog msg="test message" v=v0.0.1
```

### Logf

`func Logf(msg string)`

Log any message to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=info app=golog msg="formatted message" v=v0.0.1
```

### LogError

`func LogError(err error)`

Log an error to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=error app=golog msg="error message" file="github.com/crowdriff/golog/golog.go" line=24 v=v0.0.1
```

### LogFatal

`func LogFatal(err error)`

Log an error to standard out and exit.

Format:
```
time="2016-01-18T13:49:17-05:00" level=fatal app=golog msg="error message" file="github.com/crowdriff/golog/golog.go" line=24 v=v0.0.1
```

### LogPanic

`func LogPanic(err error)`

Log an error to standard out and panic.

Format:
```
time="2016-01-18T13:49:17-05:00" level=panic app=golog msg="error message" file="github.com/crowdriff/golog/golog.go" line=24 v=v0.0.1
```

### LogWarning

`func LogWarning(msg string)`

Log a warning message to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=warn app=golog msg="warning message" v=v0.0.1
```

### LoggingMiddleware

LoggingMiddleware is a middleware function (`func(http.Handler) http.Handler`) that logs each request to standard out.

Format:
```
time="2016-01-18T13:49:17-05:00" level=info app=api code=200 dur=115745 ip="127.0.0.1:65089" method=GET size=277 uri="/test?token=token1" v=v0.0.1
```
