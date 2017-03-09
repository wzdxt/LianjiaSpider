SELECT
  date_format(时间, '%Y-%m') 月份,
  sum(成交量)                 成交量,
  sum(成交额)                 成交额,
  sum(面积)                  总面积,
  sum(成交额) / sum(面积)       均价
FROM (
       SELECT
         date                   时间,
         count(0) - 1           成交量,
         sum(price)             成交额,
         sum(size)              面积,
         sum(price) / sum(size) 均价
       FROM chengjiao
       WHERE date BETWEEN '2016-01-01' AND now() AND (
         name = '00'
         OR (
           #     qu IN ('普陀', '长宁', '闸北', '徐汇', '静安')
           bankuai IN ('中远两湾城')
#            bankuai IN ('大宁')
         )
       )
       GROUP BY date
     ) t
GROUP BY date_format(时间, '%Y-%m');

# 插入占位数据
# INSERT INTO num (id) VALUES
#   (0),
#   (1),
#   (2),
#   (3),
#   (4),
#   (5),
#   (6),
#   (7),
#   (8),
#   (9);
#
# INSERT INTO chengjiao (date, name, page_id, pic, unit_price, price)
#   SELECT
#     adddate(now(), numlist.id) AS date,
#     '00'                       AS name,
#     numlist.id                 AS page_id,
#     '00'                       AS pic,
#     0                          AS unit_price,
#     0                          AS price
#   FROM
#     (
#       SELECT n1.id + n10.id * 10 + n100.id * 100 + n1000.id * 1000 AS id
#       FROM
#         num n1
#         CROSS JOIN num n10
#         CROSS JOIN num n100
#         CROSS JOIN num n1000
#     ) AS numlist;


