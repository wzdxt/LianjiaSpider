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


