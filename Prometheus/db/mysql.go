package db

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type MysqlDb struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	Db           string
	Charset      string
	MaxLifeTime  int64
	MaxIdleConns int64
	DbPool       *sql.DB
}
type Cur struct {
	row          *sql.Rows
	columns_name []string
	columns      []interface{}
}

func (db *MysqlDb) Start() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", db.User, db.Passwd, db.Host, int(db.Port), db.Db, db.Charset)
	log.Debug("DSN==>", dsn)
	DbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DbPool = DbPool
	//dr := time.ParseDuration(fmt.Sprintf("%ds", t))
	db.DbPool.SetConnMaxLifetime(time.Duration(db.MaxLifeTime) * time.Second)
	db.DbPool.SetMaxIdleConns(int(db.MaxIdleConns))
	return db.Ping()
}
func (db *MysqlDb) Ping() error {
	return db.DbPool.Ping()
}

func (db *MysqlDb) Close() error {
	return db.DbPool.Close()
}

func (db *MysqlDb) Get(table string, columns_name []string, conditions []string, columns ...interface{}) (*Cur, error) {
	var conditions_str = ""
	if len(conditions) != 0 {
		conditions_str = " WHERE " + strings.Join(conditions, ` AND `)
	}
	sql := fmt.Sprintf(`SELECT %s FROM %s %s`, strings.Join(columns_name, `,`), table, conditions_str)
	rows, err := db.DbPool.Query(sql)
	return &Cur{rows, columns_name, columns}, err

}
func (c *Cur) Fetch() bool {
	if c.row.Next() {
		err := c.row.Scan(c.columns...)
		if err != nil {
			log.Fatal("db fetch failed:", err.Error())
			return false
		}
		return true
	} else {
		c.row.Close()
		return false
	}
}

func (db *MysqlDb) Add(table string, columns_name []string, values [][]interface{}) error {
	values3 := []string{}
	for row := range values {
		values2 := []string{}
		for i := range columns_name {
			values2 = append(values2, _CheckTypeAndModifyString(values[row][i]))
		}
		values3 = append(values3, strings.Join(values2, `,`))
	}
	sql := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, table, strings.Join(columns_name, `,`), strings.Join(values3, `),(`))
	_, err := db.DbPool.Exec(sql)
	if err != nil {
		log.Debug("Mysql insert failed,sql:%s,err:%s", sql, err.Error())
	}
	return err

}

func (db *MysqlDb) Delete(table string, conditions []string) error {
	sql := fmt.Sprintf(`DELETE FROM %s WHERE %s`, table, strings.Join(conditions, ` AND `))
	_, err := db.DbPool.Exec(sql)
	if err != nil {
		log.Debug("Mysql delete failed,sql:%s,err:%s", sql, err.Error())
	}
	return err

}

func (db *MysqlDb) Update(table string, conditions []string, columns_name []string, values []interface{}) error {
	kv := []string{}
	conditions_str := strings.Join(conditions, ` AND `)
	for i := range columns_name {
		kv = append(kv, fmt.Sprintf(" %s = %s", columns_name[i], _CheckTypeAndModifyString(values[i])))
	}
	sql := fmt.Sprintf(`UPDATE %s SET %s WHERE %s`, table, strings.Join(kv, `,`), conditions_str)
	_, err := db.DbPool.Exec(sql)
	if err != nil {
		log.Debug("Mysql Update failed,sql:%s,err:%s", sql, err.Error())
	}
	return err
}

func _CheckTypeAndModifyString(v interface{}) string {
	switch v.(type) {
	case string:
		return `'` + v.(string) + `'`
	case int:
		return fmt.Sprintf("%d", v.(int))
	default:
		log.Debug("Mysql Dbi Can't analyis value's type")
		return "Unkonw Type"
	}
}
