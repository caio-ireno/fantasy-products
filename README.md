# QUERY 1 - Without total updated in Sales

```sql
select SUM(p.price*s.quantity) as total, c.condition 
from 
    sales s 
    join invoices i on s.invoice_id=i.id
    join customers c on c.id =i.customer_id
    join products p on s.product_id= p.id
GROUP BY c.condition;
```

Melhoria sugerida:

- Adicionar um case para gerar o Activo (1)/ Inactivo (0).
- Usar o ROUND para arredondar.

```sql
SELECT 
    CASE c.condition
        WHEN 1 THEN 'Activo ( 1 )'
        ELSE 'Inactivo ( 0 )'
    END AS Condition1,
    ROUND(SUM(s.quantity * p.price), 2) AS Total
FROM
    sales s
    JOIN invoices i ON s.invoice_id = i.id
    JOIN customers c ON c.id = i.customer_id
    JOIN products p ON s.product_id = p.id
GROUP BY c.condition;
```


# QUERY UPDATE TOTAL 

```sql
select SUM(s.quantity) as total, s.invoice_id, i.customer_id
from 
sales s 
join invoices i on s.invoice_id=i.id
join customers c on c.id =i.customer_id
join products p on s.product_id= p.id
GROUP BY s.invoice_id;
```