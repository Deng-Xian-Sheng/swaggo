package api

import (
	"time"

	"swaggo/test/pkg/api/subpackage"
)

type TypeString string

type TypeInterface interface {
	Hello()
}

type SimpleStructure struct {
	Id    float32                    `json:"id" swaggo:"true,dfsdfdsf,19"`
	Name  string                     `json:"name" swaggo:"true,,xus"`
	Age   int                        `json:"age" swaggo:"true,the user age,18"`
	CTime time.Time                  `json:"ctime" swaggo:"true,create time"`
	Sub   subpackage.SimpleStructure `json:"sub" swaggo:"true"`
	I     TypeInterface              `json:"i" swaggo:"true"`
	T     TypeString                 `json:"t"`
	Map   map[string]string          `json:"map", swaggo:",map type"`
}

type SimpleStructureWithAnnotations struct {
	Id   int    `json:"id"`
	Name string `json:"required,omitempty"`
}

type StructureWithSlice struct {
	Id   int
	Name []byte
}

// hello
type StructureWithEmbededStructure struct {
	SimpleStructure
}
type StructureWithEmbededPointer struct {
	*StructureWithSlice
}

type StructureWithAnonymousStructure struct {
	Anonymous []struct {
		Name string
		StructureWithSlice
		Anonymous []struct {
			Name string
		}
	}
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}
