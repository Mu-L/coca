package domain

import (
	"github.com/phodal/coca/core/infrastructure"
	"strings"
)

type JMethod struct {
	Name              string
	Type              string
	StartLine         int
	StartLinePosition int
	StopLine          int
	StopLinePosition  int
	Parameters        []JParameter
	MethodCalls       []JMethodCall
	Override          bool
	Annotations       []Annotation
	IsConstructor     bool
	IsReturnNull      bool
	Modifiers         []string
	Creators          []JClassNode
}

func NewJMethod() JMethod {
	return *&JMethod{
		Name:              "",
		Type:              "",
		Annotations:       nil,
		StartLine:         0,
		StartLinePosition: 0,
		StopLine:          0,
		StopLinePosition:  0,
		Parameters:        nil,
		MethodCalls:       nil,
		IsConstructor:     false,
	}
}

func (m *JMethod) IsStatic() bool {
	return infrastructure.StringArrayContains(m.Modifiers, "static")
}

func (m *JMethod) IsGetterSetter() bool {
	return strings.HasPrefix(m.Name, "set") || strings.HasPrefix(m.Name, "get")
}

type JMethodInfo struct {
	Name       string
	Type       string
	Parameters []JParameter
	Length     string
}