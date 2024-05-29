package GoTools

import (
	"encoding/json"
	"io"
	"net/http"
)

// 任意类型转Json比特数组，传入任意类型，输出比特数组
func AnyToJsonByte(input any) (result []byte, err error) {
	return json.Marshal(input)
}

// Json比特数组转相应类型，传入比特数组，输出对应类型数据
func JsonByteToAny[T any](input []byte, result *T) error {
	return json.Unmarshal(input, result)
}

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

// 从body中获取Json字比特数组
func GetJsonByteFromBody(body io.Reader) ([]byte, error) {
	var data []byte
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 从body中获取任意类型数据
func GetAnyFromBody[T any](body io.Reader, result *T) error {
	byteData, err := GetJsonByteFromBody(body)
	if err != nil {
		return err
	}
	err = JsonByteToAny(byteData, result)
	return err
}
