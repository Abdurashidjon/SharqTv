CREATE TABLE IF NOT EXISTS admin_roles (
     id uuid PRIMARY KEY,
     title VARCHAR(50) NOT NULL UNIQUE,
     control_category BOOLEAN NOT NULL DEFAULT FALSE,
     control_movie_creators BOOLEAN NOT NULL DEFAULT FALSE,
     control_favorite_movie BOOLEAN NOT NULL DEFAULT FALSE,
     control_dashboard BOOLEAN NOT NULL DEFAULT FALSE,
     control_tariff BOOLEAN NOT NULL DEFAULT FALSE,
     control_permissions BOOLEAN NOT NULL DEFAULT FALSE,
     control_movie BOOLEAN NOT NULL DEFAULT FALSE,
     control_notification BOOLEAN NOT NULL DEFAULT FALSE,
     control_admin BOOLEAN NOT NULL DEFAULT FALSE,
     control_genre BOOLEAN NOT NULL DEFAULT FALSE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS admins (
     id uuid PRIMARY KEY,
     access_token text NOT NULL UNIQUE,
     name VARCHAR(50) NOT NULL UNIQUE,
     username VARCHAR(20) NOT NULL UNIQUE,
     password VARCHAR NOT NULL,
     phone TEXT [] NOT NULL UNIQUE,
     active BOOLEAN NOT NULL DEFAULT TRUE,
     address TEXT not null,
     description TEXT,
     logo VARCHAR(50),
     role_id uuid,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP,
     CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES admin_roles(id) ON DELETE
     SET NULL
);
