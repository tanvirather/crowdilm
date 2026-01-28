create table if not exists quran_line (
  quran_id varchar(100) character set utf8 collate utf8_general_ci not null,
  line_id int not null,
  text varchar(2000) character set utf8 collate utf8_general_ci not null,
  primary key(quran_id, line_id)
) engine=innodb default charset=utf8 collate=utf8_unicode_ci;
