-- 文字列の末尾から cm を削除
UPDATE member_infos 
SET height = TRIM(TRAILING 'cm' FROM height)
WHERE height LIKE '%cm';

ALTER TABLE member_infos RENAME COLUMN height TO height_cm;

ALTER TABLE member_infos
ALTER COLUMN height_cm TYPE double precision USING height_cm::double precision;
