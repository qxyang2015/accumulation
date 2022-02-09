package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
)

type AddContextVisitor struct {
	pkgContext string // 如果有引入context包情况下，context包的别名
}

func (vi *AddContextVisitor) Visit(node ast.Node) ast.Visitor {

	switch node.(type) {
	case *ast.File:
		file := node.(*ast.File)
		// 没有任何import
		if len(file.Imports) == 0 {
			vi.addImportWithoutAnyImport(file)
		}
	case *ast.GenDecl:
		genDecl := node.(*ast.GenDecl)
		// 有import
		if genDecl.Tok == token.IMPORT {
			vi.addImport(genDecl)
		}
	case *ast.InterfaceType:
		// 遍历所有的接口类型
		interfaceType := node.(*ast.InterfaceType)
		vi.addContextAndError(interfaceType)
		return nil
	}
	return vi
}

func (vi *AddContextVisitor) addImport(genDecl *ast.GenDecl) {
	hasImportedContext := false
	for _, value := range genDecl.Specs {
		// 如果已经包含"context"
		importSpec := value.(*ast.ImportSpec)
		if importSpec.Path.Value == strconv.Quote("context") {
			hasImportedContext = true
			// 有别名的情况下记录别名
			if importSpec.Name != nil {
				vi.pkgContext = importSpec.Name.Name
			}
		}
	}
	if hasImportedContext {
		return
	}
	if !hasImportedContext {
		genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote("context"),
			},
		})
	}
}

// 没有import情况下 引入context
func (vi *AddContextVisitor) addImportWithoutAnyImport(file *ast.File) {
	genDecl := &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote("context"),
				},
			},
		},
	}
	list := []ast.Decl{genDecl}
	file.Decls = append(list, file.Decls...)
}

// 为接口方法添加参数
func (vi *AddContextVisitor) addContextAndError(interfaceType *ast.InterfaceType) {
	// 接口方法不为空是，遍历接口方法
	if interfaceType.Methods != nil || interfaceType.Methods.List != nil {
		for _, v := range interfaceType.Methods.List {
			ft := v.Type.(*ast.FuncType)
			hasContext := false
			hasError := false
			// 判断参数中是否包含context.Context类型
			for _, value := range ft.Params.List {
				if expr, ok := value.Type.(*ast.SelectorExpr); ok {
					if ident, ok := expr.X.(*ast.Ident); ok {
						if ident.Name == "context" {
							hasContext = true
						}
					}
				}
			}

			if ft.Results != nil && ft.Results.List != nil {
				// 判断返回参数中是否包含error类型
				for i, value := range ft.Results.List {
					if ident, ok := value.Type.(*ast.Ident); ok {
						if ident.Name == "error" {
							ft.Results.List[i].Names = []*ast.Ident{
								ast.NewIdent("err"),
							}
							hasError = true
						}
					}
				}
			}

			if !hasError {
				errField := &ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("err"),
					},
					Type: ast.NewIdent("error"),
				}
				if ft.Results == nil {
					ft.Results = &ast.FieldList{}
				}
				ft.Results.List = append(ft.Results.List, errField)
			}

			// 为没有context参数的方法添加context参数
			if !hasContext {
				x := "context"
				if vi.pkgContext != "" {
					x = vi.pkgContext
				}
				ctxField := &ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("ctx"),
					},
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent(x),
						Sel: ast.NewIdent("Context"),
					},
				}
				list := []*ast.Field{
					ctxField,
				}
				ft.Params.List = append(list, ft.Params.List...)
			}
		}
	}
}

func main() {
	fSet := token.NewFileSet()

	path, _ := filepath.Abs("./demo/ast_demo/demo.go")
	f, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	v := &AddContextVisitor{}
	ast.Walk(v, f)

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fSet, f)
	if err != nil {
		log.Fatal(err)
	}
	// 输出Go代码
	b, _ := format.Source(buffer.Bytes()) // fmt
	//b, _ = imports.Process("", b, nil)    // imports
	fmt.Printf("%s\n", b)
}
