CREATE TABLE
    IF NOT EXISTS public.accounts (
        id uuid NOT NULL,
        name varchar NOT NULL,
        email varchar(100) NOT NULL,
        cell_phone varchar(20) NOT NULL,        
        account_type varchar(10) NOT NULL,
        credit_limit numeric(15) DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        CONSTRAINT accounts_pkey PRIMARY KEY (id)
    );
CREATE INDEX idx_account_type ON accounts(account_type);    