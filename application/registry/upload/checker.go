/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present  Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package upload

import (
	"fmt"
	"strings"
	"time"

	"github.com/webx-top/echo"

	"github.com/admpub/color"
	"github.com/admpub/log"
	"github.com/admpub/nging/application/registry/upload/table"
)

// UploadLinkLifeTime 上传链接生存时间
var UploadLinkLifeTime int64 = 86400

// Checker 验证并生成子文件夹名称和文件名称
type Checker func(echo.Context, table.TableInfoStorer) (subdir string, name string, err error)

// DefaultChecker 默认Checker
var DefaultChecker = func(ctx echo.Context, t table.TableInfoStorer) (subdir string, name string, err error) {
	refid := ctx.Formx(`refid`).String()
	timestamp := ctx.Formx(`time`).Int64()
	if len(refid) == 0 {
		refid = `0`
	}
	// Token(ctx.Queries())
	// 验证签名（避免上传接口被滥用）
	if ctx.Form(`token`) != Token(`refid`, refid, `time`, timestamp) {
		err = ctx.E(`令牌错误`)
		return
	}
	if time.Now().Local().Unix()-timestamp > UploadLinkLifeTime {
		err = ctx.E(`上传网址已过期`)
		return
	}
	subdir = fmt.Sprint(refid) + `/`
	//subdir = time.Now().Format(`2006/01/02/`)
	t.SetTableID(refid)
	return
}

// ConfigChecker 系统配置文件上传
func ConfigChecker(ctx echo.Context, t table.TableInfoStorer) (subdir string, name string, err error) {
	group := ctx.Form(`group`)
	key := ctx.Form(`key`)
	timestamp := ctx.Formx(`time`).Int64()
	// 验证签名（避免上传接口被滥用）
	if ctx.Form(`token`) != Token(`group`, group, `key`, key, `refid`, `0`, `time`, timestamp) {
		err = ctx.E(`令牌错误`)
		return
	}
	if time.Now().Local().Unix()-timestamp > UploadLinkLifeTime {
		err = ctx.E(`上传网址已过期`)
		return
	}
	subdir = key + `/`
	t.SetTableID(group + `.` + key)
	t.SetTableName(`nging_config`)
	t.SetFieldName(`value`)
	return
}

func CheckerRegister(typ string, checker Checker, fieldNames ...string) {
	log.Info(color.GreenString(`checker.register:`), typ)
	if len(fieldNames) > 0 {
		SubdirGet(typ).SetChecker(checker, fieldNames...)
		return
	}
	tableName, fieldName, _ := GetTableInfo(typ)
	info := SubdirGet(tableName)
	info.SetChecker(checker, fieldName)
}

func CheckerGet(typ string, defaults ...string) Checker {
	s := SubdirGet(typ)
	if s != nil {
		return s.MustChecker()
	}
	if len(defaults) == 0 {
		tmp := strings.SplitN(typ, `.`, 2)
		if len(tmp) == 2 {
			return CheckerGet(tmp[0])
		}
		return DefaultChecker
	}
	return CheckerGet(defaults[0])
}
