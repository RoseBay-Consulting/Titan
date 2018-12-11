package benchmarking

import (
	"time"
	"strconv"
	"fmt"
	"os"
	"bufio"
	"github.com/ethereum/go-ethereum/node"
)

func makeTimestamp() int64 {


	return (time.Now().UnixNano()/int64(time.Millisecond)) //1000000
}

func Writeincsv(nodeid string, message string){

	writeincsv(nodeid,message, makeTimestamp() )

}


//nodeid to name the node namee"
//	"strconv"
//	"fmt"
//	"os"
//message is on which process it occured
//data is datetime
func writeincsv(nodeid string, mesasge string, data int64) {
	convertedata := strconv.FormatInt(data, 10)
	var x=nodeid+","+mesasge+","+convertedata;
	path := node.DefaultDataDir() + "/result.txt"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fileHandle, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()
		fmt.Fprintln(writer, x)
		writer.Flush()
	}	else {
		fileHandle, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		writer := bufio.NewWriter(fileHandle)
		defer fileHandle.Close()
		fmt.Fprintln(writer, x)
		writer.Flush()
	}



	//fileHandle, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)

}

