package libs

import (
	"crypto/md5"
	"fmt"
)

//字符串截取
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return str
	}
	if end < 0 || end > length {
		return str
	}
	return string(rs[start:end])
}

//md5加密
func S_md5(s string) string {
	sign := md5.Sum([]byte(s))
	ss := fmt.Sprintf("%X", sign)
	return ss
}

//读取文件内容
func Readf(name string) (int, string, error) {
	f, err := os.Open(name)
	if err != nil {
		return 0, "", err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	//	//读取@之前的字符串
	//	if result, err := br.ReadString(byte('@')); err == nil {
	//		fmt.Println("内容:", result)
	//	}

	buf := make([]byte, 1024)
	n, err := br.Read(buf)
	if err != nil {
		return 0, "", err
	}
	return n, string(buf), nil
}

//字符串写入文件
func Writef(name, content string) (int, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	w := bufio.NewWriterSize(f, 1096)
	n, err1 := w.WriteString(content)
	if err1 != nil {
		return 0, err1
	}

	if err2 := w.Flush(); err != nil {
		return 0, err2
	}
	return n, nil

	//	//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
	//	buf := []byte(content)
	//	n, err := w.Write(buf)
	//	if err != nil {
	//		return 0, err
	//	}
	//	if err := w.Flush(); err != nil {
	//		return 0, err
	//	}
	//	return n, nil

}

//http get 请求
func Httpget(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36`)
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return string(body), err
}
