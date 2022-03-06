import {ZoneApi} from "@/app/api";
import {SpireApiClient} from "@/app/api/spire-api-client";
import {SpireQueryBuilder} from "@/app/api/spire-query-builder";

export class Zones {
  public static zones = <any>[]
  public static zonesByShortName = <any>{}

  public static async getZones() {
    if (this.zones && this.zones.length > 0) {
      return this.zones
    }

    const result = await (new ZoneApi(SpireApiClient.getOpenApiConfig()))
      .listZones(
        // @ts-ignore
        (new SpireQueryBuilder())
          .orderBy(["expansion", "short_name"])
          .limit(10000)
          .get()
      )
    if (result.status === 200) {
      this.zones = result.data

      for (let zone of this.zones) {
        this.zonesByShortName[zone.short_name] = zone
      }

      return result.data
    }

    return {}
  }

  public static async getZoneLongNameByShortName(shortName: string) {
    shortName = shortName.toLowerCase()
    if (this.zonesByShortName[shortName]) {
      return this.zonesByShortName[shortName].long_name
    }

    const zones = (await this.getZones())
    for (const zone of zones) {
      if (zone.short_name === shortName) {
        return zone.long_name
      }
    }

    return ""
  }
}
