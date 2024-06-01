package GoTools

import (
	"encoding/json"
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
)

// 存入或更新数据，传入数据库名、键名、值
func DBPut(DBName string, key, value any) error {
	db, err := leveldb.OpenFile(DBName, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	kData, err := json.Marshal(key)
	if err != nil {
		return err
	}
	vData, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = db.Put(kData, vData, nil)
	return err
}

// 获取数据，传入数据库名、键名、值指针
func DBGet[T any](DBName string, key any, value *T) (err error) {
	db, err := leveldb.OpenFile(DBName, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	kData, err := json.Marshal(key)
	if err != nil {
		return err
	}
	data, err := db.Get(kData, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, value)
	return err
}

// 删除数据，传入数据库名、键名
func DBDelete(DBName string, key any) error {
	db, err := leveldb.OpenFile(DBName, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	kData, err := json.Marshal(key)
	if err != nil {
		return err
	}
	err = db.Delete(kData, nil)
	return err
}

// 成组存入或更新数据库，传入数据库名、键组、值组
func DBPutList[T1 any, T2 any](DBName string, KList []T1, VList []T2) error {
	if len(KList) != len(VList) {
		return errors.New("wrong length")
	}
	db, err := leveldb.OpenFile(DBName, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	for i, key := range KList {
		kData, err := json.Marshal(key)
		if err != nil {
			return err
		}
		vData, err := json.Marshal(VList[i])
		if err != nil {
			return err
		}
		err = db.Put(kData, vData, nil)
		if err != nil {
			return err
		}
	}
	return err
}

// 成组读取数据库,传入数据库名、键组、值组
func DBGetList[T1 any, T2 any](DBName string, KList []T1, VList []T2) error {
	if len(KList) != len(VList) {
		return errors.New("wrong length")
	}
	db, err := leveldb.OpenFile(DBName, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	for i, key := range KList {
		kData, err := json.Marshal(key)
		if err != nil {
			return err
		}
		data, err := db.Get(kData, nil)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &VList[i])
		if err != nil {
			return err
		}
	}
	return err
}
