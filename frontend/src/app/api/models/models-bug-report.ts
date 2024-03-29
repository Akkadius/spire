export interface ModelsBugReport {
    _can_duplicate?: number;
    _character_flags?: number;
    _crash_bug?: number;
    _target_info?: number;
    _unknown_value?: number;
    account_id?: number;
    bug_report?: string;
    bug_status?: number;
    category_id?: number;
    category_name?: string;
    character_id?: number;
    character_name?: string;
    client_version_id?: number;
    client_version_name?: string;
    heading?: number;
    id?: number;
    last_review?: string;
    last_reviewer?: string;
    optional_info_mask?: number;
    pos_x?: number;
    pos_y?: number;
    pos_z?: number;
    report_datetime?: string;
    reporter_name?: string;
    reporter_spoof?: number;
    reviewer_notes?: string;
    system_info?: string;
    target_id?: number;
    target_name?: string;
    time_played?: number;
    ui_path?: string;
    zone?: string;
}
