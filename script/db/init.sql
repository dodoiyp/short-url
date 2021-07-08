
create table shorturl.urls (
  short_url  varchar(20),
  url varchar(200),
  expire_at DATETIME,
  created_at DATETIME,
  primary key (short_url)
) ENGINE=INNODB;


create table shorturl.sequences (
  sequence_no int auto_increment, 
  PRIMARY KEY (sequence_no)
) ENGINE=INNODB;
