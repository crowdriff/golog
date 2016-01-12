# golog

golog is a logging package to write errors, panics, requests, or any message to standard out with a set format.

## API

### `func (l *Logger) Log(s string) error`

Log any string to standard out.

Use:
```
l := NewLogger("golog")
l.Log("this is the message")
```

Output:
```
2016/01/12 10:21:38 [golog] this is the message
```

### `func (l *Logger) LogError(err error) error`

Log an error to standard out.

Use:
```
l := NewLogger("golog")
l.LogError(errors.New("this is the error message"))
```

Output:
```
2016/01/12 10:21:38 [golog] error: this is the error message
```

### `func (l *Logger) LogPanic() `

Recover from a panic and log the panic message and stack trace to standard out.

Use:
```
l := NewLogger("golog")
l.LogPanic()
panic("this is the panic message")
```

Output:
```
2016/01/12 10:21:38 [golog] panic: this is the panic message
<stack trace>
```
