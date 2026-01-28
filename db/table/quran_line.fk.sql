alter table quran_line add constraint fk_quran_quran_id_quran_id foreign key (quran_id) references quran(id);
alter table quran_line add constraint fk_line_line_id_line_id foreign key (line_id) references line(id);
