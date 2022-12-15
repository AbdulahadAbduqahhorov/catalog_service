CREATE TABLE IF NOT EXISTS "category" (
	"id" CHAR(36) PRIMARY KEY,
	"title" VARCHAR(255) UNIQUE NOT NULL,
	"created_at" TIMESTAMP DEFAULT now() NOT NULL,
	"updated_at" TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "product" (
	"id" CHAR(36) PRIMARY KEY,
	"title" VARCHAR(255) UNIQUE NOT NULL,
    "description" VARCHAR(255)  NOT NULL,
	"quantity" INT NOT NULL CHECK (quantity >= 0),
    "price" INT NOT NULL CHECK (price > 0),
    "category_id" CHAR(36) REFERENCES category(id) ON DELETE CASCADE,
	"created_at" TIMESTAMP DEFAULT now() NOT NULL,
	"updated_at" TIMESTAMP
);
