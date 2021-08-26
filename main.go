package main

import (
	"ZheTian/server"
	"fmt"
)

func main() {

	fmt.Println(`
▒███████▒ ██░ ██ ▓█████▄▄▄█████▓ ██▓ ▄▄▄       ███▄    █ 
▒ ▒ ▒ ▄▀░▓██░ ██▒▓█   ▀▓  ██▒ ▓▒▓██▒▒████▄     ██ ▀█   █ 
░ ▒ ▄▀▒░ ▒██▀▀██░▒███  ▒ ▓██░ ▒░▒██▒▒██  ▀█▄  ▓██  ▀█ ██▒
  ▄▀▒   ░░▓█ ░██ ▒▓█  ▄░ ▓██▓ ░ ░██░░██▄▄▄▄██ ▓██▒  ▐▌██▒
▒███████▒░▓█▒░██▓░▒████▒ ▒██▒ ░ ░██░ ▓█   ▓██▒▒██░   ▓██░
░▒▒ ▓░▒░▒ ▒ ░░▒░▒░░ ▒░ ░ ▒ ░░   ░▓   ▒▒   ▓▒█░░ ▒░   ▒ ▒ 
░░▒ ▒ ░ ▒ ▒ ░▒░ ░ ░ ░  ░   ░     ▒ ░  ▒   ▒▒ ░░ ░░   ░ ▒░
░ ░ ░ ░ ░ ░  ░░ ░   ░    ░       ▒ ░  ░   ▒      ░   ░ ░ 
  ░ ░     ░  ░  ░   ░  ░         ░        ░  ░         ░ 
░
:: ZheTian Powerful remote load and execute ShellCode tool
 //[Version 1.0.0] Github [http://github.com/yqcs/ZheTian]`)
	server.Execute()

}
