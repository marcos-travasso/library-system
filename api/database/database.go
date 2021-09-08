package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Database struct {
	Dir string
}

func (dbDir *Database) CreateDatabase() {
	if _, err := os.Stat(dbDir.Dir); os.IsNotExist(err) {
		_, err := os.Create(dbDir.Dir)
		if err != nil {
			log.Printf("Failed to create: %q\n", err)
			return
		}
	}

	err := dbDir.fillDatabaseTables()
	if err != nil {
		log.Fatal(err)
	}
}

func (dbDir *Database) fillDatabaseTables() error {
	conn, err := sql.Open("sqlite3", dbDir.Dir)
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close: %q\n", err)
		}
	}(conn)

	sqlStatments := []string{
		`CREATE TABLE IF NOT EXISTS Autores(idAutor INTEGER PRIMARY key, pessoa INTEGER, FOREIGN KEY(pessoa) REFERENCES pessoas(idPessoa))`,
		`CREATE TABLE IF NOT EXISTS Enderecos(idEndereco INTEGER primary key, CEP CHAR(8), cidade VARCHAR(40), bairro VARCHAR(25), rua varchar(50), numero integer, complemento VARCHAR(20))`,
		`CREATE TABLE IF NOT EXISTS Generos(idGenero integer PRIMARY key, nome varchar(20))`,
		`CREATE TABLE IF NOT EXISTS Livros(idLivro integer primary key, titulo Varchar(50), ano int, autor integer, paginas integer, foreign key(autor) REFERENCES Autores(idAutor))`,
		`CREATE TABLE IF NOT EXISTS Pessoas(idPessoa INTEGER PRIMARY key, nome varchar(50), genero char(1), nascimento CHAR(8))`,
		`CREATE TABLE IF NOT EXISTS Usuarios(idUsuario INTEGER primary key, pessoa INTEGER, celular CHAR(11), telefone CHAR(10), endereco integer, cpf char(11), email varchar(50), responsavel integer, criacao DATE, FOREIGN KEY(pessoa) REFERENCES Pessoas(idPessoa), FOREIGN KEY(responsavel) REFERENCES Pessoas(idPessoa))`,
		`CREATE TABLE IF NOT EXISTS devolucoes(idDevolucao integer primary key, emprestimo integer, dataDeDevolucao date, foreign key(emprestimo) references emprestimos(idEmprestimo))`,
		`CREATE TABLE IF NOT EXISTS "emprestimos" ("idEmprestimo"	INTEGER, "livro" integer, "usuario" integer, "dataDoPedido"	date, "devolvido" INTEGER NOT NULL DEFAULT 0, FOREIGN KEY("livro") REFERENCES "Livros"("idLivro"), FOREIGN KEY("usuario") REFERENCES "Usuarios"("idUsuario"), PRIMARY KEY("idEmprestimo"))`,
		`CREATE TABLE IF NOT EXISTS generos_dos_livros(livro integer, genero integer, FOREIGN key(livro) REFERENCES Livros(idlivro), foreign key(genero) REFERENCES Generos(idGenero))`,
	}

	for _, sqlStmt := range sqlStatments {
		_, err = conn.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
			return err
		}
	}
	return nil
}

func (dbDir *Database) clearDatabase() error {
	if _, err := os.Stat(dbDir.Dir); !os.IsNotExist(err) {
		err := os.Remove(dbDir.Dir)
		if err != nil {
			log.Printf("Failed to remove: %q\n", err)
			return err
		}
	}
	return nil
}

func sendStatement(e entity, statementType string, db *sql.DB) error {

	statement, err := e.SQLStatement(statementType)
	if err != nil {
		log.Printf("Fail to get statement: %s", err)
		return err
	}

	_, err = db.Exec(statement)
	if err != nil {
		log.Printf("Fail to %s: %s", statementType, err)
		return err
	}

	return err
}
