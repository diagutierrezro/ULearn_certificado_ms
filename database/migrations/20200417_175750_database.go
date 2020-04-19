package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Database_20200417_175750 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Database_20200417_175750{}
	m.Created = "20200417_175750"

	migration.Register("Database_20200417_175750", m)
}

// Run the migrations
func (m *Database_20200417_175750) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SEQUENCE id_certificado_seq START WITH 1 INCREMENT BY 1 NO MINVALUE MAXVALUE 123456789 CACHE 1;")
	m.SQL("CREATE TABLE certificado ( id_certificado integer DEFAULT nextval('public.id_certificado_seq'::regclass) PRIMARY KEY, texto text NOT NULL, id_usuario integer DEFAULT 1 NOT NULL, id_curso integer NOT NULL );")
}

// Reverse the migrations
func (m *Database_20200417_175750) Down() {
	m.SQL("DROP SEQUENCE if exists id_certificado_seq CASCADE;")
	m.SQL("DROP TABLE if exists certificado CASCADE;")
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
