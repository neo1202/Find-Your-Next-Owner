CREATE TABLE message_user_groups (
  id uuid PRIMARY KEY NOT NULL,
  user_id uuid NOT NULL,
  message_partner_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (message_partner_id) REFERENCES users(id) ON DELETE CASCADE
);