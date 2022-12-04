## Backend for web-shop 'Masterium' (ITMO project)

---

Launch:
make build && make run

Stop:
make stop

---

API:

- Get all item groups:
GET /api/item_groups
<br>
- Get item group by id
GET item_groups/:id
<br>
- Get items by item group id
GET item_groups/:id/items
<br>
- Get all items
GET /items
<br>
- Get item by id
GET /items/\<id\>

---

Database schema
<img src="db/schema.png">

Todo:
- Authorization
- Carts

Remarks:
Tried to use a clean architechture