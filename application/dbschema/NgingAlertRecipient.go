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

type Slice_NgingAlertRecipient []*NgingAlertRecipient

func (s Slice_NgingAlertRecipient) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingAlertRecipient) RangeRaw(fn func(m *NgingAlertRecipient) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingAlertRecipient) GroupBy(keyField string) map[string][]*NgingAlertRecipient {
	r := map[string][]*NgingAlertRecipient{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingAlertRecipient{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingAlertRecipient) KeyBy(keyField string) map[string]*NgingAlertRecipient {
	r := map[string]*NgingAlertRecipient{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingAlertRecipient) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingAlertRecipient) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingAlertRecipient) FromList(data interface{}) Slice_NgingAlertRecipient {
	values, ok := data.([]*NgingAlertRecipient)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingAlertRecipient{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingAlertRecipient(ctx echo.Context) *NgingAlertRecipient {
	m := &NgingAlertRecipient{}
	m.SetContext(ctx)
	return m
}

// NgingAlertRecipient 报警收信人
type NgingAlertRecipient struct {
	base    factory.Base
	objects []*NgingAlertRecipient

	Id          uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name        string `db:"name" bson:"name" comment:"名称" json:"name" xml:"name"`
	Account     string `db:"account" bson:"account" comment:"账号" json:"account" xml:"account"`
	Extra       string `db:"extra" bson:"extra" comment:"扩展信息(JSON)" json:"extra" xml:"extra"`
	Type        string `db:"type" bson:"type" comment:"类型" json:"type" xml:"type"`
	Platform    string `db:"platform" bson:"platform" comment:"平台(dingding-钉钉;workwx-企业微信)" json:"platform" xml:"platform"`
	Description string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	Disabled    string `db:"disabled" bson:"disabled" comment:"是否(Y/N)禁用" json:"disabled" xml:"disabled"`
	Created     uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated     uint   `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (a *NgingAlertRecipient) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingAlertRecipient) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingAlertRecipient) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingAlertRecipient) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingAlertRecipient) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingAlertRecipient) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingAlertRecipient) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingAlertRecipient) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingAlertRecipient) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingAlertRecipient) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingAlertRecipient) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingAlertRecipient) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingAlertRecipient) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingAlertRecipient) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingAlertRecipient) Objects() []*NgingAlertRecipient {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingAlertRecipient) XObjects() Slice_NgingAlertRecipient {
	return Slice_NgingAlertRecipient(a.Objects())
}

func (a *NgingAlertRecipient) NewObjects() factory.Ranger {
	return &Slice_NgingAlertRecipient{}
}

func (a *NgingAlertRecipient) InitObjects() *[]*NgingAlertRecipient {
	a.objects = []*NgingAlertRecipient{}
	return &a.objects
}

func (a *NgingAlertRecipient) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingAlertRecipient) Short_() string {
	return "nging_alert_recipient"
}

func (a *NgingAlertRecipient) Struct_() string {
	return "NgingAlertRecipient"
}

func (a *NgingAlertRecipient) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingAlertRecipient) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingAlertRecipient) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingAlertRecipient) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingAlertRecipient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertRecipient(*v))
		case []*NgingAlertRecipient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertRecipient(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingAlertRecipient) GroupBy(keyField string, inputRows ...[]*NgingAlertRecipient) map[string][]*NgingAlertRecipient {
	var rows Slice_NgingAlertRecipient
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertRecipient(inputRows[0])
	} else {
		rows = Slice_NgingAlertRecipient(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingAlertRecipient) KeyBy(keyField string, inputRows ...[]*NgingAlertRecipient) map[string]*NgingAlertRecipient {
	var rows Slice_NgingAlertRecipient
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertRecipient(inputRows[0])
	} else {
		rows = Slice_NgingAlertRecipient(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingAlertRecipient) AsKV(keyField string, valueField string, inputRows ...[]*NgingAlertRecipient) param.Store {
	var rows Slice_NgingAlertRecipient
	if len(inputRows) > 0 {
		rows = Slice_NgingAlertRecipient(inputRows[0])
	} else {
		rows = Slice_NgingAlertRecipient(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingAlertRecipient) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingAlertRecipient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertRecipient(*v))
		case []*NgingAlertRecipient:
			err = DBI.FireReaded(a, queryParam, Slice_NgingAlertRecipient(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingAlertRecipient) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.Type) == 0 {
		a.Type = "email"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingAlertRecipient) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Type) == 0 {
		a.Type = "email"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingAlertRecipient) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Type) == 0 {
		a.Type = "email"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingAlertRecipient) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Type) == 0 {
		a.Type = "email"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingAlertRecipient) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {
	a.Updated = uint(time.Now().Unix())
	if len(a.Type) == 0 {
		a.Type = "email"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
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

func (a *NgingAlertRecipient) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingAlertRecipient) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingAlertRecipient) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "email"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
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

func (a *NgingAlertRecipient) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "email"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
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

func (a *NgingAlertRecipient) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
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

func (a *NgingAlertRecipient) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		a.Updated = uint(time.Now().Unix())
		if len(a.Type) == 0 {
			a.Type = "email"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.Type) == 0 {
			a.Type = "email"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
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

func (a *NgingAlertRecipient) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingAlertRecipient) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

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

func (a *NgingAlertRecipient) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingAlertRecipient) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingAlertRecipient) Reset() *NgingAlertRecipient {
	a.Id = 0
	a.Name = ``
	a.Account = ``
	a.Extra = ``
	a.Type = ``
	a.Platform = ``
	a.Description = ``
	a.Disabled = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *NgingAlertRecipient) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Name"] = a.Name
		r["Account"] = a.Account
		r["Extra"] = a.Extra
		r["Type"] = a.Type
		r["Platform"] = a.Platform
		r["Description"] = a.Description
		r["Disabled"] = a.Disabled
		r["Created"] = a.Created
		r["Updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Name":
			r["Name"] = a.Name
		case "Account":
			r["Account"] = a.Account
		case "Extra":
			r["Extra"] = a.Extra
		case "Type":
			r["Type"] = a.Type
		case "Platform":
			r["Platform"] = a.Platform
		case "Description":
			r["Description"] = a.Description
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "Created":
			r["Created"] = a.Created
		case "Updated":
			r["Updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingAlertRecipient) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "account":
			a.Account = param.AsString(value)
		case "extra":
			a.Extra = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "platform":
			a.Platform = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsUint(value)
		}
	}
}

func (a *NgingAlertRecipient) Set(key interface{}, value ...interface{}) {
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
		case "Name":
			a.Name = param.AsString(vv)
		case "Account":
			a.Account = param.AsString(vv)
		case "Extra":
			a.Extra = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "Platform":
			a.Platform = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsUint(vv)
		}
	}
}

func (a *NgingAlertRecipient) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["name"] = a.Name
		r["account"] = a.Account
		r["extra"] = a.Extra
		r["type"] = a.Type
		r["platform"] = a.Platform
		r["description"] = a.Description
		r["disabled"] = a.Disabled
		r["created"] = a.Created
		r["updated"] = a.Updated
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "name":
			r["name"] = a.Name
		case "account":
			r["account"] = a.Account
		case "extra":
			r["extra"] = a.Extra
		case "type":
			r["type"] = a.Type
		case "platform":
			r["platform"] = a.Platform
		case "description":
			r["description"] = a.Description
		case "disabled":
			r["disabled"] = a.Disabled
		case "created":
			r["created"] = a.Created
		case "updated":
			r["updated"] = a.Updated
		}
	}
	return r
}

func (a *NgingAlertRecipient) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingAlertRecipient) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingAlertRecipient) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingAlertRecipient) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
