# Задание

#### Реализовать утилиту копирования файлов (см man dd). Выводить в консоль прогресс копирования. 
#### Программа должна корректно обрабатывать ситуацию, когда offset или offset+limit за пределами source файла.
#### Настроить и запустить линтеры, создать Makefile для автоматизации тестирования и сборки. 
#### Должна быть возможность скачать протестировать и установить программу с помощью go get/test/install

# Пример использования:

> копирует 2К из source в dest, пропуская 1K данных<br/>
> gocopy ­-from /path/to/source ­-to /path/to/dest ­-offset 1024 -­limit 2048<br/>

## О Программе
Hello, it's my programm for copying files - gocopy.<br/>
This is flags for usage:<br/>
<br/>
-from string - path to source.<br/>
-to string - path to destination.<br/>
-offset int - offset in file in bytes.<br/>
-limit int - copy limit of file.<br/>

## Make комманды

make<br/>
make install<br/>
make build<br/>
make test<br/>
make clean<br/>
make run<br/>

make build_docker<br/>
