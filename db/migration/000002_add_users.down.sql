ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "owner_currnecy_key";

ALTER TABLE "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

DROP TABLE IF EXISTS users;