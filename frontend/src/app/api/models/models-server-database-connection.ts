import { ModelsUserServerDatabaseConnection } from './models-user-server-database-connection';
export interface ModelsServerDatabaseConnection {
    content_db_host?: string;
    content_db_name?: string;
    content_db_port?: string;
    content_db_username?: string;
    created_at?: string;
    created_by?: number;
    created_from_ip?: string;
    db_host?: string;
    db_name?: string;
    db_port?: string;
    db_username?: string;
    deleted_at?: string;
    id?: number;
    logs_db_host?: string;
    logs_db_name?: string;
    logs_db_port?: string;
    logs_db_username?: string;
    name?: string;
    updated_at?: string;
    user_server_database_connections?: Array<ModelsUserServerDatabaseConnection>;
}
