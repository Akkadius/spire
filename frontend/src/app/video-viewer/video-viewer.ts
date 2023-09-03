export default class VideoViewer {
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

  public static debounce(func, delay) {
    let debounceTimer;
    return function () {
      // @ts-ignore
      const context = this;
      const args    = arguments;
      clearTimeout(debounceTimer);
      debounceTimer = setTimeout(() => func.apply(context, args), delay);
    };
  }

  public static stopAll() {
    let videos = document.getElementsByClassName("video-preview");
    for (let i = 0; i < videos.length; i++) {
      let video = <HTMLVideoElement>videos.item(i)
      if (video) {
        video.pause()
      }
    }
  }

  public static handleRender = VideoViewer.debounce(() => {
    // @ts-ignore
    VideoViewer.log("Render")

    let unloading = []
    let playing   = []
    let loaded    = []
    let videos    = document.getElementsByClassName("video-preview");
    for (let i = 0; i < videos.length; i++) {
      let video = <HTMLVideoElement>videos.item(i)
      if (video) {
        let dataSrc = video.getAttribute("data-src")

        if (VideoViewer.videoLoaded(video)) {
          // @ts-ignore
          loaded.push(video.getAttribute("id"))
        }

        // Toggle playing
        if (VideoViewer.elementInViewport(video, -500)) {
          if (dataSrc && !VideoViewer.videoLoaded(video)) {
            let source = document.createElement("source");

            // video.setAttribute("src", dataSrc);
            // video.removeAttribute("data-src");
            // video.pause()
            video.innerHTML = "";
            video.removeAttribute("src");

            source.setAttribute("src", dataSrc);
            source.setAttribute("type", "video/mp4");
            video.appendChild(source);
            video.load();
            video.play().then(() => {
              let videoId = video.getAttribute("data-video-id")
              if (videoId) {
                let videoEl = document.getElementById('overlay-' + videoId)
                if (videoEl) {
                  videoEl.style.display = "block";
                }
              }
            });

            // @ts-ignore
            playing.push(video.getAttribute("id"))
          }
        } else if (VideoViewer.videoPlaying(video) && video.hasChildNodes()) {
          // @ts-ignore
          unloading.push(video.getAttribute("id"))
          video.pause()
          video.innerHTML = "";
          video.removeAttribute("src");
          video.load()
        }
      }
    }

    if (playing.length > 0) {
      VideoViewer.log("Playing", playing)
    }
    if (unloading.length > 0) {
      VideoViewer.log("Unloading", unloading)
    }
    if (loaded.length > 0) {
      VideoViewer.log("Loaded", loaded)
    }

  }, 10);

  public static elementInViewport(elm, threshold = 0, mode = "visible") {
    let rect       = elm.getBoundingClientRect();
    let viewHeight = Math.max(document.documentElement.clientHeight, window.innerHeight);
    let above      = rect.bottom - threshold < 0;
    let below      = rect.top - viewHeight + threshold >= 0;

    return mode === 'above' ? above : (mode === 'below' ? below : !above && !below);
  }

  public static videoPlaying(el) {
    return (el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2);
  }

  public static videoLoaded(el) {

    const src = el.getElementsByTagName("source")
    if (!src) {
      return false
    }

    if (src.length === 0) {
      return false
    }

    return el.getElementsByTagName("source")[0].src.length > 0
  }

  public static addScrollListener() {
    this.destroyScrollListener()

    window.addEventListener("scroll", VideoViewer.handleRender);
    window.addEventListener("resize", VideoViewer.handleRender);
  }

  public static destroyScrollListener() {
    window.removeEventListener("scroll", VideoViewer.handleRender, false)
    window.removeEventListener("resize", VideoViewer.handleRender, false)
  }

}
