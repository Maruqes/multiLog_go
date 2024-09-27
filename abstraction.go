package multiLog

func CreateLog(identifier string) {
	add_tab(identifier, "New log created: "+identifier+"\n")
}

func Log(identifier string, content string) {
	content = content + "\n"
	add_content(identifier, content)
}

func Error(identifier string, content string) {
	add_content(identifier, "|ERR|"+content+"|ERR|\n")
}

func Warning(identifier string, content string) {
	add_content(identifier, "|WARN|"+content+"|WARN|\n")
}

func Critical(identifier string, content string) {
	add_content(identifier, "|ERR|☠️CRITICAL ERROR☠️"+content+"|ERR|\n")
}

func Info(identifier string, content string) {
	add_content(identifier, "|INFO|"+content+"|INFO|\n")
}

func Panic(identifier string, content string) {
	add_content(identifier, "|ERR|PANIC! PANIC!"+content+"|ERR|\n")
	add_content(identifier, "|ERR|PANIC! PANIC!"+content+"|ERR|\n")
	add_content(identifier, "|ERR|PANIC! PANIC!"+content+"|ERR|\n")
	panic(content)
}

func RemoveLog(identifier string) {
	remove_tab(identifier)
}
