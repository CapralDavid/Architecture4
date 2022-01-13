package main

import (
	"bufio"
	"os"
    	"flag"
    	"github.com/CapralDavid/Architecture4/engine"
)
// bufio нужно для чтения файла 

// напоминалка для меня
// bufio scanner что бы открыть файл
// .scan рассамтривает следующую строку.нечего рассамтривать?передает false
// engine.parse разбирает строку, определяет команду,возвращает сообщение(для print, reverse)

// надо уточнить у Димы,правильно ли я понял

func main() {
	//создаем луп
	eventLoop := new(engine.EventLoop) 
	eventLoop.Start()
	flag.Parse()

	//откуда читаем команды для лупа
	inputFile := "testFile.txt"

	if input, err := os.Open(inputFile); err == nil {

		defer input.Close()
		// получаем инпут
		scanner := bufio.NewScanner(input)
		emptyFile := true
        
		// читаем инпут,проходим кажд из строк в петле
		for scanner.Scan() {
			emptyFile = false
			// .text дает доступ к строке 
			commandLine := scanner.Text()

			// в методе parse разд эту строку 
			// и рассматриваем кажд часть
			cmd := engine.Parse(commandLine)
			// вносим комманды,сообщения в луп
			eventLoop.Post(cmd)
		}

		if emptyFile {
			cmd := &engine.PrintCmd{Msg: "File is empty!!!"}
			eventLoop.Post(cmd)
		}
	}
	// вот теперь исполняем все
	eventLoop.AwaitFinish()
}
