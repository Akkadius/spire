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
      let task = r.data

      // @ts-ignore
      // sort activityids before return
      if (task.task_activities) {
        // @ts-ignore
        task.task_activities.sort((a, b) => (a.activityid > b.activityid) ? 1 : -1)
      }

      return task
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

  public static async createNewTaskActivity(task: any) {
    let taskActivity        = this.getExampleActivity()
    taskActivity.taskid     = task.id
    taskActivity.activityid = (task.task_activities ? this.getNextActivityId(task.task_activities) : 0)
    taskActivity.step       = (task.task_activities ? this.getLatestStep(task.task_activities) : 1)

    // @ts-ignore
    return await this.getTaskActivitiesApi()
      .createTaskActivity(
        {
          taskActivity: taskActivity
        }
      )
  }

  public static async updateTaskActivityId(taskActivity: any, previousId: number) {
    let request = (new SpireQueryBuilder())
      .where("taskid", "=", taskActivity.taskid)
      .where("activityid", "=", previousId)
      .get()

    // format zones
    taskActivity.zones = taskActivity.zones.toString()

    // @ts-ignore
    return await this.getTaskActivitiesApi()
      .updateTaskActivity(
        {
          id: taskActivity.taskid,
          taskActivity: taskActivity
        },
        {query: request}
      )
  }

  public static async cloneTaskActivity(task: any, sourceActivityId) {
    let taskActivity        = task.task_activities[sourceActivityId]
    taskActivity.taskid     = task.id
    taskActivity.activityid = (task.task_activities ? this.getNextActivityId(task.task_activities) : 0)
    taskActivity.step       = (task.task_activities ? this.getLatestStep(task.task_activities) : 1)

    // @ts-ignore
    return await this.getTaskActivitiesApi()
      .createTaskActivity(
        {
          taskActivity: taskActivity
        }
      )
  }

  // TODO: bubble up error handling
  public static async deleteTaskActivity(activity: any) {
    let request = (new SpireQueryBuilder())
      .where("activityid", "=", activity.activityid)
      .get()

    // @ts-ignore
    request.id = activity.taskid

    // @ts-ignore
    return await this.getTaskActivitiesApi()
      .deleteTaskActivity(
        {id: activity.taskid},
        {query: request}
      )
  }

  private static getNextActivityId(activities: any) {
    let nextId = 0
    activities.forEach((a) => {
      nextId = a.activityid
    })

    return nextId + 1
  }

  private static getLatestStep(activities: any) {
    let step = 0
    activities.forEach((a) => {
      step = a.step
    })

    return step
  }

  public static buildActivityDescription(activity: any) {
    if (activity.description_override !== "") {
      return activity.description_override;
    }

    switch (activity.activitytype) {
      case TASK_ACTIVITY_TYPE.DELIVER:
        return `Deliver ${activity.goalcount} ${activity.item_list} to ${activity.target_name}`;
      case TASK_ACTIVITY_TYPE.KILL:
        return `Kill ${activity.goalcount} ${activity.target_name}`;
      case TASK_ACTIVITY_TYPE.LOOT:
        return `Loot ${activity.goalcount} ${activity.item_list} from ${activity.target_name}`;
      case TASK_ACTIVITY_TYPE.SPEAK_WITH:
        return "Speak with " + activity.target_name;
      case TASK_ACTIVITY_TYPE.EXPLORE:
        return "Explore " + activity.target_name;
      case TASK_ACTIVITY_TYPE.TRADESKILL:
        return `Create ${activity.goalcount} ${activity.item_list} using tradeskills`
      case TASK_ACTIVITY_TYPE.FISH:
        return `Fish for ${activity.goalcount} ${activity.item_list}`;
      case TASK_ACTIVITY_TYPE.FORAGE:
        return `Forage ${activity.goalcount} ${activity.item_list}`;
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
      "step": "This is the logical step of your task activity, you can have many activities in one step, you have to complete all activities in one step to unlock the next step",
      "activitytype": "This is the type of task activity, kill, loot etc.",
      "item_list": "List of items related to this activity.",
      "target_name": "This describes what the activity is targeting, it is different depending on the activity. For example 'orcs' Would display 'Kill X orcs'",
      "optional": "Describes whether or not this activity is optional",
      "goalid": "Goal ID is different depending on your activity type as well as Goal Method",
      "goalmethod": "The type of goal used for this activity.",
      "goalcount": "Quantity of the goal to check for.",
      "delivertonpc": "NPCID of the NPC to deliver the item to.",
      "zones": "Zone ID related to this activity.",
      "description_override": "Use this to completely override other fields that otherwise influence the activity description (activity target, item list etc.)",
      "title": "The title of your task",
      "type": "The type of task you are making",
      "description": "Narrative description of your task",
      "duration_code": "Reflects the type of duration",
      "duration": "Task duration in seconds",
      "minlevel": "Minimum level to recieve the task",
      "maxlevel": "Maximum level to recieve the task",
      "completion_emote": "Emote text that is displayed upon succesful completion of the task",
      "rewardmethod": "Method used for providing reward(s)",
      "reward": "Name of reward provided for completing the task",
      "rewardid": "Item ID for the rewarded item",
      "xpreward": "Amount of experience rewarded",
      "faction_reward": "Amount of faction rewarded",
      "reward_ebon_crystals": "Number of ebon crystals rewarded",
      "reward_radiant_crystals": "Number of radiant crystals rewarded",
      "cashreward": "Amount of coin rewarded in copper",
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
      "rewardmethod": 0,
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
        "goalid": 0,
        "goalmethod": 0,
        "goalcount": 1,
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
        "goalid": 0,
        "goalmethod": 0,
        "goalcount": 1,
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
        "goalid": 0,
        "goalmethod": 0,
        "goalcount": 1,
        "delivertonpc": 0,
        "zones": "152",
        "optional": 0
      }]
    }
  }

  public static getExampleActivity() {
    return {
      "taskid": 2,
      "activityid": 0,
      "step": 1,
      "activitytype": 2,
      "target_name": "Orcs",
      "item_list": "",
      "skill_list": "-1",
      "spell_list": "0",
      "description_override": "",
      "goalid": 0,
      "goalmethod": 0,
      "goalcount": 1,
      "delivertonpc": 0,
      "zones": "21",
      "optional": 0
    }
  }

}
