import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);
dayjs.extend(duration);

export default class Time {
  public static fromNow(date: string): string {
    return dayjs(date).fromNow();
  }
  public static fromNowUnix(unix: number): string {
    return dayjs.unix(unix).fromNow();
  }
  public static toNowUnix(unix: number): string {
    return dayjs.unix(unix).toNow();
  }
  public static humanizeUnix(unix: number): string {
    let now = Date.now() / 1000;
    let diffInHours = Math.abs(unix - now) / 3600;

    return dayjs.duration(diffInHours, "hours").humanize()
  }

  static format(time, format: string) {
    return dayjs(time).format(format);
  }

  static calculateRemainingTimeServerReboot(unixTimestamp): string {
    const now = dayjs();
    const endTime = dayjs.unix(unixTimestamp);
    const duration = dayjs.duration(endTime.diff(now));

    if (duration.asMilliseconds() < 0) {
      return "";
    }

    const minutes = duration.minutes();
    const seconds = duration.seconds();

    return `${minutes}m, ${seconds}s`;
  }

  static formatUnix(unix, format: string) {
    return dayjs.unix(unix).format(format);
  }
}
