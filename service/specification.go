package service

import (
	"github.com/cloudintheking/go-criteria/utils"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type CommonDbOperation struct {
	Selects  []string               //查询字段
	Preloads map[string]interface{} //预加载字段
	Orders   []string               //排序
	Groups   []string               //分组
	Havings  map[string]interface{} //分组筛选
	Limit    *int                   //限制大小
}

/**
 *  @Description: 封装其他操作
 *  @param db
 *  @param operation
 *  @return *gorm.DB
 */
func CreateCommonOperationDb(db *gorm.DB, operation *CommonDbOperation) *gorm.DB {
	if operation != nil {
		if operation.Preloads != nil {
			for pk, pv := range operation.Preloads {
				if pv != nil {
					switch v := pv.(type) {
					case func(pdb *gorm.DB) *gorm.DB: //函数处理
						db = db.Preload(pk, v)
					case []interface{}: //条件处理
						db = db.Preload(pk, v...)
					}
				} else {
					db = db.Preload(pk)
				}
			}
		}
		if operation.Selects != nil {
			db = db.Select(operation.Selects)
		}
		if operation.Orders != nil {
			for _, order := range operation.Orders {
				db = db.Order(order)
			}
		}
		if operation.Limit != nil {
			db = db.Limit(*operation.Limit)
		}
		if operation.Groups != nil {
			for _, group := range operation.Groups {
				db = db.Group(group)
			}
		}
		if operation.Havings != nil {
			for hk, hv := range operation.Havings {
				db = db.Having(hk, utils.Interface2Slice(hv)...)
			}
		}
	}
	return db

}

/**
 * @Description:  创建整数从句
 * @param tx db操作
 * @param name 查询参数名
 * @param criteria 查询条件
 */
type ClauseConstantsEnum struct {
	Regexp    string
	Contains  string
	Equals    string
	NotEquals string
	Lt        string
	Lte       string
	Gt        string
	Gte       string
	In        string
}

var ClauseConstants = ClauseConstantsEnum{
	Regexp:    "regexp",
	Contains:  "like",
	Equals:    "=",
	NotEquals: "<>",
	Lt:        "<",
	Lte:       "<=",
	Gt:        ">",
	Gte:       ">=",
	In:        "in",
}

/**
 * @Description:  创建整数型从句
 * @param tx
 * @param name
 * @param criteria
 */
func BuildIntSpecification(db *gorm.DB, name string, criteria IntFilter) (tx *gorm.DB) {
	tx = db
	var query string
	if criteria.Equals != nil {
		query = strings.Join([]string{name, ClauseConstants.Equals, "?"}, " ")
		tx = db.Where(query, criteria.Equals)
	}
	if criteria.NotEquals != nil {
		query = strings.Join([]string{name, ClauseConstants.NotEquals, "?"}, " ")
		tx = db.Where(query, criteria.NotEquals)
	}
	if criteria.Lt != nil {
		query = strings.Join([]string{name, ClauseConstants.Lt, "?"}, " ")
		tx = db.Where(query, criteria.Lt)
	}
	if criteria.Lte != nil {
		query = strings.Join([]string{name, ClauseConstants.Lte, "?"}, " ")
		tx = db.Where(query, criteria.Lte)
	}
	if criteria.Gt != nil {
		query = strings.Join([]string{name, ClauseConstants.Gt, "?"}, " ")
		tx = db.Where(query, criteria.Gt)
	}
	if criteria.Gte != nil {
		query = strings.Join([]string{name, ClauseConstants.Gte, "?"}, " ")
		tx = db.Where(query, criteria.Gte)
	}
	if criteria.In != nil {
		query = strings.Join([]string{name, ClauseConstants.In, "(?)"}, " ")
		inStrSlice := strings.Split(*criteria.In, ",")
		var inIntSlice []int64
		for _, inStr := range inStrSlice {
			inInt, _ := strconv.ParseInt(inStr, 10, 64)
			inIntSlice = append(inIntSlice, inInt)
		}
		tx = db.Where(query, inIntSlice)
	}
	return
}

/**
 * @Description:  创建浮点型从句
 * @param tx
 * @param name
 * @param criteria
 */
func BuildFloatSpecification(db *gorm.DB, name string, criteria FloatFilter) (tx *gorm.DB) {
	tx = db
	var query string
	if criteria.Equals != nil {
		query = strings.Join([]string{name, ClauseConstants.Equals, "?"}, " ")
		tx = db.Where(query, criteria.Equals)
	}
	if criteria.NotEquals != nil {
		query = strings.Join([]string{name, ClauseConstants.NotEquals, "?"}, " ")
		tx = db.Where(query, criteria.NotEquals)
	}
	if criteria.Lt != nil {
		query = strings.Join([]string{name, ClauseConstants.Lt, "?"}, " ")
		tx = db.Where(query, criteria.Lt)
	}
	if criteria.Lte != nil {
		query = strings.Join([]string{name, ClauseConstants.Lte, "?"}, " ")
		tx = db.Where(query, criteria.Lte)
	}
	if criteria.Gt != nil {
		query = strings.Join([]string{name, ClauseConstants.Gt, "?"}, " ")
		tx = db.Where(query, criteria.Gt)
	}
	if criteria.Gte != nil {
		query = strings.Join([]string{name, ClauseConstants.Gte, "?"}, " ")
		tx = db.Where(query, criteria.Gte)
	}
	if criteria.In != nil {
		query = strings.Join([]string{name, ClauseConstants.In, "(?)"}, " ")
		inStrSlice := strings.Split(*criteria.In, ",")
		var inFloatSlice []float64
		for _, inStr := range inStrSlice {
			inInt, _ := strconv.ParseFloat(inStr, 64)
			inFloatSlice = append(inFloatSlice, inInt)
		}
		tx = db.Where(query, inFloatSlice)
	}
	return
}

/**
 * @Description:  创建字符串型从句
 * @param tx
 * @param name
 * @param criteria
 */
func BuildStringSpecification(db *gorm.DB, name string, criteria StringFilter) (tx *gorm.DB) {
	tx = db
	var query string
	if criteria.Equals != nil {
		query = strings.Join([]string{name, ClauseConstants.Equals, "?"}, " ")
		tx = db.Where(query, criteria.Equals)
	}
	if criteria.NotEquals != nil {
		query = strings.Join([]string{name, ClauseConstants.NotEquals, "?"}, " ")
		tx = db.Where(query, criteria.NotEquals)
	}
	if criteria.Contains != nil {
		query = strings.Join([]string{name, ClauseConstants.Contains, "?"}, " ")
		tx = db.Where(query, "%"+*criteria.Contains+"%")
	}
	if criteria.Regexp != nil {
		query = strings.Join([]string{name, ClauseConstants.Regexp, "?"}, " ")
		tx = db.Where(query, criteria.Contains)
	}
	if criteria.In != nil {
		query = strings.Join([]string{name, ClauseConstants.In, "(?)"}, " ")
		tx = db.Where(query, strings.Split(*criteria.In, ","))
	}
	return

}

/**
 * @Description:  创建时间型从句
 * @param tx
 * @param name
 * @param criteria
 */
func BuildTimeSpecification(db *gorm.DB, name string, criteria TimeFilter) (tx *gorm.DB) {
	tx = db
	var query string
	var timeV time.Time
	var err error

	if criteria.Lt != nil {
		timeV, err = time.ParseInLocation("2006-01-02 15:04:05", *criteria.Lt, time.Local)
		query = strings.Join([]string{name, ClauseConstants.Lt, "?"}, " ")
		tx = db.Where(query, timeV)
	}
	if criteria.Lte != nil {
		timeV, err = time.ParseInLocation("2006-01-02 15:04:05", *criteria.Lte, time.Local)
		query = strings.Join([]string{name, ClauseConstants.Lte, "?"}, " ")
		tx = db.Where(query, timeV)
	}
	if criteria.Gt != nil {
		timeV, err = time.ParseInLocation("2006-01-02 15:04:05", *criteria.Gt, time.Local)
		query = strings.Join([]string{name, ClauseConstants.Gt, "?"}, " ")
		tx = db.Where(query, timeV)
	}
	if criteria.Gte != nil {
		timeV, err = time.ParseInLocation("2006-01-02 15:04:05", *criteria.Gte, time.Local)
		query = strings.Join([]string{name, ClauseConstants.Gte, "?"}, " ")
		tx = db.Where(query, timeV)
	}
	if err != nil {
		panic("时间" + *criteria.Lt + "格式转换异常")
	}
	return
}

/**
 * @Description:  创建布尔型条件从句
 * @param tx
 * @param name
 * @param criteria
 * @return *gorm.DB
 */
func BuildBoolSpecification(db *gorm.DB, name string, criteria BoolFilter) (tx *gorm.DB) {
	tx = db
	var query string
	if criteria.Equals != nil {
		query = strings.Join([]string{name, ClauseConstants.Equals, "?"}, "")
		tx = db.Where(query, criteria.Equals)
	}
	return
}
