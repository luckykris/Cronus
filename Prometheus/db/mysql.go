package db

import (
	"database/sql"
	"fmt"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type crud interface{
	Query(string, ...interface {}) (*sql.Rows, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

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
type Tx struct{
	inher *sql.Tx
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

func (db *MysqlDb) Begin()(Txi,error){
	tx,err:=db.DbPool.Begin()
	t:=new(Tx)
	t.inher=tx
	return t,err
}
func (tx *Tx)Commit()error{
		return tx.inher.Commit()
}
func (tx *Tx)Rollback()error{
	return tx.inher.Rollback()
}

func combine_get_sql(crud_i crud,table string, groupby interface{},columns_name []string, conditions []string, columns ...interface{})(*Cur, error){
	var conditions_str = ""
	var groupby_str = ""
	if len(conditions) != 0 {
		conditions_str = " WHERE " + strings.Join(conditions, ` AND `)
	}
	if groupby !=nil{
		groupby_str=fmt.Sprintf("GROUP BY %s",groupby.(string))
	}
	sql := fmt.Sprintf("SELECT %s FROM %s %s %s", strings.Join(columns_name, ","), table, conditions_str,groupby_str)
	log.Debug(sql)
	rows, err :=crud_i.Query(sql)
	return &Cur{rows, columns_name, columns}, err
}
func combine_getjoin_sql(crud_i crud,ltable string,tables [][4]string, groupby interface{},columns_name []string, conditions []string, columns ...interface{})(*Cur, error){
	table:=ltable
	for i:=range tables{
		table+=fmt.Sprintf(" %s JOIN %s ON %s=%s " ,tables[i][0],tables[i][1],tables[i][2],tables[i][3])
	}
	return combine_get_sql(crud_i,table, groupby,columns_name, conditions, columns ...)
}
func combine_add_sql(crud_i crud ,table string, columns_name []string, values [][]interface{})error{
	values3 := []string{}
	for row := range values {
		values2 := []string{}
		for i := range columns_name {
			values2 = append(values2, _CheckTypeAndModifyString(values[row][i]))
		}
		values3 = append(values3, strings.Join(values2, `,`))
	}
	sql := fmt.Sprintf("INSERT INTO %s (`%s`) VALUES (%s)", table, strings.Join(columns_name, "`,`"), strings.Join(values3, `),(`))
	log.Debug(sql)
	_, err := crud_i.Exec(sql)
	if err != nil {
		log.Debug(fmt.Sprintf("Mysql insert failed,sql:%s,err:%s", sql, err.Error()))
	}
	return err
}
func combine_update_sql(crud_i crud,table string, conditions []string, columns_name []string, values []interface{})error{
	kv := []string{}
	conditions_str := strings.Join(conditions, ` AND `)
	for i := range columns_name {
		kv = append(kv, fmt.Sprintf(" `%s` = %s", columns_name[i], _CheckTypeAndModifyString(values[i])))
	}
	sql := fmt.Sprintf(`UPDATE %s SET %s WHERE %s`, table, strings.Join(kv, `,`), conditions_str)
	log.Debug(sql)
	_, err := crud_i.Exec(sql)
	if err != nil {
		log.Debug(fmt.Sprintf("Mysql Update failed,sql:%s,err:%s", sql, err.Error()))
	}
	return err
}
func combine_delete_sql(crud_i crud, table string,conditions []string)error{
	sql := fmt.Sprintf(`DELETE FROM %s WHERE %s`, table, strings.Join(conditions, ` AND `))
	log.Debug(sql)
	_, err := crud_i.Exec(sql)
	if err != nil {
		log.Debug(fmt.Sprintf("Mysql delete failed,sql:%s,err:%s", sql, err.Error()))
	}
	return err
}

func (db *MysqlDb) Get(table string, groupby interface{},columns_name []string, conditions []string, columns ...interface{}) (*Cur, error) {
	return combine_get_sql(db.DbPool,table, groupby,columns_name, conditions, columns ...)
}
func (db *MysqlDb) GetJoin(ltable string,tables [][4]string, groupby interface{},columns_name []string, conditions []string, columns ...interface{}) (*Cur, error) {	
	return combine_getjoin_sql(db.DbPool,ltable, tables,groupby,columns_name, conditions, columns ...)
}
func (db *MysqlDb) Add(table string, columns_name []string, values [][]interface{}) error {
	return combine_add_sql(db.DbPool,table , columns_name , values)
}
func (db *MysqlDb) Delete(table string, conditions []string) error {
		return combine_delete_sql(db.DbPool,table , conditions)
}
func (db *MysqlDb) Update(table string, conditions []string, columns_name []string, values []interface{}) error {
	return combine_update_sql(db.DbPool,table , conditions, columns_name, values )
}


func (tx *Tx) Get(table string, groupby interface{},columns_name []string, conditions []string, columns ...interface{}) (*Cur, error) {
	return combine_get_sql(tx.inher,table, groupby,columns_name, conditions, columns ...)
}
func (tx *Tx) GetJoin(ltable string,tables [][4]string, groupby interface{},columns_name []string, conditions []string, columns ...interface{}) (*Cur, error) {	
	return combine_getjoin_sql(tx.inher,ltable, tables,groupby,columns_name, conditions, columns ...)
}
func (tx *Tx)Add(table string, columns_name []string, values [][]interface{}) error {
	return combine_add_sql(tx.inher,table , columns_name , values)
}
func (tx *Tx) Delete(table string, conditions []string) error {
		return combine_delete_sql(tx.inher,table , conditions)
}
func (tx *Tx) Update(table string, conditions []string, columns_name []string, values []interface{}) error {
	return combine_update_sql(tx.inher,table , conditions, columns_name, values )
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

func (c *Cur) Close()  {
	c.row.Close()
}



func _CheckTypeAndModifyString(v interface{}) string {
	switch v.(type) {
	case string:
		return `'` + v.(string) + `'`
	case uint8:
		return fmt.Sprintf("%d", v.(uint8))
	case int:
		return fmt.Sprintf("%d", v.(int))
	case uint32:
		return fmt.Sprintf("%d", v.(uint32))
	case int64:
		return fmt.Sprintf("%d", v.(int64))
	case float64:
		return fmt.Sprintf("%f", v.(float64))
	case nil:
		return `null`
	default:
		log.Debug("Mysql Dbi Can't analyis value's type")
		return "Unkonw Type"
	}
}
