package main

import (
	"bytes"
	"os"
	"sync"
	"time"
	"io"
)

var bufPool = sync.Pool{
	New: func()any{
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time{
	return time.Unix(113614245, 0)
}

func Log(w io.Writer, key, val string){
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main(){
	Log(os.Stdout, "path", "/search?q=flowers")
}