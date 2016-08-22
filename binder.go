package tmpq

import (
	"regexp"
	"strconv"
)

type (
	// Binder provides interfaces for parameter binding by name.
	Binder interface {
		Bind(name string, v interface{})
		QueryArgs() (query string, args []interface{})
	}

	binderImpl struct {
		rawQuery          string
		parameterIndices  map[string]int
		parameters        map[string]interface{}
		nameToPlaceHolder map[string]string
	}
)

const (
	parameterGroupName     = "name"
	parameterNameRegexText = `\$\{(?P<` + parameterGroupName + `>[^\$\{\}]*?)\}`
)

var (
	parameterNameRegex = regexp.MustCompile(parameterNameRegexText)
)

// NewBinder returns a new Binder.
func NewBinder(query string) Binder {
	result := &binderImpl{
		rawQuery:          query,
		parameterIndices:  make(map[string]int, 0),
		parameters:        make(map[string]interface{}, 0),
		nameToPlaceHolder: make(map[string]string, 0),
	}
	result.parseQuery()
	return result
}

// Bind binds the parameter to the name.
func (b *binderImpl) Bind(name string, v interface{}) {
	b.parameters[name] = v
}

func (b *binderImpl) QueryArgs() (query string, args []interface{}) {
	query = parameterNameRegex.ReplaceAllStringFunc(b.rawQuery, func(placeHolder string) string {
		return "$" + strconv.Itoa(b.parameterIndices[placeHolder]+1)
	})

	args = make([]interface{}, len(b.parameters))

	for name, placeHolder := range b.nameToPlaceHolder {
		index := b.parameterIndices[placeHolder]
		args[index] = b.parameters[name]
	}

	return query, args
}

func (b *binderImpl) parseQuery() {
	query := b.rawQuery
	allMatches := parameterNameRegex.FindAllStringSubmatch(query, -1)
	if len(allMatches) <= 0 {
		return
	}

	nameIndex := b.nameIndex()
	if nameIndex < 0 {
		return
	}

	argIndex := 0
	for _, m := range allMatches {
		placeHolder := m[0]
		name := m[nameIndex]

		if _, ok := b.nameToPlaceHolder[name]; ok {
			continue
		}

		b.nameToPlaceHolder[name] = placeHolder
		b.parameters[name] = nil
		b.parameterIndices[placeHolder] = argIndex
		argIndex++
	}

}

func (b *binderImpl) nameIndex() int {
	for index, name := range parameterNameRegex.SubexpNames() {
		if name == parameterGroupName {
			return index
		}
	}
	return -1
}
