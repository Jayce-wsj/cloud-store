// 数据库工具包
package db

import (
	"database/sql"

	"golang.org/x/crypto/ssh"
)

//
var DefaultDB *Database

// Database 数据容器抽象对象定义
type Database struct {
	Type      string // 用来给 SqlBuilder 进行一些特殊的判断 (空值或 mysql 皆表示这是一个 MySQL 实例)
	DB        *sql.DB
	SSHClient *ssh.Client
}

// Close 关闭数据库连接
func (dba *Database) Close() error {
	defer func() {
		if dba.SSHClient != nil {
			_ = dba.SSHClient.Close()
		}
	}()
	return dba.DB.Close()
}

// Exec 执行语句
func (dba *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	return dba.DB.Exec(query, args...)
}

// Query 查询单条记录
func (dba *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dba.DB.Query(query, args...)
}

// QueryRow 查询单条记录
func (dba *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	return dba.DB.QueryRow(query, args...)
}

// Update 执行 UPDATE 语句并返回受影响的行数
// 返回0表示没有出错, 但没有被更新的行
// 返回-1表示出错
func (dba *Database) Update(query string, args ...interface{}) (int64, error) {
	ret, err := dba.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	aff, err := ret.RowsAffected()
	if err != nil {
		return -1, err
	}
	return aff, nil
}

// Delete 执行 DELETE 语句并返回受影响的行数
// 返回0表示没有出错, 但没有被删除的行
// 返回-1表示出错
func (dba *Database) Delete(query string, args ...interface{}) (int64, error) {
	return dba.Update(query, args...)
}

// Insert 执行 INSERT 语句并返回最后生成的自增ID
// 返回0表示没有出错, 但没生成自增ID
// 返回-1表示出错
func (dba *Database) Insert(query string, args ...interface{}) (int64, error) {
	ret, err := dba.Exec(query, args...)
	if err != nil {
		return -1, err
	}
	last, err := ret.LastInsertId()
	if err != nil {
		return -1, err

	}
	return last, nil
}

// Select 查询不定字段的结果集
func (dba *Database) Select(query string, args ...interface{}) (Results, error) {
	rows, err := dba.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rowsToResults(rows)
}

// SelectOne 查询一行不定字段的结果
func (dba *Database) SelectOne(query string, args ...interface{}) (OneRow, error) {
	ret, err := dba.Select(query, args...)
	if err != nil {
		return nil, err
	}
	if len(ret) > 0 {
		return ret[0], nil
	}
	return make(OneRow), nil
}

func rowsToResults(rows *sql.Rows) (Results, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colNum := len(cols)
	rawValues := make([][]byte, colNum)

	// query.Scan 的参数，因为每次查询出来的列是不定长的，所以传入长度固定当次查询的长度
	scans := make([]interface{}, len(cols))

	// 将每行数据填充到[][]byte里
	for i := range rawValues {
		scans[i] = &rawValues[i]
	}

	results := make(Results, 0)
	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return nil, err
		}

		row := make(map[string]string)

		for k, raw := range rawValues {
			key := cols[k]
			/*if raw == nil {
				row[key] = "\\N"
			} else {*/
			row[key] = string(raw)
			//}
		}
		results = append(results, row)
	}
	return results, nil
}

// Results 多行数据集结果
type Results []OneRow

// OneRow 单行查询结果
type OneRow map[string]string

// Set 设置值
func (row OneRow) Set(key, val string) {
	row[key] = val
}

// Exist 判断字段是否存在
func (row OneRow) Exist(field string) bool {
	if _, ok := row[field]; ok {
		return true
	}
	return false
}

// Get 获取指定字段的值
func (row OneRow) Get(field string) string {
	if v, ok := row[field]; ok {
		return v
	}
	return ""
}

// GetInt8 获取指定字段的 int8 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetInt8(field string) int8 {
	if v, ok := row[field]; ok {
		return Atoi8(v)
	}
	return 0
}

// GetInt16 获取指定字段的 int16 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetInt16(field string) int16 {
	if v, ok := row[field]; ok {
		return Atoi16(v)
	}
	return 0
}

// GetInt 获取指定字段的 int 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetInt(field string) int {
	if v, ok := row[field]; ok {
		return Atoi(v)
	}
	return 0
}

// GetInt 获取指定字段的 int32 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetInt32(field string) int32 {
	if v, ok := row[field]; ok {
		return Atoi32(v)
	}
	return 0
}

// GetInt64 获取指定字段的 int64 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetInt64(field string) int64 {
	if v, ok := row[field]; ok {
		return Atoi64(v)
	}
	return 0
}

// GetUint8 获取指定字段的 uint8 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetUint8(field string) uint8 {
	if v, ok := row[field]; ok {
		return AtoUi8(v)
	}
	return 0
}

// GetUint16 获取指定字段的 uint16 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetUint16(field string) uint16 {
	if v, ok := row[field]; ok {
		return AtoUi16(v)
	}
	return 0
}

// GetUint 获取指定字段的 uint 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetUint(field string) uint {
	if v, ok := row[field]; ok {
		return AtoUi(v)
	}
	return 0
}

// GetUint 获取指定字段的 uint32 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetUint32(field string) uint32 {
	if v, ok := row[field]; ok {
		return AtoUi32(v)
	}
	return 0
}

// GetUint64 获取指定字段的 uint64 类型值, 注意, 如果该字段不存在则会返回0
func (row OneRow) GetUint64(field string) uint64 {
	if v, ok := row[field]; ok {
		return AtoUi64(v)
	}
	return 0
}
