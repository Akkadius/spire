function scrollToTarget(containerId, targetId, offset = 300) {
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
    }
  }, 100)
}

function scrollToHash(route) {
  const hash       = route.hash.replace("#", "");
  const hashTarget = document.getElementById(hash)
  if (hashTarget) {
    const y = hashTarget.getBoundingClientRect().top + window.scrollY;
    window.scroll(
      {
        top: y,
        behavior: 'smooth'
      }
    );
  }
}

export {scrollToTarget, scrollToHash};
