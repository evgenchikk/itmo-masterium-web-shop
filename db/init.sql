CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "name" varchar(30),
  "email" varchar(30),
  "password" varchar(30),
  "created_at" timestamp DEFAULT NOW()
);

CREATE TABLE "carts" (
  "id" serial PRIMARY KEY,
  "user_id" int
);

CREATE TABLE "item_categories" (
  "id" serial PRIMARY KEY,
  "name" text
);

CREATE TABLE "item_groups" (
  "id" serial PRIMARY KEY,
  "category_id" int,
  "name" text,
  "image" text,
  "description" text
);

CREATE TABLE "items" (
  "id" serial PRIMARY KEY,
  "item_group_id" int,
  "price" numeric(12,3),
  "image" text,
  "feature" text
);

CREATE TABLE "carts_items" (
  "cart_id" int,
  "item_id" int,
  "count" int,
  PRIMARY KEY ("cart_id", "item_id")
);

CREATE TABLE "favourites" (
  "user_id" int,
  "item_id" int,
  PRIMARY KEY ("user_id", "item_id")
);

CREATE TABLE "tokens" (
  "id" serial PRIMARY KEY,
  "user_id" int,
  "token" text,
  "created_at" timestamp DEFAULT NOW()
);


ALTER TABLE "carts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "item_groups" ADD FOREIGN KEY ("category_id") REFERENCES "item_categories" ("id");

ALTER TABLE "items" ADD FOREIGN KEY ("item_group_id") REFERENCES "item_groups" ("id");

ALTER TABLE "carts_items" ADD FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "carts_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "favourites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "favourites" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


-- INSERTING VALUES --

insert into item_categories (name)
values ('Плюш'),
       ('Дерево');

insert into item_groups (name, category_id, description, image)
values ('Плюшевая игрушка в форме тыквы', 1, 'Это уникальная игрушка, которая понравится не только детям, но и взрослым. Она точно копирует внешний вид настоящего овоща, а фактурная поверхность и ручное раскрашивание делают каждую фигурку уникальной. Если вашему ребенку нравятся овощи, а в частности тыквы - он будет в полном восторге от такой игрушки! Но реалистичные фигурки предметов пользуются популярностью не только в сюжетных играх у детей - взрослые с удовольствием используют их для обустраивания интерьера помещений.', 'pumpkin-yellow.jpg'),
       ('Детский игрушечный автомобиль', 2, 'Простота этой деревянной игрушки развивает воображение вашего  ребенка, вдохновляя придумывать разные сюжеты и сценарии игр. Так как игрушка цельная, а нк состоящая из нескольких частей, то она не будет разбросана по всему дому.', 'Hedgehog.jpg'),
       ('Обучающие деревянные цветные камни', 2, 'Такие игрушки учат концентрировать свое внимание на деталях, а также развивают навыки осмысливания, ведь в отличие от высокотехнологических игрушек, разработанных так, что они отвлекают, а иногда даже раздражают ребенка, оказывая на него сильное эмоциональное воздействие. Ничто так не поможет развитию мелкой моторики вашего ребенка, как игры в деревянные камни, блоки, сбор конструкторов.', 'mint-stones.jpg'),
       ('Плюшевый мишка', 1, 'Мягкие игрушки мишки пользуются большой популярностью среди взрослых и детей. Плюшевые мишки нашего магазина заинтересуют своей внешностью, качеством, маленьким размером. За счёт лёгкости наполнителя, плюшевого малыша можно брать его в дорогу маленькому путешественнику. Для создания игрушек используется только прочный и безопасный материал. Их баятельные мордашки  излучают море нежности и тепла. Для них не сложно найти местечко в доме или квартире.  Пушистый мех достаточно густой и одновременно мягкий, прикасаться к которому одно удовольствие. Надёжней друга, чем плюшевый мишка, Вам не найти.', 'yellow-bear.jpg'),
       ('Набор деревянных строительных блоков', 2, 'Такие игрушки учат концентрировать свое внимание на деталях, а также развивают навыки осмысливания, ведь в отличие от высокотехнологических игрушек, разработанных так, что они отвлекают, а иногда даже раздражают ребенка, оказывая на него сильное эмоциональное воздействие. Ничто так не поможет развитию мелкой моторики вашего ребенка, как игры в деревянные камни, блоки, сбор конструкторов.', 'blue-blocks.jpg'),
       ('Детский деревянный конструктор', 2, 'Такие игрушки учат концентрировать свое внимание на деталях, а также развивают навыки осмысливания, ведь в отличие от высокотехнологических игрушек, разработанных так, что они отвлекают, а иногда даже раздражают ребенка, оказывая на него сильное эмоциональное воздействие. Ничто так не поможет развитию мелкой моторики вашего ребенка, как игры в деревянные камни, блоки, сбор конструкторов.', 'green-constructor.jpg');

