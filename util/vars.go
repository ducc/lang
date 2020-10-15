package util

import (
	"fmt"
	"strings"
)

type Vars struct {
	inner map[string]interface{}
}

func NewVars() *Vars {
	return &Vars{inner: make(map[string]interface{})}
}

func (v *Vars) String() string {
	pairs := make([]string, 0)
	for k, v := range v.inner {
		stringV := fmt.Sprint(v)
		if len(stringV) < 20 {
			pairs = append(pairs, fmt.Sprintf("%s=%s", k, stringV))
		} else {
			pairs = append(pairs, k+"=<present>")
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
}

func (d *Vars) Get(key string) interface{} {
	val := d.inner[key]
	if val == nil {
		panic(fmt.Sprintf("key %s does not exist in data", key))
	}

	return val
}

func (d *Vars) Clone() *Vars {
	inner := make(map[string]interface{})
	for k, v := range d.inner {
		inner[k] = v
	}
	return &Vars{inner: inner}
}

func (d *Vars) Set(key string, value interface{}) *Vars {
	clone := d.Clone()
	clone.inner[key] = value
	return clone
}
