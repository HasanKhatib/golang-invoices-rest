-- db/migrations/<timestamp>_create_invoices_table.up.sql
CREATE TABLE invoices (
  id          VARCHAR(255) PRIMARY KEY,
  description TEXT,
  amount      DECIMAL(10,2)
);
