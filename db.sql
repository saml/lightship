CREATE TABLE brands(
  id INT PRIMARY KEY,
  name VARCHAR(63) UNIQUE
);


CREATE TABLE tags(
  id INT PRIMARY KEY,
  slug VARCHAR(127) UNIQUE,
  name VARCHAR(127) NOT NULL
);

CREATE TABLE authors(
  id INT PRIMARY KEY,
  name VARCHAR(127) NOT NULL,
  twitter VARCHAR(127)
);

-- Dynamic crop.
-- Original: a.jpg
-- 100x200 around focus: a.w100.h200.jpg
-- 30x40 around 0,10,300,400 rectangle: a.c300,400@0,10.w30.h40.jpg
-- Quality 72: a.q72.jpg
CREATE TABLE images(
  id INT PRIMARY KEY,
  width INT NOT NULL,
  height INT NOT NULL,

  -- part of image url
  path VARCHAR(255) UNIQUE,

  -- where image bytes are stored
  storage_id VARCHAR(127) NOT NULL,

  -- main focal point/rectangle x,y,w,h
  focus VARCHAR(31),

  -- metadata
  credit VARCHAR(255),
  caption VARCHAR(255)
);

-- article life time
-- created in drafts table
-- drafts -> articles: when it's published first time, moved to articles table.
-- articles -> drafts: published article can be edited.
-- unpublishing an article just sets boolean flag in articles.


CREATE TABLE articles(
  id INT PRIMARY KEY,
  is_published INT NOT NULL, -- 0 for unpublished.
  brand_id INT REFERENCES brands(id) NOT NULL,
  publish_date DATETIME NOT NULL,
  path VARCHAR(255) UNIQUE, -- part of url.

  title VARCHAR(255) NOT NULL,
  primary_image_id INT REFERENCES images(id),
  excerpt TEXT, -- html.
  content TEXT -- html. includes excerpt, primary_image.
);

CREATE TABLE article_authors(
  id INT PRIMARY KEY,
  idx INT NOT NULL, -- multiple authors need to have defined order.
  article_id INT REFERENCES articles(id),
  author_id INT REFERENCES authors(id)
);

CREATE TABLE article_tags(
  id INT PRIMARY KEY,
  idx INT NOT NULL,
  article_id INT REFERENCES articles(id),
  tag_id INT REFERENCES tags(id)
);


-- Draft articles contain many serialized fields.
-- When drafts get published, those are deserialized and normalized.
-- When a published article gets modified, it's moved to drafts as well.
CREATE TABLE drafts(
  id INT PRIMARY KEY,
  article_id INT REFERENCES articles(id), -- Unless new draft, have reference to article.
  brand_id INT REFERENCES brands(id),
  publish_date DATETIME,

  author_ids VARCHAR(255), -- comma separated ids
  tag_ids VARCHAR(255),

  title VARCHAR(255),
  content TEXT -- primary image, excerpt are marked in html.
);
