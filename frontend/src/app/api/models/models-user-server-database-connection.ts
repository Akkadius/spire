import { ModelsServerDatabaseConnection } from './models-server-database-connection';
import { ModelsUser } from './models-user';
export interface ModelsUserServerDatabaseConnection {
    active?: number;
    created_at?: string;
    created_by?: number;
    database_connection?: ModelsServerDatabaseConnection;
    deleted_at?: string;
    id?: number;
    server_database_connection_id?: number;
    updated_at?: string;
    user?: ModelsUser;
    user_id?: number;
}
