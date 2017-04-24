package goutils

type ModelInterface struct {
	//todo
}

const (
	O_STRING   = 1
	O_NUMERIC  = 2
	O_FLOAT    = 3
	O_BOOL     = 4
	O_DATETIME = 5
	O_OBJECT   = 6
)

var GormModel = []string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
}

type Model struct {
	Fields  []string
	Configs map[string]*Config
}

type Config struct {
	Name       string
	Type       int32
	Length     int32
	Default    interface{}
	Required   bool
	Memo       string
	ObjectName string
	JsonName   string
	FormName   string
	OrmName    string
}

func (this *Model) Add(element string) *Model {
	for k, v := range this.Fields {
		if v == element {
			this.Fields = append(this.Fields[:k], this.Fields[k+1:]...)
			break
		}
	}
	this.Fields = append(this.Fields, element)
	return this
}

func (this *Model) SetConfig(name string, conf *Config) *Model {
	this.Configs[name] = conf
	return this
}

func (this *Model) GetConfig(name string) (ok bool, conf *Config) {
	conf, ok = this.Configs[name]
	return
}
