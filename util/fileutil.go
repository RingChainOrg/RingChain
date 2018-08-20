package util

import (
	"io/ioutil"
	"os"
	"io"
	"path/filepath"
	"log"
	"strings"
	"path"
)

func ReadFile(filePath string)  ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func ReadJson(jsonPath string) ()  {

}

func FileOrDirectoryExists(path string) (bool, error) {
	//如果返回的错误为nil,说明文件或文件夹存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	//如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	//如果返回的错误为其它类型,则不确定是否在存在
	return false, err
}

//使用之前用FileOrDirectoryExists先判断是否存在.
func IsDirectory(path string) (bool, error)   {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir(), nil
	}else {

		return false, err;
	}
}

func AppendFile(filepath string, data []byte) error  {
	return ioutil.WriteFile(filepath, data, os.ModeAppend);
}

func OverwriteFile(filepath string, data []byte) (written int, err error) {
	dst, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return dst.WriteAt(data,0);
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()//如果目标文件后面没有使用defer dst.Close(),那么一旦创建目标文件调试失败，它将直接返回错误，那么，会导致源文件一直保持打开状态，这样资源就得不到释放。
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//func CopyFile(dstName, srcName string, n int64) (written int64, err error) {
//	src, err := os.Open(srcName)
//	if err != nil {
//		return
//	}
//	defer src.Close()
//	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		return
//	}
//	defer dst.Close()
//	return io.CopyN(dst, src, n) //
//}

func RemoveFile(path string) error  {
	return os.Remove(path)
}

/*
  递归创建目录os.MkdirAll(path string, perm FileMode) error
  path  目录名及子目录
  perm  目录权限位
  error 如果成功返回nil，如果目录已经存在默认什么都不做
*/
func CreateAllDir(dir string) error{
	err := os.MkdirAll(dir, 0777);
	return err;
}


func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func JoinDir(path1 string, path2 string) string  {
	return path.Join(path1, path2);
}

func GetDirectory(dir string) (string, error){
	dir, err := filepath.Abs(filepath.Dir(dir))
	if err != nil {
		return "", err;
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
