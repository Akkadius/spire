import {TaskActivityApi, TaskApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {TASK_ACTIVITY_TYPE} from "@/app/constants/eq-task-constants";
import {HttpStatus} from "@/app/api/http-status";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Tasks {

  public static async getTasks() {
    const r = await (new TaskApi(SpireApiClient.getOpenApiConfig())).listTasks()
    if (r.status === HttpStatus.OK) {
      return r.data
    }

    return []
  }

  public static getTaskApi() {
    return (new TaskApi(SpireApiClient.getOpenApiConfig()));
  }

  public static getTaskActivitiesApi() {
    return (new TaskActivityApi(SpireApiClient.getOpenApiConfig()));
  }

  // TODO: bubble up error handling
  public static async getTask(taskId: number) {
    let request = (new SpireQueryBuilder())
      .includes(this.getRelationships())
      .get()

    // @ts-ignore
    request.id = taskId
    // @ts-ignore
    const r    = await this.getTaskApi().getTask(request)

    if (r.status === HttpStatus.OK) {
      return r.data
    }

    return {}
  }

  // TODO: bubble up error handling
  public static async getTaskWithActivities(taskId: number) {
    let request = (new SpireQueryBuilder())
      .includes(["TaskActivities"])
      .get()

    // @ts-ignore
    request.id = taskId
    // @ts-ignore

    const r = await this.getTaskApi().getTask(request)

    if (r.status === HttpStatus.OK) {
      return r.data
    }

    return {}
  }

  // TODO: bubble up error handling
  public static async createTask(task: any) {
    // @ts-ignore
    return await this.getTaskApi().createTask({task: task}, {})
  }

  // TODO: bubble up error handling
  public static async updateTask(task: any) {
    // @ts-ignore
    return await this.getTaskApi().updateTask({id: task.id, task: task}, {
      query: (new SpireQueryBuilder())
        .includes(["TaskActivities"])
        .get()
    })
  }

  // TODO: bubble up error handling
  public static async deleteTaskWithActivities(task: any) {
    try {
      const r = await this.getTaskApi().deleteTask({id: task.id}, {})
      if (r.status === HttpStatus.OK) {
        await this.getTaskActivitiesApi().deleteTaskActivity({id: task.id}, {})
      }
    } catch (err) {
      if (err.response.data.error) {
        return err.response.data.error
      }
    }

    return ""
  }

  public static buildActivityDescription(activity: any) {
    if (activity.description_override !== "") {
      return activity.description_override;
    }

    switch (activity.activitytype) {
      case TASK_ACTIVITY_TYPE.DELIVER:
        if (activity.item_list !== "") {
          return activity.item_list
        }

        return "Deliver " + activity.goalcount + " to " + activity.target_name;
      case TASK_ACTIVITY_TYPE.KILL:
        return "Kill " + activity.goalcount + " " + activity.target_name;
      case TASK_ACTIVITY_TYPE.LOOT:
        let string = "Loot " + activity.goalcount;
        if (activity.item_list !== "") {
          string += " " + activity.item_list
        }
        if (activity.target_name !== "") {
          string += " from " + activity.target_name
        }

        return string
      case TASK_ACTIVITY_TYPE.SPEAK_WITH:
        return "Speak with " + activity.target_name;
      case TASK_ACTIVITY_TYPE.EXPLORE:
        return "Explore " + activity.target_name;
      case TASK_ACTIVITY_TYPE.TRADESKILL:
        return "Create " + activity.goalcount + " " + activity.target_name;
      case TASK_ACTIVITY_TYPE.FISH:
        return "Fish " + activity.goalcount;
      case TASK_ACTIVITY_TYPE.FORAGE:
        return "Forage " + activity.goalcount;
      case TASK_ACTIVITY_TYPE.USE:
        return "Use " + activity.goalcount;
      case TASK_ACTIVITY_TYPE.USE2:
        return "Use " + activity.goalcount;
      case TASK_ACTIVITY_TYPE.TOUCH:
        return "Touch " + activity.target_name;
      case TASK_ACTIVITY_TYPE.GIVE:
        return "Give " + activity.goalcount + " to " + activity.target_name;
      case TASK_ACTIVITY_TYPE.QUEST_SCRIPT_DRIVEN:
        return "None";
      default:
        return "None";
    }
  }

  public static getRelationships() {
    return [
      "TaskActivities",
      // "TaskActivities.Goallists",
      // "TaskActivities.NpcType",
      // "TaskActivities.NpcType.Spawnentries",
      // "TaskActivities.NpcType.Spawnentries.NpcType",
      // "TaskActivities.NpcType.Spawnentries.Spawngroup",
      // "TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
      // "Tasksets",
    ]
  }

  public static getFieldDescriptions() {
    return {
      "id": "Task identifier",
    }
  }

  public static getFieldDescription(field: string) {
    // we do this because the payload we get back from spire API is
    // formatted slightly different
    let fieldLookup = field.toLowerCase().replace("_", "")

    for (let key in this.getFieldDescriptions()) {
      let keyLookup = key.toLowerCase().replace("_", "")
      if (keyLookup === fieldLookup) {
        return this.getFieldDescriptions()[key]
      }
    }

    return ''
  }

  // used to seed "new" tasks in the task editor
  public static getExampleTask() {
    return {
      "id": 2,
      "type": 2,
      "duration": 3600,
      "duration_code": 0,
      "title": "Example Task 2",
      "description": "Example Task 2",
      "reward": "XP",
      "rewardid": 0,
      "cashreward": 0,
      "xpreward": 10,
      "rewardmethod": 2,
      "reward_radiant_crystals": 0,
      "reward_ebon_crystals": 0,
      "minlevel": 0,
      "maxlevel": 0,
      "level_spread": 0,
      "min_players": 0,
      "max_players": 0,
      "repeatable": 1,
      "faction_reward": 0,
      "completion_emote": "",
      "replay_timer_seconds": 0,
      "request_timer_seconds": 0,
      "task_activities": [{
        "taskid": 2,
        "activityid": 0,
        "step": 1,
        "activitytype": 2,
        "target_name": "Orcs",
        "item_list": "",
        "skill_list": "-1",
        "spell_list": "0",
        "description_override": "",
        "goalid": 1,
        "goalmethod": 1,
        "goalcount": 5,
        "delivertonpc": 0,
        "zones": "21",
        "optional": 0
      }, {
        "taskid": 2,
        "activityid": 1,
        "step": 1,
        "activitytype": 3,
        "target_name": "any creature",
        "item_list": "Rusty Items",
        "skill_list": "-1",
        "spell_list": "0",
        "description_override": "",
        "goalid": 2,
        "goalmethod": 1,
        "goalcount": 3,
        "delivertonpc": 0,
        "zones": "0",
        "optional": 0
      }, {
        "taskid": 2,
        "activityid": 2,
        "step": 1,
        "activitytype": 5,
        "target_name": "",
        "item_list": "",
        "skill_list": "-1",
        "spell_list": "0",
        "description_override": "Locate the Antonica Spires in the Luclin Nexus",
        "goalid": 1,
        "goalmethod": 0,
        "goalcount": 1,
        "delivertonpc": 0,
        "zones": "152",
        "optional": 0
      }]
    }

  }
}
