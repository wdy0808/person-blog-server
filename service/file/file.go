package file

import "path/filepath"
import "os"
 
func CurrentDir() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("unable to get current process")
	}
	return
}

func FileExist(file string) bool {
	_, err := os.Stat(file)
    if err != nil && os.IsNotExist(err) {  
        return false  
    }  
    return true  
}
