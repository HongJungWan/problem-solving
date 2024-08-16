-- https://school.programmers.co.kr/learn/courses/30/lessons/131533 [상품 별 오프라인 매출 구하기]

SELECT p.product_code as PRODUCT_CODE, SUM(p.price * o.sales_amount) as SALES
FROM product p
join offline_sale o on p.product_id = o.product_id
GROUP BY p.product_id
ORDER BY SALES DESC, p.product_code