export class WindowManager {
  private static scale = 1;

  public static applyZoom(scale = 0.8) {
    const el = document.body;
    const inverse = 1 / scale * 100;

    WindowManager.scale = scale;

    // @ts-ignore
    el.style.zoom = `${scale}`;
    // el.style.transform = `scale(${scale})`;
    // el.style.transformOrigin = 'top left';
    // el.style.width = `${inverse}%`;
    // el.style.height = `${inverse}%`;
  }

  public static resizeFillScreenElements() {
    const visualHeight = (window.visualViewport && window.visualViewport.height) || window.innerHeight;
    const elements = document.querySelectorAll('.fill-screen');

    elements.forEach(el => {
      const rect = el.getBoundingClientRect();

      // Adjust rect.top to be in unscaled (real layout) units
      const topUnscaled = rect.top / WindowManager.scale;

      const remainingHeight = visualHeight / WindowManager.scale - (topUnscaled + 10); // adjust for margin if needed

      // @ts-ignore
      el.style.maxHeight = `${remainingHeight}px`;
      // @ts-ignore
      el.style.overflow = 'auto';
      // @ts-ignore
      el.style.boxSizing = 'border-box';
    });
  }

  static hookListeners() {
    window.addEventListener('resize', this.resizeFillScreenElements);
    window.addEventListener('DOMContentLoaded', this.resizeFillScreenElements);
    window.addEventListener('load', this.resizeFillScreenElements);
  }
}
