package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// Generate generate template project
func Generate(dstPath string, style string, pkgName string) error {
	var (
		err               error
		filePath, dirPath string
		file              *os.File
	)

	tmpls, exist := tmplFiles[style]
	if !exist {
		return fmt.Errorf("not exist style(%s) template project", style)
	}

	ctx := &tmplCtx{
		PkgName: pkgName,
	}

	tmpl := template.New("mgo")
	for fileName, assetInfo := range tmpls {
		filePath = filepath.Join(dstPath, fileName)
		dirPath = filepath.Dir(filePath)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}

		perm := os.FileMode(0644)
		if assetInfo.isExec {
			perm = 0755
		}
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			break
		}

		if assetInfo.isTmpl {
			t, err := tmpl.Parse(MustAssetString(assetInfo.name))
			if err != nil {
				break
			}
			if err = t.Execute(file, ctx); err != nil {
				break
			}
		} else {
			if _, err = file.Write(MustAsset(assetInfo.name)); err != nil {
				break
			}
		}
	}
	return err
}
