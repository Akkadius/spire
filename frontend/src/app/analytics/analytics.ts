import {SpireApi} from "../api/spire-api";

export default class Analytics {
  static trackEvent(name, value) {
    SpireApi.v1().post('/analytics/event', {event_name: name, event_value: value})
  }

  static trackCountsEvent(name, key) {
    SpireApi.v1().post('/analytics/count', {event_name: name, event_key: key})
  }

  static trackAllEvents(name, value) {
    Analytics.trackEvent(name, value)
    Analytics.trackCountsEvent(name, value)
  }
}
