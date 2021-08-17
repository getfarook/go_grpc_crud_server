Overview:
-------------------------
This is golang backend project based on gRPC protocol. It does CRUD operations for partner and product. Invoice part is imcomplete.

The front end for this project is done using flutter project. The front end flutter project repository name is "flutter_grpc_crud". 



Database used:
---------------------------
Postgresql installed in my local system was used for this project. 

SQLs used for this project:
-----------------------------
Products:

CREATE TABLE products (id SERIAL PRIMARY KEY, name TEXT UNIQUE NOT NULL, price NUMERIC, details TEXT);

INSERT INTO products (name, price, details) VALUES ('Sony Playstation', 345.56, 'VR Game consoles');

Partners:

CREATE TABLE partners (id SERIAL PRIMARY KEY, name TEXT NOT NULL, open_bal NUMERIC, curr_bal NUMERIC, contact_no TEXT, details TEXT);

INSERT INTO partners (name, open_bal, curr_bal, contact_no, details) VALUES ('Thomas K', 345.56, 90700.00, '957734622',  'Retailer Ready Cash');

Invoice:

CREATE TABLE invoice(
id SERIAL PRIMARY KEY,
invoice_no TEXT NOT NULL, 
partner_id INT,
date DATE NOT NULL,
inv_amount NUMERIC,

CONSTRAINT fk_partner
	FOREIGN KEY(partner_id)
		REFERENCES partners(id)
			ON DELETE RESTRICT
);



CREATE TABle invoice_line(
id SERIAL PRIMARY KEY,
invoice_id INT,
product_id INT,
qty INT,
price NUMERIC,
line_amount NUMERIC,

CONSTRAINT fk_invoice
	FOREIGN KEY(invoice_id)
		REFERENCES invoice(id)
ON DELETE CASCADE,

CONSTRAINT fk_product
	FOREIGN KEY(product_id)
		REFERENCES products(id)
			ON DELETE RESTRICT


);


Sample Inserts:

INSERT INTO invoice (invoice_no, partner_id, date, inv_amount) VALUES ('INV001', 1, '2020-12-23', 13400.56);

INSERT INTO invoice_line (invoice_id, product_id, qty, price, line_amount) VALUES (1, 17, 15, 250.00, 3750.00);
