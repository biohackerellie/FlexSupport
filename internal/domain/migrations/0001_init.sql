create extension if not exists citext;
create extension if not exists pgcrypto;

create table if not exists users (
  id uuid primary key default gen_random_uuid(),
  name text not null,
  email citext unique not null,
  email_verified boolean not null default false,
  password_hash text,
  is_system_admin boolean not null default false,
  created_at timestamptz not null default now(),
  last_login_at timestamptz
);

create table if not exists tenants (
  id uuid primary key default gen_random_uuid(),
  slug text not null unique,
  name text not null,
  logo_url text,
  timezone text not null default 'UTC',
  created_at timestamptz not null default now(),
  onboarding_completed_at timestamptz
);

create table if not exists tenant_memberships (
  tenant_id uuid not null references tenants (id) on delete cascade,
  user_id uuid not null references users(id) on delete cascade,
  status text not null default 'active', -- active/invited/disabled
  created_at timestamptz not null default now(),
  primary key (tenant_id, user_id)
);

create index on tenant_memberships (user_id);

create table if not exists roles (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants (id) on delete cascade,
  name text not null,
  is_system boolean not null default false,
  unique (tenant_id, name)
);

create table if not exists permissions (
  key text primary key, -- e.g. 'ticket.read', 'ticket.write', 'project.admin', etc.
  description text
);

create table if not exists role_permissions (
  role_id uuid not null references roles(id) on delete cascade,
  permission_key text not null references permissions(key) on delete cascade,
  primary key (role_id, permission_key)
);

create table if not exists membership_roles (
  tenant_id uuid not null,
  user_id uuid not null,
  role_id uuid not null references roles(id) on delete cascade,
  primary key (tenant_id, user_id, role_id),
  foreign key (tenant_id, user_id) references tenant_memberships (tenant_id, user_id) on delete cascade
);

create table if not exists projects (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants (id) on delete cascade,
  key text not null, -- IT, HR, etc
  name text not null,
  description text,
  is_archived boolean not null default false,
  created_at timestamptz not null default now(),
  unique (tenant_id, key),
  unique (tenant_id, name)
);

create index on projects (tenant_id);

create table if not exists project_portal_settings (
  project_id uuid primary key references projects(id) on delete cascade,
  enabled boolean not null default false
);

create table if not exists project_memberships (
  tenant_id uuid not null references tenants(id) on delete cascade,
  project_id uuid not null references projects(id) on delete cascade,
  user_id uuid not null references users(id) on delete cascade,
  created_at timestamptz not null default now(),
  primary key (tenant_id, project_id, user_id)
);

create index on project_memberships (tenant_id, user_id);


create table if not exists request_types (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  project_id uuid not null references projects(id) on delete cascade,

  key text not null,   -- "maintenance_request"
  name text not null,  -- "Maintenance Request"
  description text,
  is_archived boolean not null default false,

  sort_order int not null default 0,

  unique (tenant_id, project_id, key),
  unique (tenant_id, project_id, id)
);

create index on request_types (tenant_id, project_id);
create table if not exists request_type_fields (
  tenant_id uuid not null references tenants(id) on delete cascade,
  request_type_id uuid not null references request_types(id) on delete cascade,
  field_id uuid not null references custom_fields(id) on delete restrict,

  sort_order int not null default 0,

  -- per ticket type behavior
  required boolean not null default false,
  requester_visible boolean not null default true,

  primary key (tenant_id, request_type_id, field_id)
);

create index on request_type_fields (tenant_id, request_type_id);

create table if not exists tickets (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  project_id uuid not null references projects(id) on delete restrict,
  request_type_id uuid not null,

  ticket_number bigint not null,
  title text not null,
  description text,

  status text not null,
  priority text,

  created_by_user_id uuid references users(id),
  requester_user_id uuid references users(id),
  assigned_to_user_id uuid references users(id),

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  closed_at timestamptz,

  unique (project_id, ticket_number),
  constraint tickets_request_type_same_project_fk
    foreign key (tenant_id, project_id, request_type_id)
    references request_types (tenant_id, project_id, id)
    on delete restrict
);

create index if not exists tickets_tenant_project_status_idx
  on tickets (tenant_id, project_id, status);

create index if not exists tickets_tenant_requester_idx
  on tickets (tenant_id, requester_user_id);

create index if not exists tickets_tenant_assigned_idx
  on tickets (tenant_id, assigned_to_user_id);

create table if not exists custom_fields (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,

  key text not null,            -- stable programmatic key: "asset_tag"
  name text not null,           -- display name: "Asset Tag"
  description text,

  field_type text not null,     -- 'text','textarea','number','bool','date','datetime',
                                -- 'select','multiselect','user','group' (future)

  is_archived boolean not null default false,

  created_at timestamptz not null default now(),

  unique (tenant_id, key)
);


create index if not exists custom_fields_tenant_idx
  on custom_fields (tenant_id);


create table if not exists custom_field_options (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  field_id uuid not null references custom_fields(id) on delete cascade,

  value text not null,       -- stable value stored in ticket
  label text not null,       -- display label
  sort_order int not null default 0,
  is_archived boolean not null default false,

  unique (field_id, value)
);

create index on custom_field_options (tenant_id, field_id);

create table if not exists ticket_field_values (
  tenant_id uuid not null references tenants(id) on delete cascade,
  ticket_id uuid not null references tickets(id) on delete cascade,
  field_id uuid not null references custom_fields(id) on delete cascade,

  value jsonb not null, -- e.g. {"text":"abc"} or {"option":"laptop"} or {"user_id":"..."}
  updated_at timestamptz not null default now(),

  primary key (tenant_id, ticket_id, field_id)
);

create index on ticket_field_values (tenant_id, field_id);
create index ticket_field_values_value_gin
  on ticket_field_values using gin (value);

create table if not exists ticket_comments (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  ticket_id uuid not null references tickets(id) on delete cascade,
  author_user_id uuid references users(id),
  body text not null,
  is_internal boolean not null default false,
  created_at timestamptz not null default now()
);

create index on ticket_comments (tenant_id, ticket_id, created_at);

create table if not exists ticket_events (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  ticket_id uuid not null references tickets(id) on delete cascade,
  actor_user_id uuid references users(id),
  type text not null, -- 'status_changed', 'comment_added', etc
  payload jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now()
);

create index on ticket_events (tenant_id, ticket_id, created_at);


create table if not exists tenant_settings (
  tenant_id uuid primary key references tenants(id) on delete cascade,

  public_portal_enabled boolean not null default false,

  portal_title text,
  portal_welcome_text text,

  outbound_from_name text,
  outbound_from_email citext
);

create table if not exists integrations (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  integration_type text not null, -- 'shopify', 'oidc', 'webhook', 'smtp', etc
  name text not null,
  enabled boolean not null default false,
  config jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now(),
  unique (tenant_id, integration_type, name)
);

create index on integrations (tenant_id, integration_type);


create table if not exists tenant_domains (
  id uuid primary key default gen_random_uuid(),
  tenant_id uuid not null references tenants(id) on delete cascade,
  hostname text not null unique, -- "acme.flexsupport.com" or "support.acme.com"
  is_primary boolean not null default false,
  created_at timestamptz not null default now()
);

create index if not exists tenant_domains_tenant_idx
  on tenant_domains (tenant_id);

create table if not exists schema_migrations (
  version integer primary key,
  name text not null,
  applied_at timestamptz not null default now() 
);
