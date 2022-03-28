type CountEvent = {
  name: string;
  key: string;
}

type Event = {
  name: string;
  value: string;
}

export default class Analytics {
  private static eventsBuffer: Event[]           = []
  private static countEventsBuffer: CountEvent[] = []

  static trackEvent(name, value) {
    this.eventsBuffer.push(
      {
        name: name,
        value: value,
      }
    )

    console.log("events buffer [%s]", this.eventsBuffer.length)
    console.log(this.eventsBuffer)

    // SpireApiClient.v1().post('/analytics/event', {event_name: name, event_value: value})
  }

  static trackCountsEvent(name, key) {
    this.countEventsBuffer.push(
      {
        name: name,
        key: key,
      }
    )

    console.log("count events buffer [%s]", this.countEventsBuffer.length)
    console.log(this.countEventsBuffer)

    // SpireApiClient.v1().post('/analytics/count', {event_name: name, event_key: key})
  }

  static trackAllEvents(name, value) {
    Analytics.trackEvent(name, value)
    Analytics.trackCountsEvent(name, value)
  }
}