insert into items (price, image, item_group_id, feature)
values (800, 'pumpkin-yellow.jpg', 1, 'Цвет: Желтый'),
       (800, 'pumpkun-dark-green.jpg', 1, 'Цвет: Тёмно зеленый'),
       (800, 'pumpkin-white-brown.jpg', 1, 'Цвет: Бело коричневый'),
       (800, 'pumpkin-orange.png', 1, 'Цвет: Оранжевый'),
       (800, 'pumpkin-brown.jpg', 1, 'Цвет: Коричневый'),
       (800, 'pumpkin-blue.jpg', 1, 'Цвет: Синий'),
       (800, 'pumpkin-red.jpg', 1, 'Цвет: Красный'),
       (800, 'pumpkin-light-green.jpg', 1, 'Цвет: Светло зеленый'),
       (800, 'pumpkin-rose-purple.jpg', 1, 'Цвет: Пыльная роза'),
       --
       (2000, 'Hedgehog.jpg', 2, 'Форма: Ёж'),
       (2000, 'Elephant.jpg', 2, 'Форма: Слон'),
       (2000, 'Dinosaur.jpg', 2, 'Форма: Динозавр'),
       (2000, 'car-2.jpg', 2, 'Форма: Большая машина'),
       (2000, 'car-3.jpg', 2, 'Форма: Маленькая машина'),
       (2000, 'bus.jpg', 2, 'Форма: Автобус'),
       (2000, 'Truck.jpg', 2, 'Форма: Маленький трактор'),
       (2000, 'Truck-2.jpg', 2, 'Форма: Большой трактор'),
       --
       (1800, 'mint-stones.jpg', 3, 'Цвет: Мятный'),
       (1800, 'blue-stones.jpg', 3, 'Цвет: Голубой'),
       (1800, 'brigth-stones.jpg', 3, 'Цвет: Разноцветные темные'),
       (1800, 'pink-stones.jpg', 3, 'Цвет: Розовые'),
       (1800, 'light-stones.jpg', 3, 'Цвет: Разноцветные светлые'),
       --
       (2000, 'yellow-bear.jpg', 4, 'Цвет: Жёлто белый'),
       (2000, 'blue-bear.jpg', 4, 'Цвет: Голубой'),
       (2000, 'yellow-bear-2.jpg', 4, 'Цвет: Жёлтый'),
       --
       (1200, 'blue-blocks.jpg', 5, 'Цвет: Голубой'),
       (1200, 'brown-blocks.jpg', 5, 'Цвет: Коричневый'),
       (1200, 'green-blocks.jpg', 5, 'Цвет: Зеленый'),
       (1200, 'light-green-blocks.jpg', 5, 'Цвет: Салатовый'),
       (1200, 'rainbow-blocks.jpg', 5, 'Цвет: Радужный'),
       --
       (1000, 'green-constructor.jpg', 6, 'Цвет: Зеленый'),
       (1000, 'pink-constructor.jpg', 6, 'Цвет: Розовый'),
       (1000, 'red-constructor.jpg', 6, 'Цвет: Красный'),
       (1000, 'grey-constructor.jpg', 6, 'Цвет: Красный'),
       (1000, 'beige-constructor.jpg', 6, 'Цвет: Бежевый');
