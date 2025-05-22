export const TASK_TYPES = {
  0: "Task",
  1: "Shared Task",
  2: "Quest",
  3: "Expedition",
}

export const TASK_TYPE = {
  TASK: 0,
  SHARED_TASK: 1,
  QUEST: 2,
  EXPEDITION: 3,
}

export const TASK_DURATION_TYPES = {
  0: "None",
  1: "Short",
  2: "Medium",
  3: "Long",
}

export const TASK_ACTIVITY_TYPES = {
  1: "Deliver",
  2: "Kill",
  3: "Loot",
  4: "Speak With",
  5: "Explore",
  6: "Tradeskill",
  7: "Fish",
  8: "Forage",
  9: "Use",
  10: "Use",
  11: "Touch",
  100: "Give Cash",
  255: "Quest Driven",
}

export const TASK_ACTIVITY_TYPE = {
  DELIVER: 1,
  KILL: 2,
  LOOT: 3,
  SPEAK_WITH: 4,
  EXPLORE: 5,
  TRADESKILL: 6,
  FISH: 7,
  FORAGE: 8,
  USE: 9,
  USE2: 10,
  TOUCH: 11,
  GIVE: 100,
  QUEST_SCRIPT_DRIVEN: 255,
};

export const TASK_DURATION_HUMAN = {
  0: "Infinite",
  3600: "1 hour",
  7200: "2 hours",
  10800: "3 hours",
  14400: "4 hours",
  18000: "5 hours",
  21600: "6 hours",
  25200: "7 hours",
  28800: "8 hours",
  32400: "9 hours",
  36000: "10 hours",
  39600: "11 hours",
  43200: "12 hours",
  86400: "1 days",
  172800: "2 days",
  259200: "3 days",
  345600: "4 days",
  432000: "5 days",
  518400: "6 days",
  604800: "7 days",
  2629746: "1 month",
}

export const TASK_REWARD_METHOD = {
  0: "Single Item ID",
  1: "List of Items (Goal List)",
  2: "Quest Controlled",
  255: "Quest Controlled",
}

export const TASK_GOAL_METHOD_TYPE = {
  0: "Single ID",
  1: "List (of entries)",
  2: "Quest Controlled",
}

export const TASK_GOAL_METHOD_TYPES = {
  SINGLE_ID: 0,
  LIST: 1,
  QUEST_CONTROLLED: 2,
}

export const TASK_REWARD_METHOD_TYPE = {
  0: "Single ID",
  1: "List (of entries)",
  2: "Quest Controlled",
}
