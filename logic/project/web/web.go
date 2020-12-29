package web

import (
	"fmt"
	_ "github.com/changsongl/microservice/tmpl/statik"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"os"
)

type Initializer struct {
}

func NewInitializer() *Initializer {
	return &Initializer{}
}

func (i *Initializer) Init() error {
	// 生成目录结构
	statikFS, err := fs.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	file, err := statikFS.Open("/web/test.go")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("content: %s\n", content)
	return nil
}
