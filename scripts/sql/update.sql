# SET @date = '2017-03-04 23:59:59';
SET @date = now();
# 更新最新价格
UPDATE ershoufang
SET price =
(
  SELECT p.price
  FROM ershoufang_price p
  WHERE p.ershoufang_id = ershoufang.id
        AND p.created_at < @date
  ORDER BY created_at DESC
  LIMIT 1
);
# 更新最新单价
UPDATE ershoufang
SET unit_price =
(
  SELECT p.unit_price
  FROM ershoufang_price p
  WHERE p.ershoufang_id = ershoufang.id
        AND p.created_at < @date
  ORDER BY created_at DESC
  LIMIT 1
);

UPDATE xiaoqu xq
  JOIN (
         SELECT
           h.xiaoqu       xiaoqu,
           max(h.bankuai) bankuai,
           max(h.qu)      qu
         FROM ershoufang h
         WHERE h.xiaoqu IS NOT NULL
         GROUP BY h.xiaoqu) t
    ON xq.name = t.xiaoqu
SET xq.bankuai = t.bankuai, xq.qu = t.qu;

