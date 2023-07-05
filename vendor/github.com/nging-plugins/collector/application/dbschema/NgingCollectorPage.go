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

type Slice_NgingCollectorPage []*NgingCollectorPage

func (s Slice_NgingCollectorPage) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorPage) RangeRaw(fn func(m *NgingCollectorPage) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingCollectorPage) GroupBy(keyField string) map[string][]*NgingCollectorPage {
	r := map[string][]*NgingCollectorPage{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingCollectorPage{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingCollectorPage) KeyBy(keyField string) map[string]*NgingCollectorPage {
	r := map[string]*NgingCollectorPage{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingCollectorPage) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingCollectorPage) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingCollectorPage) FromList(data interface{}) Slice_NgingCollectorPage {
	values, ok := data.([]*NgingCollectorPage)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingCollectorPage{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingCollectorPage(ctx echo.Context) *NgingCollectorPage {
	m := &NgingCollectorPage{}
	m.SetContext(ctx)
	return m
}

// NgingCollectorPage 采集页面
type NgingCollectorPage struct {
	base    factory.Base
	objects []*NgingCollectorPage

	Id            uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"" json:"id" xml:"id"`
	ParentId      uint   `db:"parent_id" bson:"parent_id" comment:"父级规则" json:"parent_id" xml:"parent_id"`
	RootId        uint   `db:"root_id" bson:"root_id" comment:"根页面ID" json:"root_id" xml:"root_id"`
	HasChild      string `db:"has_child" bson:"has_child" comment:"是否有子级" json:"has_child" xml:"has_child"`
	Uid           uint   `db:"uid" bson:"uid" comment:"用户ID" json:"uid" xml:"uid"`
	GroupId       uint   `db:"group_id" bson:"group_id" comment:"规则组" json:"group_id" xml:"group_id"`
	Name          string `db:"name" bson:"name" comment:"规则名" json:"name" xml:"name"`
	Description   string `db:"description" bson:"description" comment:"说明" json:"description" xml:"description"`
	EnterUrl      string `db:"enter_url" bson:"enter_url" comment:"入口网址模板(网址一行一个)" json:"enter_url" xml:"enter_url"`
	Sort          int    `db:"sort" bson:"sort" comment:"排序" json:"sort" xml:"sort"`
	Created       uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Browser       string `db:"browser" bson:"browser" comment:"浏览器" json:"browser" xml:"browser"`
	Type          string `db:"type" bson:"type" comment:"页面类型" json:"type" xml:"type"`
	ScopeRule     string `db:"scope_rule" bson:"scope_rule" comment:"页面区域规则" json:"scope_rule" xml:"scope_rule"`
	DuplicateRule string `db:"duplicate_rule" bson:"duplicate_rule" comment:"去重规则(url-判断网址;rule-判断规则是否改过;content-判断网页内容是否改过;none-不去重)" json:"duplicate_rule" xml:"duplicate_rule"`
	ContentType   string `db:"content_type" bson:"content_type" comment:"内容类型" json:"content_type" xml:"content_type"`
	Charset       string `db:"charset" bson:"charset" comment:"字符集" json:"charset" xml:"charset"`
	Timeout       uint   `db:"timeout" bson:"timeout" comment:"超时时间(秒)" json:"timeout" xml:"timeout"`
	Waits         string `db:"waits" bson:"waits" comment:"等待时间范围(秒),例如2-8" json:"waits" xml:"waits"`
	Proxy         string `db:"proxy" bson:"proxy" comment:"代理地址" json:"proxy" xml:"proxy"`
}

// - base function

func (a *NgingCollectorPage) Trans() factory.Transactioner {
	return a.base.Trans()
}

func (a *NgingCollectorPage) Use(trans factory.Transactioner) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingCollectorPage) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingCollectorPage) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingCollectorPage) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingCollectorPage) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingCollectorPage) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingCollectorPage) ConnID() int {
	return a.base.ConnID()
}

func (a *NgingCollectorPage) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingCollectorPage) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingCollectorPage) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingCollectorPage) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

func (a *NgingCollectorPage) New(structName string, connID ...int) factory.Model {
	return a.base.New(structName, connID...)
}

func (a *NgingCollectorPage) Base_() factory.Baser {
	return &a.base
}

// - current function

func (a *NgingCollectorPage) Objects() []*NgingCollectorPage {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingCollectorPage) XObjects() Slice_NgingCollectorPage {
	return Slice_NgingCollectorPage(a.Objects())
}

func (a *NgingCollectorPage) NewObjects() factory.Ranger {
	return &Slice_NgingCollectorPage{}
}

func (a *NgingCollectorPage) InitObjects() *[]*NgingCollectorPage {
	a.objects = []*NgingCollectorPage{}
	return &a.objects
}

func (a *NgingCollectorPage) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingCollectorPage) Short_() string {
	return "nging_collector_page"
}

func (a *NgingCollectorPage) Struct_() string {
	return "NgingCollectorPage"
}

func (a *NgingCollectorPage) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingCollectorPage) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.SetConnID(source.ConnID())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingCollectorPage) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
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

func (a *NgingCollectorPage) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorPage:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorPage(*v))
		case []*NgingCollectorPage:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorPage(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorPage) GroupBy(keyField string, inputRows ...[]*NgingCollectorPage) map[string][]*NgingCollectorPage {
	var rows Slice_NgingCollectorPage
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorPage(inputRows[0])
	} else {
		rows = Slice_NgingCollectorPage(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingCollectorPage) KeyBy(keyField string, inputRows ...[]*NgingCollectorPage) map[string]*NgingCollectorPage {
	var rows Slice_NgingCollectorPage
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorPage(inputRows[0])
	} else {
		rows = Slice_NgingCollectorPage(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingCollectorPage) AsKV(keyField string, valueField string, inputRows ...[]*NgingCollectorPage) param.Store {
	var rows Slice_NgingCollectorPage
	if len(inputRows) > 0 {
		rows = Slice_NgingCollectorPage(inputRows[0])
	} else {
		rows = Slice_NgingCollectorPage(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingCollectorPage) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
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
		case *[]*NgingCollectorPage:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorPage(*v))
		case []*NgingCollectorPage:
			err = DBI.FireReaded(a, queryParam, Slice_NgingCollectorPage(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingCollectorPage) Insert() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if len(a.Type) == 0 {
		a.Type = "content"
	}
	if len(a.DuplicateRule) == 0 {
		a.DuplicateRule = "none"
	}
	if len(a.ContentType) == 0 {
		a.ContentType = "html"
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

func (a *NgingCollectorPage) Update(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if len(a.Type) == 0 {
		a.Type = "content"
	}
	if len(a.DuplicateRule) == 0 {
		a.DuplicateRule = "none"
	}
	if len(a.ContentType) == 0 {
		a.ContentType = "html"
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

func (a *NgingCollectorPage) Updatex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if len(a.Type) == 0 {
		a.Type = "content"
	}
	if len(a.DuplicateRule) == 0 {
		a.DuplicateRule = "none"
	}
	if len(a.ContentType) == 0 {
		a.ContentType = "html"
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

func (a *NgingCollectorPage) UpdateByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {

	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if len(a.Type) == 0 {
		a.Type = "content"
	}
	if len(a.DuplicateRule) == 0 {
		a.DuplicateRule = "none"
	}
	if len(a.ContentType) == 0 {
		a.ContentType = "html"
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

func (a *NgingCollectorPage) UpdatexByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {

	if len(a.HasChild) == 0 {
		a.HasChild = "N"
	}
	if len(a.Type) == 0 {
		a.Type = "content"
	}
	if len(a.DuplicateRule) == 0 {
		a.DuplicateRule = "none"
	}
	if len(a.ContentType) == 0 {
		a.ContentType = "html"
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

func (a *NgingCollectorPage) UpdateField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.UpdateFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorPage) UpdatexField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (affected int64, err error) {
	return a.UpdatexFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingCollectorPage) UpdateFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["has_child"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["has_child"] = "N"
		}
	}
	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "content"
		}
	}
	if val, ok := kvset["duplicate_rule"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["duplicate_rule"] = "none"
		}
	}
	if val, ok := kvset["content_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["content_type"] = "html"
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

func (a *NgingCollectorPage) UpdatexFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (affected int64, err error) {

	if val, ok := kvset["has_child"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["has_child"] = "N"
		}
	}
	if val, ok := kvset["type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["type"] = "content"
		}
	}
	if val, ok := kvset["duplicate_rule"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["duplicate_rule"] = "none"
		}
	}
	if val, ok := kvset["content_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["content_type"] = "html"
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

func (a *NgingCollectorPage) UpdateValues(mw func(db.Result) db.Result, keysValues *db.KeysValues, args ...interface{}) (err error) {
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

func (a *NgingCollectorPage) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.HasChild) == 0 {
			a.HasChild = "N"
		}
		if len(a.Type) == 0 {
			a.Type = "content"
		}
		if len(a.DuplicateRule) == 0 {
			a.DuplicateRule = "none"
		}
		if len(a.ContentType) == 0 {
			a.ContentType = "html"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.HasChild) == 0 {
			a.HasChild = "N"
		}
		if len(a.Type) == 0 {
			a.Type = "content"
		}
		if len(a.DuplicateRule) == 0 {
			a.DuplicateRule = "none"
		}
		if len(a.ContentType) == 0 {
			a.ContentType = "html"
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

func (a *NgingCollectorPage) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

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

func (a *NgingCollectorPage) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

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

func (a *NgingCollectorPage) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingCollectorPage) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingCollectorPage) Reset() *NgingCollectorPage {
	a.Id = 0
	a.ParentId = 0
	a.RootId = 0
	a.HasChild = ``
	a.Uid = 0
	a.GroupId = 0
	a.Name = ``
	a.Description = ``
	a.EnterUrl = ``
	a.Sort = 0
	a.Created = 0
	a.Browser = ``
	a.Type = ``
	a.ScopeRule = ``
	a.DuplicateRule = ``
	a.ContentType = ``
	a.Charset = ``
	a.Timeout = 0
	a.Waits = ``
	a.Proxy = ``
	return a
}

func (a *NgingCollectorPage) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["ParentId"] = a.ParentId
		r["RootId"] = a.RootId
		r["HasChild"] = a.HasChild
		r["Uid"] = a.Uid
		r["GroupId"] = a.GroupId
		r["Name"] = a.Name
		r["Description"] = a.Description
		r["EnterUrl"] = a.EnterUrl
		r["Sort"] = a.Sort
		r["Created"] = a.Created
		r["Browser"] = a.Browser
		r["Type"] = a.Type
		r["ScopeRule"] = a.ScopeRule
		r["DuplicateRule"] = a.DuplicateRule
		r["ContentType"] = a.ContentType
		r["Charset"] = a.Charset
		r["Timeout"] = a.Timeout
		r["Waits"] = a.Waits
		r["Proxy"] = a.Proxy
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "ParentId":
			r["ParentId"] = a.ParentId
		case "RootId":
			r["RootId"] = a.RootId
		case "HasChild":
			r["HasChild"] = a.HasChild
		case "Uid":
			r["Uid"] = a.Uid
		case "GroupId":
			r["GroupId"] = a.GroupId
		case "Name":
			r["Name"] = a.Name
		case "Description":
			r["Description"] = a.Description
		case "EnterUrl":
			r["EnterUrl"] = a.EnterUrl
		case "Sort":
			r["Sort"] = a.Sort
		case "Created":
			r["Created"] = a.Created
		case "Browser":
			r["Browser"] = a.Browser
		case "Type":
			r["Type"] = a.Type
		case "ScopeRule":
			r["ScopeRule"] = a.ScopeRule
		case "DuplicateRule":
			r["DuplicateRule"] = a.DuplicateRule
		case "ContentType":
			r["ContentType"] = a.ContentType
		case "Charset":
			r["Charset"] = a.Charset
		case "Timeout":
			r["Timeout"] = a.Timeout
		case "Waits":
			r["Waits"] = a.Waits
		case "Proxy":
			r["Proxy"] = a.Proxy
		}
	}
	return r
}

func (a *NgingCollectorPage) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "parent_id":
			a.ParentId = param.AsUint(value)
		case "root_id":
			a.RootId = param.AsUint(value)
		case "has_child":
			a.HasChild = param.AsString(value)
		case "uid":
			a.Uid = param.AsUint(value)
		case "group_id":
			a.GroupId = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "description":
			a.Description = param.AsString(value)
		case "enter_url":
			a.EnterUrl = param.AsString(value)
		case "sort":
			a.Sort = param.AsInt(value)
		case "created":
			a.Created = param.AsUint(value)
		case "browser":
			a.Browser = param.AsString(value)
		case "type":
			a.Type = param.AsString(value)
		case "scope_rule":
			a.ScopeRule = param.AsString(value)
		case "duplicate_rule":
			a.DuplicateRule = param.AsString(value)
		case "content_type":
			a.ContentType = param.AsString(value)
		case "charset":
			a.Charset = param.AsString(value)
		case "timeout":
			a.Timeout = param.AsUint(value)
		case "waits":
			a.Waits = param.AsString(value)
		case "proxy":
			a.Proxy = param.AsString(value)
		}
	}
}

func (a *NgingCollectorPage) Set(key interface{}, value ...interface{}) {
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
		case "ParentId":
			a.ParentId = param.AsUint(vv)
		case "RootId":
			a.RootId = param.AsUint(vv)
		case "HasChild":
			a.HasChild = param.AsString(vv)
		case "Uid":
			a.Uid = param.AsUint(vv)
		case "GroupId":
			a.GroupId = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "Description":
			a.Description = param.AsString(vv)
		case "EnterUrl":
			a.EnterUrl = param.AsString(vv)
		case "Sort":
			a.Sort = param.AsInt(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Browser":
			a.Browser = param.AsString(vv)
		case "Type":
			a.Type = param.AsString(vv)
		case "ScopeRule":
			a.ScopeRule = param.AsString(vv)
		case "DuplicateRule":
			a.DuplicateRule = param.AsString(vv)
		case "ContentType":
			a.ContentType = param.AsString(vv)
		case "Charset":
			a.Charset = param.AsString(vv)
		case "Timeout":
			a.Timeout = param.AsUint(vv)
		case "Waits":
			a.Waits = param.AsString(vv)
		case "Proxy":
			a.Proxy = param.AsString(vv)
		}
	}
}

func (a *NgingCollectorPage) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["parent_id"] = a.ParentId
		r["root_id"] = a.RootId
		r["has_child"] = a.HasChild
		r["uid"] = a.Uid
		r["group_id"] = a.GroupId
		r["name"] = a.Name
		r["description"] = a.Description
		r["enter_url"] = a.EnterUrl
		r["sort"] = a.Sort
		r["created"] = a.Created
		r["browser"] = a.Browser
		r["type"] = a.Type
		r["scope_rule"] = a.ScopeRule
		r["duplicate_rule"] = a.DuplicateRule
		r["content_type"] = a.ContentType
		r["charset"] = a.Charset
		r["timeout"] = a.Timeout
		r["waits"] = a.Waits
		r["proxy"] = a.Proxy
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "parent_id":
			r["parent_id"] = a.ParentId
		case "root_id":
			r["root_id"] = a.RootId
		case "has_child":
			r["has_child"] = a.HasChild
		case "uid":
			r["uid"] = a.Uid
		case "group_id":
			r["group_id"] = a.GroupId
		case "name":
			r["name"] = a.Name
		case "description":
			r["description"] = a.Description
		case "enter_url":
			r["enter_url"] = a.EnterUrl
		case "sort":
			r["sort"] = a.Sort
		case "created":
			r["created"] = a.Created
		case "browser":
			r["browser"] = a.Browser
		case "type":
			r["type"] = a.Type
		case "scope_rule":
			r["scope_rule"] = a.ScopeRule
		case "duplicate_rule":
			r["duplicate_rule"] = a.DuplicateRule
		case "content_type":
			r["content_type"] = a.ContentType
		case "charset":
			r["charset"] = a.Charset
		case "timeout":
			r["timeout"] = a.Timeout
		case "waits":
			r["waits"] = a.Waits
		case "proxy":
			r["proxy"] = a.Proxy
		}
	}
	return r
}

func (a *NgingCollectorPage) ListPage(cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, nil, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingCollectorPage) ListPageAs(recv interface{}, cond *db.Compounds, sorts ...interface{}) error {
	_, err := pagination.NewLister(a, recv, func(r db.Result) db.Result {
		return r.OrderBy(sorts...)
	}, cond.And()).Paging(a.Context())
	return err
}

func (a *NgingCollectorPage) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingCollectorPage) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
