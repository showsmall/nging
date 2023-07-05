// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/db/lib/factory/pagination"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingTask []*NgingTask

func (s Slice_NgingTask) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingTask) RangeRaw(fn func(m *NgingTask) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingTask) GroupBy(keyField string) map[string][]*NgingTask {
	r := map[string][]*NgingTask{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingTask{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingTask) KeyBy(keyField string) map[string]*NgingTask {
	r := map[string]*NgingTask{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingTask) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingTask) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingTask) FromList(data interface{}) Slice_NgingTask {
	values, ok := data.([]*NgingTask)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingTask{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingTask(ctx echo.Context) *NgingTask {
	m := &NgingTask{}
	m.SetContext(ctx)
	return m
}

// NgingTask 任务
type NgingTask struct {
	base    factory.Base
	objects []*NgingTask

	Id            uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	Uid           uint   `db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	GroupId       uint   `db:"group_id" bson:"group_id" comment:"分组ID" json:"group_id" xml:"group_id"`
	Name          string `db:"name" bson:"name" comment:"任务名称" json:"name" xml:"name"`
	Type          int    `db:"type" bson:"type" comment:"任务类型" json:"type" xml:"type"`
	Description   string `db:"description" bson:"description" comment:"任务描述" json:"description" xml:"description"`
	CronSpec      string `db:"cron_spec" bson:"cron_spec" comment:"时间表达式" json:"cron_spec" xml:"cron_spec"`
	Concurrent    uint   `db:"concurrent" bson:"concurrent" comment:"是否支持多实例" json:"concurrent" xml:"concurrent"`
	Command       string `db:"command" bson:"command" comment:"命令详情" json:"command" xml:"command"`
	WorkDirectory string `db:"work_directory" bson:"work_directory" comment:"工作目录" json:"work_directory" xml:"work_directory"`
	Env           string `db:"env" bson:"env" comment:"环境变量(一行一个，格式为：var1=val1)" json:"env" xml:"env"`
	Disabled      string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	EnableNotify  uint   `db:"enable_notify" bson:"enable_notify" comment:"是否启用通知" json:"enable_notify" xml:"enable_notify"`
	NotifyEmail   string `db:"notify_email" bson:"notify_email" comment:"通知人列表" json:"notify_email" xml:"notify_email"`
	Timeout       uint64 `db:"timeout" bson:"timeout" comment:"超时设置" json:"timeout" xml:"timeout"`
	ExecuteTimes  uint   `db:"execute_times" bson:"execute_times" comment:"累计执行次数" json:"execute_times" xml:"execute_times"`
	PrevTime      uint   `db:"prev_time" bson:"prev_time" comment:"上次执行时间" json:"prev_time" xml:"prev_time"`
	Created       uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated       uint   `db:"updated" bson:"updated" comment:"修改时间" json:"updated" xml:"updated"`
	ClosedLog     string `db:"closed_log" bson:"closed_log" comment:"是否(Y/N)关闭日志" json:"closed_log" xml:"closed_log"`
}

// - base function

func (a *NgingTask) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingTask) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingTask) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingTask) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingTask) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingTask) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingTask) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingTask) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingTask) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingTask) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingTask) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingTask) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingTask) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingTask) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingTask) Objects() []*NgingTask {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingTask) XObjects() Slice_NgingTask {
	return Slice_NgingTask(a.Objects())
}

func (a *NgingTask) NewObjects() factory.Ranger {
	return &Slice_NgingTask{}
}

func (a *NgingTask) InitObjects() *[]*NgingTask {
	a.objects = []*NgingTask{}
	return &a.objects
}

func (a *NgingTask) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingTask) Short_() string {
	return "nging_task"
}

func (a *NgingTask) Struct_() string {
	return "NgingTask"
}

func (a *NgingTask) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingTask) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingTask) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingTask) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingTask:
			err = DBI.FireReaded(a, queryParam, Slice_NgingTask(*v))
		case []*NgingTask:
			err = DBI.FireReaded(a, queryParam, Slice_NgingTask(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingTask) GroupBy(keyField string, inputRows ...[]*NgingTask) map[string][]*NgingTask {
	var rows Slice_NgingTask
	if len(inputRows) > 0 {
		rows = Slice_NgingTask(inputRows[0])
	} else {
		rows = Slice_NgingTask(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingTask) KeyBy(keyField string, inputRows ...[]*NgingTask) map[string]*NgingTask {
	var rows Slice_NgingTask
	if len(inputRows) > 0 {
		rows = Slice_NgingTask(inputRows[0])
	} else {
		rows = Slice_NgingTask(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingTask) AsKV(keyField string, valueField string, inputRows ...[]*NgingTask) param.Store {
	var rows Slice_NgingTask
	if len(inputRows) > 0 {
		rows = Slice_NgingTask(inputRows[0])
	} else {
		rows = Slice_NgingTask(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingTask) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingTask:
			err = DBI.FireReaded(a, queryParam, Slice_NgingTask(*v))
		case []*NgingTask:
			err = DBI.FireReaded(a, queryParam, Slice_NgingTask(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingTask) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ClosedLog) == 0 {
		a.ClosedLog = "N"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingTask) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ClosedLog) == 0 {
		a.ClosedLog = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingTask) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ClosedLog) == 0 {
		a.ClosedLog = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Updatex()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingTask) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ClosedLog) == 0 {
		a.ClosedLog = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingTask) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.ClosedLog) == 0 {
		a.ClosedLog = "N"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingTask) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingTask) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingTask) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["closed_log"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["closed_log"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingTask) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["closed_log"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["closed_log"] = "N"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Updatex()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(kvset).Updatex(); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", &m, editColumns, mw, args...)
	return
}

func (a *NgingTask) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(keysValues).Update()
	}
	m := *a
	m.FromRow(keysValues.Map())
	if err = DBI.FireUpdate("updating", &m, keysValues.Keys(), mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(keysValues).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, keysValues.Keys(), mw, args...)
}

func (a *NgingTask) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.ClosedLog) == 0 {
			a.ClosedLog = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.ClosedLog) == 0 {
			a.ClosedLog = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingTask) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingTask) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingTask) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingTask) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingTask) Reset() *NgingTask {
	a.Id = 0
	a.Uid = 0
	a.GroupId = 0
	a.Name = ``
	a.Type = 0
	a.Description = ``
	a.CronSpec = ``
	a.Concurrent = 0
	a.Command = ``
	a.WorkDirectory = ``
	a.Env = ``
	a.Disabled = ``
	a.EnableNotify = 0
	a.NotifyEmail = ``
	a.Timeout = 0
	a.ExecuteTimes = 0
	a.PrevTime = 0
	a.Created = 0
	a.Updated = 0
	a.ClosedLog = ``
	return a
}

func (a *NgingTask) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Uid"] = a.Uid
		r["GroupId"] = a.GroupId
		r["Name"] = a.Name
		r["Type"] = a.Type
		r["Description"] = a.Description
		r["CronSpec"] = a.CronSpec
		r["Concurrent"] = a.Concurrent
		r["Command"] = a.Command
		r["WorkDirectory"] = a.WorkDirectory
		r["Env"] = a.Env
		r["Disabled"] = a.Disabled
		r["EnableNotify"] = a.EnableNotify
		r["NotifyEmail"] = a.NotifyEmail
		r["Timeout"] = a.Timeout
		r["ExecuteTimes"] = a.ExecuteTimes
		r["PrevTime"] = a.PrevTime
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		r["ClosedLog"] = a.ClosedLog
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Uid":
			r["Uid"] = a.Uid
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Name":
			r["Name"] = a.Name
		case "Type":
			r["Type"] = a.Type
		case "Description":
			r["Description"] = a.Description
		case "CronSpec":
			r["CronSpec"] = a.CronSpec
		case "Concurrent":
			r["Concurrent"] = a.Concurrent
		case "Command":
			r["Command"] = a.Command
		case "WorkDirectory":
			r["WorkDirectory"] = a.WorkDirectory
		case "Env":
			r["Env"] = a.Env
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "EnableNotify":
			r["EnableNotify"] = a.EnableNotify
		case "NotifyEmail":
			r["NotifyEmail"] = a.NotifyEmail
		case "Timeout":
			r["Timeout"] = a.Timeout
		case "ExecuteTimes":
			r["ExecuteTimes"] = a.ExecuteTimes
		case "PrevTime":
			r["PrevTime"] = a.PrevTime
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		case "ClosedLog":
			r["ClosedLog"] = a.ClosedLog
		}
	}
	return r
}

func (a *NgingTask) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "type":
			a.Type = param.AsInt(value)
		case "description":
			a.Description = param.AsString(value)
		case "cron_spec":
			a.CronSpec = param.AsString(value)
		case "concurrent":
			a.Concurrent = param.AsUint(value)
		case "command":
			a.Command = param.AsString(value)
		case "work_directory":
			a.WorkDirectory = param.AsString(value)
		case "env":
			a.Env = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "enable_notify":
			a.EnableNotify = param.AsUint(value)
		case "notify_email":
			a.NotifyEmail = param.AsString(value)
		case "timeout":
			a.Timeout = param.AsUint64(value)
		case "execute_times":
			a.ExecuteTimes = param.AsUint(value)
		case "prev_time":
			a.PrevTime = param.AsUint(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		case "closed_log":
			a.ClosedLog = param.AsString(value)
		}
	}
}

func (a *NgingTask) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Type":
			a.Type = param.AsInt(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "CronSpec":
			a.CronSpec = param.AsString(vv)
		case "Concurrent":
			a.Concurrent = param.AsUint(vv)
		case "Command":
			a.Command = param.AsString(vv)
		case "WorkDirectory":
			a.WorkDirectory = param.AsString(vv)
		case "Env":
			a.Env = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "EnableNotify":
			a.EnableNotify = param.AsUint(vv)
		case "NotifyEmail":
			a.NotifyEmail = param.AsString(vv)
		case "Timeout":
			a.Timeout = param.AsUint64(vv)
		case "ExecuteTimes":
			a.ExecuteTimes = param.AsUint(vv)
		case "PrevTime":
			a.PrevTime = param.AsUint(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		case "ClosedLog":
			a.ClosedLog = param.AsString(vv)
		}
	}
}

func (a *NgingTask) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["uid"] = a.Uid
		r["group_id"] = a.GroupId
		r["name"] = a.Name
		r["type"] = a.Type
		r["description"] = a.Description
		r["cron_spec"] = a.CronSpec
		r["concurrent"] = a.Concurrent
		r["command"] = a.Command
		r["work_directory"] = a.WorkDirectory
		r["env"] = a.Env
		r["disabled"] = a.Disabled
		r["enable_notify"] = a.EnableNotify
		r["notify_email"] = a.NotifyEmail
		r["timeout"] = a.Timeout
		r["execute_times"] = a.ExecuteTimes
		r["prev_time"] = a.PrevTime
		r["created"] = a.Created
		r["updated"] = a.Updated
		r["closed_log"] = a.ClosedLog
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "uid":
			r["uid"] = a.Uid
		case "group_id":
			r["group_id"] = a.GroupId
		case "name":
			r["name"] = a.Name
		case "type":
			r["type"] = a.Type
		case "description":
			r["description"] = a.Description
		case "cron_spec":
			r["cron_spec"] = a.CronSpec
		case "concurrent":
			r["concurrent"] = a.Concurrent
		case "command":
			r["command"] = a.Command
		case "work_directory":
			r["work_directory"] = a.WorkDirectory
		case "env":
			r["env"] = a.Env
		case "disabled":
			r["disabled"] = a.Disabled
		case "enable_notify":
			r["enable_notify"] = a.EnableNotify
		case "notify_email":
			r["notify_email"] = a.NotifyEmail
		case "timeout":
			r["timeout"] = a.Timeout
		case "execute_times":
			r["execute_times"] = a.ExecuteTimes
		case "prev_time":
			r["prev_time"] = a.PrevTime
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		case "closed_log":
			r["closed_log"] = a.ClosedLog
		}
	}
	return r
}

func (a *NgingTask) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingTask) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingTask) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingTask) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
