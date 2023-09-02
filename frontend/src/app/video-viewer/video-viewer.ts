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
    let videos    = document.getElementsByClassName("video-preview");
    for (let i = 0; i < videos.length; i++) {
      let video = <HTMLVideoElement>videos.item(i)
      if (video) {
        let dataSrc = video.getAttribute("data-src")

        // Toggle playing
        if (VideoViewer.elementInViewport(video)) {
          if (dataSrc && !VideoViewer.videoLoaded(video) && !VideoViewer.videoPlaying(video)) {
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
            video.play();

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
  }, 10);

  public static elementInViewport(elem) {
    // if (!(elem instanceof Element)) throw Error('DomUtil: elem is not an element.');


    const style = getComputedStyle(elem);
    if (style.display === 'none') return false;
    if (style.visibility !== 'visible') return false;
    // @ts-ignore
    if (style.opacity < 0.1) return false;
    // @ts-ignore
    if (elem.offsetWidth + elem.offsetHeight + elem.getBoundingClientRect().height +
      elem.getBoundingClientRect().width === 0) {
      return false;
    }
    const elemCenter = {
      x: elem.getBoundingClientRect().left + elem.offsetWidth / 2,
      y: elem.getBoundingClientRect().top + elem.offsetHeight / 4
    };
    if (elemCenter.x < 0) return false;
    if (elemCenter.x > (document.documentElement.clientWidth || window.innerWidth)) return false;
    if (elemCenter.y < 0) return false;
    if (elemCenter.y > (document.documentElement.clientHeight || window.innerHeight)) return false;
    let pointContainer = document.elementFromPoint(elemCenter.x, elemCenter.y);
    // @ts-ignore
    do {
      if (pointContainer === elem) return true;
    } while (pointContainer && pointContainer === pointContainer.parentNode);

    return false;
  }

  public static videoPlaying(el) {
    return (el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2);
  }

  public static videoLoaded(el) {
    return el.readyState === 4
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
