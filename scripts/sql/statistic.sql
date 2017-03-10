# 小区均价
SELECT
  h.xiaoqu                         小区,
  sum(h.price) / sum(h.size / 100) 均价
FROM ershoufang h
GROUP BY h.xiaoqu
ORDER BY 均价 DESC;

# 今日调价
SELECT *
FROM ershoufang_price
WHERE date_format(created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d');

# 调价历史
SELECT
  h.xiaoqu               小区,
  h.name                 房产,
  p.created_at           时间,
  p_prev.price           调价前,
  p.price - p_prev.price 差价,
  p.price                调价后,
  p.unit_price           单价
FROM ershoufang h
  JOIN ershoufang_price pp ON h.id = pp.ershoufang_id AND pp.prev_id > 0
                              AND date_format(pp.created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d')
  LEFT JOIN ershoufang_price p ON h.id = p.ershoufang_id
  JOIN ershoufang_price p_prev ON p_prev.id = p.prev_id
  JOIN xiaoqu xq ON h.xiaoqu_page_id = xq.page_id
ORDER BY xq.id, h.id, p.id;

# 今日调价趋势
SELECT
  sum(调幅),
  sum(abs(调幅) + 调幅) / 2   上调,
  sum(abs(调幅) - 调幅) / 2   下调,
  sum(面积),
  sum(调幅) / sum(面积) * 100 单价差
FROM (
       SELECT
         xq.name                小区,
         h.name                 房产,
         h.size                 面积,
         p.price - p_prev.price 调幅
       FROM ershoufang_price p
         JOIN ershoufang_price p_prev ON p.prev_id = p_prev.id
         JOIN ershoufang h ON p.ershoufang_id = h.id
         JOIN xiaoqu xq ON h.xiaoqu_page_id = xq.page_id
       WHERE date_format(p.created_at, '%y-%m-%d') = date_format(now(), '%y-%m-%d')
     ) t;

# 每日调价趋势
SELECT
  date_format(时间, '%y-%m-%d') 时间,
  sum(调幅)                     总差额,
  sum(abs(调幅) + 调幅) / 2       上调,
  sum(abs(调幅) - 调幅) / 2       下调,
  sum(面积)                     总面积,
  sum(调幅) / sum(面积) * 100     单价差
FROM (
       SELECT
         h.xiaoqu               小区,
         p.created_at           时间,
         p.price - p_prev.price 调幅,
         h.size                 面积
       FROM ershoufang_price p
         JOIN ershoufang_price p_prev ON p.prev_id = p_prev.id
         JOIN ershoufang h ON p.ershoufang_id = h.id
         JOIN xiaoqu xq ON h.xiaoqu_page_id = xq.page_id
       WHERE (h.xiaoqu = '普陀' OR h.bankuai = '普陀' OR h.qu = '普陀')
     ) t
GROUP BY date_format(时间, '%y-%m-%d');

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
