import {TaskApi} from "@/app/api";
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

  public static async getTask(taskId: number) {
    let request = (new SpireQueryBuilder())
      .includes(this.getRelationships())
      .get()

    // @ts-ignore
    request.id = taskId
    // @ts-ignore
    const r    = await (new TaskApi(SpireApiClient.getOpenApiConfig())).getTask(request)

    if (r.status === HttpStatus.OK) {
      return r.data
    }

    return {}
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
      "TaskActivities.Goallists",
      "TaskActivities.NpcType",
      "TaskActivities.NpcType.Spawnentries",
      "TaskActivities.NpcType.Spawnentries.NpcType",
      "TaskActivities.NpcType.Spawnentries.Spawngroup",
      "TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
      "Tasksets",
    ]
  }
}
