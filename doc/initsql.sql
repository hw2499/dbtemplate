create table if not exists  t_device_info (
  d_id varchar(32) not null,
  d_name varchar(50),
  d_setup_address varchar(100),
  d_ip varchar(20),
  d_desc varchar(100),
  d_width float, 
  d_height float,
  d_writetime varchar(19),
  PRIMARY KEY (d_id)
);
create table if not exists  t_customer_info (
  c_id varchar(32) not null, 
  c_name varchar(50),
  c_address varchar(100),
  c_phone varchar(20),
  c_writetime varchar(19),
  PRIMARY KEY (c_id)
);
create table if not exists  t_c_d_info (
  cd_id varchar(32) not null,
  cd_c_id varchar(32),
  cd_d_id varchar(32),
  cd_setup_writetime varchar(19),
  PRIMARY KEY (cd_id)
);