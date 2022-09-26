import { ModelsRuleValue } from './models-rule-value';
export interface ModelsRuleSet {
    name?: string;
    rule_values?: Array<ModelsRuleValue>;
    ruleset_id?: number;
}
