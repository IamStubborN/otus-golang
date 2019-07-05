#Задание

#### Реализовать утилиту копирования файлов (см man dd). Выводить в консоль прогресс копирования. Программа
#### должна корректно обрабатывать ситуацию, когда offset или offset+limit за пределами source файла.
#### Настроить и запустить линтеры, создать Makefile для автоматизации тестирования и сборки. Должна быть
#### возможность скачать протестировать и установить программу с помощью go get/test/install

#Пример использования:

> копирует 2К из source в dest, пропуская 1K данных
> gocopy ­-from /path/to/source ­-to /path/to/dest ­-offset 1024 -­limit 2048

## О Программе
Hello, it's my programm for copying files - gocopy.
This is flags for usage:
	-from string - path to source.
	-to string - path to destination.
	-offset int - offset in file in bytes.
	-limit int - copy limit of file.

## Go комманды

go get 
go install
go test

