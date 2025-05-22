export default class LazyImageLoader {
  public static debug = true;

  public static log(...args: any[]) {
    if (this.debug) {
      // @ts-ignore
      // 1. Convert args to a normal array
      let args = Array.prototype.slice.call(arguments);

      // 2. Prepend log prefix log string
      args.unshift("[LazyImageLoader]");

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

  public static handleRender = LazyImageLoader.debounce(() => {
    // @ts-ignore
    let images   = document.getElementsByClassName("lazy-image");
    for (let i = 0; i < images.length; i++) {
      let image = images.item(i)
      if (image) {
        let dataSrc = image.getAttribute("data-src");
        if (dataSrc && LazyImageLoader.elementInViewport(image)) {
          // @ts-ignore
          image.src = dataSrc;
          image.classList.add('fade-in')
          image.classList.remove('lazy-image-unloaded')
        }
      }
    }

    // @ts-ignore
    let videos   = document.getElementsByClassName("lazy-video");
    for (let i = 0; i < videos.length; i++) {
      let video = videos.item(i)
      if (video) {
        let dataSrc = video.getAttribute("data-src");
        if (dataSrc && LazyImageLoader.elementInViewport(video)) {
          // @ts-ignore
          video.src = dataSrc;
          video.classList.add('fade-in')
          video.classList.remove('lazy-video')
        }
      }
    }
  }, 10);

  public static elementInViewport(elem) {
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

  public static imagePlaying(el) {
    return (el.currentTime > 0 && !el.paused && !el.ended && el.readyState > 2);
  }

  public static imageLoaded(el) {
    return el.readyState === 4
  }

  public static addScrollListener() {
    this.destroyScrollListener()

    window.addEventListener("scroll", LazyImageLoader.handleRender);
    window.addEventListener("resize", LazyImageLoader.handleRender);
  }

  public static destroyScrollListener() {
    window.removeEventListener("scroll", LazyImageLoader.handleRender, false)
    window.removeEventListener("resize", LazyImageLoader.handleRender, false)
  }

}
