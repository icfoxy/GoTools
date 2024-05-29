package GoTools

import (
	"encoding/json"
	"io"
	"net/http"
)

// 数据转json并响应
func RespondByJSON(w http.ResponseWriter, code int, input any) error {
	//各类型数据转json比特数组
	data, err := ToJsonByte(input)
	if err != nil {
		return err
	}
	//写入Header中Content-Type信息
	w.Header().Add("Content-Type", "application/json")
	//响应Header
	w.WriteHeader(code)
	//响应body
	w.Write(data)
	return nil
}

// json格式报错
func RespondByErr(w http.ResponseWriter, code int, errInfo string) {
	RespondByJSON(w, code, struct {
		Err string
	}{
		Err: errInfo,
	})
}

// 从Body中获取json
func GetJson(body io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
