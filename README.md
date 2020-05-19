* Для работы с mysql используется orm [GORM](https://gorm.io/)
* На таблицу fias_address вешаем индекс aoguid
* на таблицу fias_house вешаем индекс houseguid
* Запуск выгрузки go build -o fiascli && ./fiascli checkupdates
* Генерация по спеке proto - protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. domain/address/delivery/grpc/address/address.proto && protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:.  domain/address/delivery/grpc/address/address.proto && protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:.  domain/address/delivery/grpc/address/address.proto

TODO:
* Автоматически накинуть индексы
* Проработать методы апи(пока сделаны для теста)
* Не сохраняется версия
* Проработать домены, вынести interfaces