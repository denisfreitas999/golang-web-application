create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
);

select * from produtos;

insert into produtos (nome, descricao, preco, quantidade) values
	('Camiseta', 'Preta', 19.99, 10),
	('Fone', 'Headset', 99.29, 5);

insert into produtos (nome, descricao, preco, quantidade) values
	('Produto Novo', '2.0 teste', 1.99, 1);