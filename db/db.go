package db

import (
    "database/sql"
    "fmt"
    "gitee.com/clearluo/gotools/util"
    "reflect"
    "sync"

    "gitee.com/clearluo/gotools/log"

    "xorm.io/xorm"
)
type SqlData struct {
    Sql   string
    Param []interface{}
}

// 加载静态数据到缓存
func InitMap() {
    defer util.Profiling("db.InitMap")()

    loadTables := []interface{}{}
    w := sync.WaitGroup{}
    for i, v := range loadTables {
        w.Add(1)
        go func(tableStruct interface{}, i int) {
            defer w.Done()
            defer func() {
                if err := recover(); err != nil {
                    log.Errorf("Recover in IninRedisNew:%v", i)
                }
            }()
            m := reflect.ValueOf(tableStruct).MethodByName("LoadDatas")
            if m.IsValid() && m.Kind() == reflect.Func {
                m.Call(nil)
            }
        }(v, i)
        w.Wait() // 并行运行
    }
    // w.Wait() // 并发运行
}

// 通过事务提交sql
// isCheck，更新有效记录是否必须不为0
func ExeSqlByTransaction(engine xorm.Engine, isCheck bool, sqlSlice ...*SqlData) error {
    var sqlResult sql.Result
    session := engine.NewSession()
    defer session.Close()
    err := session.Begin()
    for i := range sqlSlice {
        if sqlSlice[i].Sql == "" {
            continue
        }
        execSql := []interface{}{sqlSlice[i].Sql}
        execSql = append(execSql, sqlSlice[i].Param...)
        log.Infof("execute sql:%v", execSql)
        if sqlResult, err = session.Exec(execSql...); err != nil {
            log.Error(err)
            break
        } else {
            num, _ := sqlResult.RowsAffected()
            log.Infof("num:%v", num)
            if isCheck && num < 1 {
                err = fmt.Errorf("rowAffect num:%v", num)
                log.Warn(err)
                break
            }
        }
    }
    if err != nil {
        err = fmt.Errorf("[ExeSqlByTransaction.fail]:%v", err)
        log.Error(err)
        session.Rollback()
        return err
    }
    err = session.Commit()
    if err != nil {
        err = fmt.Errorf("[ExeSqlByTransaction.fail]:%v", err)
        log.Error(err)
        return err
    }
    return nil
}

// 通过反射调用结构对应的TableName函数,达到返回表名的目的
func GetTableName(tableStruct interface{}) string {
    m := reflect.ValueOf(tableStruct).MethodByName("TableName")
    if m.IsValid() && m.Kind() == reflect.Func {
        re := m.Call(nil)
        for _, v := range re {
            if v.IsValid() {
                return v.String()
            }
        }
    }
    return "unknown"
}

// 根据主键 id通用更新表数据
// tableStruct:为表映射后的结构指针
// updateMap:为更新表数据的 map 结构，期中必须包含主键 id
func UpdateTableById(engine *xorm.Engine, tableStruct interface{}, updateMap map[string]interface{}) error {
    if reflect.TypeOf(tableStruct).Kind() != reflect.Ptr {
        err := fmt.Errorf("tableStruct must ptr")
        log.Error(err)
        return err
    }
    id, ok := updateMap["id"]
    if !ok {
        err := fmt.Errorf("updateMap not found id")
        log.Warn(err)
        return err
    }
    num, err := engine.Table(tableStruct).ID(id).Update(updateMap)
    if err != nil {
        log.Warn(err)
    }
    _ = num
    //log.Infof("[UpdateTableById.%v] id:%v effectNum:%v\n", GetTableName(tableStruct), id, num)
    return nil
}

func GetTableRowById(engine *xorm.Engine,o interface{}, id int) error {
    // 注意！！！o为对应表映射结构体指针，(必须是指针)
    tableName := GetTableName(o)
    if tableName == "unknown" {
        err := fmt.Errorf("GetTableRowById param err")
        log.Warn(err)
        return err
    }
    sql := fmt.Sprintf("SELECT * FROM %v WHERE id = ?", tableName)
    ok, err := engine.SQL(sql, id).Get(o)
    if err != nil {
        log.Warn("err: " + err.Error())
        return err
    }
    if !ok {
        err = fmt.Errorf("%v not found by id:%v", tableName, id)
        log.Warn(err)
        return err
    }
    return nil
}
