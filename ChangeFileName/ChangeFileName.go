package ChangeFileName

import (
"flag"
"fmt"
"os"
"path/filepath"
)

//renamedirfiles -d "E:\shared\图片素材\ps123_20121120_01\背景图片打包下载" -p "bg%d"

func ChangeFileName() {
	// 解析命令行参数
	var dir string
	flag.StringVar(&dir, "d", "", "directory path")
	var pattern string
	flag.StringVar(&pattern, "p", "", "name pattern, eg. newname%d")
	flag.Parse()
	if dir == "" || pattern == "" {
		flag.Usage()
		return
	}
	// 遍历文件夹，获取文件路径
	paths := make([]string, 0)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	// 遍历文件路径，修改文件名
	for i, path := range paths {
		newPath := filepath.Join(filepath.Dir(path), fmt.Sprintf(pattern, i)+filepath.Ext(path))
		os.Rename(path, newPath)
	}
}