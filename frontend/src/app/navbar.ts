export class Navbar {
  public static collapsed = false;

  public static expand() {
    if (this.collapsed) {

      let elements = [
        "main-content",
        "navbar",
        "navbar-contents"
      ]

      for (let e of elements) {
        let t = document.getElementsByClassName(e);
        if (t && t[0]) {
          t[0].classList.remove("navbar-collapsed");
        }
      }

      // chevron
      let t = document.getElementById("collapse-nav-chevron")
      if (t) {
        t.style.display = "none"
      }

      this.collapsed = false
    }
  }

  public static collapse() {
    // check if we're on mobile
    if (window.innerWidth < 768) {
      return
    }

    let elements = [
      "main-content",
      "navbar",
      "navbar-contents"
    ]

    for (let e of elements) {
      let t = document.getElementsByClassName(e);
      if (t && t[0]) {
        t[0].classList.add("navbar-collapsed");
      }
    }

    // chevron
    let t = document.getElementById("collapse-nav-chevron")
    if (t) {
      t.style.display = "block"
    }

    this.collapsed = true
  }

  public static toggleCollapse() {
    if (this.collapsed) {
      this.expand()
      return
    }
    this.collapse()
  }
}
