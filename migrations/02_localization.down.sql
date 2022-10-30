ALTER TABLE member_infos
ALTER COLUMN generation TYPE VARCHAR (10),
ALTER COLUMN blood_type TYPE VARCHAR (3);

ALTER TABLE member_infos
DROP CONSTRAINT member_info_id_locale_unique;

ALTER TABLE member_infos
DROP CONSTRAINT fk_locale;

ALTER TABLE member_infos
DROP COLUMN locale_id;

DROP TABLE locales;
