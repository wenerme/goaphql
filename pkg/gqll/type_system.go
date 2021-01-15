package gqll

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type TypeSystem struct {
	definitions []Definition
	nameMap     map[string]Definition
	typeMap     map[string][]Definition
}

func (self *TypeSystem) GetNamedDefinition(name string) Definition {
	return self.nameMap[name]
}
func (self *TypeSystem) GetTypedDefinitions(typ string) []Definition {
	return self.typeMap[typ]
}
func (self *TypeSystem) GetDefinitions() []Definition {
	return self.definitions
}

func NewTypeSystem(defs []Definition) (*TypeSystem, error) {
	ts := &TypeSystem{
		nameMap: make(map[string]Definition),
		typeMap: make(map[string][]Definition),
	}
	namesMap := make(map[string][]Definition)

	for _, v := range defs {
		// name index
		// for extension
		//if vn, ok := v.(HasExtendTypeName); ok {
		//	name := vn.GetExtendTypeName()
		//	namesMap[name] = append(namesMap[name], v)
		//} else
		if vn, ok := v.(HasName); ok {
			name := vn.GetName()
			namesMap[name] = append(namesMap[name], v)
		}

		// type index
		typeName := TypeOf(v).Name()
		ts.typeMap[typeName] = append(ts.typeMap[typeName], v)
		ts.definitions = append(ts.definitions, v)
	}

	// aggregate
	for k, v := range namesMap {
		if k == "" || len(v) == 0 {
			continue
		}
		if len(v) == 1 {
			ts.nameMap[k] = v[0]
		} else {
			// Doto merge
			//ts.nameMap[k], err = MergeDefinitions(v)
			//if err != nil {
			//	return nil, err
			//}

			var target Definition
			for _, d := range v {
				if TypeOf(d).IsTypeDefinition() {
					if target != nil {
						return nil, errors.Errorf("multi definition for %s", NameOf(d))
					}
					target = d
				}
			}
			ts.nameMap[k] = target
		}
	}

	return ts, nil
}
func MergeDefinitions(definitions []Definition) (Definition, error) {
	var target Definition
	var targetIdx = 0
	for i, v := range definitions {
		if strings.HasSuffix(TypeOf(v).Name(), "TypeDefinition") {
			target = v
			targetIdx = i
			break
		}
	}
	if target == nil {
		for _, v := range definitions {
			logrus.WithFields(logrus.Fields{
				"Name": NameOf(v),
				"Type": TypeOf(definitions[0]).Name(),
			}).Warn("Error report")
		}
		return nil, errors.Errorf("gqll: merge target not found")
	}
	for i, v := range definitions {
		if i == targetIdx {
			continue
		}
		if err := MergeDefinition(target, v); err != nil {
			return nil, err
		}
	}

	return definitions[0], nil
}

func MergeDefinition(target Definition, ext Definition) error {

	if feat, ok := target.(HasFieldDefinitions); ok {
		origin := feat.GetFieldDefinitions()
		extra := ext.(HasFieldDefinitions).GetFieldDefinitions()

		feat.SetFieldDefinitions(append(origin, extra...))
	}
	if feat, ok := target.(HasInterfaces); ok {
		origin := feat.GetInterfaces()
		extra := ext.(HasInterfaces).GetInterfaces()

		feat.SetInterfaces(append(origin, extra...))
	}

	if feat, ok := target.(HasDirectives); ok {
		origin := feat.GetDirectives()
		extra := ext.(HasDirectives).GetDirectives()

		feat.SetDirectives(append(origin, extra...))
	}

	return nil
}
