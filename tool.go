package tool

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func Properties2Map(path string) (map[string]string, error) {
	var (
		m   map[string]string
		err error
		f   *os.File
		b   []byte
	)
	f, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m = make(map[string]string, 0)
	rf := bufio.NewReader(f)
	for {
		b, _, err = rf.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		l := string(b)
		if strings.HasPrefix(l, "#") {
			continue
		}
		sl := strings.Split(l, "=")
		if len(sl) > 1 {
			m[sl[0]] = strings.Join(sl[1:], ",")
		}
	}
	return m, nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CopyFile(source, target string) {
	if FileExist(source) && !FileExist(target) {
		if sf, err := os.Open(source); err == nil {
			defer sf.Close()
			if tf, err := os.Create(target); err == nil {
				tf.Close()
				io.Copy(bufio.NewWriter(tf), bufio.NewReader(sf))
			}
		}
	}
}

func Bytes2File(data []byte, target string) {
	if FileExist(target) {
		return
	}
	if tf, err := os.Create(target); err == nil {
		defer tf.Close()
		wf := bufio.NewWriter(tf)
		wf.Write(data)
		wf.Flush()
	}
}
