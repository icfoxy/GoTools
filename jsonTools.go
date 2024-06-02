package GoTools

import (
	"encoding/json"
	"io"
	"net/http"
)

// 各类数据转json并响应
func RespondByJSON(w http.ResponseWriter, code int, input any) error {
	//各类型数据转json比特数组
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	// 写入Header中Content - Type信息
	w.Header().Add("Content-Type", "application/json")
	//响应Header
	w.WriteHeader(code)
	//响应body
	w.Write(data)
	return nil
}

// json格式报错
func RespondByErr(w http.ResponseWriter, code int, err error) error {
	err := RespondByJSON(w, code, err)
	return err
}

// 从body中获取Json，并转为任意类型数据
func GetAnyFromBody[T any](body io.Reader, result *T) error {
	err := json.NewDecoder(body).Decode(result)
	return err
}
