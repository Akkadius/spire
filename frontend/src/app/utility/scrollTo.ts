function scrollTo(containerId, targetId, offset = 300) {
  // bring focus to the selected model
  // we queue this on a timeout because elements haven't been rendered yet
  setTimeout(() => {
    // @ts-ignore
    const container = document.getElementById(containerId);
    // @ts-ignore
    const target    = document.getElementById(targetId)
    if (container && target) {
      // @ts-ignore
      container.scrollTop = target.offsetTop - offset;
      console.log("scroll trigger", container, target)
    }
  }, 100)
}

export {scrollTo}
