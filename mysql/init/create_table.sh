CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

$CMD_MYSQL -e "create table todos (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    task varchar(50) NOT NULL,
    description varchar(1000),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );"
$CMD_MYSQL -e "insert into todos values (
    1, 
    'washing dishes', 
    'a lot of dishes because of the party at my house last night', 
    now(), 
    now()
    );"
$CMD_MYSQL -e "insert into todos values (
    2,
    'contact with some travel agencies', 
    'I''m planning summer trip with my family. But, we don''t have enough money to trip very much. I need more information.', 
    now(), 
    now()
    );"
