package main

import "log"

type YceLogger struct {
	Logger log.logger
	Err    myerror.YceError
}

// plan 1: call log.Printf("")
// the formatted string is made by ourself, call Sprintf()

// plan 2: call log.New()
// the formatted string is made by logger

//out: os.Stderr
//prefix: time.Now, Who[controller]
//flag: Lshortfile
func New(out io.Writer, prefix string, flag int) {
	return &YceLogger{Logger: log.New(out, prefix, flag)}
}

//2016-08-23 07:55:32 [ListDeployController] getApiServer error: error=Operation time out
func (y *YceLogger) Printf() {
	Logger.Printf(y.Err.Error)
}

func (y *YceLogger) WriteResponse() {
	//TODO: iris.Write(). But this iris is not the same with one in Controllers
	//TODO: Response.code, Response.Message, Response.Data
	// Response.Data = Err.Encode()
}
