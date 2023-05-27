#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table article (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    task varchar(50) NOT NULL,
    description varchar(1000)
    );"
$CMD_MYSQL -e  "insert into article values (1, 'washing dishes', 'a lot of dishes because of the party at my house last night');"
$CMD_MYSQL -e  "insert into article values (2, 'contact with some travel agencies', 'I'm planning summer trip with my family. But, we don't have enough money to trip very much. I need more information.');"
