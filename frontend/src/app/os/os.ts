export class OS {
  static get() {
    // @ts-ignore
    let userAgent = window.navigator.userAgent;
    // @ts-ignore
    let platform  = window.navigator?.userAgentData?.platform || window.navigator.platform;
    let macosPlatforms   = ['Macintosh', 'MacIntel', 'MacPPC', 'Mac68K']
    let windowsPlatforms = ['Win32', 'Win64', 'Windows', 'WinCE']
    let iosPlatforms     = ['iPhone', 'iPad', 'iPod']
    let os               = ''

    if (macosPlatforms.indexOf(platform) !== -1) {
      os = 'Mac OS';
    } else if (iosPlatforms.indexOf(platform) !== -1) {
      os = 'iOS';
    } else if (windowsPlatforms.indexOf(platform) !== -1) {
      os = 'Windows';
    } else if (/Android/.test(userAgent)) {
      os = 'Android';
    } else if (/Linux/.test(platform)) {
      os = 'Linux';
    }

    return os;
  }
}
