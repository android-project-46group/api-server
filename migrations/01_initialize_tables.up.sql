CREATE TABLE IF NOT EXISTS groups(
    group_id    serial PRIMARY KEY,
    group_name  VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS members(
    member_id   serial PRIMARY KEY,
    group_id    INT NOT NULL,
    name_ja     VARCHAR (20) NOT NULL,
    joined_at   TIMESTAMP,
    left_at     TIMESTAMP,
    CONSTRAINT fk_group
        FOREIGN KEY(group_id)
            REFERENCES groups(group_id)
);

CREATE TABLE IF NOT EXISTS member_infos(
    member_info_id  serial PRIMARY KEY,
    member_id       INT NOT NULL,
    birthday        VARCHAR (10),
    blood_type      VARCHAR (3),
    height          SMALLINT,
    generation      SMALLINT,
    blog_url        VARCHAR (100),
    img_url         VARCHAR (100),
    CONSTRAINT fk_member
        FOREIGN KEY(member_id)
            REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS formations(
    formation_id    serial PRIMARY KEY,
    group_id        INT NOT NULL,
    member_id       INT NOT NULL,
    single          VARCHAR (10),
    song_title      VARCHAR (20),
    position        VARCHAR (10),
    CONSTRAINT fk_formation_g
        FOREIGN KEY(group_id)
            REFERENCES groups(group_id),
    CONSTRAINT fk_formation_m
        FOREIGN KEY(member_id)
            REFERENCES members(member_id)
);

CREATE TABLE IF NOT EXISTS tags(
    tag_id    serial PRIMARY KEY,
    tag_name  VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS member_tags(
    member_tag_id   serial PRIMARY KEY,
    member_id       INT NOT NULL,
    tag_id          INT NOT NULL,
    CONSTRAINT fk_member_tag_m
        FOREIGN KEY(member_id)
            REFERENCES members(member_id),
    CONSTRAINT fk_member_tag_t
        FOREIGN KEY(tag_id)
            REFERENCES tags(tag_id)
);
