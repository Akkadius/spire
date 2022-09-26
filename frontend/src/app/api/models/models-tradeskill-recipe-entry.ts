import { ModelsTradeskillRecipe } from './models-tradeskill-recipe';
export interface ModelsTradeskillRecipeEntry {
    componentcount?: number;
    failcount?: number;
    id?: number;
    iscontainer?: number;
    item_id?: number;
    recipe_id?: number;
    salvagecount?: number;
    successcount?: number;
    tradeskill_recipe?: ModelsTradeskillRecipe;
}
