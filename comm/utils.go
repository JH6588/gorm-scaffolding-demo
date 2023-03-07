package comm

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"gorm.io/gorm"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// SQL查询条件构建
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		//if len(ks) > 3 {
		//	return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		//}

		if whereSQL != "" {
			whereSQL += " AND "
		}
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
			break
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
				break
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
				break
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
				break
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
				break
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
				break
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
				break
			case "like":
				whereSQL += fmt.Sprint(k, " like ? ")
				vals = append(vals, v)
				break
			case "likeEscaped":
				whereSQL += fmt.Sprint(k, " LIKE  ? escape '/' ")
				vals = append(vals, v)
				break
			}
			break

		default:
			whereSQL += k
			//超过两个问号
			if reflect.ValueOf(v).Kind() == reflect.Slice && strings.Count(k, "?") > 1 {
				for _, x := range reflect.ValueOf(v).Interface().([]string) {
					vals = append(vals, x)
				}
			} else {
				vals = append(vals, v)
			}

		}
	}

	return
}

// 排序设置
type OrderInfo struct {
	SortKey  string `form:"sortKey"`
	SortType uint8  `form:"sortType"`
}

func (t *OrderInfo) GetOrderSql(db *gorm.DB) *gorm.DB {
	if t == nil {
		return db
	}
	orderStr := GetJsonKeyMappingKeyFromStruct(t, t.SortKey)
	if t.SortType == 1 {
		orderStr += " DESC"
	}
	return db.Order(orderStr)
}

// 分页设置
type PageInfo struct {
	Page int32
	Size int32
}

func (t *PageInfo) GetPageSql(db *gorm.DB) *gorm.DB {
	if t == nil {
		return db
	}

	if t.Page <= 1 {
		t.Page = 1

	}
	return db.Limit(int(t.Size)).Offset(int(t.Size) * int(t.Page-1))
}

//-------------反射---------------

// 将请求结构体直接生成map[string]interface{} 的sql搜索条件
func GetSqlBuildMapFromStruct(p interface{}) map[string]interface{} {
	item := make(map[string]interface{})
	typeInfo := reflect.TypeOf(p)
	valueInfo := reflect.ValueOf(p)
	if typeInfo.Kind() == reflect.Ptr {
		typeInfo = typeInfo.Elem()
		valueInfo = valueInfo.Elem()
	}
	for i := 0; i < typeInfo.NumField(); i++ {
		field := valueInfo.Field(i)
		filedValue := field.Interface()
		//零值过滤
		if field.IsZero() {
			continue
		}

		if field.Kind() == reflect.Slice && field.Len() == 0 {
			continue
		}
		//嵌套类型处理
		if field.Kind() == reflect.Struct {
			m := GetSqlBuildMapFromStruct(filedValue)
			for k, v := range m {
				item[k] = v
			}
		}

		sqlK := strings.TrimSpace(typeInfo.Field(i).Tag.Get("sqlK"))
		sqlV := strings.TrimSpace(typeInfo.Field(i).Tag.Get("sqlV"))
		//sqlK,sqlV 为空过滤
		if sqlK == "" && sqlV == "" {
			continue
		} else if sqlV == "" {
			item[sqlK] = filedValue
		} else {
			sqlV = strings.ReplaceAll(sqlV, "{%v}", fmt.Sprintf("%v", field))
			item[sqlK] = sqlV
		}

	}
	return item
}

// 映射tag  如 column:xxx  提取xx
func GetJsonKeyMappingKeyFromStruct[T any](t T, key string) string {
	ele := reflect.TypeOf(t)

	if ele.Kind() == reflect.Ptr {
		ele = ele.Elem()

	}

	for i := 0; i < ele.NumField(); i++ {
		js := strings.TrimSpace(ele.Field(i).Tag.Get("json"))
		fmt.Println(js)
		if js == key {
			ck := strings.TrimSpace(ele.Field(i).Tag.Get("gorm"))

			idx := strings.Index(ck, "column:")
			if idx == -1 {
				break

			}

			return strings.TrimSpace(strings.Split(ck[strings.Index(ck, "column:")+len("column:"):], ",")[0])
		}

	}
	return ToSnake(key)
}

//----------字符串操作----------

// 获取随机长度纯英文字符串
func GetRandomString(dataLen int) string {

	rand.Seed(time.Now().UnixNano())
	var rs []byte
	// String
	charset := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < dataLen; i++ {
		rs = append(rs, charset[rand.Intn(len(charset))])
	}

	return string(rs)
}

// 驼峰转蛇形
func ToSnake(str string) string {
	var rs []rune
	for _, r := range []rune(str) {
		if r > 126 {
			panic("non ascii insied!")
		}

		if 65 <= r && r < 91 {
			rs = append(rs, 95, r+32)
		} else {
			rs = append(rs, r)
		}
	}

	if rs[0] == 95 {
		rs = rs[1:]
	}

	return string(rs)
}
