export class Npcs {
  public static getCleanName(name) {
    name = name.replace(/[^a-z0-9 _]/gi, '')
    name = name.replace("_", " ")
    return name
  }
}
