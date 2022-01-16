export default class VideoViewer {
  public static rendererEventListener: null;

  public static debug = true;

  public static log(...args: any[]) {
    if (this.debug) {
      // @ts-ignore
      // 1. Convert args to a normal array
      let args = Array.prototype.slice.call(arguments);

      // 2. Prepend log prefix log string
      args.unshift("[VideoViewer]");

      // 3. Pass along arguments to console.log
      // @ts-ignore
      console.log.apply(console, args);
    }
  }

  public static handleRender() {
    // @ts-ignore
    VideoViewer.log("Render")

    let playing  = []
    let stopping = []
    let videos   = document.getElementsByClassName("video-preview");
    for (let i = 0; i < videos.length; i++) {
      let video = <HTMLVideoElement>videos.item(i)
      if (video) {
        let source  = document.createElement("source");
        let dataSrc = video.getAttribute("data-src")

        // Toggle playing
        if (VideoViewer.elementInViewport(video)) {
          if (dataSrc) {

            // video.setAttribute("src", dataSrc);
            video.removeAttribute("data-src");
            video.pause()
            video.innerHTML = "";
            video.removeAttribute("src");

            source.setAttribute("src", dataSrc);
            source.setAttribute("type", "video/mp4");
            video.appendChild(source);
            video.load();
            video.play();
          }

          if (!VideoViewer.videoPlaying(video) && VideoViewer.videoLoaded(video)) {
            video.play()
            // @ts-ignore
            playing.push(video.getAttribute("id"))
          }
        } else {
          if (VideoViewer.videoPlaying(video) && VideoViewer.videoLoaded(video)) {
            video.pause()
            // @ts-ignore
            stopping.push(video.getAttribute("id"))
          }
        }
      }
    }

    if (playing.length > 0) {
      VideoViewer.log("Playing", playing)
    }
    if (stopping.length > 0) {
      VideoViewer.log("Stopping", stopping)
    }
  }

  public static elementInViewport(el) {
    let top    = el.offsetTop;
    let left   = el.offsetLeft;
    let width  = el.offsetWidth;
    let height = el.offsetHeight;

    while (el.offsetParent) {
      el = el.offsetParent;
      top += el.offsetTop;
      left += el.offsetLeft;
    }

    return (
      top < (window.pageYOffset + window.innerHeight) &&
      left < (window.pageXOffset + window.innerWidth) &&
      (top + height) > window.pageYOffset &&
      (left + width) > window.pageXOffset
    );
  }

  public static videoPlaying(el) {
    return !!(el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2);
  }

  public static videoLoaded(el) {
    return el.readyState === 4
  }

  public static addScrollListener() {
    this.destroyScrollListener()

    window.addEventListener("scroll", VideoViewer.handleRender);
  }

  public static destroyScrollListener() {
    window.removeEventListener("scroll", VideoViewer.handleRender, false)
  }

}
