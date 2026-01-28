create table if not exists quran(
  id varchar(100) character set utf8 collate utf8_general_ci not null,
  language varchar(100) character set utf8 collate utf8_general_ci not null,
  name varchar(100) character set utf8 collate utf8_general_ci not null,
  name_english varchar(100) character set utf8 collate utf8_general_ci not null,
  quran_type varchar(100) character set utf8 collate utf8_general_ci not null,
  primary key(id)
) engine=innodb default charset=utf8 collate=utf8_unicode_ci;
