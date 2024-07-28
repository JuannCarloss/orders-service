package domain

type Orders struct {
	Customer_email string  `json:"email"`
	Timestamp      string  `json:"date"`
	Status         string  `json:"status"`
	Total          float32 `json:"total"`
}

func NewOrder() *Orders {
	return &Orders{}
}

/*
CASSANDRA KEYSPACE

CREATE KEYSPACE orderms
WITH replication =
{'class': 'SimpleStrategy',
'replication_factor': '1'}
 AND durable_writes = true;

CASSANDRA TABLE

CREATE TABLE orders (
    email text,
    date text,
    status text,
    total_price float,
    PRIMARY KEY (email, date)
);
*/
