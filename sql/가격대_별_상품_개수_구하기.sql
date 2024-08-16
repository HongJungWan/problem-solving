-- https://school.programmers.co.kr/learn/courses/30/lessons/131530 [가격대 별 상품 개수 구하기]

SELECT TRUNCATE(PRICE, -4) as PRICE_GROUP, COUNT(PRODUCT_ID) as PRODUCTS
FROM PRODUCT
GROUP BY PRICE_GROUP
ORDER BY PRICE ASC;