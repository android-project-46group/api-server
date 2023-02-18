ALTER TABLE member_infos ALTER COLUMN height_cm TYPE varchar;
UPDATE member_infos SET height_cm = height_cm::varchar;

ALTER TABLE member_infos RENAME COLUMN height_cm TO height;

-- 文字列の末尾から cm を追加
UPDATE member_infos
SET height = height || 'cm'
WHERE height NOT LIKE '%cm';
