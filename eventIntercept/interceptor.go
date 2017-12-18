package eventIntercept

import (
	"strconv"
	"time"
	"os/exec"
	"os"
	"fmt"
	//"github.com/ethereum/go-ethereum/logger"
	//"github.com/ethereum/go-ethereum/logger/glog"
	"bufio"
	"log"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/logger/glog"
)

type Interceptor struct {
	sendNode int
	recvNode int
	fileDir string
	filename string
	messageType string
	state int
	info int
	eventId int
	exist chan bool

}
/*important global variable to decide whether intercept or not in the ethereum system*/
var IsIntercept bool = true

func NewIntercept(sendNode int, recvNode int, messageType string, state int, info int){
	interceptor := &Interceptor{
					sendNode:		sendNode,
					recvNode:		recvNode,
					messageType:    messageType,
					state:			state,
					info:			info}

	interceptor.eventId = interceptor.gethash()
	interceptor.fileDir = "/tmp/ipc"
	interceptor.filename = interceptor.createFilename()
	interceptor.exist = make(chan bool)

	//start intercepting messages into the file inside /new/ folder
	newFileName := interceptor.fileDir+"/new/"+interceptor.filename
	newFile, err := os.Create(newFileName) // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	writer := bufio.NewWriter(newFile)
	defer newFile.Close()
	fmt.Fprintln(writer, "sendNode=",sendNode)
	fmt.Fprintln(writer,"recvNode=",recvNode)
	fmt.Fprintln(writer,"messageType=",messageType)
	fmt.Fprintln(writer, "state=",state)
	fmt.Fprintln(writer, "info=",info)
	fmt.Fprintln(writer, "eventId=",interceptor.eventId)
	writer.Flush()

	interceptor.commitEvent()
	ackFileName := interceptor.fileDir+"/ack/"+interceptor.filename
	go interceptor.wait(ackFileName)
	interceptor.waitAck(ackFileName)

}

func UpdateIntercept(sendNode int, recvNode int, messageType string, state int, info int) {
	interceptor := &Interceptor{
		sendNode:		sendNode,
		recvNode:		recvNode,
		messageType:    messageType,
		state:			state,
		info:			info}

	interceptor.eventId = interceptor.gethash()
	interceptor.fileDir = "/tmp/ipc"
	interceptor.filename = interceptor.createFilename()

	//start intercepting messages into the file inside /new/ folder
	newFileName := interceptor.fileDir+"/new/"+interceptor.filename
	newFile, err := os.Create(newFileName) // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	writer := bufio.NewWriter(newFile)
	defer newFile.Close()
	fmt.Fprintln(writer, "sendNode=",sendNode)
	fmt.Fprintln(writer,"recvNode=",recvNode)
	fmt.Fprintln(writer,"messageType=",messageType)
	fmt.Fprintln(writer, "state=",state)
	fmt.Fprintln(writer, "info=",info)
	fmt.Fprintln(writer, "eventId=",interceptor.eventId)
	writer.Flush()

	interceptor.commitEvent()
}


func (self *Interceptor) commitEvent()  {
	glog.V(logger.Info).Infoln("@huanke mv from new to send ",self.filename)
	cmd := exec.Command("mv", self.fileDir+"/new/"+self.filename, self.fileDir+"/send/"+self.filename)
	cmd.Run()
}


func (self *Interceptor) wait(ackFileName string)  {
	glog.V(logger.Info).Infoln("@huanke waitAck ", ackFileName)
	for  {
		_, err := os.Stat(ackFileName)
		if err==nil {
			self.exist<-true
			return
		}
		time.Sleep(time.Millisecond*250)
	}
}

func (self *Interceptor) waitAck(ackFileName string)  {
	<-self.exist
	glog.V(logger.Info).Infoln("@huanke waitExist ", ackFileName)

	//open and read the ack file
	file, err1 := os.Open(ackFileName)
	if err1!=nil {
		panic(err1)
	}
	content := make([]byte, 10)
	file.Read(content)
	file.Close()
	glog.V(logger.Info).Infoln("@huanke ackFile : %s",content)

	//remove the ack file
	cmd := exec.Command("rm", ackFileName)
	cmd.Run()

}

func (self *Interceptor) createFilename() string {
	var filename string
	now := time.Now()
	nanos := now.Nanosecond()
	millis := nanos / 1000000
	filename = "eth-" + strconv.Itoa(self.eventId)+ strconv.Itoa(millis)
	return filename
}


func (self *Interceptor) gethash() int {
	var prime int = 19
	var hash int = 1
	//sendNode, _:= strconv.Atoi(self.sendNode) how to convert string to int
	hash = prime*hash + self.sendNode
	hash = prime*hash +  self.recvNode
	hash = prime*hash + self.state
	hash = prime*hash + self.info
	return hash
}