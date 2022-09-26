import { ModelsServerDatabaseConnection } from './models-server-database-connection';
import { ModelsUserServerDatabaseConnection } from './models-user-server-database-connection';
export interface ModelsUser {
    avatar?: string;
    created_at?: string;
    email?: string;
    first_name?: string;
    full_name?: string;
    id?: number;
    last_name?: string;
    owned_connections?: Array<ModelsServerDatabaseConnection>;
    provider?: string;
    updated_at?: string;
    user_connections?: Array<ModelsUserServerDatabaseConnection>;
    user_name?: string;
}
