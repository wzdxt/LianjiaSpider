# 小区均价
SELECT
  xq.name                          小区,
  sum(h.price) / sum(h.size / 100) 均价
FROM xiaoqu xq
  JOIN ershoufang h ON h.xiaoqu_page_id = xq.page_id
GROUP BY xq.id
ORDER BY 均价;

# 今日调价
SELECT *
FROM ershoufang_price
WHERE date_format(created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d');

# 调价历史
SELECT
  xq.name      小区,
  h.name       房产,
  p.created_at 时间,
  p.price      总价,
  p.unit_price 单价
FROM ershoufang h
  JOIN ershoufang_price pp ON h.id = pp.ershoufang_id AND pp.prev_id > 0
                              AND date_format(pp.created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d')
  LEFT JOIN ershoufang_price p ON h.id = p.ershoufang_id
  JOIN xiaoqu xq ON h.xiaoqu_page_id = xq.page_id
ORDER BY xq.id, h.id, p.id;

# 今日调价趋势
SELECT
  sum(调幅),
  sum(面积),
  sum(调幅) / sum(面积) * 100 单价差
FROM (
       SELECT
         xq.name            小区,
         h.name             房产,
         h.size             面积,
         p.price - pp.price 调幅
       FROM ershoufang_price p
         JOIN ershoufang_price pp ON p.prev_id = pp.id
         JOIN ershoufang h ON p.ershoufang_id = h.id
         JOIN xiaoqu xq ON h.xiaoqu_page_id = xq.page_id
       WHERE date_format(p.created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d')
     ) t;

# 今日新房源
SELECT
  xq.name      小区,
  h.name       房产,
  h.price      总价,
  h.unit_price 单价
FROM ershoufang h
  JOIN xiaoqu xq ON xq.page_id = h.xiaoqu_page_id
WHERE date_format(h.created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d');

# 今日出售
SELECT
  xq.name      小区,
  h.name       房产,
  h.page_id    房产编号,
  h.price      总价,
  h.unit_price 单价,
  h.sold_date  成交日期
FROM xiaoqu xq
  JOIN ershoufang h ON h.xiaoqu_page_id = xq.page_id AND h.sold_date IS NOT NULL
ORDER BY xq.id;

# 成交总价分布
SET @jiange = 200;
SET @sum = (SELECT count(*)
            FROM ershoufang h
            WHERE h.sold_date IS NOT NULL);
SELECT
  concat(h.price DIV @jiange * @jiange, ' - ', h.price DIV @jiange * @jiange + @jiange) 价位,
  count(*)                                                                              数量,
  concat(count(*) / @sum * 100, '%')                                                    比例
FROM
  ershoufang h
WHERE h.sold_date IS NOT NULL
GROUP BY h.price DIV @jiange
ORDER BY h.price DIV @jiange;


;
SELECT date_format(now(), '%Y-%m-%d')
FROM dual;
