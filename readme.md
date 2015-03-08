# Golang external services block
### Info
Package allows store errors count from external services. 
Useful when any of external services is down / busy / answers with timeout and you need to slow down work for some period.

### Usage
```go
...

var (
    services = &goesblock.Services{}
)

func init() {
    services = goesblock.Get()
}

func main() {
    ...
    
    for _, task := range tasks {
        if services.Down() {
            time.Sleep(time.Second * 30)
        }
        
        go task.do()
    }
    
    ...
}

func (t *task) do() {
    var err error
    
    if err = callExternalServiceOne(); err != nil {
        services.ExternalServiceOne.IncError()
        return
    }
    
    ...

    if err = callExternalServiceTwo(); err != nil {
        services.ExternalServiceTwo.IncError()
        return
    }

    ...
}

...
```
![](ave-go.png?raw=true "For the glory of go of course >:)")